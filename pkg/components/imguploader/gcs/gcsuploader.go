// Package gcs provides an image uploader for GCS.
package gcs

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"
	"github.com/grafana/grafana/pkg/ifaces/gcsifaces"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/util"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
)

// NewUploader returns a new Uploader.
func NewUploader(keyFile, bucket, path string, enableSignedURLs bool, signedURLExpiration time.Duration) (*Uploader, error) {
	if signedURLExpiration <= 0 {
		return nil, fmt.Errorf("invalid signed URL expiration: %q", signedURLExpiration)
	}
	uploader := &Uploader{
		KeyFile:             keyFile,
		Bucket:              bucket,
		path:                path,
		log:                 log.New("gcsuploader"),
		enableSignedURLs:    enableSignedURLs,
		signedURLExpiration: signedURLExpiration,
	}

	uploader.log.Debug("Created uploader", "key", keyFile, "bucket", bucket, "path", path, "enableSignedUrls",
		enableSignedURLs, "signedUrlExpiration", signedURLExpiration.String())

	return uploader, nil
}

// newClient returns a new GCS client.
// Stubbable by tests.
var newClient = func(ctx context.Context, opts ...option.ClientOption) (gcsifaces.StorageClient, error) {
	client, err := storage.NewClient(ctx, opts...)
	return clientWrapper{client}, err
}

// Uploader supports uploading images to GCS.
type Uploader struct {
	KeyFile             string
	Bucket              string
	path                string
	log                 log.Logger
	enableSignedURLs    bool
	signedURLExpiration time.Duration
}

// Upload uploads an image to GCS.
func (u *Uploader) Upload(ctx context.Context, imageDiskPath string) (string, error) {
	fileName, err := util.GetRandomString(20)
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(imageDiskPath)
	if ext == "" {
		ext = ".png"
	}
	fileName += ext

	key := path.Join(u.path, fileName)

	var keyData []byte
	if u.KeyFile != "" {
		u.log.Debug("Opening key file ", u.KeyFile)
		keyData, err = os.ReadFile(u.KeyFile)
		if err != nil {
			return "", err
		}
	}

	const scope = storage.ScopeReadWrite

	var client gcsifaces.StorageClient
	if u.KeyFile != "" {
		u.log.Debug("Creating Google credentials from JSON")
		creds, err := google.CredentialsFromJSON(ctx, keyData, scope)
		if err != nil {
			return "", err
		}

		u.log.Debug("Creating GCS client")
		client, err = newClient(ctx, option.WithCredentials(creds))
		if err != nil {
			return "", err
		}
	} else {
		u.log.Debug("Creating GCS client with default application credentials")
		client, err = newClient(ctx, option.WithScopes(scope))
		if err != nil {
			return "", err
		}
	}

	if err := u.uploadFile(ctx, client, imageDiskPath, key); err != nil {
		return "", err
	}

	if !u.enableSignedURLs {
		return fmt.Sprintf("https://storage.googleapis.com/%s/%s", u.Bucket, key), nil
	}

	u.log.Debug("Signing GCS URL")
	var jwtData []byte
	if u.KeyFile != "" {
		jwtData = keyData
	} else {
		creds, err := client.FindDefaultCredentials(ctx, scope)
		if err != nil {
			return "", fmt.Errorf("failed to find default Google credentials: %s", err)
		}
		jwtData = creds.JSON
	}
	conf, err := client.JWTConfigFromJSON(jwtData)
	if err != nil {
		return "", err
	}
	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         "GET",
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().Add(u.signedURLExpiration),
	}
	signedURL, err := client.SignedURL(u.Bucket, key, opts)
	if err != nil {
		return "", err
	}

	return signedURL, nil
}

func (u *Uploader) uploadFile(
	ctx context.Context,
	client gcsifaces.StorageClient,
	imageDiskPath,
	key string,
) error {
	u.log.Debug("Opening image file", "path", imageDiskPath)

	// We can ignore the gosec G304 warning on this one because `imageDiskPath` comes
	// from alert notifiers and is only used to upload images generated by alerting.
	// nolint:gosec
	fileReader, err := os.Open(imageDiskPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := fileReader.Close(); err != nil {
			u.log.Warn("Failed to close file", "err", err, "path", imageDiskPath)
		}
	}()

	// Set public access if not generating a signed URL
	pubAcc := !u.enableSignedURLs

	u.log.Debug("Uploading to GCS bucket using SDK", "bucket", u.Bucket, "key", key, "public", pubAcc)

	uri := fmt.Sprintf("gs://%s/%s", u.Bucket, key)

	wc := client.Bucket(u.Bucket).Object(key).NewWriter(ctx)
	if pubAcc {
		wc.SetACL("publicRead")
	}
	if _, err := io.Copy(wc, fileReader); err != nil {
		_ = wc.Close()
		return fmt.Errorf("failed to upload to %s: %s", uri, err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("failed to upload to %s: %s", uri, err)
	}

	return nil
}

type clientWrapper struct {
	client *storage.Client
}

func (c clientWrapper) Bucket(key string) gcsifaces.StorageBucket {
	return bucketWrapper{c.client.Bucket(key)}
}

func (c clientWrapper) FindDefaultCredentials(ctx context.Context, scope string) (*google.Credentials, error) {
	return google.FindDefaultCredentials(ctx, scope)
}

func (c clientWrapper) JWTConfigFromJSON(keyJSON []byte) (*jwt.Config, error) {
	return google.JWTConfigFromJSON(keyJSON)
}

func (c clientWrapper) SignedURL(bucket, name string, opts *storage.SignedURLOptions) (string, error) {
	return storage.SignedURL(bucket, name, opts)
}

type bucketWrapper struct {
	bucket *storage.BucketHandle
}

func (b bucketWrapper) Object(key string) gcsifaces.StorageObject {
	return objectWrapper{b.bucket.Object(key)}
}

type objectWrapper struct {
	object *storage.ObjectHandle
}

func (o objectWrapper) NewWriter(ctx context.Context) gcsifaces.StorageWriter {
	return writerWrapper{o.object.NewWriter(ctx)}
}

type writerWrapper struct {
	*storage.Writer
}

func (w writerWrapper) SetACL(acl string) {
	w.PredefinedACL = acl
}
