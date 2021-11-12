package main

import (
	"fmt"
	"os"
	"os/exec"
)

func KubectlApply(kubeConfigPath string, resourceConfig string) error {
	// get current directory
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Get current directory failed:" + err.Error())
	}

	// prepare kubectl arguments
	args := "apply --kubeconfig " + kubeConfigPath + " -f - " + resourceConfig

	// execute kubectl apply
	cmd := exec.Command(currentPath+"/assets/kubectl", args)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Execute command failed:" + err.Error())
	}

	fmt.Println("Execute Command finished.")
	return err
}
