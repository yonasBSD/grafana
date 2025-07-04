package setting

import (
	"regexp"

	"github.com/grafana/grafana/pkg/services/kmsproviders"
)

type SecretsManagerSettings struct {
	SecretKey          string
	EncryptionProvider string
	AvailableProviders []string
}

func (cfg *Cfg) readSecretsManagerSettings() {
	secretsMgmt := cfg.Raw.Section("secrets_manager")
	cfg.SecretsManagement.EncryptionProvider = secretsMgmt.Key("encryption_provider").MustString(kmsproviders.Default)

	// TODO: These are not used yet by the secrets manager because we need to distentagle the dependencies with OSS.
	cfg.SecretsManagement.SecretKey = secretsMgmt.Key("secret_key").MustString("")
	cfg.SecretsManagement.AvailableProviders = regexp.MustCompile(`\s*,\s*`).Split(secretsMgmt.Key("available_encryption_providers").MustString(""), -1) // parse comma separated list
}
