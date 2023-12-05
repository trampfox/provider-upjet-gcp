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

type ProjectMetadataInitParameters struct {

	// A series of key value pairs.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ProjectMetadataObservation struct {

	// an identifier for the resource with format {{project}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A series of key value pairs.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ProjectMetadataParameters struct {

	// A series of key value pairs.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// ProjectMetadataSpec defines the desired state of ProjectMetadata
type ProjectMetadataSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectMetadataParameters `json:"forProvider"`
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
	InitProvider ProjectMetadataInitParameters `json:"initProvider,omitempty"`
}

// ProjectMetadataStatus defines the observed state of ProjectMetadata.
type ProjectMetadataStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectMetadataObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectMetadata is the Schema for the ProjectMetadatas API. Manages common instance metadata
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ProjectMetadata struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metadata) || (has(self.initProvider) && has(self.initProvider.metadata))",message="spec.forProvider.metadata is a required parameter"
	Spec   ProjectMetadataSpec   `json:"spec"`
	Status ProjectMetadataStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectMetadataList contains a list of ProjectMetadatas
type ProjectMetadataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectMetadata `json:"items"`
}

// Repository type metadata.
var (
	ProjectMetadata_Kind             = "ProjectMetadata"
	ProjectMetadata_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectMetadata_Kind}.String()
	ProjectMetadata_KindAPIVersion   = ProjectMetadata_Kind + "." + CRDGroupVersion.String()
	ProjectMetadata_GroupVersionKind = CRDGroupVersion.WithKind(ProjectMetadata_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectMetadata{}, &ProjectMetadataList{})
}
