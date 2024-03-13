package main

import (
	"github.com/AmitGujar/terraform-boilerplate/scripts/pkgs/terraform"
)

func main() {
	terraform.Check()
	terraform.Init()
	terraform.Plan("secret.tfvars")
	terraform.Apply()
}
