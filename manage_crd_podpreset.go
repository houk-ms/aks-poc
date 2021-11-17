package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

type PodPresetParams struct {
	Name          string
	Namespace     string
	ConfigMapRefs []string
	SecretRefs    []string
}

type PodPresetEnvFromSource struct {
	Name string `yaml:"name"`
}

type PodPresetSelector struct {
	MatchedLabels map[string]string `yaml:"matchLabels"`
}

type PodPresetSpec struct {
	EnvFrom  []map[string]PodPresetEnvFromSource `yaml:"envFrom"`
	Selector PodPresetSelector                   `yaml:"selector"`
}

type PodPreset struct {
	APIVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   map[string]string `yaml:"metadata"`
	Spec       PodPresetSpec     `yaml:"spec"`
}

func CreatePodPreset(client dynamic.Interface, params *PodPresetParams) {
	podPreset := PodPreset{
		APIVersion: "redhatcop.redhat.io/v1alpha1",
		Kind:       "PodPreset",
		Metadata: map[string]string{
			"name":      params.Name,
			"namespace": params.Namespace,
		},
		Spec: PodPresetSpec{
			EnvFrom: []map[string]PodPresetEnvFromSource{},
			Selector: PodPresetSelector{
				MatchedLabels: map[string]string{
					"role": "service_connector",
				},
			},
		},
	}

	for _, name := range params.ConfigMapRefs {
		podPreset.Spec.EnvFrom = append(podPreset.Spec.EnvFrom, map[string]PodPresetEnvFromSource{
			"configMapRef": {
				Name: name,
			},
		})
	}

	for _, name := range params.SecretRefs {
		podPreset.Spec.EnvFrom = append(podPreset.Spec.EnvFrom, map[string]PodPresetEnvFromSource{
			"secretRef": {
				Name: name,
			},
		})
	}

	yamlData, err := yaml.Marshal(podPreset)
	if err != nil {
		panic(err)
	}

	gvrPodPreset := schema.GroupVersionResource{
		Group:    "redhatcop.redhat.io",
		Version:  "v1alpha1",
		Resource: "podpresets",
	}

	gvkPodPreset := schema.GroupVersionKind{
		Group:   "redhatcop.redhat.io",
		Version: "v1alpha1",
		Kind:    "PodPreset",
	}
	CreateCRDResource(client, gvrPodPreset, gvkPodPreset, params.Namespace, string(yamlData))
	fmt.Printf("podpreset/%s created", params.Name)
}
