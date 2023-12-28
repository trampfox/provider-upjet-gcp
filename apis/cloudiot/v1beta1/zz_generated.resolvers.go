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
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-gcp/apis/pubsub/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Device.
func (mg *Device) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Registry),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.RegistryRef,
		Selector:     mg.Spec.ForProvider.RegistrySelector,
		To: reference.To{
			List:    &RegistryList{},
			Managed: &Registry{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Registry")
	}
	mg.Spec.ForProvider.Registry = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RegistryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Registry.
func (mg *Registry) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.EventNotificationConfigs); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventNotificationConfigs[i3].PubsubTopicName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.EventNotificationConfigs[i3].PubsubTopicNameRef,
			Selector:     mg.Spec.ForProvider.EventNotificationConfigs[i3].PubsubTopicNameSelector,
			To: reference.To{
				List:    &v1beta1.TopicList{},
				Managed: &v1beta1.Topic{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EventNotificationConfigs[i3].PubsubTopicName")
		}
		mg.Spec.ForProvider.EventNotificationConfigs[i3].PubsubTopicName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EventNotificationConfigs[i3].PubsubTopicNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.EventNotificationConfigs); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventNotificationConfigs[i3].PubsubTopicName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.EventNotificationConfigs[i3].PubsubTopicNameRef,
			Selector:     mg.Spec.InitProvider.EventNotificationConfigs[i3].PubsubTopicNameSelector,
			To: reference.To{
				List:    &v1beta1.TopicList{},
				Managed: &v1beta1.Topic{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EventNotificationConfigs[i3].PubsubTopicName")
		}
		mg.Spec.InitProvider.EventNotificationConfigs[i3].PubsubTopicName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EventNotificationConfigs[i3].PubsubTopicNameRef = rsp.ResolvedReference

	}

	return nil
}
