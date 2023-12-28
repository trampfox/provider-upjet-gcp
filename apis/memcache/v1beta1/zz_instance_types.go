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

	// The full name of the GCE network to connect the instance to.  If not provided,
	// 'default' will be used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/servicenetworking/v1beta1.Connection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("network",false)
	AuthorizedNetwork *string `json:"authorizedNetwork,omitempty" tf:"authorized_network,omitempty"`

	// Reference to a Connection in servicenetworking to populate authorizedNetwork.
	// +kubebuilder:validation:Optional
	AuthorizedNetworkRef *v1.Reference `json:"authorizedNetworkRef,omitempty" tf:"-"`

	// Selector for a Connection in servicenetworking to populate authorizedNetwork.
	// +kubebuilder:validation:Optional
	AuthorizedNetworkSelector *v1.Selector `json:"authorizedNetworkSelector,omitempty" tf:"-"`

	// A user-visible name for the instance.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Resource labels to represent user-provided metadata.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Maintenance policy for an instance.
	// Structure is documented below.
	MaintenancePolicy []MaintenancePolicyInitParameters `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	// User-specified parameters for this memcache instance.
	// Structure is documented below.
	MemcacheParameters []MemcacheParametersInitParameters `json:"memcacheParameters,omitempty" tf:"memcache_parameters,omitempty"`

	// The major version of Memcached software. If not provided, latest supported version will be used.
	// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
	// determined by our system based on the latest supported minor version.
	// Default value is MEMCACHE_1_5.
	// Possible values are: MEMCACHE_1_5.
	MemcacheVersion *string `json:"memcacheVersion,omitempty" tf:"memcache_version,omitempty"`

	// The resource name of the instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration for memcache nodes.
	// Structure is documented below.
	NodeConfig []NodeConfigInitParameters `json:"nodeConfig,omitempty" tf:"node_config,omitempty"`

	// Number of nodes in the memcache instance.
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the Memcache instance. If it is not provided, the provider region is used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Zones where memcache nodes should be provisioned.  If not
	// provided, all zones will be used.
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type InstanceObservation struct {

	// The full name of the GCE network to connect the instance to.  If not provided,
	// 'default' will be used.
	AuthorizedNetwork *string `json:"authorizedNetwork,omitempty" tf:"authorized_network,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Endpoint for Discovery API
	DiscoveryEndpoint *string `json:"discoveryEndpoint,omitempty" tf:"discovery_endpoint,omitempty"`

	// A user-visible name for the instance.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{region}}/instances/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resource labels to represent user-provided metadata.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Maintenance policy for an instance.
	// Structure is documented below.
	MaintenancePolicy []MaintenancePolicyObservation `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	// Output only. Published maintenance schedule.
	// Structure is documented below.
	MaintenanceSchedule []MaintenanceScheduleObservation `json:"maintenanceSchedule,omitempty" tf:"maintenance_schedule,omitempty"`

	// The full version of memcached server running on this instance.
	MemcacheFullVersion *string `json:"memcacheFullVersion,omitempty" tf:"memcache_full_version,omitempty"`

	// Additional information about the instance state, if available.
	// Structure is documented below.
	MemcacheNodes []MemcacheNodesObservation `json:"memcacheNodes,omitempty" tf:"memcache_nodes,omitempty"`

	// User-specified parameters for this memcache instance.
	// Structure is documented below.
	MemcacheParameters []MemcacheParametersObservation `json:"memcacheParameters,omitempty" tf:"memcache_parameters,omitempty"`

	// The major version of Memcached software. If not provided, latest supported version will be used.
	// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
	// determined by our system based on the latest supported minor version.
	// Default value is MEMCACHE_1_5.
	// Possible values are: MEMCACHE_1_5.
	MemcacheVersion *string `json:"memcacheVersion,omitempty" tf:"memcache_version,omitempty"`

	// The resource name of the instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration for memcache nodes.
	// Structure is documented below.
	NodeConfig []NodeConfigObservation `json:"nodeConfig,omitempty" tf:"node_config,omitempty"`

	// Number of nodes in the memcache instance.
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the Memcache instance. If it is not provided, the provider region is used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Zones where memcache nodes should be provisioned.  If not
	// provided, all zones will be used.
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type InstanceParameters struct {

	// The full name of the GCE network to connect the instance to.  If not provided,
	// 'default' will be used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/servicenetworking/v1beta1.Connection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("network",false)
	// +kubebuilder:validation:Optional
	AuthorizedNetwork *string `json:"authorizedNetwork,omitempty" tf:"authorized_network,omitempty"`

	// Reference to a Connection in servicenetworking to populate authorizedNetwork.
	// +kubebuilder:validation:Optional
	AuthorizedNetworkRef *v1.Reference `json:"authorizedNetworkRef,omitempty" tf:"-"`

	// Selector for a Connection in servicenetworking to populate authorizedNetwork.
	// +kubebuilder:validation:Optional
	AuthorizedNetworkSelector *v1.Selector `json:"authorizedNetworkSelector,omitempty" tf:"-"`

	// A user-visible name for the instance.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Resource labels to represent user-provided metadata.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Maintenance policy for an instance.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MaintenancePolicy []MaintenancePolicyParameters `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	// User-specified parameters for this memcache instance.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MemcacheParameters []MemcacheParametersParameters `json:"memcacheParameters,omitempty" tf:"memcache_parameters,omitempty"`

	// The major version of Memcached software. If not provided, latest supported version will be used.
	// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
	// determined by our system based on the latest supported minor version.
	// Default value is MEMCACHE_1_5.
	// Possible values are: MEMCACHE_1_5.
	// +kubebuilder:validation:Optional
	MemcacheVersion *string `json:"memcacheVersion,omitempty" tf:"memcache_version,omitempty"`

	// The resource name of the instance.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration for memcache nodes.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	NodeConfig []NodeConfigParameters `json:"nodeConfig,omitempty" tf:"node_config,omitempty"`

	// Number of nodes in the memcache instance.
	// +kubebuilder:validation:Optional
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the Memcache instance. If it is not provided, the provider region is used.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Zones where memcache nodes should be provisioned.  If not
	// provided, all zones will be used.
	// +kubebuilder:validation:Optional
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type MaintenancePolicyInitParameters struct {

	// Optional. Description of what this policy is for.
	// Create/Update methods return INVALID_ARGUMENT if the
	// length is greater than 512.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Required. Maintenance window that is applied to resources covered by this policy.
	// Minimum 1. For the current version, the maximum number of weekly_maintenance_windows
	// is expected to be one.
	// Structure is documented below.
	WeeklyMaintenanceWindow []WeeklyMaintenanceWindowInitParameters `json:"weeklyMaintenanceWindow,omitempty" tf:"weekly_maintenance_window,omitempty"`
}

