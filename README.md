# DEPRECATED: FLUX-VAULT-DEMO
This project is designed to give an overview on how to integrate Vault secrets into your Flux deployed applications.

## Setup
Before starting make sure to have the following applications installed:
- [Flux](https://fluxcd.io/docs/get-started)
- [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
- [Vault](https://www.vaultproject.io/downloads) - Optional (the demo does not use the cli, but you'll need it if you want to play around on your own)

Set the following environment variables:
- `GITHUB_TOKEN` - personal access token used to authenticate with GitHub
- `GITHUB_USER` - the GitHub username that will be set as the `owner` during the Flux bootstrap process

The easiest setup is to fork this project into your own repo, clone the fork, then run
```bash
make
```

This will create a local Kind cluster called `flux-vault-demo`, bootstrap Flux to your fork, and configure Vault with all the permissions necessary to run this demo.

If you have a different k8s cluster you would like to use instead, run
```bash
make flux
```
to skip the Kind cluster creation and only bootstrap Flux.

If you are not using GitHub for your repository, you can run
```bash
make kind
```
to create a local Kind cluster, but you will need to follow the [Flux bootstrap](https://fluxcd.io/docs/installation/#bootstrap) instructions on how to properly setup Flux on your repository.

## Cleanup
When you are done with the demo you can run
```bash
make cleanup
```
to remove the demo Kind cluster from your machine.
