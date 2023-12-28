//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEngineHTTPTargetInitParameters) DeepCopyInto(out *AppEngineHTTPTargetInitParameters) {
	*out = *in
	if in.AppEngineRouting != nil {
		in, out := &in.AppEngineRouting, &out.AppEngineRouting
		*out = make([]AppEngineRoutingInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = new(string)
		**out = **in
	}
	if in.HTTPMethod != nil {
		in, out := &in.HTTPMethod, &out.HTTPMethod
		*out = new(string)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.RelativeURI != nil {
		in, out := &in.RelativeURI, &out.RelativeURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEngineHTTPTargetInitParameters.
func (in *AppEngineHTTPTargetInitParameters) DeepCopy() *AppEngineHTTPTargetInitParameters {
	if in == nil {
		return nil
	}
	out := new(AppEngineHTTPTargetInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEngineHTTPTargetObservation) DeepCopyInto(out *AppEngineHTTPTargetObservation) {
	*out = *in
	if in.AppEngineRouting != nil {
		in, out := &in.AppEngineRouting, &out.AppEngineRouting
		*out = make([]AppEngineRoutingObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = new(string)
		**out = **in
	}
	if in.HTTPMethod != nil {
		in, out := &in.HTTPMethod, &out.HTTPMethod
		*out = new(string)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.RelativeURI != nil {
		in, out := &in.RelativeURI, &out.RelativeURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEngineHTTPTargetObservation.
func (in *AppEngineHTTPTargetObservation) DeepCopy() *AppEngineHTTPTargetObservation {
	if in == nil {
		return nil
	}
	out := new(AppEngineHTTPTargetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEngineHTTPTargetParameters) DeepCopyInto(out *AppEngineHTTPTargetParameters) {
	*out = *in
	if in.AppEngineRouting != nil {
		in, out := &in.AppEngineRouting, &out.AppEngineRouting
		*out = make([]AppEngineRoutingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = new(string)
		**out = **in
	}
	if in.HTTPMethod != nil {
		in, out := &in.HTTPMethod, &out.HTTPMethod
		*out = new(string)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.RelativeURI != nil {
		in, out := &in.RelativeURI, &out.RelativeURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEngineHTTPTargetParameters.
func (in *AppEngineHTTPTargetParameters) DeepCopy() *AppEngineHTTPTargetParameters {
	if in == nil {
		return nil
	}
	out := new(AppEngineHTTPTargetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEngineRoutingInitParameters) DeepCopyInto(out *AppEngineRoutingInitParameters) {
	*out = *in
	if in.Instance != nil {
		in, out := &in.Instance, &out.Instance
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEngineRoutingInitParameters.
func (in *AppEngineRoutingInitParameters) DeepCopy() *AppEngineRoutingInitParameters {
	if in == nil {
		return nil
	}
	out := new(AppEngineRoutingInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEngineRoutingObservation) DeepCopyInto(out *AppEngineRoutingObservation) {
	*out = *in
	if in.Instance != nil {
		in, out := &in.Instance, &out.Instance
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEngineRoutingObservation.
func (in *AppEngineRoutingObservation) DeepCopy() *AppEngineRoutingObservation {
	if in == nil {
		return nil
	}
	out := new(AppEngineRoutingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEngineRoutingParameters) DeepCopyInto(out *AppEngineRoutingParameters) {
	*out = *in
	if in.Instance != nil {
		in, out := &in.Instance, &out.Instance
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEngineRoutingParameters.
func (in *AppEngineRoutingParameters) DeepCopy() *AppEngineRoutingParameters {
	if in == nil {
		return nil
	}
	out := new(AppEngineRoutingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPTargetInitParameters) DeepCopyInto(out *HTTPTargetInitParameters) {
	*out = *in
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = new(string)
		**out = **in
	}
	if in.HTTPMethod != nil {
		in, out := &in.HTTPMethod, &out.HTTPMethod
		*out = new(string)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.OAuthToken != nil {
		in, out := &in.OAuthToken, &out.OAuthToken
		*out = make([]OAuthTokenInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OidcToken != nil {
		in, out := &in.OidcToken, &out.OidcToken
		*out = make([]OidcTokenInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPTargetInitParameters.
func (in *HTTPTargetInitParameters) DeepCopy() *HTTPTargetInitParameters {
	if in == nil {
		return nil
	}
	out := new(HTTPTargetInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPTargetObservation) DeepCopyInto(out *HTTPTargetObservation) {
	*out = *in
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = new(string)
		**out = **in
	}
	if in.HTTPMethod != nil {
		in, out := &in.HTTPMethod, &out.HTTPMethod
		*out = new(string)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.OAuthToken != nil {
		in, out := &in.OAuthToken, &out.OAuthToken
		*out = make([]OAuthTokenObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OidcToken != nil {
		in, out := &in.OidcToken, &out.OidcToken
		*out = make([]OidcTokenObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPTargetObservation.
func (in *HTTPTargetObservation) DeepCopy() *HTTPTargetObservation {
	if in == nil {
		return nil
	}
	out := new(HTTPTargetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPTargetParameters) DeepCopyInto(out *HTTPTargetParameters) {
	*out = *in
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = new(string)
		**out = **in
	}
	if in.HTTPMethod != nil {
		in, out := &in.HTTPMethod, &out.HTTPMethod
		*out = new(string)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.OAuthToken != nil {
		in, out := &in.OAuthToken, &out.OAuthToken
		*out = make([]OAuthTokenParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OidcToken != nil {
		in, out := &in.OidcToken, &out.OidcToken
		*out = make([]OidcTokenParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPTargetParameters.
func (in *HTTPTargetParameters) DeepCopy() *HTTPTargetParameters {
	if in == nil {
		return nil
	}
	out := new(HTTPTargetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Job) DeepCopyInto(out *Job) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Job.
func (in *Job) DeepCopy() *Job {
	if in == nil {
		return nil
	}
	out := new(Job)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Job) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobInitParameters) DeepCopyInto(out *JobInitParameters) {
	*out = *in
	if in.AppEngineHTTPTarget != nil {
		in, out := &in.AppEngineHTTPTarget, &out.AppEngineHTTPTarget
		*out = make([]AppEngineHTTPTargetInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AttemptDeadline != nil {
		in, out := &in.AttemptDeadline, &out.AttemptDeadline
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.HTTPTarget != nil {
		in, out := &in.HTTPTarget, &out.HTTPTarget
		*out = make([]HTTPTargetInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.PubsubTarget != nil {
		in, out := &in.PubsubTarget, &out.PubsubTarget
		*out = make([]PubsubTargetInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RetryConfig != nil {
		in, out := &in.RetryConfig, &out.RetryConfig
		*out = make([]RetryConfigInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobInitParameters.
func (in *JobInitParameters) DeepCopy() *JobInitParameters {
	if in == nil {
		return nil
	}
	out := new(JobInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobList) DeepCopyInto(out *JobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Job, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobList.
func (in *JobList) DeepCopy() *JobList {
	if in == nil {
		return nil
	}
	out := new(JobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobObservation) DeepCopyInto(out *JobObservation) {
	*out = *in
	if in.AppEngineHTTPTarget != nil {
		in, out := &in.AppEngineHTTPTarget, &out.AppEngineHTTPTarget
		*out = make([]AppEngineHTTPTargetObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AttemptDeadline != nil {
		in, out := &in.AttemptDeadline, &out.AttemptDeadline
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.HTTPTarget != nil {
		in, out := &in.HTTPTarget, &out.HTTPTarget
		*out = make([]HTTPTargetObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.PubsubTarget != nil {
		in, out := &in.PubsubTarget, &out.PubsubTarget
		*out = make([]PubsubTargetObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RetryConfig != nil {
		in, out := &in.RetryConfig, &out.RetryConfig
		*out = make([]RetryConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobObservation.
func (in *JobObservation) DeepCopy() *JobObservation {
	if in == nil {
		return nil
	}
	out := new(JobObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobParameters) DeepCopyInto(out *JobParameters) {
	*out = *in
	if in.AppEngineHTTPTarget != nil {
		in, out := &in.AppEngineHTTPTarget, &out.AppEngineHTTPTarget
		*out = make([]AppEngineHTTPTargetParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AttemptDeadline != nil {
		in, out := &in.AttemptDeadline, &out.AttemptDeadline
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.HTTPTarget != nil {
		in, out := &in.HTTPTarget, &out.HTTPTarget
		*out = make([]HTTPTargetParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.PubsubTarget != nil {
		in, out := &in.PubsubTarget, &out.PubsubTarget
		*out = make([]PubsubTargetParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RetryConfig != nil {
		in, out := &in.RetryConfig, &out.RetryConfig
		*out = make([]RetryConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobParameters.
func (in *JobParameters) DeepCopy() *JobParameters {
	if in == nil {
		return nil
	}
	out := new(JobParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobSpec) DeepCopyInto(out *JobSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobSpec.
func (in *JobSpec) DeepCopy() *JobSpec {
	if in == nil {
		return nil
	}
	out := new(JobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobStatus) DeepCopyInto(out *JobStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobStatus.
func (in *JobStatus) DeepCopy() *JobStatus {
	if in == nil {
		return nil
	}
	out := new(JobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OAuthTokenInitParameters) DeepCopyInto(out *OAuthTokenInitParameters) {
	*out = *in
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountEmail != nil {
		in, out := &in.ServiceAccountEmail, &out.ServiceAccountEmail
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OAuthTokenInitParameters.
func (in *OAuthTokenInitParameters) DeepCopy() *OAuthTokenInitParameters {
	if in == nil {
		return nil
	}
	out := new(OAuthTokenInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OAuthTokenObservation) DeepCopyInto(out *OAuthTokenObservation) {
	*out = *in
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountEmail != nil {
		in, out := &in.ServiceAccountEmail, &out.ServiceAccountEmail
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OAuthTokenObservation.
func (in *OAuthTokenObservation) DeepCopy() *OAuthTokenObservation {
	if in == nil {
		return nil
	}
	out := new(OAuthTokenObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OAuthTokenParameters) DeepCopyInto(out *OAuthTokenParameters) {
	*out = *in
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountEmail != nil {
		in, out := &in.ServiceAccountEmail, &out.ServiceAccountEmail
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OAuthTokenParameters.
func (in *OAuthTokenParameters) DeepCopy() *OAuthTokenParameters {
	if in == nil {
		return nil
	}
	out := new(OAuthTokenParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OidcTokenInitParameters) DeepCopyInto(out *OidcTokenInitParameters) {
	*out = *in
	if in.Audience != nil {
		in, out := &in.Audience, &out.Audience
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountEmail != nil {
		in, out := &in.ServiceAccountEmail, &out.ServiceAccountEmail
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcTokenInitParameters.
func (in *OidcTokenInitParameters) DeepCopy() *OidcTokenInitParameters {
	if in == nil {
		return nil
	}
	out := new(OidcTokenInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OidcTokenObservation) DeepCopyInto(out *OidcTokenObservation) {
	*out = *in
	if in.Audience != nil {
		in, out := &in.Audience, &out.Audience
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountEmail != nil {
		in, out := &in.ServiceAccountEmail, &out.ServiceAccountEmail
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcTokenObservation.
func (in *OidcTokenObservation) DeepCopy() *OidcTokenObservation {
	if in == nil {
		return nil
	}
	out := new(OidcTokenObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OidcTokenParameters) DeepCopyInto(out *OidcTokenParameters) {
	*out = *in
	if in.Audience != nil {
		in, out := &in.Audience, &out.Audience
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountEmail != nil {
		in, out := &in.ServiceAccountEmail, &out.ServiceAccountEmail
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcTokenParameters.
func (in *OidcTokenParameters) DeepCopy() *OidcTokenParameters {
	if in == nil {
		return nil
	}
	out := new(OidcTokenParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PubsubTargetInitParameters) DeepCopyInto(out *PubsubTargetInitParameters) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(string)
		**out = **in
	}
	if in.TopicName != nil {
		in, out := &in.TopicName, &out.TopicName
		*out = new(string)
		**out = **in
	}
	if in.TopicNameRef != nil {
		in, out := &in.TopicNameRef, &out.TopicNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TopicNameSelector != nil {
		in, out := &in.TopicNameSelector, &out.TopicNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubsubTargetInitParameters.
func (in *PubsubTargetInitParameters) DeepCopy() *PubsubTargetInitParameters {
	if in == nil {
		return nil
	}
	out := new(PubsubTargetInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PubsubTargetObservation) DeepCopyInto(out *PubsubTargetObservation) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(string)
		**out = **in
	}
	if in.TopicName != nil {
		in, out := &in.TopicName, &out.TopicName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubsubTargetObservation.
func (in *PubsubTargetObservation) DeepCopy() *PubsubTargetObservation {
	if in == nil {
		return nil
	}
	out := new(PubsubTargetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PubsubTargetParameters) DeepCopyInto(out *PubsubTargetParameters) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(string)
		**out = **in
	}
	if in.TopicName != nil {
		in, out := &in.TopicName, &out.TopicName
		*out = new(string)
		**out = **in
	}
	if in.TopicNameRef != nil {
		in, out := &in.TopicNameRef, &out.TopicNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TopicNameSelector != nil {
		in, out := &in.TopicNameSelector, &out.TopicNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubsubTargetParameters.
func (in *PubsubTargetParameters) DeepCopy() *PubsubTargetParameters {
	if in == nil {
		return nil
	}
	out := new(PubsubTargetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryConfigInitParameters) DeepCopyInto(out *RetryConfigInitParameters) {
	*out = *in
	if in.MaxBackoffDuration != nil {
		in, out := &in.MaxBackoffDuration, &out.MaxBackoffDuration
		*out = new(string)
		**out = **in
	}
	if in.MaxDoublings != nil {
		in, out := &in.MaxDoublings, &out.MaxDoublings
		*out = new(float64)
		**out = **in
	}
	if in.MaxRetryDuration != nil {
		in, out := &in.MaxRetryDuration, &out.MaxRetryDuration
		*out = new(string)
		**out = **in
	}
	if in.MinBackoffDuration != nil {
		in, out := &in.MinBackoffDuration, &out.MinBackoffDuration
		*out = new(string)
		**out = **in
	}
	if in.RetryCount != nil {
		in, out := &in.RetryCount, &out.RetryCount
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryConfigInitParameters.
func (in *RetryConfigInitParameters) DeepCopy() *RetryConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(RetryConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryConfigObservation) DeepCopyInto(out *RetryConfigObservation) {
	*out = *in
	if in.MaxBackoffDuration != nil {
		in, out := &in.MaxBackoffDuration, &out.MaxBackoffDuration
		*out = new(string)
		**out = **in
	}
	if in.MaxDoublings != nil {
		in, out := &in.MaxDoublings, &out.MaxDoublings
		*out = new(float64)
		**out = **in
	}
	if in.MaxRetryDuration != nil {
		in, out := &in.MaxRetryDuration, &out.MaxRetryDuration
		*out = new(string)
		**out = **in
	}
	if in.MinBackoffDuration != nil {
		in, out := &in.MinBackoffDuration, &out.MinBackoffDuration
		*out = new(string)
		**out = **in
	}
	if in.RetryCount != nil {
		in, out := &in.RetryCount, &out.RetryCount
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryConfigObservation.
func (in *RetryConfigObservation) DeepCopy() *RetryConfigObservation {
	if in == nil {
		return nil
	}
	out := new(RetryConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryConfigParameters) DeepCopyInto(out *RetryConfigParameters) {
	*out = *in
	if in.MaxBackoffDuration != nil {
		in, out := &in.MaxBackoffDuration, &out.MaxBackoffDuration
		*out = new(string)
		**out = **in
	}
	if in.MaxDoublings != nil {
		in, out := &in.MaxDoublings, &out.MaxDoublings
		*out = new(float64)
		**out = **in
	}
	if in.MaxRetryDuration != nil {
		in, out := &in.MaxRetryDuration, &out.MaxRetryDuration
		*out = new(string)
		**out = **in
	}
	if in.MinBackoffDuration != nil {
		in, out := &in.MinBackoffDuration, &out.MinBackoffDuration
		*out = new(string)
		**out = **in
	}
	if in.RetryCount != nil {
		in, out := &in.RetryCount, &out.RetryCount
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryConfigParameters.
func (in *RetryConfigParameters) DeepCopy() *RetryConfigParameters {
	if in == nil {
		return nil
	}
	out := new(RetryConfigParameters)
	in.DeepCopyInto(out)
	return out
}
