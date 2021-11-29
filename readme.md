# Interface

## Prerequisite
- Enable addon `azure-keyvault-secrets-provider` on your AKS
- Enable user MI for your AKS
- Assign AKS user MI permission to get secrets in your KeyVault

## AKS + KeyVault
```
./main --kube-config C:/Users/houk/.kube/config --connection houk-conn --kv-name houk-kv --kv-tenantid 72f988bf-86f1-41af-91ab-2d7cd011db47 --umi-clientid 4b0601b3-2087-4421-80b2-d393ba2fc257
```

To verify, the following resource will appear in your AKS cluster.
-  A `SecretProviderClass` named `serviceconnector-{kv_name}-secretprovider`

## AKS + OtherTargets
```
./main --kube-config C:/Users/houk/.kube/config --connection houk-conn --config-map ENV1=houk-val1 ENV2=houk-val2 --secret ENV3=some-secrets1 ENV4=some-secrets2
```
To verify, the following resource will appear in your AKS cluster.
- A config map named `serviceconnector-{conn_name}-configmap`
- A secret named `serviceconnector-{conn_name}-secret`

## AKS + OtherTargets + KeyVault
### Case1: AKS + OtherTargets + KeyVault in one connection
```
./main --kube-config C:/Users/houk/.kube/config --connection houk-conn --config-map ENV1=houk-val1 ENV2=houk-val2 --kv-name houk-kv --kv-tenantid 72f988bf-86f1-41af-91ab-2d7cd011db47 --umi-clientid 4b0601b3-2087-4421-80b2-d393ba2fc257 --kv-secret ENV3=houk-secret ENV4=houk-secret2
```
To verify, the following resource will appear in your AKS cluster.
- A `SecretProviderClass` named `serviceconnector-{conn_name}-secretprovider`
- A pod named `serviceconnector-{conn_name}-pod`
- A config map named `serviceconnector-{conn_name}-configmap`
- A secret named `serviceconnector-{conn_name}-secret`


### Case2: AKS + OtherTargets + KeyVault in two connections

**1. Create AKS + KeyVault Connection**
```
./main --kube-config C:/Users/houk/.kube/config --connection houk-conn --kv-name houk-kv --kv-tenantid 72f988bf-86f1-41af-91ab-2d7cd011db47 --umi-clientid 4b0601b3-2087-4421-80b2-d393ba2fc257
```

To verify, the following resource will appear in your AKS cluster.
-  A `SecretProviderClass` named `serviceconnector-{kv_name}-secretprovider`


**2. Create AKS + OtherTarget Connection**
```
./main --kube-config C:/Users/houk/.kube/config --connection houk-conn --config-map ENV1=houk-val1 ENV2=houk-val2 --kv-name houk-kv --kv-secret ENV3=houk-secret ENV4=houk-secret2
```
To verify, the following resource will appear in your AKS cluster.
- A `SecretProviderClass` named `serviceconnector-{conn_name}-secretprovider`
- A pod named `serviceconnector-{conn_name}-pod`
- A config map named `serviceconnector-{conn_name}-configmap`
- A secret named `serviceconnector-{conn_name}-secret`
