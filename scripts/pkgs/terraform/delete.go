package terraform

import (
	"fmt"
	"os/exec"
)

func Destroy(tfvarsfile string) {
	fmt.Println("Starting destruction of resources...")
	cmd := exec.Command("terraform", "destroy", "-var-file", tfvarsfile)
	callBack(cmd)
}
