import { Trans, t } from '@grafana/i18n';
import { Tooltip, Button, Stack } from '@grafana/ui';
import { DashboardInteractions } from 'app/features/dashboard-scene/utils/interactions';

type VersionsButtonsType = {
  hasMore: boolean;
  canCompare: boolean;
  getVersions: (append: boolean) => void;
  getDiff: () => void;
  isLastPage: boolean;
};
export const VersionsHistoryButtons = ({
  hasMore,
  canCompare,
  getVersions,
  getDiff,
  isLastPage,
}: VersionsButtonsType) => {
  return (
    <Stack>
      {hasMore && (
        <Button
          type="button"
          onClick={() => {
            getVersions(true);
            DashboardInteractions.showMoreVersionsClicked();
          }}
          variant="secondary"
          disabled={isLastPage}
        >
          <Trans i18nKey="dashboard-scene.versions-history-buttons.show-more-versions">Show more versions</Trans>
        </Button>
      )}
      <Tooltip
        content={t(
          'dashboard-scene.versions-history-buttons.content-select-two-versions-to-start-comparing',
          'Select two versions to start comparing'
        )}
        placement="bottom"
      >
        <Button type="button" disabled={!canCompare} onClick={getDiff} icon="code-branch">
          <Trans i18nKey="dashboard-scene.versions-history-buttons.compare-versions">Compare versions</Trans>
        </Button>
      </Tooltip>
    </Stack>
  );
};
