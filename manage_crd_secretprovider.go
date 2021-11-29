package main

import (
	"context"
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

type SecretProviderParams struct {
	Name                   string
	Namespace              string
	K8sSecretName          string
	KeyvaultName           string
	KeyvaultTenantId       string
	UserAssignedIdentityID string
	KeyvaultSecretPairs    map[string]string
}

type SpecParameters struct {
	UsePodIdentity         string `yaml:"usePodIdentity"`
	UseVMManagedIdentity   string `yaml:"useVMManagedIdentity"`
	UserAssignedIdentityID string `yaml:"userAssignedIdentityID"`
	KeyvaultName           string `yaml:"keyvaultName"`
	Objects                string `yaml:"objects"`
	TenantId               string `yaml:"tenantId"`
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
		Kind:       "SecretProviderClass",
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
				UsePodIdentity:         "false",
				UseVMManagedIdentity:   "true",
				UserAssignedIdentityID: params.UserAssignedIdentityID,
				KeyvaultName:           params.KeyvaultName,
				TenantId:               params.KeyvaultTenantId,
				Objects:                "",
			},
		},
	}

	objectsValue := "array:\n"
	for key, secret := range params.KeyvaultSecretPairs {
		secretProvider.Spec.SecretObjects[0].Data = append(
			secretProvider.Spec.SecretObjects[0].Data, SecretObjectData{
				ObjectName: secret,
				Key:        key,
			})

		objectsValue += fmt.Sprintf(
			`  - |
    objectName: %s
    objectType: secret
`, secret)
	}
	secretProvider.Spec.Parameters.Objects = objectsValue

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
	fmt.Printf("secretprovider/%s created\n", params.Name)
}

func UpdateSecretProvider(client dynamic.Interface, params *SecretProviderParams) {
	gvrSecretProvider := schema.GroupVersionResource{
		Group:    "secrets-store.csi.x-k8s.io",
		Version:  "v1alpha1",
		Resource: "secretproviderclasses",
	}

	utd, _ := client.Resource(gvrSecretProvider).Namespace(params.Namespace).Get(context.TODO(), "akv-secretprovider", metav1.GetOptions{})
	data, _ := utd.MarshalJSON()

	var secretProvider SecretProviderClass
	json.Unmarshal(data, &secretProvider)

	params.KeyvaultTenantId = secretProvider.Spec.Parameters.TenantId
	params.UserAssignedIdentityID = secretProvider.Spec.Parameters.UserAssignedIdentityID

	CreateSecretProvider(client, params)
}
