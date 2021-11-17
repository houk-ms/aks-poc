package main

import (
	"fmt"
	"os"
	"os/exec"
)

func KubectlApply(kubeConfigPath string, yamlPath string) error {
	// get current directory
	// currentPath, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }

	// prepare kubectl arguments
	args := []string{"apply", "--kubeconfig", kubeConfigPath, "-f", yamlPath}

	// execute kubectl apply
	cmd := exec.Command("C:/Users/houk/Desktop/msws/aks-poc/kubectl.exe", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("Execute Command finished.")
	return err
}
