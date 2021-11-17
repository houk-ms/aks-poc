package main

import "os"

func DeployPodpreset(useSecretProvider bool) {
	kubeConfigPath := "C:/Users/houk/Desktop/msws/aks-poc/assets/config" // external param
	currentPath, _ := os.Getwd()

	// Deploy infrastucture
	KubectlApply(kubeConfigPath, currentPath+"/assets/infra_cert_manager.yaml")
	KubectlApply(kubeConfigPath, currentPath+"/assets/infra_pod_preset.yaml")

	client := GetClient(kubeConfigPath)
	dclient := GetDynamiClient(kubeConfigPath)

	// create a nameapce
	// nsParams := &NamespaceParams{
	// 	Name: "serviceconnector",
	// }
	// CreateNamespace(client, nsParams)

	if useSecretProvider {
		// create a secret provider
		spParams := &SecretProviderParams{
			Name:                "serviceconnector-secretprovider",
			Namespace:           "serviceconnector",
			K8sSecretName:       "serviceconnector-secret",
			KeyvaultName:        "houk-kv",                              // external param
			KeyvaultTenantId:    "72f988bf-86f1-41af-91ab-2d7cd011db47", // external param
			KeyvaultSecretPairs: map[string]string{},
		}
		CreateSecretProvider(dclient, spParams)

		// create a pod to mount secrets
		pparams := &PodParams{
			Name:               "serviceconnector-pod",
			Namespace:          "default",
			ContainerName:      "serviceconnector-container",
			SecretProviderName: "serviceconnector-secretprovider",
			K8sSecretName:      "serviceconnector-secret",
		}
		CreatePod(client, pparams)
	} else {
		sParams := &SecretParams{
			Name:      "serviceconnector-secret",
			Namespace: "default",
			SecretPairs: map[string]string{
				"houk-key": "houk-value",
			},
		}
		CreateSecret(client, sParams)
	}

	// create a config map
	params := &ConfigMapParams{
		Name:      "serviceconnector-configmap",
		Namespace: "default",
		ConfigMapPairs: map[string]string{
			"houk-config": "houk-value", // external param
		},
	}
	CreateConfigMap(client, params)

	// create a pod preset
	presetParams := &PodPresetParams{
		Name:          "serviceconnector-podpreset",
		Namespace:     "default",
		ConfigMapRefs: []string{"houk-config"}, // external param
		SecretRefs:    []string{"houk-secret"}, // external param
	}
	CreatePodPreset(dclient, presetParams)
}
