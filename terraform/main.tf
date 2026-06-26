terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

provider "digitalocean" {
  token = var.do_token
}

resource "digitalocean_app" "imvinhnguyen_web" {
  spec {
    name   = "imvinhnguyen-com"
    region = "sfo"

    alert {
      rule = "DEPLOYMENT_FAILED"
    }

    alert {
      rule = "DOMAIN_FAILED"
    }

    domain {
      name = "imvinhnguyen.com"
      type = "PRIMARY"
      zone = "imvinhnguyen.com"
    }

    domain {
      name = "www.imvinhnguyen.com"
      type = "ALIAS"
      zone = "imvinhnguyen.com"
    }

    # Route all traffic to our service. Declared explicitly because the
    # imported app's ingress still referenced the old component name.
    ingress {
      rule {
        component {
          name = "sudovinh-imvinhnguyen"
        }
        match {
          path {
            prefix = "/"
          }
        }
      }
    }

    service {
      name               = "sudovinh-imvinhnguyen"
      instance_count     = 1
      instance_size_slug = "basic-xxs"

      github {
        repo           = "sudovinh/imvinhnguyen"
        branch         = "main"
        deploy_on_push = false
      }

      dockerfile_path = "Dockerfile"

      http_port = 8080

      health_check {
        http_path             = "/"
        initial_delay_seconds = 10
        period_seconds        = 30
      }
    }
  }
}

resource "digitalocean_domain" "imvinhnguyen" {
  name = "imvinhnguyen.com"
}

# Existing app (was "imvinhnguyen-page") imported into state, then renamed
# to "imvinhnguyen-com" via the spec above on the next apply.
import {
  to = digitalocean_app.imvinhnguyen_web
  id = "5f98d41f-4ca1-4033-8acd-22254375f31a"
}

import {
  to = digitalocean_domain.imvinhnguyen
  id = "imvinhnguyen.com"
}

output "app_url" {
  value       = digitalocean_app.imvinhnguyen_web.live_url
  description = "The live URL of the deployed app"
}

output "default_ingress" {
  value       = digitalocean_app.imvinhnguyen_web.default_ingress
  description = "The default ondigitalocean.app ingress for the app"
}
