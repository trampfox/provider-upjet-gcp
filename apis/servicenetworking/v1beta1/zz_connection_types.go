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

type ConnectionInitParameters struct {

	// Name of VPC network connected with service producers using VPC peering.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// Named IP address range(s) of PEERING type reserved for
	// this service provider. Note that invoking this method with a different range when connection
	// is already established will not reallocate already provisioned service producer subnetworks.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.GlobalAddress
	ReservedPeeringRanges []*string `json:"reservedPeeringRanges,omitempty" tf:"reserved_peering_ranges,omitempty"`

	// References to GlobalAddress in compute to populate reservedPeeringRanges.
	// +kubebuilder:validation:Optional
	ReservedPeeringRangesRefs []v1.Reference `json:"reservedPeeringRangesRefs,omitempty" tf:"-"`

	// Selector for a list of GlobalAddress in compute to populate reservedPeeringRanges.
	// +kubebuilder:validation:Optional
	ReservedPeeringRangesSelector *v1.Selector `json:"reservedPeeringRangesSelector,omitempty" tf:"-"`

	// Provider peering service that is managing peering connectivity for a
	// service provider organization. For Google services that support this functionality it is
	// 'servicenetworking.googleapis.com'.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type ConnectionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of VPC network connected with service producers using VPC peering.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// (Computed) The name of the VPC Network Peering connection that was created by the service producer.
	Peering *string `json:"peering,omitempty" tf:"peering,omitempty"`

	// Named IP address range(s) of PEERING type reserved for
	// this service provider. Note that invoking this method with a different range when connection
	// is already established will not reallocate already provisioned service producer subnetworks.
	ReservedPeeringRanges []*string `json:"reservedPeeringRanges,omitempty" tf:"reserved_peering_ranges,omitempty"`

	// Provider peering service that is managing peering connectivity for a
	// service provider organization. For Google services that support this functionality it is
	// 'servicenetworking.googleapis.com'.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type ConnectionParameters struct {

	// Name of VPC network connected with service producers using VPC peering.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// Named IP address range(s) of PEERING type reserved for
	// this service provider. Note that invoking this method with a different range when connection
	// is already established will not reallocate already provisioned service producer subnetworks.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.GlobalAddress
	// +kubebuilder:validation:Optional
	ReservedPeeringRanges []*string `json:"reservedPeeringRanges,omitempty" tf:"reserved_peering_ranges,omitempty"`

	// References to GlobalAddress in compute to populate reservedPeeringRanges.
	// +kubebuilder:validation:Optional
	ReservedPeeringRangesRefs []v1.Reference `json:"reservedPeeringRangesRefs,omitempty" tf:"-"`

	// Selector for a list of GlobalAddress in compute to populate reservedPeeringRanges.
	// +kubebuilder:validation:Optional
	ReservedPeeringRangesSelector *v1.Selector `json:"reservedPeeringRangesSelector,omitempty" tf:"-"`

	// Provider peering service that is managing peering connectivity for a
	// service provider organization. For Google services that support this functionality it is
	// 'servicenetworking.googleapis.com'.
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

// ConnectionSpec defines the desired state of Connection
type ConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionParameters `json:"forProvider"`
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
	InitProvider ConnectionInitParameters `json:"initProvider,omitempty"`
}

// ConnectionStatus defines the observed state of Connection.
type ConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Connection is the Schema for the Connections API. Manages creating a private VPC connection to a service provider.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Connection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.service) || (has(self.initProvider) && has(self.initProvider.service))",message="spec.forProvider.service is a required parameter"
	Spec   ConnectionSpec   `json:"spec"`
	Status ConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionList contains a list of Connections
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connection `json:"items"`
}

// Repository type metadata.
var (
	Connection_Kind             = "Connection"
	Connection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Connection_Kind}.String()
	Connection_KindAPIVersion   = Connection_Kind + "." + CRDGroupVersion.String()
	Connection_GroupVersionKind = CRDGroupVersion.WithKind(Connection_Kind)
)

func init() {
	SchemeBuilder.Register(&Connection{}, &ConnectionList{})
}
