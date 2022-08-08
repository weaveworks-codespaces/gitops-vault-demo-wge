provider "vault" {
  address = var.vault_address
  token   = var.vault_token
}

resource "vault_auth_backend" "kubernetes" {
  type = "kubernetes"
}

resource "vault_kubernetes_auth_backend_config" "example" {
  backend                = vault_auth_backend.kubernetes.path
  kubernetes_host        = var.kubernetes_host
  disable_iss_validation = true
}

resource "vault_mount" "demo" {
  path        = "demo"
  type        = "kv-v2"
  description = "demo secret storage"
}

resource "vault_kv_secret_v2" "demo_creds" {
  mount = vault_mount.demo.path
  name  = "creds"
  data_json = jsonencode(
    {
      username = "demo",
      password = "myPassword",
    }
  )

  lifecycle {
    ignore_changes = [data_json]
  }
}

data "vault_policy_document" "read_demo_secrets" {
  rule {
    path         = "demo/*"
    capabilities = ["read", "list"]
    description  = "read secrets"
  }
}

resource "vault_policy" "read_demo" {
  name   = "read-demo"
  policy = data.vault_policy_document.read_demo_secrets.hcl
}

resource "vault_kubernetes_auth_backend_role" "flux_vault_demo" {
  backend                          = vault_auth_backend.kubernetes.path
  role_name                        = "flux-vault-demo"
  bound_service_account_names      = ["flux-vault-demo-injector", "flux-vault-demo-csi"]
  bound_service_account_namespaces = ["default"]
  token_ttl                        = 3600
  token_policies                   = ["read-demo"]
}
