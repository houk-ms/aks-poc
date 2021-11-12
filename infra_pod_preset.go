package main

var PodPresetYaml = `
apiVersion: v1
kind: Namespace
metadata:
  labels:
    control-plane: controller-manager
  name: podpreset-webhook
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.4.1
  creationTimestamp: null
  name: podpresets.redhatcop.redhat.io
spec:
  group: redhatcop.redhat.io
  names:
    kind: PodPreset
    listKind: PodPresetList
    plural: podpresets
    singular: podpreset
  scope: Namespaced
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: PodPreset is the Schema for the podpresets API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: PodPresetSpec defines the desired state of PodPreset
            properties:
              env:
                items:
                  description: EnvVar represents an environment variable present in
                    a Container.
                  properties:
                    name:
                      description: Name of the environment variable. Must be a C_IDENTIFIER.
                      type: string
                    value:
                      description: 'Variable references $(VAR_NAME) are expanded using
                        the previous defined environment variables in the container
                        and any service environment variables. If a variable cannot
                        be resolved, the reference in the input string will be unchanged.
                        The $(VAR_NAME) syntax can be escaped with a double $$, ie:
                        $$(VAR_NAME). Escaped references will never be expanded, regardless
                        of whether the variable exists or not. Defaults to "".'
                      type: string
                    valueFrom:
                      description: Source for the environment variable's value. Cannot
                        be used if value is not empty.
                      properties:
                        configMapKeyRef:
                          description: Selects a key of a ConfigMap.
                          properties:
                            key:
                              description: The key to select.
                              type: string
                            name:
                              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                TODO: Add other useful fields. apiVersion, kind, uid?'
                              type: string
                            optional:
                              description: Specify whether the ConfigMap or its key
                                must be defined
                              type: boolean
                          required:
                          - key
                          type: object
                        fieldRef:
                          description: 'Selects a field of the pod: supports metadata.name,
                            metadata.namespace, metadata.labels[''<KEY>''], metadata.annotations[''<KEY>''],
                            spec.nodeName, spec.serviceAccountName, status.hostIP,
                            status.podIP, status.podIPs.'
                          properties:
                            apiVersion:
                              description: Version of the schema the FieldPath is
                                written in terms of, defaults to "v1".
                              type: string
                            fieldPath:
                              description: Path of the field to select in the specified
                                API version.
                              type: string
                          required:
                          - fieldPath
                          type: object
                        resourceFieldRef:
                          description: 'Selects a resource of the container: only
                            resources limits and requests (limits.cpu, limits.memory,
                            limits.ephemeral-storage, requests.cpu, requests.memory
                            and requests.ephemeral-storage) are currently supported.'
                          properties:
                            containerName:
                              description: 'Container name: required for volumes,
                                optional for env vars'
                              type: string
                            divisor:
                              anyOf:
                              - type: integer
                              - type: string
                              description: Specifies the output format of the exposed
                                resources, defaults to "1"
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            resource:
                              description: 'Required: resource to select'
                              type: string
                          required:
                          - resource
                          type: object
                        secretKeyRef:
                          description: Selects a key of a secret in the pod's namespace
                          properties:
                            key:
                              description: The key of the secret to select from.  Must
                                be a valid secret key.
                              type: string
                            name:
                              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                TODO: Add other useful fields. apiVersion, kind, uid?'
                              type: string
                            optional:
                              description: Specify whether the Secret or its key must
                                be defined
                              type: boolean
                          required:
                          - key
                          type: object
                      type: object
                  required:
                  - name
                  type: object
                type: array
              envFrom:
                items:
                  description: EnvFromSource represents the source of a set of ConfigMaps
                  properties:
                    configMapRef:
                      description: The ConfigMap to select from
                      properties:
                        name:
                          description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                            TODO: Add other useful fields. apiVersion, kind, uid?'
                          type: string
                        optional:
                          description: Specify whether the ConfigMap must be defined
                          type: boolean
                      type: object
                    prefix:
                      description: An optional identifier to prepend to each key in
                        the ConfigMap. Must be a C_IDENTIFIER.
                      type: string
                    secretRef:
                      description: The Secret to select from
                      properties:
                        name:
                          description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                            TODO: Add other useful fields. apiVersion, kind, uid?'
                          type: string
                        optional:
                          description: Specify whether the Secret must be defined
                          type: boolean
                      type: object
                  type: object
                type: array
              selector:
                description: A label selector is a label query over a set of resources.
                  The result of matchLabels and matchExpressions are ANDed. An empty
                  label selector matches all objects. A null label selector matches
                  no objects.
                properties:
                  matchExpressions:
                    description: matchExpressions is a list of label selector requirements.
                      The requirements are ANDed.
                    items:
                      description: A label selector requirement is a selector that
                        contains values, a key, and an operator that relates the key
                        and values.
                      properties:
                        key:
                          description: key is the label key that the selector applies
                            to.
                          type: string
                        operator:
                          description: operator represents a key's relationship to
                            a set of values. Valid operators are In, NotIn, Exists
                            and DoesNotExist.
                          type: string
                        values:
                          description: values is an array of string values. If the
                            operator is In or NotIn, the values array must be non-empty.
                            If the operator is Exists or DoesNotExist, the values
                            array must be empty. This array is replaced during a strategic
                            merge patch.
                          items:
                            type: string
                          type: array
                      required:
                      - key
                      - operator
                      type: object
                    type: array
                  matchLabels:
                    additionalProperties:
                      type: string
                    description: matchLabels is a map of {key,value} pairs. A single
                      {key,value} in the matchLabels map is equivalent to an element
                      of matchExpressions, whose key field is "key", the operator
                      is "In", and the values array contains only "value". The requirements
                      are ANDed.
                    type: object
                type: object
              volumeMounts:
                items:
                  description: VolumeMount describes a mounting of a Volume within
                    a container.
                  properties:
                    mountPath:
                      description: Path within the container at which the volume should
                        be mounted.  Must not contain ':'.
                      type: string
                    mountPropagation:
                      description: mountPropagation determines how mounts are propagated
                        from the host to container and the other way around. When
                        not set, MountPropagationNone is used. This field is beta
                        in 1.10.
                      type: string
                    name:
                      description: This must match the Name of a Volume.
                      type: string
                    readOnly:
                      description: Mounted read-only if true, read-write otherwise
                        (false or unspecified). Defaults to false.
                      type: boolean
                    subPath:
                      description: Path within the volume from which the container's
                        volume should be mounted. Defaults to "" (volume's root).
                      type: string
                    subPathExpr:
                      description: Expanded path within the volume from which the
                        container's volume should be mounted. Behaves similarly to
                        SubPath but environment variable references $(VAR_NAME) are
                        expanded using the container's environment. Defaults to ""
                        (volume's root). SubPathExpr and SubPath are mutually exclusive.
                      type: string
                  required:
                  - mountPath
                  - name
                  type: object
                type: array
              volumes:
                items:
                  description: Volume represents a named volume in a pod that may
                    be accessed by any container in the pod.
                  properties:
                    awsElasticBlockStore:
                      description: 'AWSElasticBlockStore represents an AWS Disk resource
                        that is attached to a kubelet''s host machine and then exposed
                        to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore'
                      properties:
                        fsType:
                          description: 'Filesystem type of the volume that you want
                            to mount. Tip: Ensure that the filesystem type is supported
                            by the host operating system. Examples: "ext4", "xfs",
                            "ntfs". Implicitly inferred to be "ext4" if unspecified.
                            More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
                            TODO: how do we prevent errors in the filesystem from
                            compromising the machine'
                          type: string
                        partition:
                          description: 'The partition in the volume that you want
                            to mount. If omitted, the default is to mount by volume
                            name. Examples: For volume /dev/sda1, you specify the
                            partition as "1". Similarly, the volume partition for
                            /dev/sda is "0" (or you can leave the property empty).'
                          format: int32
                          type: integer
                        readOnly:
                          description: 'Specify "true" to force and set the ReadOnly
                            property in VolumeMounts to "true". If omitted, the default
                            is "false". More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore'
                          type: boolean
                        volumeID:
                          description: 'Unique ID of the persistent disk resource
                            in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore'
                          type: string
                      required:
                      - volumeID
                      type: object
                    azureDisk:
                      description: AzureDisk represents an Azure Data Disk mount on
                        the host and bind mount to the pod.
                      properties:
                        cachingMode:
                          description: 'Host Caching mode: None, Read Only, Read Write.'
                          type: string
                        diskName:
                          description: The Name of the data disk in the blob storage
                          type: string
                        diskURI:
                          description: The URI the data disk in the blob storage
                          type: string
                        fsType:
                          description: Filesystem type to mount. Must be a filesystem
                            type supported by the host operating system. Ex. "ext4",
                            "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
                          type: string
                        kind:
                          description: 'Expected values Shared: multiple blob disks
                            per storage account  Dedicated: single blob disk per storage
                            account  Managed: azure managed data disk (only in managed
                            availability set). defaults to shared'
                          type: string
                        readOnly:
                          description: Defaults to false (read/write). ReadOnly here
                            will force the ReadOnly setting in VolumeMounts.
                          type: boolean
                      required:
                      - diskName
                      - diskURI
                      type: object
                    azureFile:
                      description: AzureFile represents an Azure File Service mount
                        on the host and bind mount to the pod.
                      properties:
                        readOnly:
                          description: Defaults to false (read/write). ReadOnly here
                            will force the ReadOnly setting in VolumeMounts.
                          type: boolean
                        secretName:
                          description: the name of secret that contains Azure Storage
                            Account Name and Key
                          type: string
                        shareName:
                          description: Share Name
                          type: string
                      required:
                      - secretName
                      - shareName
                      type: object
                    cephfs:
                      description: CephFS represents a Ceph FS mount on the host that
                        shares a pod's lifetime
                      properties:
                        monitors:
                          description: 'Required: Monitors is a collection of Ceph
                            monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it'
                          items:
                            type: string
                          type: array
                        path:
                          description: 'Optional: Used as the mounted root, rather
                            than the full Ceph tree, default is /'
                          type: string
                        readOnly:
                          description: 'Optional: Defaults to false (read/write).
                            ReadOnly here will force the ReadOnly setting in VolumeMounts.
                            More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it'
                          type: boolean
                        secretFile:
                          description: 'Optional: SecretFile is the path to key ring
                            for User, default is /etc/ceph/user.secret More info:
                            https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it'
                          type: string
                        secretRef:
                          description: 'Optional: SecretRef is reference to the authentication
                            secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it'
                          properties:
                            name:
                              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                TODO: Add other useful fields. apiVersion, kind, uid?'
                              type: string
                          type: object
                        user:
                          description: 'Optional: User is the rados user name, default
                            is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it'
                          type: string
                      required:
                      - monitors
                      type: object
                    cinder:
                      description: 'Cinder represents a cinder volume attached and
                        mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md'
                      properties:
                        fsType:
                          description: 'Filesystem type to mount. Must be a filesystem
                            type supported by the host operating system. Examples:
                            "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4"
                            if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md'
                          type: string
                        readOnly:
                          description: 'Optional: Defaults to false (read/write).
                            ReadOnly here will force the ReadOnly setting in VolumeMounts.
                            More info: https://examples.k8s.io/mysql-cinder-pd/README.md'
                          type: boolean
                        secretRef:
                          description: 'Optional: points to a secret object containing
                            parameters used to connect to OpenStack.'
                          properties:
                            name:
                              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                TODO: Add other useful fields. apiVersion, kind, uid?'
                              type: string
                          type: object
                        volumeID:
                          description: 'volume id used to identify the volume in cinder.
                            More info: https://examples.k8s.io/mysql-cinder-pd/README.md'
                          type: string
                      required:
                      - volumeID
                      type: object
                    configMap:
                      description: ConfigMap represents a configMap that should populate
                        this volume
                      properties:
                        defaultMode:
                          description: 'Optional: mode bits used to set permissions
                            on created files by default. Must be an octal value between
                            0000 and 0777 or a decimal value between 0 and 511. YAML
                            accepts both octal and decimal values, JSON requires decimal
                            values for mode bits. Defaults to 0644. Directories within
                            the path are not affected by this setting. This might
                            be in conflict with other options that affect the file
                            mode, like fsGroup, and the result can be other mode bits
                            set.'
                          format: int32
                          type: integer
                        items:
                          description: If unspecified, each key-value pair in the
                            Data field of the referenced ConfigMap will be projected
                            into the volume as a file whose name is the key and content
                            is the value. If specified, the listed keys will be projected
                            into the specified paths, and unlisted keys will not be
                            present. If a key is specified which is not present in
                            the ConfigMap, the volume setup will error unless it is
                            marked optional. Paths must be relative and may not contain
                            the '..' path or start with '..'.
                          items:
                            description: Maps a string key to a path within a volume.
                            properties:
                              key:
                                description: The key to project.
                                type: string
                              mode:
                                description: 'Optional: mode bits used to set permissions
                                  on this file. Must be an octal value between 0000
                                  and 0777 or a decimal value between 0 and 511. YAML
                                  accepts both octal and decimal values, JSON requires
                                  decimal values for mode bits. If not specified,
                                  the volume defaultMode will be used. This might
                                  be in conflict with other options that affect the
                                  file mode, like fsGroup, and the result can be other
                                  mode bits set.'
                                format: int32
                                type: integer
                              path:
                                description: The relative path of the file to map
                                  the key to. May not be an absolute path. May not
                                  contain the path element '..'. May not start with
                                  the string '..'.
                                type: string
                            required:
                            - key
                            - path
                            type: object
                          type: array
                        name:
                          description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                            TODO: Add other useful fields. apiVersion, kind, uid?'
                          type: string
                        optional:
                          description: Specify whether the ConfigMap or its keys must
                            be defined
                          type: boolean
                      type: object
                    csi:
                      description: CSI (Container Storage Interface) represents ephemeral
                        storage that is handled by certain external CSI drivers (Beta
                        feature).
                      properties:
                        driver:
                          description: Driver is the name of the CSI driver that handles
                            this volume. Consult with your admin for the correct name
                            as registered in the cluster.
                          type: string
                        fsType:
                          description: Filesystem type to mount. Ex. "ext4", "xfs",
                            "ntfs". If not provided, the empty value is passed to
                            the associated CSI driver which will determine the default
                            filesystem to apply.
                          type: string
                        nodePublishSecretRef:
                          description: NodePublishSecretRef is a reference to the
                            secret object containing sensitive information to pass
                            to the CSI driver to complete the CSI NodePublishVolume
                            and NodeUnpublishVolume calls. This field is optional,
                            and  may be empty if no secret is required. If the secret
                            object contains more than one secret, all secret references
                            are passed.
                          properties:
                            name:
                              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                TODO: Add other useful fields. apiVersion, kind, uid?'
                              type: string
                          type: object
                        readOnly:
                          description: Specifies a read-only configuration for the
                            volume. Defaults to false (read/write).
                          type: boolean
                        volumeAttributes:
                          additionalProperties:
                            type: string
                          description: VolumeAttributes stores driver-specific properties
                            that are passed to the CSI driver. Consult your driver's
                            documentation for supported values.
                          type: object
                      required:
                      - driver
                      type: object
                    downwardAPI:
                      description: DownwardAPI represents downward API about the pod
                        that should populate this volume
                      properties:
                        defaultMode:
                          description: 'Optional: mode bits to use on created files
                            by default. Must be a Optional: mode bits used to set
                            permissions on created files by default. Must be an octal
                            value between 0000 and 0777 or a decimal value between
                            0 and 511. YAML accepts both octal and decimal values,
                            JSON requires decimal values for mode bits. Defaults to
                            0644. Directories within the path are not affected by
                            this setting. This might be in conflict with other options
                            that affect the file mode, like fsGroup, and the result
                            can be other mode bits set.'
                          format: int32
                          type: integer
                        items:
                          description: Items is a list of downward API volume file
                          items:
                            description: DownwardAPIVolumeFile represents information
                              to create the file containing the pod field
                            properties:
                              fieldRef:
                                description: 'Required: Selects a field of the pod:
                                  only annotations, labels, name and namespace are
                                  supported.'
                                properties:
                                  apiVersion:
                                    description: Version of the schema the FieldPath
                                      is written in terms of, defaults to "v1".
                                    type: string
                                  fieldPath:
                                    description: Path of the field to select in the
                                      specified API version.
                                    type: string
                                required:
                                - fieldPath
                                type: object
                              mode:
                                description: 'Optional: mode bits used to set permissions
                                  on this file, must be an octal value between 0000
                                  and 0777 or a decimal value between 0 and 511. YAML
                                  accepts both octal and decimal values, JSON requires
                                  decimal values for mode bits. If not specified,
                                  the volume defaultMode will be used. This might
                                  be in conflict with other options that affect the
                                  file mode, like fsGroup, and the result can be other
                                  mode bits set.'
                                format: int32
                                type: integer
                              path:
                                description: 'Required: Path is  the relative path
                                  name of the file to be created. Must not be absolute
                                  or contain the ''..'' path. Must be utf-8 encoded.
                                  The first item of the relative path must not start
                                  with ''..'''
                                type: string
                              resourceFieldRef:
                                description: 'Selects a resource of the container:
                                  only resources limits and requests (limits.cpu,
                                  limits.memory, requests.cpu and requests.memory)
                                  are currently supported.'
                                properties:
                                  containerName:
                                    description: 'Container name: required for volumes,
                                      optional for env vars'
                                    type: string
                                  divisor:
                                    anyOf:
                                    - type: integer
                                    - type: string
                                    description: Specifies the output format of the
                                      exposed resources, defaults to "1"
                                    pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                    x-kubernetes-int-or-string: true
                                  resource:
                                    description: 'Required: resource to select'
                                    type: string
                                required:
                                - resource
                                type: object
                            required:
                            - path
                            type: object
                          type: array
                      type: object
                    emptyDir:
                      description: 'EmptyDir represents a temporary directory that
                        shares a pod''s lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir'
                      properties:
                        medium:
                          description: 'What type of storage medium should back this
                            directory. The default is "" which means to use the node''s
                            default medium. Must be an empty string (default) or Memory.
                            More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir'
                          type: string
                        sizeLimit:
                          anyOf:
                          - type: integer
                          - type: string
                          description: 'Total amount of local storage required for
                            this EmptyDir volume. The size limit is also applicable
                            for memory medium. The maximum usage on memory medium
                            EmptyDir would be the minimum value between the SizeLimit
                            specified here and the sum of memory limits of all containers
                            in a pod. The default is nil which means that the limit
                            is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir'
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                      type: object
                    ephemeral:
                      description: "Ephemeral represents a volume that is handled
                        by a cluster storage driver (Alpha feature). The volume's
                        lifecycle is tied to the pod that defines it - it will be
                        created before the pod starts, and deleted when the pod is
                        removed. \n Use this if: a) the volume is only needed while
                        the pod runs, b) features of normal volumes like restoring
                        from snapshot or capacity    tracking are needed, c) the storage
                        driver is specified through a storage class, and d) the storage
                        driver supports dynamic volume provisioning through    a PersistentVolumeClaim
                        (see EphemeralVolumeSource for more    information on the
                        connection between this volume type    and PersistentVolumeClaim).
                        \n Use PersistentVolumeClaim or one of the vendor-specific
                        APIs for volumes that persist for longer than the lifecycle
                        of an individual pod. \n Use CSI for light-weight local ephemeral
                        volumes if the CSI driver is meant to be used that way - see
                        the documentation of the driver for more information. \n A
                        pod can use both types of ephemeral volumes and persistent
                        volumes at the same time."
                      properties:
                        readOnly:
                          description: Specifies a read-only configuration for the
                            volume. Defaults to false (read/write).
                          type: boolean
                        volumeClaimTemplate:
                          description: "Will be used to create a stand-alone PVC to
                            provision the volume. The pod in which this EphemeralVolumeSource
                            is embedded will be the owner of the PVC, i.e. the PVC
                            will be deleted together with the pod.  The name of the
                            PVC will be <pod name>-<volume name> where <volume
                            name> is the name from the PodSpec.Volumes array entry.
                            Pod validation will reject the pod if the concatenated
                            name is not valid for a PVC (for example, too long). \n
                            An existing PVC with that name that is not owned by the
                            pod will *not* be used for the pod to avoid using an unrelated
                            volume by mistake. Starting the pod is then blocked until
                            the unrelated PVC is removed. If such a pre-created PVC
                            is meant to be used by the pod, the PVC has to updated
                            with an owner reference to the pod once the pod exists.
                            Normally this should not be necessary, but it may be useful
                            when manually reconstructing a broken cluster. \n This
                            field is read-only and no changes will be made by Kubernetes
                            to the PVC after it has been created. \n Required, must
                            not be nil."
                          properties:
                            metadata:
                              description: May contain labels and annotations that
                                will be copied into the PVC when creating it. No other
                                fields are allowed and will be rejected during validation.
                              type: object
                            spec:
                              description: The specification for the PersistentVolumeClaim.
                                The entire content is copied unchanged into the PVC
                                that gets created from this template. The same fields
                                as in a PersistentVolumeClaim are also valid here.
                              properties:
                                accessModes:
                                  description: 'AccessModes contains the desired access
                                    modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1'
                                  items:
                                    type: string
                                  type: array
                                dataSource:
                                  description: 'This field can be used to specify
                                    either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot
                                    - Beta) * An existing PVC (PersistentVolumeClaim)
                                    * An existing custom resource/object that implements
                                    data population (Alpha) In order to use VolumeSnapshot
                                    object types, the appropriate feature gate must
                                    be enabled (VolumeSnapshotDataSource or AnyVolumeDataSource)
                                    If the provisioner or an external controller can
                                    support the specified data source, it will create
                                    a new volume based on the contents of the specified
                                    data source. If the specified data source is not
                                    supported, the volume will not be created and
                                    the failure will be reported as an event. In the
                                    future, we plan to support more data source types
                                    and the behavior of the provisioner may change.'
                                  properties:
                                    apiGroup:
                                      description: APIGroup is the group for the resource
                                        being referenced. If APIGroup is not specified,
                                        the specified Kind must be in the core API
                                        group. For any other third-party types, APIGroup
                                        is required.
                                      type: string
                                    kind:
                                      description: Kind is the type of resource being
                                        referenced
                                      type: string
                                    name:
                                      description: Name is the name of resource being
                                        referenced
                                      type: string
                                  required:
                                  - kind
                                  - name
                                  type: object
                                resources:
                                  description: 'Resources represents the minimum resources
                                    the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources'
                                  properties:
                                    limits:
                                      additionalProperties:
                                        anyOf:
                                        - type: integer
                                        - type: string
                                        pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                        x-kubernetes-int-or-string: true
                                      description: 'Limits describes the maximum amount
                                        of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/'
                                      type: object
                                    requests:
                                      additionalProperties:
                                        anyOf:
                                        - type: integer
                                        - type: string
                                        pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                        x-kubernetes-int-or-string: true
                                      description: 'Requests describes the minimum
                                        amount of compute resources required. If Requests
                                        is omitted for a container, it defaults to
                                        Limits if that is explicitly specified, otherwise
                                        to an implementation-defined value. More info:
                                        https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/'
                                      type: object
                                  type: object
                                selector:
                                  description: A label query over volumes to consider
                                    for binding.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                storageClassName:
                                  description: 'Name of the StorageClass required
                                    by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1'
                                  type: string
                                volumeMode:
                                  description: volumeMode defines what type of volume
                                    is required by the claim. Value of Filesystem
                                    is implied when not included in claim spec.
                                  type: string
                                volumeName:
                                  description: VolumeName is the binding reference
                                    to the PersistentVolume backing this claim.
                                  type: string
                              type: object
                          required:
                          - spec
                          type: object
                      type: object
                    fc:
                      description: FC represents a Fibre Channel resource that is
                        attached to a kubelet's host machine and then exposed to the
                        pod.
                      properties:
                        fsType:
                          description: 'Filesystem type to mount. Must be a filesystem
                            type supported by the host operating system. Ex. "ext4",
                            "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
                            TODO: how do we prevent errors in the filesystem from
                            compromising the machine'
                          type: string
                        lun:
                          description: 'Optional: FC target lun number'
                          format: int32
                          type: integer
                        readOnly:
                          description: 'Optional: Defaults to false (read/write).
                            ReadOnly here will force the ReadOnly setting in VolumeMounts.'
                          type: boolean
                        targetWWNs:
                          description: 'Optional: FC target worldwide names (WWNs)'
                          items:
                            type: string
                          type: array
                        wwids:
                          description: 'Optional: FC volume world wide identifiers
                            (wwids) Either wwids or combination of targetWWNs and
                            lun must be set, but not both simultaneously.'
                          items:
                            type: string
                          type: array
                      type: object
                    flexVolume:
                      description: FlexVolume represents a generic volume resource
                        that is provisioned/attached using an exec based plugin.
                      properties:
                        driver:
                          description: Driver is the name of the driver to use for
                            this volume.
                          type: string
                        fsType:
                          description: Filesystem type to mount. Must be a filesystem
                            type supported by the host operating system. Ex. "ext4",
                            "xfs", "ntfs". The default filesystem depends on FlexVolume
                            script.
                          type: string
                        options:
                          additionalProperties:
                            type: string
                          description: 'Optional: Extra command options if any.'
                          type: object
                        readOnly:
                          description: 'Optional: Defaults to false (read/write).
                            ReadOnly here will force the ReadOnly setting in VolumeMounts.'
                          type: boolean
                        secretRef:
                          description: 'Optional: SecretRef is reference to the secret
                            object containing sensitive information to pass to the
                            plugin scripts. This may be empty if no secret object
                            is specified. If the secret object contains more than
                            one secret, all secrets are passed to the plugin scripts.'
                          properties:
                            name:
                              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                TODO: Add other useful fields. apiVersion, kind, uid?'
                              type: string
                          type: object
                      required:
                      - driver
                      type: object
                    flocker:
                      description: Flocker represents a Flocker volume attached to
                        a kubelet's host machine. This depends on the Flocker control
                        service being running
                      properties:
                        datasetName:
                          description: Name of the dataset stored as metadata -> name
                            on the dataset for Flocker should be considered as deprecated
                          type: string
                        datasetUUID:
                          description: UUID of the dataset. This is unique identifier
                            of a Flocker dataset
                          type: string
                      type: object
                    gcePersistentDisk:
                      description: 'GCEPersistentDisk represents a GCE Disk resource
                        that is attached to a kubelet''s host machine and then exposed
                        to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk'
                      properties:
                        fsType:
                          description: 'Filesystem type of the volume that you want
                            to mount. Tip: Ensure that the filesystem type is supported
                            by the host operating system. Examples: "ext4", "xfs",
                            "ntfs". Implicitly inferred to be "ext4" if unspecified.
                            More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
                            TODO: how do we prevent errors in the filesystem from
                            compromising the machine'
                          type: string
                        partition:
                          description: 'The partition in the volume that you want
                            to mount. If omitted, the default is to mount by volume
                            name. Examples: For volume /dev/sda1, you specify the
                            partition as "1". Similarly, the volume partition for
                            /dev/sda is "0" (or you can leave the property empty).
                            More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk'
                          format: int32
                          type: integer
                        pdName:
                          description: 'Unique name of the PD resource in GCE. Used
                            to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk'
                          type: string
                        readOnly:
                          description: 'ReadOnly here will force the ReadOnly setting
                            in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk'
                          type: boolean
                      required:
                      - pdName
                      type: object
                    gitRepo:
                      description: 'GitRepo represents a git repository at a particular
                        revision. DEPRECATED: GitRepo is deprecated. To provision
                        a container with a git repo, mount an EmptyDir into an InitContainer
                        that clones the repo using git, then mount the EmptyDir into
                        the Pod''s container.'
                      properties:
                        directory:
                          description: Target directory name. Must not contain or
                            start with '..'.  If '.' is supplied, the volume directory
                            will be the git repository.  Otherwise, if specified,
                            the volume will contain the git repository in the subdirectory
                            with the given name.
                          type: string
                        repository:
                          description: Repository URL
                          type: string
                        revision:
                          description: Commit hash for the specified revision.
                          type: string
                      required:
                      - repository
                      type: object
                    glusterfs:
                      description: 'Glusterfs represents a Glusterfs mount on the
                        host that shares a pod''s lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md'
                      properties:
                        endpoints:
                          description: 'EndpointsName is the endpoint name that details
                            Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod'
                          type: string
                        path:
                          description: 'Path is the Glusterfs volume path. More info:
                            https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod'
                          type: string
                        readOnly:
                          description: 'ReadOnly here will force the Glusterfs volume
                            to be mounted with read-only permissions. Defaults to
                            false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod'
                          type: boolean
                      required:
                      - endpoints
                      - path
                      type: object
                    hostPath:
                      description: 'HostPath represents a pre-existing file or directory
                        on the host machine that is directly exposed to the container.
                        This is generally used for system agents or other privileged
                        things that are allowed to see the host machine. Most containers
                        will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
                        --- TODO(jonesdl) We need to restrict who can use host directory
                        mounts and who can/can not mount host directories as read/write.'
                      properties:
                        path:
                          description: 'Path of the directory on the host. If the
                            path is a symlink, it will follow the link to the real
                            path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath'
                          type: string
                        type:
                          description: 'Type for HostPath Volume Defaults to "" More
                            info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath'
                          type: string
                      required:
                      - path
                      type: object
                    iscsi:
                      description: 'ISCSI represents an ISCSI Disk resource that is
                        attached to a kubelet''s host machine and then exposed to
                        the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md'
                      properties:
                        chapAuthDiscovery:
                          description: whether support iSCSI Discovery CHAP authentication
                          type: boolean
                        chapAuthSession:
                          description: whether support iSCSI Session CHAP authentication
                          type: boolean
                        fsType:
                          description: 'Filesystem type of the volume that you want
                            to mount. Tip: Ensure that the filesystem type is supported
                            by the host operating system. Examples: "ext4", "xfs",
                            "ntfs". Implicitly inferred to be "ext4" if unspecified.
                            More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
                            TODO: how do we prevent errors in the filesystem from
                            compromising the machine'
                          type: string
                        initiatorName:
                          description: Custom iSCSI Initiator Name. If initiatorName
                            is specified with iscsiInterface simultaneously, new iSCSI
                            interface <target portal>:<volume name> will be created
                            for the connection.
                          type: string
                        iqn:
                          description: Target iSCSI Qualified Name.
                          type: string
                        iscsiInterface:
                          description: iSCSI Interface Name that uses an iSCSI transport.
                            Defaults to 'default' (tcp).
                          type: string
                        lun:
                          description: iSCSI Target Lun number.
                          format: int32
                          type: integer
                        portals:
                          description: iSCSI Target Portal List. The portal is either
                            an IP or ip_addr:port if the port is other than default
                            (typically TCP ports 860 and 3260).
                          items:
                            type: string
                          type: array
                        readOnly:
                          description: ReadOnly here will force the ReadOnly setting
                            in VolumeMounts. Defaults to false.
                          type: boolean
                        secretRef:
                          description: CHAP Secret for iSCSI target and initiator
                            authentication
                          properties:
                            name:
                              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                TODO: Add other useful fields. apiVersion, kind, uid?'
                              type: string
                          type: object
                        targetPortal:
                          description: iSCSI Target Portal. The Portal is either an
                            IP or ip_addr:port if the port is other than default (typically
                            TCP ports 860 and 3260).
                          type: string
                      required:
                      - iqn
                      - lun
                      - targetPortal
                      type: object
                    name:
                      description: 'Volume''s name. Must be a DNS_LABEL and unique
                        within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                      type: string
                    nfs:
                      description: 'NFS represents an NFS mount on the host that shares
                        a pod''s lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs'
                      properties:
                        path:
                          description: 'Path that is exported by the NFS server. More
                            info: https://kubernetes.io/docs/concepts/storage/volumes#nfs'
                          type: string
                        readOnly:
                          description: 'ReadOnly here will force the NFS export to
                            be mounted with read-only permissions. Defaults to false.
                            More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs'
                          type: boolean
                        server:
                          description: 'Server is the hostname or IP address of the
                            NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs'
                          type: string
                      required:
                      - path
                      - server
                      type: object
                    persistentVolumeClaim:
                      description: 'PersistentVolumeClaimVolumeSource represents a
                        reference to a PersistentVolumeClaim in the same namespace.
                        More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims'
                      properties:
                        claimName:
                          description: 'ClaimName is the name of a PersistentVolumeClaim
                            in the same namespace as the pod using this volume. More
                            info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims'
                          type: string
                        readOnly:
                          description: Will force the ReadOnly setting in VolumeMounts.
                            Default false.
                          type: boolean
                      required:
                      - claimName
                      type: object
                    photonPersistentDisk:
                      description: PhotonPersistentDisk represents a PhotonController
                        persistent disk attached and mounted on kubelets host machine
                      properties:
                        fsType:
                          description: Filesystem type to mount. Must be a filesystem
                            type supported by the host operating system. Ex. "ext4",
                            "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
                          type: string
                        pdID:
                          description: ID that identifies Photon Controller persistent
                            disk
                          type: string
                      required:
                      - pdID
                      type: object
                    portworxVolume:
                      description: PortworxVolume represents a portworx volume attached
                        and mounted on kubelets host machine
                      properties:
                        fsType:
                          description: FSType represents the filesystem type to mount
                            Must be a filesystem type supported by the host operating
                            system. Ex. "ext4", "xfs". Implicitly inferred to be "ext4"
                            if unspecified.
                          type: string
                        readOnly:
                          description: Defaults to false (read/write). ReadOnly here
                            will force the ReadOnly setting in VolumeMounts.
                          type: boolean
                        volumeID:
                          description: VolumeID uniquely identifies a Portworx volume
                          type: string
                      required:
                      - volumeID
                      type: object
                    projected:
                      description: Items for all in one resources secrets, configmaps,
                        and downward API
                      properties:
                        defaultMode:
                          description: Mode bits used to set permissions on created
                            files by default. Must be an octal value between 0000
                            and 0777 or a decimal value between 0 and 511. YAML accepts
                            both octal and decimal values, JSON requires decimal values
                            for mode bits. Directories within the path are not affected
                            by this setting. This might be in conflict with other
                            options that affect the file mode, like fsGroup, and the
                            result can be other mode bits set.
                          format: int32
                          type: integer
                        sources:
                          description: list of volume projections
                          items:
                            description: Projection that may be projected along with
                              other supported volume types
                            properties:
                              configMap:
                                description: information about the configMap data
                                  to project
                                properties:
                                  items:
                                    description: If unspecified, each key-value pair
                                      in the Data field of the referenced ConfigMap
                                      will be projected into the volume as a file
                                      whose name is the key and content is the value.
                                      If specified, the listed keys will be projected
                                      into the specified paths, and unlisted keys
                                      will not be present. If a key is specified which
                                      is not present in the ConfigMap, the volume
                                      setup will error unless it is marked optional.
                                      Paths must be relative and may not contain the
                                      '..' path or start with '..'.
                                    items:
                                      description: Maps a string key to a path within
                                        a volume.
                                      properties:
                                        key:
                                          description: The key to project.
                                          type: string
                                        mode:
                                          description: 'Optional: mode bits used to
                                            set permissions on this file. Must be
                                            an octal value between 0000 and 0777 or
                                            a decimal value between 0 and 511. YAML
                                            accepts both octal and decimal values,
                                            JSON requires decimal values for mode
                                            bits. If not specified, the volume defaultMode
                                            will be used. This might be in conflict
                                            with other options that affect the file
                                            mode, like fsGroup, and the result can
                                            be other mode bits set.'
                                          format: int32
                                          type: integer
                                        path:
                                          description: The relative path of the file
                                            to map the key to. May not be an absolute
                                            path. May not contain the path element
                                            '..'. May not start with the string '..'.
                                          type: string
                                      required:
                                      - key
                                      - path
                                      type: object
                                    type: array
                                  name:
                                    description: 'Name of the referent. More info:
                                      https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                      TODO: Add other useful fields. apiVersion, kind,
                                      uid?'
                                    type: string
                                  optional:
                                    description: Specify whether the ConfigMap or
                                      its keys must be defined
                                    type: boolean
                                type: object
                              downwardAPI:
                                description: information about the downwardAPI data
                                  to project
                                properties:
                                  items:
                                    description: Items is a list of DownwardAPIVolume
                                      file
                                    items:
                                      description: DownwardAPIVolumeFile represents
                                        information to create the file containing
                                        the pod field
                                      properties:
                                        fieldRef:
                                          description: 'Required: Selects a field
                                            of the pod: only annotations, labels,
                                            name and namespace are supported.'
                                          properties:
                                            apiVersion:
                                              description: Version of the schema the
                                                FieldPath is written in terms of,
                                                defaults to "v1".
                                              type: string
                                            fieldPath:
                                              description: Path of the field to select
                                                in the specified API version.
                                              type: string
                                          required:
                                          - fieldPath
                                          type: object
                                        mode:
                                          description: 'Optional: mode bits used to
                                            set permissions on this file, must be
                                            an octal value between 0000 and 0777 or
                                            a decimal value between 0 and 511. YAML
                                            accepts both octal and decimal values,
                                            JSON requires decimal values for mode
                                            bits. If not specified, the volume defaultMode
                                            will be used. This might be in conflict
                                            with other options that affect the file
                                            mode, like fsGroup, and the result can
                                            be other mode bits set.'
                                          format: int32
                                          type: integer
                                        path:
                                          description: 'Required: Path is  the relative
                                            path name of the file to be created. Must
                                            not be absolute or contain the ''..''
                                            path. Must be utf-8 encoded. The first
                                            item of the relative path must not start
                                            with ''..'''
                                          type: string
                                        resourceFieldRef:
                                          description: 'Selects a resource of the
                                            container: only resources limits and requests
                                            (limits.cpu, limits.memory, requests.cpu
                                            and requests.memory) are currently supported.'
                                          properties:
                                            containerName:
                                              description: 'Container name: required
                                                for volumes, optional for env vars'
                                              type: string
                                            divisor:
                                              anyOf:
                                              - type: integer
                                              - type: string
                                              description: Specifies the output format
                                                of the exposed resources, defaults
                                                to "1"
                                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                              x-kubernetes-int-or-string: true
                                            resource:
                                              description: 'Required: resource to
                                                select'
                                              type: string
                                          required:
                                          - resource
                                          type: object
                                      required:
                                      - path
                                      type: object
                                    type: array
                                type: object
                              secret:
                                description: information about the secret data to
                                  project
                                properties:
                                  items:
                                    description: If unspecified, each key-value pair
                                      in the Data field of the referenced Secret will
                                      be projected into the volume as a file whose
                                      name is the key and content is the value. If
                                      specified, the listed keys will be projected
                                      into the specified paths, and unlisted keys
                                      will not be present. If a key is specified which
                                      is not present in the Secret, the volume setup
                                      will error unless it is marked optional. Paths
                                      must be relative and may not contain the '..'
                                      path or start with '..'.
                                    items:
                                      description: Maps a string key to a path within
                                        a volume.
                                      properties:
                                        key:
                                          description: The key to project.
                                          type: string
                                        mode:
                                          description: 'Optional: mode bits used to
                                            set permissions on this file. Must be
                                            an octal value between 0000 and 0777 or
                                            a decimal value between 0 and 511. YAML
                                            accepts both octal and decimal values,
                                            JSON requires decimal values for mode
                                            bits. If not specified, the volume defaultMode
                                            will be used. This might be in conflict
                                            with other options that affect the file
                                            mode, like fsGroup, and the result can
                                            be other mode bits set.'
                                          format: int32
                                          type: integer
                                        path:
                                          description: The relative path of the file
                                            to map the key to. May not be an absolute
                                            path. May not contain the path element
                                            '..'. May not start with the string '..'.
                                          type: string
                                      required:
                                      - key
                                      - path
                                      type: object
                                    type: array
                                  name:
                                    description: 'Name of the referent. More info:
                                      https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                      TODO: Add other useful fields. apiVersion, kind,
                                      uid?'
                                    type: string
                                  optional:
                                    description: Specify whether the Secret or its
                                      key must be defined
                                    type: boolean
                                type: object
                              serviceAccountToken:
                                description: information about the serviceAccountToken
                                  data to project
                                properties:
                                  audience:
                                    description: Audience is the intended audience
                                      of the token. A recipient of a token must identify
                                      itself with an identifier specified in the audience
                                      of the token, and otherwise should reject the
                                      token. The audience defaults to the identifier
                                      of the apiserver.
                                    type: string
                                  expirationSeconds:
                                    description: ExpirationSeconds is the requested
                                      duration of validity of the service account
                                      token. As the token approaches expiration, the
                                      kubelet volume plugin will proactively rotate
                                      the service account token. The kubelet will
                                      start trying to rotate the token if the token
                                      is older than 80 percent of its time to live
                                      or if the token is older than 24 hours.Defaults
                                      to 1 hour and must be at least 10 minutes.
                                    format: int64
                                    type: integer
                                  path:
                                    description: Path is the path relative to the
                                      mount point of the file to project the token
                                      into.
                                    type: string
                                required:
                                - path
                                type: object
                            type: object
                          type: array
                      required:
                      - sources
                      type: object
                    quobyte:
                      description: Quobyte represents a Quobyte mount on the host
                        that shares a pod's lifetime
                      properties:
                        group:
                          description: Group to map volume access to Default is no
                            group
                          type: string
                        readOnly:
                          description: ReadOnly here will force the Quobyte volume
                            to be mounted with read-only permissions. Defaults to
                            false.
                          type: boolean
                        registry:
                          description: Registry represents a single or multiple Quobyte
                            Registry services specified as a string as host:port pair
                            (multiple entries are separated with commas) which acts
                            as the central registry for volumes
                          type: string
                        tenant:
                          description: Tenant owning the given Quobyte volume in the
                            Backend Used with dynamically provisioned Quobyte volumes,
                            value is set by the plugin
                          type: string
                        user:
                          description: User to map volume access to Defaults to serivceaccount
                            user
                          type: string
                        volume:
                          description: Volume is a string that references an already
                            created Quobyte volume by name.
                          type: string
                      required:
                      - registry
                      - volume
                      type: object
                    rbd:
                      description: 'RBD represents a Rados Block Device mount on the
                        host that shares a pod''s lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md'
                      properties:
                        fsType:
                          description: 'Filesystem type of the volume that you want
                            to mount. Tip: Ensure that the filesystem type is supported
                            by the host operating system. Examples: "ext4", "xfs",
                            "ntfs". Implicitly inferred to be "ext4" if unspecified.
                            More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
                            TODO: how do we prevent errors in the filesystem from
                            compromising the machine'
                          type: string
                        image:
                          description: 'The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it'
                          type: string
                        keyring:
                          description: 'Keyring is the path to key ring for RBDUser.
                            Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it'
                          type: string
                        monitors:
                          description: 'A collection of Ceph monitors. More info:
                            https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it'
                          items:
                            type: string
                          type: array
                        pool:
                          description: 'The rados pool name. Default is rbd. More
                            info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it'
                          type: string
                        readOnly:
                          description: 'ReadOnly here will force the ReadOnly setting
                            in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it'
                          type: boolean
                        secretRef:
                          description: 'SecretRef is name of the authentication secret
                            for RBDUser. If provided overrides keyring. Default is
                            nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it'
                          properties:
                            name:
                              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                TODO: Add other useful fields. apiVersion, kind, uid?'
                              type: string
                          type: object
                        user:
                          description: 'The rados user name. Default is admin. More
                            info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it'
                          type: string
                      required:
                      - image
                      - monitors
                      type: object
                    scaleIO:
                      description: ScaleIO represents a ScaleIO persistent volume
                        attached and mounted on Kubernetes nodes.
                      properties:
                        fsType:
                          description: Filesystem type to mount. Must be a filesystem
                            type supported by the host operating system. Ex. "ext4",
                            "xfs", "ntfs". Default is "xfs".
                          type: string
                        gateway:
                          description: The host address of the ScaleIO API Gateway.
                          type: string
                        protectionDomain:
                          description: The name of the ScaleIO Protection Domain for
                            the configured storage.
                          type: string
                        readOnly:
                          description: Defaults to false (read/write). ReadOnly here
                            will force the ReadOnly setting in VolumeMounts.
                          type: boolean
                        secretRef:
                          description: SecretRef references to the secret for ScaleIO
                            user and other sensitive information. If this is not provided,
                            Login operation will fail.
                          properties:
                            name:
                              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                TODO: Add other useful fields. apiVersion, kind, uid?'
                              type: string
                          type: object
                        sslEnabled:
                          description: Flag to enable/disable SSL communication with
                            Gateway, default false
                          type: boolean
                        storageMode:
                          description: Indicates whether the storage for a volume
                            should be ThickProvisioned or ThinProvisioned. Default
                            is ThinProvisioned.
                          type: string
                        storagePool:
                          description: The ScaleIO Storage Pool associated with the
                            protection domain.
                          type: string
                        system:
                          description: The name of the storage system as configured
                            in ScaleIO.
                          type: string
                        volumeName:
                          description: The name of a volume already created in the
                            ScaleIO system that is associated with this volume source.
                          type: string
                      required:
                      - gateway
                      - secretRef
                      - system
                      type: object
                    secret:
                      description: 'Secret represents a secret that should populate
                        this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret'
                      properties:
                        defaultMode:
                          description: 'Optional: mode bits used to set permissions
                            on created files by default. Must be an octal value between
                            0000 and 0777 or a decimal value between 0 and 511. YAML
                            accepts both octal and decimal values, JSON requires decimal
                            values for mode bits. Defaults to 0644. Directories within
                            the path are not affected by this setting. This might
                            be in conflict with other options that affect the file
                            mode, like fsGroup, and the result can be other mode bits
                            set.'
                          format: int32
                          type: integer
                        items:
                          description: If unspecified, each key-value pair in the
                            Data field of the referenced Secret will be projected
                            into the volume as a file whose name is the key and content
                            is the value. If specified, the listed keys will be projected
                            into the specified paths, and unlisted keys will not be
                            present. If a key is specified which is not present in
                            the Secret, the volume setup will error unless it is marked
                            optional. Paths must be relative and may not contain the
                            '..' path or start with '..'.
                          items:
                            description: Maps a string key to a path within a volume.
                            properties:
                              key:
                                description: The key to project.
                                type: string
                              mode:
                                description: 'Optional: mode bits used to set permissions
                                  on this file. Must be an octal value between 0000
                                  and 0777 or a decimal value between 0 and 511. YAML
                                  accepts both octal and decimal values, JSON requires
                                  decimal values for mode bits. If not specified,
                                  the volume defaultMode will be used. This might
                                  be in conflict with other options that affect the
                                  file mode, like fsGroup, and the result can be other
                                  mode bits set.'
                                format: int32
                                type: integer
                              path:
                                description: The relative path of the file to map
                                  the key to. May not be an absolute path. May not
                                  contain the path element '..'. May not start with
                                  the string '..'.
                                type: string
                            required:
                            - key
                            - path
                            type: object
                          type: array
                        optional:
                          description: Specify whether the Secret or its keys must
                            be defined
                          type: boolean
                        secretName:
                          description: 'Name of the secret in the pod''s namespace
                            to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret'
                          type: string
                      type: object
                    storageos:
                      description: StorageOS represents a StorageOS volume attached
                        and mounted on Kubernetes nodes.
                      properties:
                        fsType:
                          description: Filesystem type to mount. Must be a filesystem
                            type supported by the host operating system. Ex. "ext4",
                            "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
                          type: string
                        readOnly:
                          description: Defaults to false (read/write). ReadOnly here
                            will force the ReadOnly setting in VolumeMounts.
                          type: boolean
                        secretRef:
                          description: SecretRef specifies the secret to use for obtaining
                            the StorageOS API credentials.  If not specified, default
                            values will be attempted.
                          properties:
                            name:
                              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                                TODO: Add other useful fields. apiVersion, kind, uid?'
                              type: string
                          type: object
                        volumeName:
                          description: VolumeName is the human-readable name of the
                            StorageOS volume.  Volume names are only unique within
                            a namespace.
                          type: string
                        volumeNamespace:
                          description: VolumeNamespace specifies the scope of the
                            volume within StorageOS.  If no namespace is specified
                            then the Pod's namespace will be used.  This allows the
                            Kubernetes name scoping to be mirrored within StorageOS
                            for tighter integration. Set VolumeName to any name to
                            override the default behaviour. Set to "default" if you
                            are not using namespaces within StorageOS. Namespaces
                            that do not pre-exist within StorageOS will be created.
                          type: string
                      type: object
                    vsphereVolume:
                      description: VsphereVolume represents a vSphere volume attached
                        and mounted on kubelets host machine
                      properties:
                        fsType:
                          description: Filesystem type to mount. Must be a filesystem
                            type supported by the host operating system. Ex. "ext4",
                            "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
                          type: string
                        storagePolicyID:
                          description: Storage Policy Based Management (SPBM) profile
                            ID associated with the StoragePolicyName.
                          type: string
                        storagePolicyName:
                          description: Storage Policy Based Management (SPBM) profile
                            name.
                          type: string
                        volumePath:
                          description: Path that identifies vSphere volume vmdk
                          type: string
                      required:
                      - volumePath
                      type: object
                  required:
                  - name
                  type: object
                type: array
            type: object
          status:
            description: PodPresetStatus defines the observed state of PodPreset
            properties:
              conditions:
                items:
                  description: "Condition contains details for one aspect of the current
                    state of this API Resource. --- This struct is intended for direct
                    use as an array at the field path .status.conditions.  For example,
                    type FooStatus struct{     // Represents the observations of a
                    foo's current state.     // Known .status.conditions.type are:
                    \"Available\", \"Progressing\", and \"Degraded\"     // +patchMergeKey=type
                    \    // +patchStrategy=merge     // +listType=map     // +listMapKey=type
                    \    Conditions []metav1.Condition json:\"conditions,omitempty\"
                    patchStrategy:\"merge\" patchMergeKey:\"type\" protobuf:\"bytes,1,rep,name=conditions\"
                    \n     // other fields }"
                  properties:
                    lastTransitionTime:
                      description: lastTransitionTime is the last time the condition
                        transitioned from one status to another. This should be when
                        the underlying condition changed.  If that is not known, then
                        using the time when the API field changed is acceptable.
                      format: date-time
                      type: string
                    message:
                      description: message is a human readable message indicating
                        details about the transition. This may be an empty string.
                      maxLength: 32768
                      type: string
                    observedGeneration:
                      description: observedGeneration represents the .metadata.generation
                        that the condition was set based upon. For instance, if .metadata.generation
                        is currently 12, but the .status.conditions[x].observedGeneration
                        is 9, the condition is out of date with respect to the current
                        state of the instance.
                      format: int64
                      minimum: 0
                      type: integer
                    reason:
                      description: reason contains a programmatic identifier indicating
                        the reason for the condition's last transition. Producers
                        of specific condition types may define expected values and
                        meanings for this field, and whether the values are considered
                        a guaranteed API. The value should be a CamelCase string.
                        This field may not be empty.
                      maxLength: 1024
                      minLength: 1
                      pattern: ^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$
                      type: string
                    status:
                      description: status of the condition, one of True, False, Unknown.
                      enum:
                      - "True"
                      - "False"
                      - Unknown
                      type: string
                    type:
                      description: type of condition in CamelCase or in foo.example.com/CamelCase.
                        --- Many .condition.type values are consistent across resources
                        like Available, but because arbitrary conditions can be useful
                        (see .node.status.conditions), the ability to deconflict is
                        important. The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)
                      maxLength: 316
                      pattern: ^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$
                      type: string
                  required:
                  - lastTransitionTime
                  - message
                  - reason
                  - status
                  - type
                  type: object
                type: array
                x-kubernetes-list-map-keys:
                - type
                x-kubernetes-list-type: map
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: podpreset-webhook-leader-election-role
  namespace: podpreset-webhook
rules:
- apiGroups:
  - ""
  - coordination.k8s.io
  resources:
  - configmaps
  - leases
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
  - patch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: null
  name: podpreset-webhook-manager-role
rules:
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - get
  - list
  - patch
  - update
- apiGroups:
  - redhatcop.redhat.io
  resources:
  - podpresets
  verbs:
  - create
  - get
  - list
  - patch
  - update
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: podpreset-webhook-leader-election-rolebinding
  namespace: podpreset-webhook
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: podpreset-webhook-leader-election-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: podpreset-webhook
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: podpreset-webhook-manager-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: podpreset-webhook-manager-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: podpreset-webhook
---
apiVersion: v1
data:
  controller_manager_config.yaml: |
    apiVersion: controller-runtime.sigs.k8s.io/v1alpha1
    kind: ControllerManagerConfig
    health:
      healthProbeBindAddress: :8081
    metrics:
      bindAddress: 127.0.0.1:8080
    webhook:
      port: 9443
    leaderElection:
      leaderElect: true
      resourceName: 7256067a.redhat.io
kind: ConfigMap
metadata:
  name: podpreset-webhook-manager-config
  namespace: podpreset-webhook
---
apiVersion: v1
kind: Service
metadata:
  name: podpreset-webhook-webhook-service
  namespace: podpreset-webhook
spec:
  ports:
  - port: 443
    targetPort: 9443
  selector:
    control-plane: controller-manager
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: controller-manager
  name: podpreset-webhook-controller-manager
  namespace: podpreset-webhook
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: controller-manager
  template:
    metadata:
      labels:
        control-plane: controller-manager
    spec:
      containers:
      - args:
        - --leader-elect
        command:
        - /manager
        image: quay.io/redhat-cop/podpreset-webhook:latest
        livenessProbe:
          httpGet:
            path: /healthz
            port: 8081
          initialDelaySeconds: 15
          periodSeconds: 20
        name: manager
        ports:
        - containerPort: 9443
          name: webhook-server
          protocol: TCP
        readinessProbe:
          httpGet:
            path: /readyz
            port: 8081
          initialDelaySeconds: 5
          periodSeconds: 10
        resources:
          limits:
            cpu: 100m
            memory: 300Mi
          requests:
            cpu: 100m
            memory: 100Mi
        securityContext:
          allowPrivilegeEscalation: false
        volumeMounts:
        - mountPath: /apiserver.local.config/certificates
          name: apiservice-cert
          readOnly: true
      terminationGracePeriodSeconds: 10
      volumes:
      - name: apiservice-cert
        secret:
          defaultMode: 420
          items:
          - key: tls.key
            path: apiserver.key
          - key: tls.crt
            path: apiserver.crt
          secretName: webhook-server-cert
---
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: podpreset-webhook-serving-cert
  namespace: podpreset-webhook
spec:
  dnsNames:
  - podpreset-webhook-webhook-service.podpreset-webhook.svc
  - podpreset-webhook-webhook-service.podpreset-webhook.svc.cluster.local
  issuerRef:
    kind: Issuer
    name: podpreset-webhook-selfsigned-issuer
  secretName: webhook-server-cert
---
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: podpreset-webhook-selfsigned-issuer
  namespace: podpreset-webhook
spec:
  selfSigned: {}
---
apiVersion: admissionregistration.k8s.io/v1
kind: MutatingWebhookConfiguration
metadata:
  annotations:
    cert-manager.io/inject-ca-from: podpreset-webhook/podpreset-webhook-serving-cert
  name: podpreset-webhook-mutating-webhook-configuration
webhooks:
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: podpreset-webhook-webhook-service
      namespace: podpreset-webhook
      path: /mutate
  failurePolicy: Ignore
  name: mpod.redhatcop.redhat.io
  rules:
  - apiGroups:
    - ""
    apiVersions:
    - v1
    operations:
    - CREATE
    resources:
    - pods
  sideEffects: None
`