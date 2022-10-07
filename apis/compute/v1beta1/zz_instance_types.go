/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccessConfigObservation struct {
}

type AccessConfigParameters struct {

	// If the instance has an access config, either the given external ip (in the nat_ip field) or the ephemeral (generated) ip (if you didn't provide one).
	// +kubebuilder:validation:Optional
	NATIP *string `json:"natIp,omitempty" tf:"nat_ip,omitempty"`

	// The service-level to be provided for IPv6 traffic when the
	// subnet has an external subnet. Only PREMIUM or STANDARD tier is valid for IPv6.
	// +kubebuilder:validation:Optional
	NetworkTier *string `json:"networkTier,omitempty" tf:"network_tier,omitempty"`

	// The domain name to be used when creating DNSv6
	// records for the external IPv6 ranges..
	// +kubebuilder:validation:Optional
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty" tf:"public_ptr_domain_name,omitempty"`
}

type AdvancedMachineFeaturesObservation struct {
}

type AdvancedMachineFeaturesParameters struct {

	// Defines whether the instance should have nested virtualization  enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableNestedVirtualization *bool `json:"enableNestedVirtualization,omitempty" tf:"enable_nested_virtualization,omitempty"`

	// he number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1.
	// +kubebuilder:validation:Optional
	ThreadsPerCore *float64 `json:"threadsPerCore,omitempty" tf:"threads_per_core,omitempty"`
}

type AliasIPRangeObservation struct {
}

type AliasIPRangeParameters struct {

	// The IP CIDR range represented by this alias IP range. This IP CIDR range
	// must belong to the specified subnetwork and cannot contain IP addresses reserved by
	// system or used by other network interfaces. This range may be a single IP address
	// (e.g. 10.2.3.4), a netmask (e.g. /24) or a CIDR format string (e.g. 10.1.2.0/24).
	// +kubebuilder:validation:Required
	IPCidrRange *string `json:"ipCidrRange" tf:"ip_cidr_range,omitempty"`

	// The subnetwork secondary range name specifying
	// the secondary range from which to allocate the IP CIDR range for this alias IP
	// range. If left unspecified, the primary range of the subnetwork will be used.
	// +kubebuilder:validation:Optional
	SubnetworkRangeName *string `json:"subnetworkRangeName,omitempty" tf:"subnetwork_range_name,omitempty"`
}

type BootDiskObservation struct {

	// The RFC 4648 base64
	// encoded SHA-256 hash of the [customer-supplied encryption key]
	// (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this resource.
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty" tf:"disk_encryption_key_sha256,omitempty"`
}

type BootDiskParameters struct {

	// Whether the disk will be auto-deleted when the instance
	// is deleted. Defaults to true.
	// +kubebuilder:validation:Optional
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// Name with which attached disk will be accessible.
	// On the instance, this device will be /dev/disk/by-id/google-{{device_name}}.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// A 256-bit [customer-supplied encryption key]
	// (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption),
	// encoded in RFC 4648 base64
	// to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw
	// may be set.
	// +kubebuilder:validation:Optional
	DiskEncryptionKeyRawSecretRef *v1.SecretKeySelector `json:"diskEncryptionKeyRawSecretRef,omitempty" tf:"-"`

	// Parameters for a new disk that will be created
	// alongside the new instance. Either initialize_params or source must be set.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	InitializeParams []InitializeParamsParameters `json:"initializeParams,omitempty" tf:"initialize_params,omitempty"`

	// The self_link of the encryption key that is
	// stored in Google Cloud KMS to encrypt this disk. Only one of kms_key_self_link
	// and disk_encryption_key_raw may be set.
	// +kubebuilder:validation:Optional
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// The mode in which to attach this disk, either READ_WRITE
	// or READ_ONLY. If not specified, the default is to attach the disk in READ_WRITE mode.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name or self_link of the existing disk (such as those managed by
	// google_compute_disk) or disk image. To create an instance from a snapshot, first create a
	// google_compute_disk from a snapshot and reference it here.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type ConfidentialInstanceConfigObservation struct {
}

type ConfidentialInstanceConfigParameters struct {

	// Defines whether the instance should have confidential compute enabled. on_host_maintenance has to be set to TERMINATE or this will fail to create the VM.
	// +kubebuilder:validation:Required
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute" tf:"enable_confidential_compute,omitempty"`
}

