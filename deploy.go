package main

func main() {
	kubeConfigPath := "C:/Users/houk/Desktop/msws/akspoc/config"

	// Deploy infrastucture
	KubectlApply(kubeConfigPath, CertManagerYaml)
	KubectlApply(kubeConfigPath, PodPresetYaml)

	// Deploy CRD
	// create a nameapce
	client := GetClient(kubeConfigPath)
	nsParams := &NamespaceParams{
		Name: "serviceconnector",
	}
	CreateNamespace(client, nsParams)

	// create a secret provider
	dclient := GetDynamiClient("C:/Users/houk/Desktop/msws/akspoc/config")
	spParams := &SecretProviderParams{
		Name:                "serviceconnector-secretprovider",
		Namespace:           "serviceconnector",
		K8sSecretName:       "serviceconnector-secret",
		KeyvaultName:        "houk-kv",                              // external param
		KeyvaultTenantId:    "72f988bf-86f1-41af-91ab-2d7cd011db47", // external param
		KeyvaultSecretNames: []string{"houk-secret"},
	}
	CreateSecretProvider(dclient, spParams)

	// create a pod to mouny secrets
	pparams := &PodParams{
		Name:               "serviceconnector-pod",
		Namespace:          "default",
		ContainerName:      "serviceconnector-container",
		SecretProviderName: "serviceconnector-secretprovider",
	}
	CreatePod(client, pparams)

	// create a config map
	params := &ConfigMapParams{
		Name:      "serviceconnector-configmap",
		Namespace: "default",
		Data: map[string]string{
			"houk-config": "houk-value", // external param
		},
	}
	CreateConfigMap(client, params)

	// create a pod preset
	presetParams := &PodPresetParams{
		Name:          "serviceconnector-podpreset",
		Namespace:     "serviceconnector",
		ConfigMapRefs: []string{"houk-config"}, // external param
		SecretRefs:    []string{"houk-secret"}, // external param
	}
	CreatePodPreset(dclient, presetParams)
}
