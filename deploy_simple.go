package main

import (
	"fmt"
	"strings"
)

func DeploySimple(useKeyVault bool, withTenant bool, params DeployParams) {
	client := GetClient(params.KubeConfigPath)
	dclient := GetDynamiClient(params.KubeConfigPath)

	secretName := fmt.Sprintf("serviceconnector-%s-secret", strings.ToLower(params.ConnectionName))
	podName := fmt.Sprintf("serviceconnector-%s-pod", strings.ToLower(params.ConnectionName))
	configMapName := fmt.Sprintf("serviceconnector-%s-configmap", strings.ToLower(params.ConnectionName))

	if useKeyVault {
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

		if withTenant {
			CreateSecretProvider(dclient, spParams)
		} else {
			// get KeyvaultTenantId and UserAssignedIdentityID from existing secretProvider
			UpdateSecretProvider(dclient, spParams)
		}

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
