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

type NotificationChannelInitParameters struct {

	// An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// If true, the notification channel will be deleted regardless
	// of its use in alert policies (the policies will be updated
	// to remove the channel). If false, channels that are still
	// referenced by an existing alerting policy will fail to be
	// deleted in a delete operation.
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// Configuration fields that define the channel and its behavior. The
	// permissible and required labels are specified in the
	// NotificationChannelDescriptor corresponding to the type field. They can also be configured via
	// the sensitive_labels block, but cannot be configured in both places.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Different notification type behaviors are configured primarily using the the labels field on this
	// resource. This block contains the labels which contain secrets or passwords so that they can be marked
	// sensitive and hidden from plan output. The name of the field, eg: password, will be the key
	// in the labels map in the api request.
	// Credentials may not be specified in both locations and will cause an error. Changing from one location
	// to a different credential configuration in the config will require an apply to update state.
	// Structure is documented below.
	SensitiveLabels []SensitiveLabelsInitParameters `json:"sensitiveLabels,omitempty" tf:"sensitive_labels,omitempty"`

	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels map[string]*string `json:"userLabels,omitempty" tf:"user_labels,omitempty"`
}

type NotificationChannelObservation struct {

	// An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// If true, the notification channel will be deleted regardless
	// of its use in alert policies (the policies will be updated
	// to remove the channel). If false, channels that are still
	// referenced by an existing alerting policy will fail to be
	// deleted in a delete operation.
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration fields that define the channel and its behavior. The
	// permissible and required labels are specified in the
	// NotificationChannelDescriptor corresponding to the type field. They can also be configured via
	// the sensitive_labels block, but cannot be configured in both places.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The full REST resource name for this channel. The syntax is:
	// projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]
	// The [CHANNEL_ID] is automatically assigned by the server on creation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Different notification type behaviors are configured primarily using the the labels field on this
	// resource. This block contains the labels which contain secrets or passwords so that they can be marked
	// sensitive and hidden from plan output. The name of the field, eg: password, will be the key
	// in the labels map in the api request.
	// Credentials may not be specified in both locations and will cause an error. Changing from one location
	// to a different credential configuration in the config will require an apply to update state.
	// Structure is documented below.
	SensitiveLabels []SensitiveLabelsParameters `json:"sensitiveLabels,omitempty" tf:"sensitive_labels,omitempty"`

	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels map[string]*string `json:"userLabels,omitempty" tf:"user_labels,omitempty"`

	// Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require verification or that this specific channel has been exempted from verification because it was created prior to verification being required for channels of this type.This field cannot be modified using a standard UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
	VerificationStatus *string `json:"verificationStatus,omitempty" tf:"verification_status,omitempty"`
}

type NotificationChannelParameters struct {

	// An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// If true, the notification channel will be deleted regardless
	// of its use in alert policies (the policies will be updated
	// to remove the channel). If false, channels that are still
	// referenced by an existing alerting policy will fail to be
	// deleted in a delete operation.
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// Configuration fields that define the channel and its behavior. The
	// permissible and required labels are specified in the
	// NotificationChannelDescriptor corresponding to the type field. They can also be configured via
	// the sensitive_labels block, but cannot be configured in both places.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Different notification type behaviors are configured primarily using the the labels field on this
	// resource. This block contains the labels which contain secrets or passwords so that they can be marked
	// sensitive and hidden from plan output. The name of the field, eg: password, will be the key
	// in the labels map in the api request.
	// Credentials may not be specified in both locations and will cause an error. Changing from one location
	// to a different credential configuration in the config will require an apply to update state.
	// Structure is documented below.
	SensitiveLabels []SensitiveLabelsParameters `json:"sensitiveLabels,omitempty" tf:"sensitive_labels,omitempty"`

	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels map[string]*string `json:"userLabels,omitempty" tf:"user_labels,omitempty"`
}

type SensitiveLabelsInitParameters struct {
}

type SensitiveLabelsObservation struct {
}

type SensitiveLabelsParameters struct {

	// An authorization token for a notification channel. Channel types that support this field include: slack
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	AuthTokenSecretRef *v1.SecretKeySelector `json:"authTokenSecretRef,omitempty" tf:"-"`

	// An password for a notification channel. Channel types that support this field include: webhook_basicauth
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// An servicekey token for a notification channel. Channel types that support this field include: pagerduty
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	ServiceKeySecretRef *v1.SecretKeySelector `json:"serviceKeySecretRef,omitempty" tf:"-"`
}

// NotificationChannelSpec defines the desired state of NotificationChannel
type NotificationChannelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationChannelParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider NotificationChannelInitParameters `json:"initProvider,omitempty"`
}

// NotificationChannelStatus defines the observed state of NotificationChannel.
type NotificationChannelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationChannelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationChannel is the Schema for the NotificationChannels API. A NotificationChannel is a medium through which an alert is delivered when a policy violation is detected.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type NotificationChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || has(self.initProvider.type)",message="type is a required parameter"
	Spec   NotificationChannelSpec   `json:"spec"`
	Status NotificationChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationChannelList contains a list of NotificationChannels
type NotificationChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationChannel `json:"items"`
}

// Repository type metadata.
var (
	NotificationChannel_Kind             = "NotificationChannel"
	NotificationChannel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotificationChannel_Kind}.String()
	NotificationChannel_KindAPIVersion   = NotificationChannel_Kind + "." + CRDGroupVersion.String()
	NotificationChannel_GroupVersionKind = CRDGroupVersion.WithKind(NotificationChannel_Kind)
)

func init() {
	SchemeBuilder.Register(&NotificationChannel{}, &NotificationChannelList{})
}
