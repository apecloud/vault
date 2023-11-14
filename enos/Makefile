.PHONY: default
default: check-fmt shellcheck

.PHONY: check-fmt
check-fmt: check-fmt-enos check-fmt-modules

.PHONY: fmt
fmt: fmt-enos fmt-modules

.PHONY: check-fmt-enos
check-fmt-enos:
	enos fmt --check --diff .

.PHONY: fmt-enos
fmt-enos:
	enos fmt .

.PHONY: check-fmt-modules
check-fmt-modules:
	terraform fmt -check -diff -recursive ./modules

.PHONY: fmt-modules
fmt-modules:
	terraform fmt -diff -recursive ./modules

.PHONY: validate-enos
validate-enos:
	enos scenario validate agent --timeout 30m0s
	enos scenario validate autopilot --timeout 30m0s
	enos scenario validate proxy --timeout 30m0s
	enos scenario validate replication --timeout 45m0s
	enos scenario validate seal_ha --timeout 30m0s
	enos scenario validate smoke --timeout 30m0s
	enos scenario validate upgrade --timeout 30m0s

.PHONY: lint
lint: check-fmt shellcheck validate-enos

.PHONY: shellcheck
shellcheck:
	find ./modules/ -type f -name '*.sh' | xargs shellcheck
