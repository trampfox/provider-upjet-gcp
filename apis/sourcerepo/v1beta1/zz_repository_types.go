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

type PubsubConfigsInitParameters struct {

	// The format of the Cloud Pub/Sub messages.
	MessageFormat *string `json:"messageFormat,omitempty" tf:"message_format,omitempty"`

	// Email address of the service account used for publishing Cloud Pub/Sub messages.
	// This service account needs to be in the same project as the PubsubConfig. When added,
	// the caller needs to have iam.serviceAccounts.actAs permission on this service account.
	// If unspecified, it defaults to the compute engine default service account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("email",true)
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// Reference to a ServiceAccount in cloudplatform to populate serviceAccountEmail.
	// +kubebuilder:validation:Optional
	ServiceAccountEmailRef *v1.Reference `json:"serviceAccountEmailRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount in cloudplatform to populate serviceAccountEmail.
	// +kubebuilder:validation:Optional
	ServiceAccountEmailSelector *v1.Selector `json:"serviceAccountEmailSelector,omitempty" tf:"-"`

	// The identifier for this object. Format specified above.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/pubsub/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`

	// Reference to a Topic in pubsub to populate topic.
	// +kubebuilder:validation:Optional
	TopicRef *v1.Reference `json:"topicRef,omitempty" tf:"-"`

	// Selector for a Topic in pubsub to populate topic.
	// +kubebuilder:validation:Optional
	TopicSelector *v1.Selector `json:"topicSelector,omitempty" tf:"-"`
}

type PubsubConfigsObservation struct {

	// The format of the Cloud Pub/Sub messages.
	MessageFormat *string `json:"messageFormat,omitempty" tf:"message_format,omitempty"`

	// Email address of the service account used for publishing Cloud Pub/Sub messages.
	// This service account needs to be in the same project as the PubsubConfig. When added,
	// the caller needs to have iam.serviceAccounts.actAs permission on this service account.
	// If unspecified, it defaults to the compute engine default service account.
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// The identifier for this object. Format specified above.
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type PubsubConfigsParameters struct {

	// The format of the Cloud Pub/Sub messages.
	// +kubebuilder:validation:Optional
	MessageFormat *string `json:"messageFormat" tf:"message_format,omitempty"`

	// Email address of the service account used for publishing Cloud Pub/Sub messages.
	// This service account needs to be in the same project as the PubsubConfig. When added,
	// the caller needs to have iam.serviceAccounts.actAs permission on this service account.
	// If unspecified, it defaults to the compute engine default service account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("email",true)
	// +kubebuilder:validation:Optional
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// Reference to a ServiceAccount in cloudplatform to populate serviceAccountEmail.
	// +kubebuilder:validation:Optional
	ServiceAccountEmailRef *v1.Reference `json:"serviceAccountEmailRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount in cloudplatform to populate serviceAccountEmail.
	// +kubebuilder:validation:Optional
	ServiceAccountEmailSelector *v1.Selector `json:"serviceAccountEmailSelector,omitempty" tf:"-"`

	// The identifier for this object. Format specified above.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/pubsub/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`

	// Reference to a Topic in pubsub to populate topic.
	// +kubebuilder:validation:Optional
	TopicRef *v1.Reference `json:"topicRef,omitempty" tf:"-"`

	// Selector for a Topic in pubsub to populate topic.
	// +kubebuilder:validation:Optional
	TopicSelector *v1.Selector `json:"topicSelector,omitempty" tf:"-"`
}

type RepositoryInitParameters struct {

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// How this repository publishes a change in the repository through Cloud Pub/Sub.
	// Keyed by the topic names.
	// Structure is documented below.
	PubsubConfigs []PubsubConfigsInitParameters `json:"pubsubConfigs,omitempty" tf:"pubsub_configs,omitempty"`
}

type RepositoryObservation struct {

	// an identifier for the resource with format projects/{{project}}/repos/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// How this repository publishes a change in the repository through Cloud Pub/Sub.
	// Keyed by the topic names.
	// Structure is documented below.
	PubsubConfigs []PubsubConfigsObservation `json:"pubsubConfigs,omitempty" tf:"pubsub_configs,omitempty"`

	// The disk usage of the repo, in bytes.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// URL to clone the repository from Google Cloud Source Repositories.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type RepositoryParameters struct {

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// How this repository publishes a change in the repository through Cloud Pub/Sub.
	// Keyed by the topic names.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PubsubConfigs []PubsubConfigsParameters `json:"pubsubConfigs,omitempty" tf:"pubsub_configs,omitempty"`
}

// RepositorySpec defines the desired state of Repository
type RepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryParameters `json:"forProvider"`
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
	InitProvider RepositoryInitParameters `json:"initProvider,omitempty"`
}

// RepositoryStatus defines the observed state of Repository.
type RepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Repository is the Schema for the Repositorys API. A repository (or repo) is a Git repository storing versioned source content.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RepositorySpec   `json:"spec"`
	Status            RepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryList contains a list of Repositorys
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}

// Repository type metadata.
var (
	Repository_Kind             = "Repository"
	Repository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Repository_Kind}.String()
	Repository_KindAPIVersion   = Repository_Kind + "." + CRDGroupVersion.String()
	Repository_GroupVersionKind = CRDGroupVersion.WithKind(Repository_Kind)
)

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}
