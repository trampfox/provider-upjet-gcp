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

type ObjectACLInitParameters struct {

	// The name of the bucket the object is stored in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.Bucket
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in storage to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in storage to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The name of the object to apply the acl to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.BucketObject
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("output_name",true)
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// Reference to a BucketObject in storage to populate object.
	// +kubebuilder:validation:Optional
	ObjectRef *v1.Reference `json:"objectRef,omitempty" tf:"-"`

	// Selector for a BucketObject in storage to populate object.
	// +kubebuilder:validation:Optional
	ObjectSelector *v1.Selector `json:"objectSelector,omitempty" tf:"-"`

	// The "canned" predefined ACL to apply. Must be set if role_entity is not.
	PredefinedACL *string `json:"predefinedAcl,omitempty" tf:"predefined_acl,omitempty"`

	// List of role/entity pairs in the form ROLE:entity. See GCS Object ACL documentation for more details.
	// Must be set if predefined_acl is not.
	// +listType=set
	RoleEntity []*string `json:"roleEntity,omitempty" tf:"role_entity,omitempty"`
}

type ObjectACLObservation struct {

	// The name of the bucket the object is stored in.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the object to apply the acl to.
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// The "canned" predefined ACL to apply. Must be set if role_entity is not.
	PredefinedACL *string `json:"predefinedAcl,omitempty" tf:"predefined_acl,omitempty"`

	// List of role/entity pairs in the form ROLE:entity. See GCS Object ACL documentation for more details.
	// Must be set if predefined_acl is not.
	// +listType=set
	RoleEntity []*string `json:"roleEntity,omitempty" tf:"role_entity,omitempty"`
}

type ObjectACLParameters struct {

	// The name of the bucket the object is stored in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in storage to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in storage to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The name of the object to apply the acl to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.BucketObject
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("output_name",true)
	// +kubebuilder:validation:Optional
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// Reference to a BucketObject in storage to populate object.
	// +kubebuilder:validation:Optional
	ObjectRef *v1.Reference `json:"objectRef,omitempty" tf:"-"`

	// Selector for a BucketObject in storage to populate object.
	// +kubebuilder:validation:Optional
	ObjectSelector *v1.Selector `json:"objectSelector,omitempty" tf:"-"`

	// The "canned" predefined ACL to apply. Must be set if role_entity is not.
	// +kubebuilder:validation:Optional
	PredefinedACL *string `json:"predefinedAcl,omitempty" tf:"predefined_acl,omitempty"`

	// List of role/entity pairs in the form ROLE:entity. See GCS Object ACL documentation for more details.
	// Must be set if predefined_acl is not.
	// +kubebuilder:validation:Optional
	// +listType=set
	RoleEntity []*string `json:"roleEntity,omitempty" tf:"role_entity,omitempty"`
}

// ObjectACLSpec defines the desired state of ObjectACL
type ObjectACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectACLParameters `json:"forProvider"`
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
	InitProvider ObjectACLInitParameters `json:"initProvider,omitempty"`
}

// ObjectACLStatus defines the observed state of ObjectACL.
type ObjectACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectACL is the Schema for the ObjectACLs API. Creates a new object ACL in Google Cloud Storage.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ObjectACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ObjectACLSpec   `json:"spec"`
	Status            ObjectACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectACLList contains a list of ObjectACLs
type ObjectACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectACL `json:"items"`
}

// Repository type metadata.
var (
	ObjectACL_Kind             = "ObjectACL"
	ObjectACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ObjectACL_Kind}.String()
	ObjectACL_KindAPIVersion   = ObjectACL_Kind + "." + CRDGroupVersion.String()
	ObjectACL_GroupVersionKind = CRDGroupVersion.WithKind(ObjectACL_Kind)
)

func init() {
	SchemeBuilder.Register(&ObjectACL{}, &ObjectACLList{})
}
