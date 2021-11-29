package main

import (
	"os"
	"strings"
)

type DeployParams struct {
	KubeConfigPath         string
	ConnectionName         string
	KeyvaultName           string
	KeyvaultTenantId       string
	UserAssignedIdentityID string
	KeyvaultSecretPairs    map[string]string
	ConfigMapPairs         map[string]string
	SecretPairs            map[string]string
}

func main() {
	dParams := DeployParams{
		KubeConfigPath:         "",
		KeyvaultName:           "",
		KeyvaultTenantId:       "",
		UserAssignedIdentityID: "",
		KeyvaultSecretPairs:    map[string]string{},
		ConfigMapPairs:         map[string]string{},
		SecretPairs:            map[string]string{},
	}

	argName := "unknow"
	for _, arg := range os.Args {
		if arg[:2] == "--" {
			argName = arg[2:]
			continue
		}
		switch argName {
		case "kube-config":
			dParams.KubeConfigPath = arg
		case "connection":
			dParams.ConnectionName = arg
		case "config-map":
			items := strings.Split(arg, "=")
			dParams.ConfigMapPairs[items[0]] = items[1]
		case "secret":
			items := strings.Split(arg, "=")
			dParams.SecretPairs[items[0]] = items[1]
		case "kv-name":
			dParams.KeyvaultName = arg
		case "kv-tenantid":
			dParams.KeyvaultTenantId = arg
		case "umi-clientid":
			dParams.UserAssignedIdentityID = arg
		case "kv-secret":
			items := strings.Split(arg, "=")
			dParams.KeyvaultSecretPairs[items[0]] = items[1]
		case "unkown":
			panic("Args parsing error!")
		}
	}

	useKeyVault := false
	withTenant := true
	if len(dParams.KeyvaultSecretPairs) > 0 {
		useKeyVault = true
	}
	if dParams.KeyvaultTenantId == "" {
		withTenant = false
	}

	DeploySimple(useKeyVault, withTenant, dParams)
}
