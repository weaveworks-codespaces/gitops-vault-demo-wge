.PHONY: all
all: kind flux

.PHONY: kind
kind:
	kind create cluster --name=flux-vault-demo --config=./kind/config.yaml

.PHONY: flux
flux:
	@test $${GITHUB_TOKEN?Please set environment variable GITHUB_TOKEN}
	@test $${GITHUB_USER?Please set environment variable GITHUB_USER}
	flux bootstrap github \
		--owner="$(GITHUB_USER)" \
		--repository=flux-vault-demo \
		--branch=main \
		--path=./clusters/kind \
		--personal

.PHONY: cleanup
cleanup:
	kind delete cluster --name=flux-vault-demo