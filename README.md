# FLUX-VAULT-DEMO

## Setup
Before starting make sure to have the following applications installed:
- [Flux](https://fluxcd.io/docs/get-started)
- [Vault](https://www.vaultproject.io/downloads)
- [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)

### Create Kind Cluster
Run the following command to create a local Kind cluster:
```bash
make kind
```

### Bootstrap Flux
Bootstap Flux onto the local Kind cluster:
```bash
make flux
```
