DEFAULT_BUMP_STAGE=beta # final|alpha|beta|candidate
DEFAULT_BUMP_SCOPE=minor # major|minor|patch
DEFAULT_BUMP_DRY_RUN=true # true|false

STAGE=$(DEFAULT_BUMP_STAGE)
SCOPE=$(DEFAULT_BUMP_SCOPE)
DRY_RUN=$(DEFAULT_BUMP_DRY_RUN)
.PHONY: bump
bump: check-git-status ## Bump version
	(bash ./hack/bump.sh --stage ${STAGE} --scope ${SCOPE} --dry-run ${DRY_RUN})


.PHONY: version
version:
	cat VERSION
