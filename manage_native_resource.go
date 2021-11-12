package main

import (
	"context"

	appsv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type ConfigMapParams struct {
	Name      string
	Namespace string
	Data      map[string]string
}

type NamespaceParams struct {
	Name string
}

type PodParams struct {
	Name               string
	Namespace          string
	ContainerName      string
	SecretProviderName string
}

func GetClient(kubeConfigPath string) *kubernetes.Clientset {
	// get kube config
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		panic(err)
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return clientset
}

func CreateConfigMap(client *kubernetes.Clientset, params *ConfigMapParams) {
	configMapClient := client.CoreV1().ConfigMaps(params.Namespace)
	configMap := &appsv1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: params.Name,
		},
		Data: params.Data,
	}

	// create or update
	_, err := configMapClient.Get(context.TODO(), params.Name, metav1.GetOptions{})
	if err == nil {
		_, err = configMapClient.Update(context.TODO(), configMap, metav1.UpdateOptions{})
	} else {
		_, err = configMapClient.Create(context.TODO(), configMap, metav1.CreateOptions{})
	}

	if err != nil {
		panic(err)
	}
}

func CreateNamespace(client *kubernetes.Clientset, params *NamespaceParams) {
	namespaceClient := client.CoreV1().Namespaces()
	namespace := &appsv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: params.Name,
		},
	}

	// create or update
	_, err := namespaceClient.Get(context.TODO(), params.Name, metav1.GetOptions{})
	if err == nil {
		_, err = namespaceClient.Update(context.TODO(), namespace, metav1.UpdateOptions{})
	} else {
		_, err = namespaceClient.Create(context.TODO(), namespace, metav1.CreateOptions{})
	}

	if err != nil {
		panic(err)
	}
}

func CreatePod(client *kubernetes.Clientset, params *PodParams) {
	podClient := client.CoreV1().Pods(params.Namespace)
	readOnly := true
	pod := &appsv1.Pod{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Pod",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      params.Name,
			Namespace: params.Namespace,
		},
		Spec: appsv1.PodSpec{
			Containers: []appsv1.Container{
				{
					Name:  "ServiceConnectorContainer",
					Image: "k8s.gcr.io/e2e-test-images/busybox:1.29",
					VolumeMounts: []appsv1.VolumeMount{
						{
							Name:      "ServiceConnectorSecretStoreVolume",
							MountPath: "/mnt/secret-store",
							ReadOnly:  true,
						},
					},
					EnvFrom: []appsv1.EnvFromSource{
						{
							ConfigMapRef: &appsv1.ConfigMapEnvSource{},
							SecretRef:    &appsv1.SecretEnvSource{},
						},
					},
				},
			},
			Volumes: []appsv1.Volume{
				{
					Name: "ServiceConnectorSecretStoreVolume",
					VolumeSource: appsv1.VolumeSource{
						CSI: &appsv1.CSIVolumeSource{
							Driver:   "secrets-store.csi.k8s.io",
							ReadOnly: &readOnly,
							VolumeAttributes: map[string]string{
								"secretProviderClass": params.SecretProviderName,
							},
						},
					},
				},
			},
		},
	}

	// create or update
	_, err := podClient.Get(context.TODO(), params.Name, metav1.GetOptions{})
	if err == nil {
		_, err = podClient.Update(context.TODO(), pod, metav1.UpdateOptions{})
	} else {
		_, err = podClient.Create(context.TODO(), pod, metav1.CreateOptions{})
	}

	if err != nil {
		panic(err)
	}
}
