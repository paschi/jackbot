terraform {
  required_providers {
    scaleway = {
      source  = "opentofu/scaleway"
      version = "2.73.0"
    }
  }
}

provider "scaleway" {
  project_id = var.project_id
  region     = var.region
  zone       = var.zone
  access_key = var.access_key
  secret_key = var.secret_key
}
