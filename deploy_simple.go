package main

import (
	"fmt"
	"regexp"
	"strings"
)

func DeploySimple(connectKeyVault bool, useKeyVault bool, withTenant bool, params DeployParams) {
	client := GetClient(params.KubeConfigPath)
	dclient := GetDynamiClient(params.KubeConfigPath)

	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	processedConnName := strings.ToLower(reg.ReplaceAllString(params.ConnectionName, ""))
	secretName := fmt.Sprintf("serviceconnector-%s-secret", processedConnName)
	podName := fmt.Sprintf("serviceconnector-%s-pod", processedConnName)
	configMapName := fmt.Sprintf("serviceconnector-%s-configmap", processedConnName)

	if connectKeyVault {
		processedKvName := strings.ToLower(reg.ReplaceAllString(params.KeyvaultName, ""))
		secretProviderName := fmt.Sprintf("serviceconnector-%s-secretprovider", processedKvName)

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

	} else if useKeyVault {
		secretProviderName := fmt.Sprintf("serviceconnector-%s-secretprovider", processedConnName)

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

	if len(params.ConfigMapPairs) > 0 {
		// create a config map
		cmParams := &ConfigMapParams{
			Name:           configMapName,
			Namespace:      "default",
			ConfigMapPairs: params.ConfigMapPairs,
		}
		CreateConfigMap(client, cmParams)
	}
}
