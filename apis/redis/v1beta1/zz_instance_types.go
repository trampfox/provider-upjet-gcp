// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type InstanceInitParameters struct {

	// Only applicable to STANDARD_HA tier which protects the instance
	// against zonal failures by provisioning it across two zones.
	// If provided, it must be a different zone from the one provided in
	// [locationId].
	AlternativeLocationID *string `json:"alternativeLocationId,omitempty" tf:"alternative_location_id,omitempty"`

	// Optional. Indicates whether OSS Redis AUTH is enabled for the
	// instance. If set to "true" AUTH is enabled on the instance.
	// Default value is "false" meaning AUTH is disabled.
	AuthEnabled *bool `json:"authEnabled,omitempty" tf:"auth_enabled,omitempty"`

	// The full name of the Google Compute Engine network to which the
	// instance is connected. If left unspecified, the default network
	// will be used.
	AuthorizedNetwork *string `json:"authorizedNetwork,omitempty" tf:"authorized_network,omitempty"`

	// The connection mode of the Redis instance.
	// Default value is DIRECT_PEERING.
	// Possible values are: DIRECT_PEERING, PRIVATE_SERVICE_ACCESS.
	ConnectMode *string `json:"connectMode,omitempty" tf:"connect_mode,omitempty"`

	// An arbitrary and optional user-provided name for the instance.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Resource labels to represent user provided metadata.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The zone where the instance will be provisioned. If not provided,
	// the service will choose a zone for the instance. For STANDARD_HA tier,
	// instances will be created across two zones for protection against
	// zonal failures. If [alternativeLocationId] is also provided, it must
	// be different from [locationId].
	LocationID *string `json:"locationId,omitempty" tf:"location_id,omitempty"`

	// Maintenance policy for an instance.
	// Structure is documented below.
	MaintenancePolicy []MaintenancePolicyInitParameters `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	// Redis memory size in GiB.
	MemorySizeGb *float64 `json:"memorySizeGb,omitempty" tf:"memory_size_gb,omitempty"`

	// Persistence configuration for an instance.
	// Structure is documented below.
	PersistenceConfig []PersistenceConfigInitParameters `json:"persistenceConfig,omitempty" tf:"persistence_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Optional. Read replica mode. Can only be specified when trying to create the instance.
	// If not set, Memorystore Redis backend will default to READ_REPLICAS_DISABLED.
	ReadReplicasMode *string `json:"readReplicasMode,omitempty" tf:"read_replicas_mode,omitempty"`

	// Redis configuration parameters, according to http://redis.io/topics/config.
	// Please check Memorystore documentation for the list of supported parameters:
	// https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs
	// +mapType=granular
	RedisConfigs map[string]*string `json:"redisConfigs,omitempty" tf:"redis_configs,omitempty"`

	// The version of Redis software. If not provided, latest supported
	// version will be used. Please check the API documentation linked
	// at the top for the latest valid values.
	RedisVersion *string `json:"redisVersion,omitempty" tf:"redis_version,omitempty"`

	// Optional. The number of replica nodes. The valid range for the Standard Tier with
	// read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled
	// for a Standard Tier instance, the only valid value is 1 and the default is 1.
	// The valid value for basic tier is 0 and the default is also 0.
	ReplicaCount *float64 `json:"replicaCount,omitempty" tf:"replica_count,omitempty"`

	// The CIDR range of internal addresses that are reserved for this
	// instance. If not provided, the service will choose an unused /29
	// block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
	// unique and non-overlapping with existing subnets in an authorized
	// network.
	ReservedIPRange *string `json:"reservedIpRange,omitempty" tf:"reserved_ip_range,omitempty"`

	// Optional. Additional IP range for node placement. Required when enabling read replicas on
	// an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or
	// "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address
	// range associated with the private service access connection, or "auto".
	SecondaryIPRange *string `json:"secondaryIpRange,omitempty" tf:"secondary_ip_range,omitempty"`

	// The service tier of the instance. Must be one of these values:
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// The TLS mode of the Redis instance, If not provided, TLS is disabled for the instance.
	TransitEncryptionMode *string `json:"transitEncryptionMode,omitempty" tf:"transit_encryption_mode,omitempty"`
}