type MaintenancePolicyObservation struct {

	// (Output)
	// Output only. The time when the policy was created.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Optional. Description of what this policy is for.
	// Create/Update methods return INVALID_ARGUMENT if the
	// length is greater than 512.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Output)
	// Output only. The time when the policy was updated.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`

	// Required. Maintenance window that is applied to resources covered by this policy.
	// Minimum 1. For the current version, the maximum number of weekly_maintenance_windows
	// is expected to be one.
	// Structure is documented below.
	WeeklyMaintenanceWindow []WeeklyMaintenanceWindowObservation `json:"weeklyMaintenanceWindow,omitempty" tf:"weekly_maintenance_window,omitempty"`
}

type MaintenancePolicyParameters struct {

	// Optional. Description of what this policy is for.
	// Create/Update methods return INVALID_ARGUMENT if the
	// length is greater than 512.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Required. Maintenance window that is applied to resources covered by this policy.
	// Minimum 1. For the current version, the maximum number of weekly_maintenance_windows
	// is expected to be one.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	WeeklyMaintenanceWindow []WeeklyMaintenanceWindowParameters `json:"weeklyMaintenanceWindow" tf:"weekly_maintenance_window,omitempty"`
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

type MemcacheNodesInitParameters struct {
}

type MemcacheNodesObservation struct {

	// (Output)
	// Hostname or IP address of the Memcached node used by the clients to connect to the Memcached server on this node.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// (Output)
	// Identifier of the Memcached node. The node id does not include project or location like the Memcached instance name.
	NodeID *string `json:"nodeId,omitempty" tf:"node_id,omitempty"`

	// (Output)
	// The port number of the Memcached server on this node.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (Output)
	// Current state of the Memcached node.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// (Output)
	// Location (GCP Zone) for the Memcached node.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type MemcacheNodesParameters struct {
}

type MemcacheParametersInitParameters struct {

	// User-defined set of parameters to use in the memcache process.
	// +mapType=granular
	Params map[string]*string `json:"params,omitempty" tf:"params,omitempty"`
}

type MemcacheParametersObservation struct {

	// (Output)
	// This is a unique ID associated with this set of parameters.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// User-defined set of parameters to use in the memcache process.
	// +mapType=granular
	Params map[string]*string `json:"params,omitempty" tf:"params,omitempty"`
}

type MemcacheParametersParameters struct {

	// User-defined set of parameters to use in the memcache process.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Params map[string]*string `json:"params,omitempty" tf:"params,omitempty"`
}

type NodeConfigInitParameters struct {

	// Number of CPUs per node.
	CPUCount *float64 `json:"cpuCount,omitempty" tf:"cpu_count,omitempty"`

	// Memory size in Mebibytes for each memcache node.
	MemorySizeMb *float64 `json:"memorySizeMb,omitempty" tf:"memory_size_mb,omitempty"`
}

type NodeConfigObservation struct {

	// Number of CPUs per node.
	CPUCount *float64 `json:"cpuCount,omitempty" tf:"cpu_count,omitempty"`

	// Memory size in Mebibytes for each memcache node.
	MemorySizeMb *float64 `json:"memorySizeMb,omitempty" tf:"memory_size_mb,omitempty"`
}

type NodeConfigParameters struct {

	// Number of CPUs per node.
	// +kubebuilder:validation:Optional
	CPUCount *float64 `json:"cpuCount" tf:"cpu_count,omitempty"`

	// Memory size in Mebibytes for each memcache node.
	// +kubebuilder:validation:Optional
	MemorySizeMb *float64 `json:"memorySizeMb" tf:"memory_size_mb,omitempty"`
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

	// Required. The length of the maintenance window, ranging from 3 hours to 8 hours.
	// A duration in seconds with up to nine fractional digits,
	// terminated by 's'. Example: "3.5s".
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Required. Start time of the window in UTC time.
	// Structure is documented below.
	StartTime []StartTimeInitParameters `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type WeeklyMaintenanceWindowObservation struct {

	// Required. The day of week that maintenance updates occur.
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// Required. The length of the maintenance window, ranging from 3 hours to 8 hours.
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

	// Required. The length of the maintenance window, ranging from 3 hours to 8 hours.
	// A duration in seconds with up to nine fractional digits,
	// terminated by 's'. Example: "3.5s".
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration" tf:"duration,omitempty"`

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

// Instance is the Schema for the Instances API. A Google Cloud Memcache instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nodeConfig) || (has(self.initProvider) && has(self.initProvider.nodeConfig))",message="spec.forProvider.nodeConfig is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nodeCount) || (has(self.initProvider) && has(self.initProvider.nodeCount))",message="spec.forProvider.nodeCount is a required parameter"
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
