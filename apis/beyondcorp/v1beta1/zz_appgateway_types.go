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

type AllocatedConnectionsInitParameters struct {
}

type AllocatedConnectionsObservation struct {

	// The ingress port of an allocated connection.
	IngressPort *float64 `json:"ingressPort,omitempty" tf:"ingress_port,omitempty"`

	// The PSC uri of an allocated connection.
	PscURI *string `json:"pscUri,omitempty" tf:"psc_uri,omitempty"`
}

type AllocatedConnectionsParameters struct {
}

type AppGatewayInitParameters struct {

	// An arbitrary user-provided name for the AppGateway.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The type of hosting used by the AppGateway.
	// Default value is HOST_TYPE_UNSPECIFIED.
	// Possible values are: HOST_TYPE_UNSPECIFIED, GCP_REGIONAL_MIG.
	HostType *string `json:"hostType,omitempty" tf:"host_type,omitempty"`

	// Resource labels to represent user provided metadata.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The type of network connectivity used by the AppGateway.
	// Default value is TYPE_UNSPECIFIED.
	// Possible values are: TYPE_UNSPECIFIED, TCP_PROXY.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AppGatewayObservation struct {

	// A list of connections allocated for the Gateway.
	// Structure is documented below.
	AllocatedConnections []AllocatedConnectionsObservation `json:"allocatedConnections,omitempty" tf:"allocated_connections,omitempty"`

	// An arbitrary user-provided name for the AppGateway.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The type of hosting used by the AppGateway.
	// Default value is HOST_TYPE_UNSPECIFIED.
	// Possible values are: HOST_TYPE_UNSPECIFIED, GCP_REGIONAL_MIG.
	HostType *string `json:"hostType,omitempty" tf:"host_type,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{region}}/appGateways/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resource labels to represent user provided metadata.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the AppGateway.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Represents the different states of a AppGateway.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The type of network connectivity used by the AppGateway.
	// Default value is TYPE_UNSPECIFIED.
	// Possible values are: TYPE_UNSPECIFIED, TCP_PROXY.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Server-defined URI for this resource.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type AppGatewayParameters struct {

	// An arbitrary user-provided name for the AppGateway.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The type of hosting used by the AppGateway.
	// Default value is HOST_TYPE_UNSPECIFIED.
	// Possible values are: HOST_TYPE_UNSPECIFIED, GCP_REGIONAL_MIG.
	// +kubebuilder:validation:Optional
	HostType *string `json:"hostType,omitempty" tf:"host_type,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the AppGateway.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The type of network connectivity used by the AppGateway.
	// Default value is TYPE_UNSPECIFIED.
	// Possible values are: TYPE_UNSPECIFIED, TCP_PROXY.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// AppGatewaySpec defines the desired state of AppGateway
type AppGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppGatewayParameters `json:"forProvider"`
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
	InitProvider AppGatewayInitParameters `json:"initProvider,omitempty"`
}

// AppGatewayStatus defines the observed state of AppGateway.
type AppGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppGateway is the Schema for the AppGateways API. A BeyondCorp AppGateway resource represents a BeyondCorp protected AppGateway to a remote application.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AppGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppGatewaySpec   `json:"spec"`
	Status            AppGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppGatewayList contains a list of AppGateways
type AppGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppGateway `json:"items"`
}

// Repository type metadata.
var (
	AppGateway_Kind             = "AppGateway"
	AppGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppGateway_Kind}.String()
	AppGateway_KindAPIVersion   = AppGateway_Kind + "." + CRDGroupVersion.String()
	AppGateway_GroupVersionKind = CRDGroupVersion.WithKind(AppGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&AppGateway{}, &AppGatewayList{})
}