type InstanceObservation struct {

	// Only applicable to STANDARD_HA tier which protects the instance
	// against zonal failures by provisioning it across two zones.
	// If provided, it must be a different zone from the one provided in
	// [locationId].
	AlternativeLocationID *string `json:"alternativeLocationId,omitempty" tf:"alternative_location_id,omitempty"`

	// Optional. Indicates whether OSS Redis AUTH is enabled for the
	// instance. If set to "true" AUTH is enabled on the instance.
	// Default value is "false" meaning AUTH is disabled.
	AuthEnabled *bool `json:"authEnabled,omitempty" tf:"auth_enabled,omitempty"`

	// The full name of the Google Compute Engine network to which the
	// instance is connected. If left unspecified, the default network
	// will be used.
	AuthorizedNetwork *string `json:"authorizedNetwork,omitempty" tf:"authorized_network,omitempty"`

	// The connection mode of the Redis instance.
	// Default value is DIRECT_PEERING.
	// Possible values are: DIRECT_PEERING, PRIVATE_SERVICE_ACCESS.
	ConnectMode *string `json:"connectMode,omitempty" tf:"connect_mode,omitempty"`

	// The time the instance was created in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The current zone where the Redis endpoint is placed.
	// For Basic Tier instances, this will always be the same as the
	// [locationId] provided by the user at creation time. For Standard Tier
	// instances, this can be either [locationId] or [alternativeLocationId]
	// and can change after a failover event.
	CurrentLocationID *string `json:"currentLocationId,omitempty" tf:"current_location_id,omitempty"`

	// Optional. The KMS key reference that you want to use to encrypt the data at rest for this Redis
	// instance. If this is provided, CMEK is enabled.
	CustomerManagedKey *string `json:"customerManagedKey,omitempty" tf:"customer_managed_key,omitempty"`

	// An arbitrary and optional user-provided name for the instance.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Hostname or IP address of the exposed Redis endpoint used by clients
	// to connect to the service.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{region}}/instances/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resource labels to represent user provided metadata.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The zone where the instance will be provisioned. If not provided,
	// the service will choose a zone for the instance. For STANDARD_HA tier,
	// instances will be created across two zones for protection against
	// zonal failures. If [alternativeLocationId] is also provided, it must
	// be different from [locationId].
	LocationID *string `json:"locationId,omitempty" tf:"location_id,omitempty"`

	// Maintenance policy for an instance.
	// Structure is documented below.
	MaintenancePolicy []MaintenancePolicyObservation `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	// Upcoming maintenance schedule.
	// Structure is documented below.
	MaintenanceSchedule []MaintenanceScheduleObservation `json:"maintenanceSchedule,omitempty" tf:"maintenance_schedule,omitempty"`

	// Redis memory size in GiB.
	MemorySizeGb *float64 `json:"memorySizeGb,omitempty" tf:"memory_size_gb,omitempty"`

	// Output only. Info per node.
	// Structure is documented below.
	Nodes []NodesObservation `json:"nodes,omitempty" tf:"nodes,omitempty"`

	// Persistence configuration for an instance.
	// Structure is documented below.
	PersistenceConfig []PersistenceConfigObservation `json:"persistenceConfig,omitempty" tf:"persistence_config,omitempty"`

	// Output only. Cloud IAM identity used by import / export operations
	// to transfer data to/from Cloud Storage. Format is "serviceAccount:".
	// The value may change over time for a given instance so should be
	// checked before each import/export operation.
	PersistenceIAMIdentity *string `json:"persistenceIamIdentity,omitempty" tf:"persistence_iam_identity,omitempty"`

	// The port number of the exposed Redis endpoint.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Output only. Hostname or IP address of the exposed readonly Redis endpoint. Standard tier only.
	// Targets all healthy replica nodes in instance. Replication is asynchronous and replica nodes
	// will exhibit some lag behind the primary. Write requests must target 'host'.
	ReadEndpoint *string `json:"readEndpoint,omitempty" tf:"read_endpoint,omitempty"`

	// Output only. The port number of the exposed readonly redis endpoint. Standard tier only.
	// Write requests should target 'port'.
	ReadEndpointPort *float64 `json:"readEndpointPort,omitempty" tf:"read_endpoint_port,omitempty"`

	// Optional. Read replica mode. Can only be specified when trying to create the instance.
	// If not set, Memorystore Redis backend will default to READ_REPLICAS_DISABLED.
	ReadReplicasMode *string `json:"readReplicasMode,omitempty" tf:"read_replicas_mode,omitempty"`

	// Redis configuration parameters, according to http://redis.io/topics/config.
	// Please check Memorystore documentation for the list of supported parameters:
	// https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs
	// +mapType=granular
	RedisConfigs map[string]*string `json:"redisConfigs,omitempty" tf:"redis_configs,omitempty"`

	// The version of Redis software. If not provided, latest supported
	// version will be used. Please check the API documentation linked
	// at the top for the latest valid values.
	RedisVersion *string `json:"redisVersion,omitempty" tf:"redis_version,omitempty"`

	// The name of the Redis region of the instance.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Optional. The number of replica nodes. The valid range for the Standard Tier with
	// read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled
	// for a Standard Tier instance, the only valid value is 1 and the default is 1.
	// The valid value for basic tier is 0 and the default is also 0.
	ReplicaCount *float64 `json:"replicaCount,omitempty" tf:"replica_count,omitempty"`

	// The CIDR range of internal addresses that are reserved for this
	// instance. If not provided, the service will choose an unused /29
	// block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
	// unique and non-overlapping with existing subnets in an authorized
	// network.
	ReservedIPRange *string `json:"reservedIpRange,omitempty" tf:"reserved_ip_range,omitempty"`

	// Optional. Additional IP range for node placement. Required when enabling read replicas on
	// an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or
	// "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address
	// range associated with the private service access connection, or "auto".
	SecondaryIPRange *string `json:"secondaryIpRange,omitempty" tf:"secondary_ip_range,omitempty"`

	// List of server CA certificates for the instance.
	// Structure is documented below.
	ServerCACerts []ServerCACertsObservation `json:"serverCaCerts,omitempty" tf:"server_ca_certs,omitempty"`

	// The service tier of the instance. Must be one of these values:
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// The TLS mode of the Redis instance, If not provided, TLS is disabled for the instance.
	TransitEncryptionMode *string `json:"transitEncryptionMode,omitempty" tf:"transit_encryption_mode,omitempty"`
}

type InstanceParameters struct {

	// Only applicable to STANDARD_HA tier which protects the instance
	// against zonal failures by provisioning it across two zones.
	// If provided, it must be a different zone from the one provided in
	// [locationId].
	// +kubebuilder:validation:Optional
	AlternativeLocationID *string `json:"alternativeLocationId,omitempty" tf:"alternative_location_id,omitempty"`

	// Optional. Indicates whether OSS Redis AUTH is enabled for the
	// instance. If set to "true" AUTH is enabled on the instance.
	// Default value is "false" meaning AUTH is disabled.
	// +kubebuilder:validation:Optional
	AuthEnabled *bool `json:"authEnabled,omitempty" tf:"auth_enabled,omitempty"`

	// The full name of the Google Compute Engine network to which the
	// instance is connected. If left unspecified, the default network
	// will be used.
	// +kubebuilder:validation:Optional
	AuthorizedNetwork *string `json:"authorizedNetwork,omitempty" tf:"authorized_network,omitempty"`

	// The connection mode of the Redis instance.
	// Default value is DIRECT_PEERING.
	// Possible values are: DIRECT_PEERING, PRIVATE_SERVICE_ACCESS.
	// +kubebuilder:validation:Optional
	ConnectMode *string `json:"connectMode,omitempty" tf:"connect_mode,omitempty"`

	// Optional. The KMS key reference that you want to use to encrypt the data at rest for this Redis
	// instance. If this is provided, CMEK is enabled.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CustomerManagedKey *string `json:"customerManagedKey,omitempty" tf:"customer_managed_key,omitempty"`

	// Reference to a CryptoKey in kms to populate customerManagedKey.
	// +kubebuilder:validation:Optional
	CustomerManagedKeyRef *v1.Reference `json:"customerManagedKeyRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate customerManagedKey.
	// +kubebuilder:validation:Optional
	CustomerManagedKeySelector *v1.Selector `json:"customerManagedKeySelector,omitempty" tf:"-"`

	// An arbitrary and optional user-provided name for the instance.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The zone where the instance will be provisioned. If not provided,
	// the service will choose a zone for the instance. For STANDARD_HA tier,
	// instances will be created across two zones for protection against
	// zonal failures. If [alternativeLocationId] is also provided, it must
	// be different from [locationId].
	// +kubebuilder:validation:Optional
	LocationID *string `json:"locationId,omitempty" tf:"location_id,omitempty"`

	// Maintenance policy for an instance.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MaintenancePolicy []MaintenancePolicyParameters `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	// Redis memory size in GiB.
	// +kubebuilder:validation:Optional
	MemorySizeGb *float64 `json:"memorySizeGb,omitempty" tf:"memory_size_gb,omitempty"`

	// Persistence configuration for an instance.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PersistenceConfig []PersistenceConfigParameters `json:"persistenceConfig,omitempty" tf:"persistence_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Optional. Read replica mode. Can only be specified when trying to create the instance.
	// If not set, Memorystore Redis backend will default to READ_REPLICAS_DISABLED.
	// +kubebuilder:validation:Optional
	ReadReplicasMode *string `json:"readReplicasMode,omitempty" tf:"read_replicas_mode,omitempty"`

	// Redis configuration parameters, according to http://redis.io/topics/config.
	// Please check Memorystore documentation for the list of supported parameters:
	// https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs
	// +kubebuilder:validation:Optional
	// +mapType=granular
	RedisConfigs map[string]*string `json:"redisConfigs,omitempty" tf:"redis_configs,omitempty"`

	// The version of Redis software. If not provided, latest supported
	// version will be used. Please check the API documentation linked
	// at the top for the latest valid values.
	// +kubebuilder:validation:Optional
	RedisVersion *string `json:"redisVersion,omitempty" tf:"redis_version,omitempty"`

	// The name of the Redis region of the instance.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// Optional. The number of replica nodes. The valid range for the Standard Tier with
	// read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled
	// for a Standard Tier instance, the only valid value is 1 and the default is 1.
	// The valid value for basic tier is 0 and the default is also 0.
	// +kubebuilder:validation:Optional
	ReplicaCount *float64 `json:"replicaCount,omitempty" tf:"replica_count,omitempty"`

	// The CIDR range of internal addresses that are reserved for this
	// instance. If not provided, the service will choose an unused /29
	// block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
	// unique and non-overlapping with existing subnets in an authorized
	// network.
	// +kubebuilder:validation:Optional
	ReservedIPRange *string `json:"reservedIpRange,omitempty" tf:"reserved_ip_range,omitempty"`

	// Optional. Additional IP range for node placement. Required when enabling read replicas on
	// an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or
	// "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address
	// range associated with the private service access connection, or "auto".
	// +kubebuilder:validation:Optional
	SecondaryIPRange *string `json:"secondaryIpRange,omitempty" tf:"secondary_ip_range,omitempty"`

	// The service tier of the instance. Must be one of these values:
	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// The TLS mode of the Redis instance, If not provided, TLS is disabled for the instance.
	// +kubebuilder:validation:Optional
	TransitEncryptionMode *string `json:"transitEncryptionMode,omitempty" tf:"transit_encryption_mode,omitempty"`
}