type GuestAcceleratorObservation struct {
}

type GuestAcceleratorParameters struct {

	// The number of the guest accelerator cards exposed to this instance.
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count,omitempty" tf:"count"`

	// The accelerator type resource to expose to this instance. E.g. nvidia-tesla-k80.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type IPv6AccessConfigObservation struct {

	// The first IPv6 address of the external IPv6 range
	// associated with this instance, prefix length is stored in externalIpv6PrefixLength in ipv6AccessConfig.
	// The field is output only, an IPv6 address from a subnetwork associated with the instance will be allocated dynamically.
	ExternalIPv6 *string `json:"externalIpv6,omitempty" tf:"external_ipv6,omitempty"`

	// The prefix length of the external IPv6 range.
	ExternalIPv6PrefixLength *string `json:"externalIpv6PrefixLength,omitempty" tf:"external_ipv6_prefix_length,omitempty"`
}

type IPv6AccessConfigParameters struct {

	// The service-level to be provided for IPv6 traffic when the
	// subnet has an external subnet. Only PREMIUM or STANDARD tier is valid for IPv6.
	// +kubebuilder:validation:Required
	NetworkTier *string `json:"networkTier" tf:"network_tier,omitempty"`

	// The domain name to be used when creating DNSv6
	// records for the external IPv6 ranges..
	// +kubebuilder:validation:Optional
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty" tf:"public_ptr_domain_name,omitempty"`
}

type InitializeParamsObservation struct {
}

type InitializeParamsParameters struct {

	// The image from which to initialize this disk. This can be
	// one of: the image's self_link, projects/{project}/global/images/{image},
	// projects/{project}/global/images/family/{family}, global/images/{image},
	// global/images/family/{family}, family/{family}, {project}/{family},
	// {project}/{image}, {family}, or {image}. If referred by family, the
	// images names must include the family name. If they don't, use the
	// google_compute_image data source.
	// For instance, the image centos-6-v20180104 includes its family name centos-6.
	// These images can be referred by family name here.
	// +crossplane:generate:reference:type=Image
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Reference to a Image to populate image.
	// +kubebuilder:validation:Optional
	ImageRef *v1.Reference `json:"imageRef,omitempty" tf:"-"`

	// Selector for a Image to populate image.
	// +kubebuilder:validation:Optional
	ImageSelector *v1.Selector `json:"imageSelector,omitempty" tf:"-"`

	// A map of key/value label pairs to assign to the instance.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The size of the image in gigabytes. If not specified, it
	// will inherit the size of its base image.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The type of reservation from which this instance can consume resources.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InstanceAttachedDiskObservation struct {

	// The RFC 4648 base64
	// encoded SHA-256 hash of the [customer-supplied encryption key]
	// (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this resource.
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty" tf:"disk_encryption_key_sha256,omitempty"`
}

