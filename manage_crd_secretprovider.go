package main

import (
	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

type SecretProviderParams struct {
	Name                string
	Namespace           string
	K8sSecretName       string
	KeyvaultName        string
	KeyvaultTenantId    string
	KeyvaultSecretNames []string
}

type ParameterObjectArrayItem struct {
	ObjectName string `yaml:"objectName"`
	ObjectType string `yaml:"objectType"`
}

type SpecParametersObjects struct {
	Array []ParameterObjectArrayItem `yaml:"array"`
}

type SpecParameters struct {
	UsePodIdentity string                `yaml:"usePodIdentity"`
	KeyvaultName   string                `yaml:"keyvaultName"`
	Objects        SpecParametersObjects `yaml:"objects"`
	TenantId       string                `yaml:"tenantId"`
}

type SecretObjectData struct {
	ObjectName string `yaml:"objectName"`
	Key        string `yaml:"key"`
}

type SecretObject struct {
	SecretName string             `yaml:"secretName"`
	Type       string             `yaml:"type"`
	Data       []SecretObjectData `yaml:"data"`
}

type SecretProviderSpec struct {
	Provider      string         `yaml:"provider"`
	SecretObjects []SecretObject `yaml:"secretObjects"`
	Parameters    SpecParameters `yaml:"parameters"`
}

type SecretProviderClass struct {
	APIVersion string             `yaml:"apiVersion"`
	Kind       string             `yaml:"kind"`
	Metadata   map[string]string  `yaml:"metadata"`
	Spec       SecretProviderSpec `yaml:"spec"`
}

func CreateSecretProvider(client dynamic.Interface, params *SecretProviderParams) {
	secretProvider := SecretProviderClass{
		APIVersion: "secrets-store.csi.x-k8s.io/v1alpha1",
		Kind:       "PodPreset",
		Metadata: map[string]string{
			"name":      params.Name,
			"namespace": params.Namespace,
		},
		Spec: SecretProviderSpec{
			Provider: "azure",
			SecretObjects: []SecretObject{
				{
					SecretName: params.K8sSecretName,
					Type:       "Opaque",
					Data:       []SecretObjectData{},
				},
			},
			Parameters: SpecParameters{
				UsePodIdentity: "false",
				KeyvaultName:   params.KeyvaultName,
				TenantId:       params.KeyvaultTenantId,
				Objects: SpecParametersObjects{
					Array: []ParameterObjectArrayItem{},
				},
			},
		},
	}

	for _, name := range params.KeyvaultSecretNames {
		secretProvider.Spec.SecretObjects[0].Data = append(
			secretProvider.Spec.SecretObjects[0].Data, SecretObjectData{
				ObjectName: name,
				Key:        name,
			})
		secretProvider.Spec.Parameters.Objects.Array = append(
			secretProvider.Spec.Parameters.Objects.Array, ParameterObjectArrayItem{
				ObjectName: name,
				ObjectType: "secret",
			})
	}

	yamlData, err := yaml.Marshal(secretProvider)
	if err != nil {
		panic(err)
	}

	gvrSecretProvider := schema.GroupVersionResource{
		Group:    "secrets-store.csi.x-k8s.io",
		Version:  "v1alpha1",
		Resource: "secretproviderclasses",
	}

	gvkSecretProvider := schema.GroupVersionKind{
		Group:   "secrets-store.csi.x-k8s.io",
		Version: "v1alpha1",
		Kind:    "SecretProviderClass",
	}
	CreateCRDResource(client, gvrSecretProvider, gvkSecretProvider, params.Namespace, string(yamlData))
}