type MaintenancePolicyInitParameters struct {

	// Optional. Description of what this policy is for.
	// Create/Update methods return INVALID_ARGUMENT if the
	// length is greater than 512.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Optional. Maintenance window that is applied to resources covered by this policy.
	// Minimum 1. For the current version, the maximum number
	// of weekly_window is expected to be one.
	// Structure is documented below.
	WeeklyMaintenanceWindow []WeeklyMaintenanceWindowInitParameters `json:"weeklyMaintenanceWindow,omitempty" tf:"weekly_maintenance_window,omitempty"`
}

type MaintenancePolicyObservation struct {

	// (Output)
	// Output only. The time when the policy was created.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Optional. Description of what this policy is for.
	// Create/Update methods return INVALID_ARGUMENT if the
	// length is greater than 512.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Output)
	// Output only. The time when the policy was last updated.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`

	// Optional. Maintenance window that is applied to resources covered by this policy.
	// Minimum 1. For the current version, the maximum number
	// of weekly_window is expected to be one.
	// Structure is documented below.
	WeeklyMaintenanceWindow []WeeklyMaintenanceWindowObservation `json:"weeklyMaintenanceWindow,omitempty" tf:"weekly_maintenance_window,omitempty"`
}

type MaintenancePolicyParameters struct {

	// Optional. Description of what this policy is for.
	// Create/Update methods return INVALID_ARGUMENT if the
	// length is greater than 512.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Optional. Maintenance window that is applied to resources covered by this policy.
	// Minimum 1. For the current version, the maximum number
	// of weekly_window is expected to be one.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	WeeklyMaintenanceWindow []WeeklyMaintenanceWindowParameters `json:"weeklyMaintenanceWindow,omitempty" tf:"weekly_maintenance_window,omitempty"`
}

