# Interface

## Prerequisite
- Enable addon `azure-keyvault-secrets-provider` on your AKS
- Enable user MI for your AKS
- Assign AKS user MI permission to get secrets in your KeyVault

## Usgae

### Case1: Secrets are directly provided
```
./main --kube-config C:/Users/houk/Desktop/msws/aks-poc/assets/config --connection houk-conn --config-map ENV1=houk-val1 ENV2=houk-val2 --secret ENV3=some-secrets1 ENV4=some-secrets2
```

### Case2: Secrets are in Keyvault
```
./main --kube-config C:/Users/houk/Desktop/msws/aks-poc/assets/config --connection houk-conn --config-map ENV1=houk-val1 ENV2=houk-val2 --kv-name houk-kv --kv-tenantid 72f988bf-86f1-41af-91ab-2d7cd011db47 --umi-clientid 4b0601b3-2087-4421-80b2-d393ba2fc257 --kv-secret ENV3=houk-secret ENV4=houk-secret2
```

### Case3: Secrets are in Keyvault, TenantID and UserMI not provided
This only works as a temparary solution.
Please make sure a `SecretProviderClass` named `akv-secretprovider` exists in your cluster, providing `tenantId` and `userMI` info before running the command. Use this yaml file as an [example](./test_yamls/secret_provider_existing.yaml).

```
./main --kube-config C:/Users/houk/Desktop/msws/aks-poc/assets/config --connection houk-conn --config-map ENV1=houk-val1 ENV2=houk-val2 --kv-name houk-kv --kv-secret ENV3=houk-secret ENV4=houk-secret2
```

## Verification
The following resources will be generated in `default` namespace of your K8S cluster.
- A pod named `serviceconnector-{conn_name}-pod`
- A config map named `serviceconnector-{conn_name}-configmap`
- A secret named `serviceconnector-{conn_name}-secret`
