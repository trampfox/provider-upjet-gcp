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

type BackendServiceSignedURLKeyInitParameters struct {

	// The backend service this signed URL key belongs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.BackendService
	BackendService *string `json:"backendService,omitempty" tf:"backend_service,omitempty"`

	// Reference to a BackendService in compute to populate backendService.
	// +kubebuilder:validation:Optional
	BackendServiceRef *v1.Reference `json:"backendServiceRef,omitempty" tf:"-"`

	// Selector for a BackendService in compute to populate backendService.
	// +kubebuilder:validation:Optional
	BackendServiceSelector *v1.Selector `json:"backendServiceSelector,omitempty" tf:"-"`

	// Name of the signed URL key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type BackendServiceSignedURLKeyObservation struct {

	// The backend service this signed URL key belongs.
	BackendService *string `json:"backendService,omitempty" tf:"backend_service,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/backendServices/{{backend_service}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the signed URL key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type BackendServiceSignedURLKeyParameters struct {

	// The backend service this signed URL key belongs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.BackendService
	// +kubebuilder:validation:Optional
	BackendService *string `json:"backendService,omitempty" tf:"backend_service,omitempty"`

	// Reference to a BackendService in compute to populate backendService.
	// +kubebuilder:validation:Optional
	BackendServiceRef *v1.Reference `json:"backendServiceRef,omitempty" tf:"-"`

	// Selector for a BackendService in compute to populate backendService.
	// +kubebuilder:validation:Optional
	BackendServiceSelector *v1.Selector `json:"backendServiceSelector,omitempty" tf:"-"`

	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	KeyValueSecretRef v1.SecretKeySelector `json:"keyValueSecretRef" tf:"-"`

	// Name of the signed URL key.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// BackendServiceSignedURLKeySpec defines the desired state of BackendServiceSignedURLKey
type BackendServiceSignedURLKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackendServiceSignedURLKeyParameters `json:"forProvider"`
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
	InitProvider BackendServiceSignedURLKeyInitParameters `json:"initProvider,omitempty"`
}

// BackendServiceSignedURLKeyStatus defines the observed state of BackendServiceSignedURLKey.
type BackendServiceSignedURLKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackendServiceSignedURLKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BackendServiceSignedURLKey is the Schema for the BackendServiceSignedURLKeys API. A key for signing Cloud CDN signed URLs for Backend Services.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type BackendServiceSignedURLKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.keyValueSecretRef)",message="spec.forProvider.keyValueSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   BackendServiceSignedURLKeySpec   `json:"spec"`
	Status BackendServiceSignedURLKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackendServiceSignedURLKeyList contains a list of BackendServiceSignedURLKeys
type BackendServiceSignedURLKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackendServiceSignedURLKey `json:"items"`
}

// Repository type metadata.
var (
	BackendServiceSignedURLKey_Kind             = "BackendServiceSignedURLKey"
	BackendServiceSignedURLKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackendServiceSignedURLKey_Kind}.String()
	BackendServiceSignedURLKey_KindAPIVersion   = BackendServiceSignedURLKey_Kind + "." + CRDGroupVersion.String()
	BackendServiceSignedURLKey_GroupVersionKind = CRDGroupVersion.WithKind(BackendServiceSignedURLKey_Kind)
)

func init() {
	SchemeBuilder.Register(&BackendServiceSignedURLKey{}, &BackendServiceSignedURLKeyList{})
}