type MaintenanceScheduleInitParameters struct {
}

type MaintenanceScheduleObservation struct {

	// (Output)
	// Output only. The end time of any upcoming scheduled maintenance for this instance.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// (Output)
	// Output only. The deadline that the maintenance schedule start time
	// can not go beyond, including reschedule.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits.
	ScheduleDeadlineTime *string `json:"scheduleDeadlineTime,omitempty" tf:"schedule_deadline_time,omitempty"`

	// (Output)
	// Output only. The start time of any upcoming scheduled maintenance for this instance.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type MaintenanceScheduleParameters struct {
}

type NodesInitParameters struct {
}

type NodesObservation struct {

	// (Output)
	// Node identifying string. e.g. 'node-0', 'node-1'
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Output)
	// Location of the node.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type NodesParameters struct {
}

type PersistenceConfigInitParameters struct {

	// Optional. Controls whether Persistence features are enabled. If not provided, the existing value will be used.
	PersistenceMode *string `json:"persistenceMode,omitempty" tf:"persistence_mode,omitempty"`

	// Optional. Available snapshot periods for scheduling.
	RdbSnapshotPeriod *string `json:"rdbSnapshotPeriod,omitempty" tf:"rdb_snapshot_period,omitempty"`

	// Optional. Date and time that the first snapshot was/will be attempted,
	// and to which future snapshots will be aligned. If not provided,
	// the current time will be used.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution
	// and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	RdbSnapshotStartTime *string `json:"rdbSnapshotStartTime,omitempty" tf:"rdb_snapshot_start_time,omitempty"`
}

