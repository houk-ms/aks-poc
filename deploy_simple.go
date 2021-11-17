package main

import (
	"fmt"
	"strings"
)

func DeploySimple(useSecretProvider bool, params DeployParams) {
	client := GetClient(params.KubeConfigPath)
	dclient := GetDynamiClient(params.KubeConfigPath)

	secretName := fmt.Sprintf("serviceconnector-%s-secret", strings.ToLower(params.ConnectionName))
	podName := fmt.Sprintf("serviceconnector-%s-pod", strings.ToLower(params.ConnectionName))
	configMapName := fmt.Sprintf("serviceconnector-%s-configmap", strings.ToLower(params.ConnectionName))

	if useSecretProvider {
		secretProviderName := "serviceconnector-secretprovider"

		// create a secret provider
		spParams := &SecretProviderParams{
			Name:                   secretProviderName,
			Namespace:              "default",
			K8sSecretName:          secretName,
			KeyvaultSecretPairs:    params.KeyvaultSecretPairs,
			KeyvaultName:           params.KeyvaultName,
			KeyvaultTenantId:       params.KeyvaultTenantId,
			UserAssignedIdentityID: params.UserAssignedIdentityID,
		}
		CreateSecretProvider(dclient, spParams)

		// create a pod to mount secrets
		pparams := &PodParams{
			Name:               podName,
			Namespace:          "default",
			ContainerName:      "serviceconnector-container",
			SecretProviderName: secretProviderName,
			K8sSecretName:      secretName,
		}
		CreatePod(client, pparams)
	} else {
		sParams := &SecretParams{
			Name:        secretName,
			Namespace:   "default",
			SecretPairs: params.SecretPairs,
		}
		CreateSecret(client, sParams)
	}

	// create a config map
	cmParams := &ConfigMapParams{
		Name:           configMapName,
		Namespace:      "default",
		ConfigMapPairs: params.ConfigMapPairs,
	}
	CreateConfigMap(client, cmParams)
}