type InstanceAttachedDiskParameters struct {

	// Name with which the attached disk will be accessible
	// under /dev/disk/by-id/google-*
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// A 256-bit [customer-supplied encryption key]
	// (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption),
	// encoded in RFC 4648 base64
	// to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	// +kubebuilder:validation:Optional
	DiskEncryptionKeyRawSecretRef *v1.SecretKeySelector `json:"diskEncryptionKeyRawSecretRef,omitempty" tf:"-"`

	// The self_link of the encryption key that is
	// stored in Google Cloud KMS to encrypt this disk. Only one of kms_key_self_link
	// and disk_encryption_key_raw may be set.
	// +kubebuilder:validation:Optional
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// Either "READ_ONLY" or "READ_WRITE", defaults to "READ_WRITE"
	// If you have a persistent disk with data that you want to share
	// between multiple instances, detach it from any read-write instances and
	// attach it to one or more instances in read-only mode.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name or self_link of the disk to attach to this instance.
	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`
}

type InstanceObservation struct {

	// Additional disks to attach to the instance. Can be repeated multiple times for multiple disks. Structure is documented below.
	// +kubebuilder:validation:Optional
	AttachedDisk []InstanceAttachedDiskObservation `json:"attachedDisk,omitempty" tf:"attached_disk,omitempty"`

	// The boot disk for the instance.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	BootDisk []BootDiskObservation `json:"bootDisk,omitempty" tf:"boot_disk,omitempty"`

	// The CPU platform used by this instance.
	CPUPlatform *string `json:"cpuPlatform,omitempty" tf:"cpu_platform,omitempty"`

	CurrentStatus *string `json:"currentStatus,omitempty" tf:"current_status,omitempty"`

	// an identifier for the resource with format projects/{{project}}/zones/{{zone}}/instances/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The server-assigned unique identifier of this instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The unique fingerprint of the labels.
	LabelFingerprint *string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`

	// The unique fingerprint of the metadata.
	MetadataFingerprint *string `json:"metadataFingerprint,omitempty" tf:"metadata_fingerprint,omitempty"`

	// Networks to attach to the instance. This can
	// be specified multiple times. Structure is documented below.
	// +kubebuilder:validation:Required
	NetworkInterface []NetworkInterfaceObservation `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The unique fingerprint of the tags.
	TagsFingerprint *string `json:"tagsFingerprint,omitempty" tf:"tags_fingerprint,omitempty"`
}

type InstanceParameters struct {

	// Configure Nested Virtualisation and Simultaneous Hyper Threading  on this VM. Structure is documented below
	// +kubebuilder:validation:Optional
	AdvancedMachineFeatures []AdvancedMachineFeaturesParameters `json:"advancedMachineFeatures,omitempty" tf:"advanced_machine_features,omitempty"`

	// If you try to update a property that requires stopping the instance without setting this field, the update will fail.
	// +kubebuilder:validation:Optional
	AllowStoppingForUpdate *bool `json:"allowStoppingForUpdate,omitempty" tf:"allow_stopping_for_update,omitempty"`

	// Additional disks to attach to the instance. Can be repeated multiple times for multiple disks. Structure is documented below.
	// +kubebuilder:validation:Optional
	AttachedDisk []InstanceAttachedDiskParameters `json:"attachedDisk,omitempty" tf:"attached_disk,omitempty"`

	// The boot disk for the instance.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	BootDisk []BootDiskParameters `json:"bootDisk" tf:"boot_disk,omitempty"`

	// Whether to allow sending and receiving of
	// packets with non-matching source or destination IPs.
	// This defaults to false.
	// +kubebuilder:validation:Optional
	CanIPForward *bool `json:"canIpForward,omitempty" tf:"can_ip_forward,omitempty"`

	// Enable Confidential Mode on this VM. Structure is documented below
	// +kubebuilder:validation:Optional
	ConfidentialInstanceConfig []ConfidentialInstanceConfigParameters `json:"confidentialInstanceConfig,omitempty" tf:"confidential_instance_config,omitempty"`

	// Enable deletion protection on this instance. Defaults to false.
	// Note: you must disable deletion protection before removing the resource (e.g.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// A brief description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Desired status of the instance. Either
	// "RUNNING" or "TERMINATED".
	// +kubebuilder:validation:Optional
	DesiredStatus *string `json:"desiredStatus,omitempty" tf:"desired_status,omitempty"`

	// Enable Virtual Displays on this instance.
	// Note: allow_stopping_for_update must be set to true or your instance must have a desired_status of TERMINATED in order to update this field.
	// +kubebuilder:validation:Optional
	EnableDisplay *bool `json:"enableDisplay,omitempty" tf:"enable_display,omitempty"`

	// List of the type and count of accelerator cards attached to the instance. Structure documented below.
	// Note: GPU accelerators can only be used with on_host_maintenance option set to TERMINATE.
	// Note: This field uses attr-as-block mode to avoid
	// breaking users during the 0.12 upgrade. To explicitly send a list
	// of zero objects you must use the following syntax:
	// example=[]
	// For more details about this behavior, see this section.
	// +kubebuilder:validation:Optional
	GuestAccelerator []GuestAcceleratorParameters `json:"guestAccelerator,omitempty" tf:"guest_accelerator,omitempty"`

	// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid.
	// Valid format is a series of labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods.
	// The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// A map of key/value label pairs to assign to the instance.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The machine type to create.
	// +kubebuilder:validation:Required
	MachineType *string `json:"machineType" tf:"machine_type,omitempty"`

	// Metadata key/value pairs to make available from
	// within the instance. Ssh keys attached in the Cloud Console will be removed.
	// Add them to your config in order to keep them attached to your instance.
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// An alternative to using the
	// startup-script metadata key, except this one forces the instance to be recreated
	// (thus re-running the script) if it is changed. This replaces the startup-script
	// metadata key on the created instance and thus the two mechanisms are not
	// allowed to be used simultaneously.  Users are free to use either mechanism - the
	// only distinction is that this separate attribute will cause a recreate on
	// modification.  On import, metadata_startup_script will not be set - if you
	// choose to specify it you will see a diff immediately after import causing a
	// destroy/recreate operation.
	// +kubebuilder:validation:Optional
	MetadataStartupScript *string `json:"metadataStartupScript,omitempty" tf:"metadata_startup_script,omitempty"`

	// Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as
	// Intel Haswell or Intel Skylake. See the complete list here.
	// Note: allow_stopping_for_update must be set to true or your instance must have a desired_status of TERMINATED in order to update this field.
	// +kubebuilder:validation:Optional
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty" tf:"min_cpu_platform,omitempty"`

	// Networks to attach to the instance. This can
	// be specified multiple times. Structure is documented below.
	// +kubebuilder:validation:Required
	NetworkInterface []NetworkInterfaceParameters `json:"networkInterface" tf:"network_interface,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Specifies the reservations that this instance can consume from.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ReservationAffinity []ReservationAffinityParameters `json:"reservationAffinity,omitempty" tf:"reservation_affinity,omitempty"`

	// - A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the instance to recreate. Currently a max of 1 resource policy is supported.
	// +kubebuilder:validation:Optional
	ResourcePolicies []*string `json:"resourcePolicies,omitempty" tf:"resource_policies,omitempty"`

	// The scheduling strategy to use. More details about
	// this configuration option are detailed below.
	// +kubebuilder:validation:Optional
	Scheduling []SchedulingParameters `json:"scheduling,omitempty" tf:"scheduling,omitempty"`

	// Scratch disks to attach to the instance. This can be
	// specified multiple times for multiple scratch disks. Structure is documented below.
	// +kubebuilder:validation:Optional
	ScratchDisk []ScratchDiskParameters `json:"scratchDisk,omitempty" tf:"scratch_disk,omitempty"`

	// Service account to attach to the instance.
	// Structure is documented below.
	// Note: allow_stopping_for_update must be set to true or your instance must have a desired_status of TERMINATED in order to update this field.
	// +kubebuilder:validation:Optional
	ServiceAccount []ServiceAccountParameters `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// Enable Shielded VM on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
	// Note: shielded_instance_config can only be used with boot images with shielded vm support. See the complete list here.
	// Note: allow_stopping_for_update must be set to true or your instance must have a desired_status of TERMINATED in order to update this field.
	// +kubebuilder:validation:Optional
	ShieldedInstanceConfig []ShieldedInstanceConfigParameters `json:"shieldedInstanceConfig,omitempty" tf:"shielded_instance_config,omitempty"`

	// A list of network tags to attach to the instance.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The zone that the machine should be created in. If it is not provided, the provider zone is used.
	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type NetworkInterfaceObservation struct {

	// An array of IPv6 access configurations for this interface.
	// Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig
	// specified, then this instance will have no external IPv6 Internet access. Structure documented below.
	// +kubebuilder:validation:Optional
	IPv6AccessConfig []IPv6AccessConfigObservation `json:"ipv6AccessConfig,omitempty" tf:"ipv6_access_config,omitempty"`

	// One of EXTERNAL, INTERNAL to indicate whether the IP can be accessed from the Internet.
	// This field is always inherited from its subnetwork.
	IPv6AccessType *string `json:"ipv6AccessType,omitempty" tf:"ipv6_access_type,omitempty"`

	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type NetworkInterfaceParameters struct {

	// Access configurations, i.e. IPs via which this
	// instance can be accessed via the Internet. Omit to ensure that the instance
	// is not accessible from the Internet.g. via
	// tunnel or because it is running on another cloud instance on that network).
	// This block can be repeated multiple times. Structure documented below.
	// +kubebuilder:validation:Optional
	AccessConfig []AccessConfigParameters `json:"accessConfig,omitempty" tf:"access_config,omitempty"`

	// An
	// array of alias IP ranges for this network interface. Can only be specified for network
	// interfaces on subnet-mode networks. Structure documented below.
	// +kubebuilder:validation:Optional
	AliasIPRange []AliasIPRangeParameters `json:"aliasIpRange,omitempty" tf:"alias_ip_range,omitempty"`

	// An array of IPv6 access configurations for this interface.
	// Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig
	// specified, then this instance will have no external IPv6 Internet access. Structure documented below.
	// +kubebuilder:validation:Optional
	IPv6AccessConfig []IPv6AccessConfigParameters `json:"ipv6AccessConfig,omitempty" tf:"ipv6_access_config,omitempty"`

	// The name or self_link of the network to attach this interface to.
	// Either network or subnetwork must be provided. If network isn't provided it will
	// be inferred from the subnetwork.
	// +crossplane:generate:reference:type=Network
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The private IP address to assign to the instance. If
	// empty, the address will be automatically assigned.
	// +kubebuilder:validation:Optional
	NetworkIP *string `json:"networkIp,omitempty" tf:"network_ip,omitempty"`

	// Reference to a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The type of vNIC to be used on this interface. Possible values: GVNIC, VIRTIO_NET.
	// +kubebuilder:validation:Optional
	NicType *string `json:"nicType,omitempty" tf:"nic_type,omitempty"`

	// The networking queue count that's specified by users for the network interface. Both Rx and Tx queues will be set to this number. It will be empty if not specified.
	// +kubebuilder:validation:Optional
	QueueCount *float64 `json:"queueCount,omitempty" tf:"queue_count,omitempty"`

	// The stack type for this network interface to identify whether the IPv6 feature is enabled or not. Values are IPV4_IPV6 or IPV4_ONLY. If not specified, IPV4_ONLY will be used.
	// +kubebuilder:validation:Optional
	StackType *string `json:"stackType,omitempty" tf:"stack_type,omitempty"`

	// The name or self_link of the subnetwork to attach this
	// interface to. Either network or subnetwork must be provided. If network isn't provided
	// it will be inferred from the subnetwork. The subnetwork must exist in the same region this
	// instance will be created in. If the network resource is in
	// legacy mode, do not specify this field. If the
	// network is in auto subnet mode, specifying the subnetwork is optional. If the network is
	// in custom subnet mode, specifying the subnetwork is required.
	// +crossplane:generate:reference:type=Subnetwork
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// The project in which the subnetwork belongs.
	// If the subnetwork is a self_link, this field is ignored in favor of the project
	// defined in the subnetwork self_link. If the subnetwork is a name and this
	// field is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	SubnetworkProject *string `json:"subnetworkProject,omitempty" tf:"subnetwork_project,omitempty"`

	// Reference to a Subnetwork to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkRef *v1.Reference `json:"subnetworkRef,omitempty" tf:"-"`

	// Selector for a Subnetwork to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkSelector *v1.Selector `json:"subnetworkSelector,omitempty" tf:"-"`
}

type NodeAffinitiesObservation struct {
}

type NodeAffinitiesParameters struct {

	// Corresponds to the label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify compute.googleapis.com/reservation-name as the key and specify the name of your reservation as the only value.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// The operator. Can be IN for node-affinities
	// or NOT_IN for anti-affinities.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// Corresponds to the label values of a reservation resource.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ReservationAffinityObservation struct {
}

type ReservationAffinityParameters struct {

	// Specifies the label selector for the reservation to use..
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SpecificReservation []SpecificReservationParameters `json:"specificReservation,omitempty" tf:"specific_reservation,omitempty"`

	// The type of reservation from which this instance can consume resources.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SchedulingObservation struct {
}

type SchedulingParameters struct {

	// Specifies if the instance should be
	// restarted if it was terminated by Compute Engine (not a user).
	// Defaults to true.
	// +kubebuilder:validation:Optional
	AutomaticRestart *bool `json:"automaticRestart,omitempty" tf:"automatic_restart,omitempty"`

	// The minimum number of virtual CPUs this instance will consume when running on a sole-tenant node.
	// +kubebuilder:validation:Optional
	MinNodeCpus *float64 `json:"minNodeCpus,omitempty" tf:"min_node_cpus,omitempty"`

	// Specifies node affinities or anti-affinities
	// to determine which sole-tenant nodes your instances and managed instance
	// groups will use as host systems. Read more on sole-tenant node creation
	// here.
	// Structure documented below.
	// +kubebuilder:validation:Optional
	NodeAffinities []NodeAffinitiesParameters `json:"nodeAffinities,omitempty" tf:"node_affinities,omitempty"`

	// Describes maintenance behavior for the
	// instance. Can be MIGRATE or TERMINATE, for more info, read
	// here.
	// +kubebuilder:validation:Optional
	OnHostMaintenance *string `json:"onHostMaintenance,omitempty" tf:"on_host_maintenance,omitempty"`

	// Specifies if the instance is preemptible.
	// If this field is set to true, then automatic_restart must be
	// set to false.  Defaults to false.
	// +kubebuilder:validation:Optional
	Preemptible *bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`

	// Describe the type of preemptible VM. This field accepts the value STANDARD or SPOT. If the value is STANDARD, there will be no discount. If this   is set to SPOT,
	// preemptible should be true and auto_restart should be
	// false. For more info about
	// SPOT, read here
	// +kubebuilder:validation:Optional
	ProvisioningModel *string `json:"provisioningModel,omitempty" tf:"provisioning_model,omitempty"`
}

type ScratchDiskObservation struct {
}

type ScratchDiskParameters struct {

	// The disk interface to use for attaching this disk; either SCSI or NVME.
	// +kubebuilder:validation:Required
	Interface *string `json:"interface" tf:"interface,omitempty"`
}

type ServiceAccountObservation struct {
}

type ServiceAccountParameters struct {

	// The service account e-mail address. If not given, the
	// default Google Compute Engine service account is used.
	// Note: allow_stopping_for_update must be set to true or your instance must have a desired_status of TERMINATED in order to update this field.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("email",true)
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Reference to a ServiceAccount in cloudplatform to populate email.
	// +kubebuilder:validation:Optional
	EmailRef *v1.Reference `json:"emailRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount in cloudplatform to populate email.
	// +kubebuilder:validation:Optional
	EmailSelector *v1.Selector `json:"emailSelector,omitempty" tf:"-"`

	// A list of service scopes. Both OAuth2 URLs and gcloud
	// short names are supported. To allow full access to all Cloud APIs, use the
	// cloud-platform scope. See a complete list of scopes here.
	// Note: allow_stopping_for_update must be set to true or your instance must have a desired_status of TERMINATED in order to update this field.
	// +kubebuilder:validation:Required
	Scopes []*string `json:"scopes" tf:"scopes,omitempty"`
}

type ShieldedInstanceConfigObservation struct {
}

type ShieldedInstanceConfigParameters struct {

	// - Compare the most recent boot measurements to the integrity policy baseline and return a pair of pass/fail results depending on whether they match or not. Defaults to true.
	// Note: allow_stopping_for_update must be set to true or your instance must have a desired_status of TERMINATED in order to update this field.
	// +kubebuilder:validation:Optional
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty" tf:"enable_integrity_monitoring,omitempty"`

	// - Verify the digital signature of all boot components, and halt the boot process if signature verification fails. Defaults to false.
	// Note: allow_stopping_for_update must be set to true or your instance must have a desired_status of TERMINATED in order to update this field.
	// +kubebuilder:validation:Optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty" tf:"enable_secure_boot,omitempty"`

	// - Use a virtualized trusted platform module, which is a specialized computer chip you can use to encrypt objects like keys and certificates. Defaults to true.
	// Note: allow_stopping_for_update must be set to true or your instance must have a desired_status of TERMINATED in order to update this field.
	// +kubebuilder:validation:Optional
	EnableVtpm *bool `json:"enableVtpm,omitempty" tf:"enable_vtpm,omitempty"`
}

type SpecificReservationObservation struct {
}

type SpecificReservationParameters struct {

	// Corresponds to the label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify compute.googleapis.com/reservation-name as the key and specify the name of your reservation as the only value.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Corresponds to the label values of a reservation resource.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. Manages a VM instance resource within GCE.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}