type PersistenceConfigObservation struct {

	// Optional. Controls whether Persistence features are enabled. If not provided, the existing value will be used.
	PersistenceMode *string `json:"persistenceMode,omitempty" tf:"persistence_mode,omitempty"`

	// (Output)
	// Output only. The next time that a snapshot attempt is scheduled to occur.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	RdbNextSnapshotTime *string `json:"rdbNextSnapshotTime,omitempty" tf:"rdb_next_snapshot_time,omitempty"`

	// Optional. Available snapshot periods for scheduling.
	RdbSnapshotPeriod *string `json:"rdbSnapshotPeriod,omitempty" tf:"rdb_snapshot_period,omitempty"`

	// Optional. Date and time that the first snapshot was/will be attempted,
	// and to which future snapshots will be aligned. If not provided,
	// the current time will be used.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution
	// and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	RdbSnapshotStartTime *string `json:"rdbSnapshotStartTime,omitempty" tf:"rdb_snapshot_start_time,omitempty"`
}

type PersistenceConfigParameters struct {

	// Optional. Controls whether Persistence features are enabled. If not provided, the existing value will be used.
	// +kubebuilder:validation:Optional
	PersistenceMode *string `json:"persistenceMode,omitempty" tf:"persistence_mode,omitempty"`

	// Optional. Available snapshot periods for scheduling.
	// +kubebuilder:validation:Optional
	RdbSnapshotPeriod *string `json:"rdbSnapshotPeriod,omitempty" tf:"rdb_snapshot_period,omitempty"`

	// Optional. Date and time that the first snapshot was/will be attempted,
	// and to which future snapshots will be aligned. If not provided,
	// the current time will be used.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution
	// and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	// +kubebuilder:validation:Optional
	RdbSnapshotStartTime *string `json:"rdbSnapshotStartTime,omitempty" tf:"rdb_snapshot_start_time,omitempty"`
}

