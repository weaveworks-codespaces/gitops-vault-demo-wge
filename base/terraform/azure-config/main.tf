terraform {
  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
      version = "~>2.0"
    }
  }

provider "azurerm" {
  features {}

  subscription_id   = "${env.ARM_SUBSCRIPTION_ID}"
  tenant_id         = "${env.ARM_TENANT_ID}"
  client_id         = "${env.ARM_CLIENT_ID}"
  client_secret     = "${env.ARM_CLIENT_SECRET}"
  use_msi = true
}
  
  resource "random_pet" "rg_name" {
  prefix = var.resource_group_name_prefix
}

resource "azurerm_resource_group" "rg" {
  location = var.resource_group_location
  name     = random_pet.rg_name.id
}