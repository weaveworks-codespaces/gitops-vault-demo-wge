variable "vault_address" {
  type        = string
  description = "Vault server address"
}

variable "vault_token" {
  type        = string
  description = "Vault auth token"
  sensitive   = true
}

variable "kubernetes_host" {
  type        = string
  description = "Server address for kubernetes host"
}