type ServerCACertsInitParameters struct {
}

type ServerCACertsObservation struct {

	// (Output)
	// The certificate data in PEM format.
	Cert *string `json:"cert,omitempty" tf:"cert,omitempty"`

	// (Output)
	// The time when the certificate was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// (Output)
	// The time when the certificate expires.
	ExpireTime *string `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	// (Output)
	// Serial number, as extracted from the certificate.
	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number,omitempty"`

	// (Output)
	// Sha1 Fingerprint of the certificate.
	Sha1Fingerprint *string `json:"sha1Fingerprint,omitempty" tf:"sha1_fingerprint,omitempty"`
}

type ServerCACertsParameters struct {
}

type StartTimeInitParameters struct {

	// Hours of day in 24 hour format. Should be from 0 to 23.
	// An API may choose to allow the value "24:00:00" for scenarios like business closing time.
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// Minutes of hour of day. Must be from 0 to 59.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	Nanos *float64 `json:"nanos,omitempty" tf:"nanos,omitempty"`

	// Seconds of minutes of the time. Must normally be from 0 to 59.
	// An API may allow the value 60 if it allows leap-seconds.
	Seconds *float64 `json:"seconds,omitempty" tf:"seconds,omitempty"`
}

type StartTimeObservation struct {

	// Hours of day in 24 hour format. Should be from 0 to 23.
	// An API may choose to allow the value "24:00:00" for scenarios like business closing time.
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// Minutes of hour of day. Must be from 0 to 59.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	Nanos *float64 `json:"nanos,omitempty" tf:"nanos,omitempty"`

	// Seconds of minutes of the time. Must normally be from 0 to 59.
	// An API may allow the value 60 if it allows leap-seconds.
	Seconds *float64 `json:"seconds,omitempty" tf:"seconds,omitempty"`
}

type StartTimeParameters struct {

	// Hours of day in 24 hour format. Should be from 0 to 23.
	// An API may choose to allow the value "24:00:00" for scenarios like business closing time.
	// +kubebuilder:validation:Optional
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// Minutes of hour of day. Must be from 0 to 59.
	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	// +kubebuilder:validation:Optional
	Nanos *float64 `json:"nanos,omitempty" tf:"nanos,omitempty"`

	// Seconds of minutes of the time. Must normally be from 0 to 59.
	// An API may allow the value 60 if it allows leap-seconds.
	// +kubebuilder:validation:Optional
	Seconds *float64 `json:"seconds,omitempty" tf:"seconds,omitempty"`
}

type WeeklyMaintenanceWindowInitParameters struct {

	// Required. The day of week that maintenance updates occur.
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// Required. Start time of the window in UTC time.
	// Structure is documented below.
	StartTime []StartTimeInitParameters `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type WeeklyMaintenanceWindowObservation struct {

	// Required. The day of week that maintenance updates occur.
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// (Output)
	// Output only. Duration of the maintenance window.
	// The current window is fixed at 1 hour.
	// A duration in seconds with up to nine fractional digits,
	// terminated by 's'. Example: "3.5s".
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Required. Start time of the window in UTC time.
	// Structure is documented below.
	StartTime []StartTimeObservation `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type WeeklyMaintenanceWindowParameters struct {

	// Required. The day of week that maintenance updates occur.
	// +kubebuilder:validation:Optional
	Day *string `json:"day" tf:"day,omitempty"`

	// Required. Start time of the window in UTC time.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	StartTime []StartTimeParameters `json:"startTime" tf:"start_time,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider InstanceInitParameters `json:"initProvider,omitempty"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. A Google Cloud Redis instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.memorySizeGb) || (has(self.initProvider) && has(self.initProvider.memorySizeGb))",message="spec.forProvider.memorySizeGb is a required parameter"
	Spec   InstanceSpec   `json:"spec"`
	Status InstanceStatus `json:"status,omitempty"`
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
