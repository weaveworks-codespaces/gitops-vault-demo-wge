variable "vault_address" {
  type        = string
  description = "Vault server address"
  default     = "http://vault.vault.svc.cluster.local:8200"
}

variable "vault_token" {
  type        = string
  description = "Vault auth token"
  default     = "root"
  sensitive   = true
}

variable "kubernetes_host" {
  type        = string
  description = "Server address for kubernetes host"
  default     = "https://kubernetes.default.svc.cluster.local"
}
