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
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-gcp/apis/compute/v1beta1"
	v1beta1 "github.com/upbound/provider-gcp/apis/kms/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Envgroup.
func (mg *Envgroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.OrgIDRef,
		Selector:     mg.Spec.ForProvider.OrgIDSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrgIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Environment.
func (mg *Environment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.OrgIDRef,
		Selector:     mg.Spec.ForProvider.OrgIDSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrgIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Instance.
func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DiskEncryptionKeyName),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DiskEncryptionKeyNameRef,
		Selector:     mg.Spec.ForProvider.DiskEncryptionKeyNameSelector,
		To: reference.To{
			List:    &v1beta1.CryptoKeyList{},
			Managed: &v1beta1.CryptoKey{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DiskEncryptionKeyName")
	}
	mg.Spec.ForProvider.DiskEncryptionKeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DiskEncryptionKeyNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.OrgIDRef,
		Selector:     mg.Spec.ForProvider.OrgIDSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrgIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NATAddress.
func (mg *NATAddress) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Organization.
func (mg *Organization) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AuthorizedNetwork),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.AuthorizedNetworkRef,
		Selector:     mg.Spec.ForProvider.AuthorizedNetworkSelector,
		To: reference.To{
			List:    &v1beta11.NetworkList{},
			Managed: &v1beta11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AuthorizedNetwork")
	}
	mg.Spec.ForProvider.AuthorizedNetwork = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AuthorizedNetworkRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RuntimeDatabaseEncryptionKeyName),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.RuntimeDatabaseEncryptionKeyNameRef,
		Selector:     mg.Spec.ForProvider.RuntimeDatabaseEncryptionKeyNameSelector,
		To: reference.To{
			List:    &v1beta1.CryptoKeyList{},
			Managed: &v1beta1.CryptoKey{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RuntimeDatabaseEncryptionKeyName")
	}
	mg.Spec.ForProvider.RuntimeDatabaseEncryptionKeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RuntimeDatabaseEncryptionKeyNameRef = rsp.ResolvedReference

	return nil
}
