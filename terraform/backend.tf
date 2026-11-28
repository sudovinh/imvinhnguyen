# State lives in the shared Spaces bucket (see terraform-root/digitalocean/README.md),
# encrypted client-side (OpenTofu only — run `tofu`, not `terraform`). Spaces creds +
# passphrase come from .env.terraform via `op run` (Makefile).
variable "state_passphrase" {
  description = "OpenTofu state/plan encryption passphrase (TF_VAR_state_passphrase via op run)"
  type        = string
  sensitive   = true
}

terraform {
  required_version = ">= 1.8"

  backend "s3" {
    endpoints = {
      s3 = "https://sfo3.digitaloceanspaces.com"
    }
    bucket       = "sv-terraform-root"
    key          = "apps/imvinhnguyen/terraform.tfstate"
    region       = "us-east-1"
    use_lockfile = true

    skip_credentials_validation = true
    skip_requesting_account_id  = true
    skip_metadata_api_check     = true
    skip_region_validation      = true
    skip_s3_checksum            = true
  }

  encryption {
    key_provider "pbkdf2" "state" {
      passphrase = var.state_passphrase
    }
    method "aes_gcm" "state" {
      keys = key_provider.pbkdf2.state
    }
    state {
      method   = method.aes_gcm.state
      enforced = true
    }
    plan {
      method   = method.aes_gcm.state
      enforced = true
    }
  }
}
