packer {
  required_plugins {
    inspec = {
      version = ">=v0.1.0"
      source  = "github.com/hashicorp/inspec"
    }
  }
}

source "digitalocean" "example"{
    api_token = "<digital ocean api token>"
    image     = "ubuntu-14-04-x64"
    region    = "sfo1"
}

build {
    sources = [
        "source.digitalocean.example"
    ]
    provisioner "inspec" {
        profile = "https://github.com/dev-sec/linux-baseline"
    }
}