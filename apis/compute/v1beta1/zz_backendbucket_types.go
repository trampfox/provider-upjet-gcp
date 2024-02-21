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

type BackendBucketInitParameters struct {

	// Cloud Storage bucket name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.Bucket
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Reference to a Bucket in storage to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameRef *v1.Reference `json:"bucketNameRef,omitempty" tf:"-"`

	// Selector for a Bucket in storage to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameSelector *v1.Selector `json:"bucketNameSelector,omitempty" tf:"-"`

	// Cloud CDN configuration for this Backend Bucket.
	// Structure is documented below.
	CdnPolicy []CdnPolicyInitParameters `json:"cdnPolicy,omitempty" tf:"cdn_policy,omitempty"`

	// Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
	// Possible values are: AUTOMATIC, DISABLED.
	CompressionMode *string `json:"compressionMode,omitempty" tf:"compression_mode,omitempty"`

	// Headers that the HTTP/S load balancer should add to proxied responses.
	CustomResponseHeaders []*string `json:"customResponseHeaders,omitempty" tf:"custom_response_headers,omitempty"`

	// An optional textual description of the resource; provided by the
	// client when the resource is created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The security policy associated with this backend bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.SecurityPolicy
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	EdgeSecurityPolicy *string `json:"edgeSecurityPolicy,omitempty" tf:"edge_security_policy,omitempty"`

	// Reference to a SecurityPolicy in compute to populate edgeSecurityPolicy.
	// +kubebuilder:validation:Optional
	EdgeSecurityPolicyRef *v1.Reference `json:"edgeSecurityPolicyRef,omitempty" tf:"-"`

	// Selector for a SecurityPolicy in compute to populate edgeSecurityPolicy.
	// +kubebuilder:validation:Optional
	EdgeSecurityPolicySelector *v1.Selector `json:"edgeSecurityPolicySelector,omitempty" tf:"-"`

	// If true, enable Cloud CDN for this BackendBucket.
	EnableCdn *bool `json:"enableCdn,omitempty" tf:"enable_cdn,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type BackendBucketObservation struct {

	// Cloud Storage bucket name.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Cloud CDN configuration for this Backend Bucket.
	// Structure is documented below.
	CdnPolicy []CdnPolicyObservation `json:"cdnPolicy,omitempty" tf:"cdn_policy,omitempty"`

	// Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
	// Possible values are: AUTOMATIC, DISABLED.
	CompressionMode *string `json:"compressionMode,omitempty" tf:"compression_mode,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// Headers that the HTTP/S load balancer should add to proxied responses.
	CustomResponseHeaders []*string `json:"customResponseHeaders,omitempty" tf:"custom_response_headers,omitempty"`

	// An optional textual description of the resource; provided by the
	// client when the resource is created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The security policy associated with this backend bucket.
	EdgeSecurityPolicy *string `json:"edgeSecurityPolicy,omitempty" tf:"edge_security_policy,omitempty"`

	// If true, enable Cloud CDN for this BackendBucket.
	EnableCdn *bool `json:"enableCdn,omitempty" tf:"enable_cdn,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/backendBuckets/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type BackendBucketParameters struct {

	// Cloud Storage bucket name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Reference to a Bucket in storage to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameRef *v1.Reference `json:"bucketNameRef,omitempty" tf:"-"`

	// Selector for a Bucket in storage to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameSelector *v1.Selector `json:"bucketNameSelector,omitempty" tf:"-"`

	// Cloud CDN configuration for this Backend Bucket.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	CdnPolicy []CdnPolicyParameters `json:"cdnPolicy,omitempty" tf:"cdn_policy,omitempty"`

	// Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
	// Possible values are: AUTOMATIC, DISABLED.
	// +kubebuilder:validation:Optional
	CompressionMode *string `json:"compressionMode,omitempty" tf:"compression_mode,omitempty"`

	// Headers that the HTTP/S load balancer should add to proxied responses.
	// +kubebuilder:validation:Optional
	CustomResponseHeaders []*string `json:"customResponseHeaders,omitempty" tf:"custom_response_headers,omitempty"`

	// An optional textual description of the resource; provided by the
	// client when the resource is created.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The security policy associated with this backend bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.SecurityPolicy
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	EdgeSecurityPolicy *string `json:"edgeSecurityPolicy,omitempty" tf:"edge_security_policy,omitempty"`

	// Reference to a SecurityPolicy in compute to populate edgeSecurityPolicy.
	// +kubebuilder:validation:Optional
	EdgeSecurityPolicyRef *v1.Reference `json:"edgeSecurityPolicyRef,omitempty" tf:"-"`

	// Selector for a SecurityPolicy in compute to populate edgeSecurityPolicy.
	// +kubebuilder:validation:Optional
	EdgeSecurityPolicySelector *v1.Selector `json:"edgeSecurityPolicySelector,omitempty" tf:"-"`

	// If true, enable Cloud CDN for this BackendBucket.
	// +kubebuilder:validation:Optional
	EnableCdn *bool `json:"enableCdn,omitempty" tf:"enable_cdn,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type BypassCacheOnRequestHeadersInitParameters struct {

	// The header field name to match on when bypassing cache. Values are case-insensitive.
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`
}

type BypassCacheOnRequestHeadersObservation struct {

	// The header field name to match on when bypassing cache. Values are case-insensitive.
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`
}

type BypassCacheOnRequestHeadersParameters struct {

	// The header field name to match on when bypassing cache. Values are case-insensitive.
	// +kubebuilder:validation:Optional
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`
}

type CacheKeyPolicyInitParameters struct {

	// Allows HTTP request headers (by name) to be used in the
	// cache key.
	IncludeHTTPHeaders []*string `json:"includeHttpHeaders,omitempty" tf:"include_http_headers,omitempty"`

	// Names of query string parameters to include in cache keys.
	// Default parameters are always included. '&' and '=' will
	// be percent encoded and not treated as delimiters.
	QueryStringWhitelist []*string `json:"queryStringWhitelist,omitempty" tf:"query_string_whitelist,omitempty"`
}

type CacheKeyPolicyObservation struct {

	// Allows HTTP request headers (by name) to be used in the
	// cache key.
	IncludeHTTPHeaders []*string `json:"includeHttpHeaders,omitempty" tf:"include_http_headers,omitempty"`

	// Names of query string parameters to include in cache keys.
	// Default parameters are always included. '&' and '=' will
	// be percent encoded and not treated as delimiters.
	QueryStringWhitelist []*string `json:"queryStringWhitelist,omitempty" tf:"query_string_whitelist,omitempty"`
}

type CacheKeyPolicyParameters struct {

	// Allows HTTP request headers (by name) to be used in the
	// cache key.
	// +kubebuilder:validation:Optional
	IncludeHTTPHeaders []*string `json:"includeHttpHeaders,omitempty" tf:"include_http_headers,omitempty"`

	// Names of query string parameters to include in cache keys.
	// Default parameters are always included. '&' and '=' will
	// be percent encoded and not treated as delimiters.
	// +kubebuilder:validation:Optional
	QueryStringWhitelist []*string `json:"queryStringWhitelist,omitempty" tf:"query_string_whitelist,omitempty"`
}

type CdnPolicyInitParameters struct {

	// Bypass the cache when the specified request headers are matched - e.g. Pragma or Authorization headers. Up to 5 headers can be specified. The cache is bypassed for all cdnPolicy.cacheMode settings.
	// Structure is documented below.
	BypassCacheOnRequestHeaders []BypassCacheOnRequestHeadersInitParameters `json:"bypassCacheOnRequestHeaders,omitempty" tf:"bypass_cache_on_request_headers,omitempty"`

	// The CacheKeyPolicy for this CdnPolicy.
	// Structure is documented below.
	CacheKeyPolicy []CacheKeyPolicyInitParameters `json:"cacheKeyPolicy,omitempty" tf:"cache_key_policy,omitempty"`

	// Specifies the cache setting for all responses from this backend.
	// The possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL and CACHE_ALL_STATIC
	// Possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL, CACHE_ALL_STATIC.
	CacheMode *string `json:"cacheMode,omitempty" tf:"cache_mode,omitempty"`

	// Specifies the maximum allowed TTL for cached content served by this origin.
	ClientTTL *float64 `json:"clientTtl,omitempty" tf:"client_ttl,omitempty"`

	// Specifies the default TTL for cached content served by this origin for responses
	// that do not have an existing valid TTL (max-age or s-max-age).
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// Specifies the maximum allowed TTL for cached content served by this origin.
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects.
	NegativeCaching *bool `json:"negativeCaching,omitempty" tf:"negative_caching,omitempty"`

	// Sets a cache TTL for the specified HTTP status code. negativeCaching must be enabled to configure negativeCachingPolicy.
	// Omitting the policy and leaving negativeCaching enabled will use Cloud CDN's default cache TTLs.
	// Structure is documented below.
	NegativeCachingPolicy []NegativeCachingPolicyInitParameters `json:"negativeCachingPolicy,omitempty" tf:"negative_caching_policy,omitempty"`

	// If true then Cloud CDN will combine multiple concurrent cache fill requests into a small number of requests to the origin.
	RequestCoalescing *bool `json:"requestCoalescing,omitempty" tf:"request_coalescing,omitempty"`

	// Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache.
	ServeWhileStale *float64 `json:"serveWhileStale,omitempty" tf:"serve_while_stale,omitempty"`

	// Maximum number of seconds the response to a signed URL request will
	// be considered fresh. After this time period,
	// the response will be revalidated before being served.
	// When serving responses to signed URL requests,
	// Cloud CDN will internally behave as though
	// all responses from this backend had a "Cache-Control: public,
	// max-age=[TTL]" header, regardless of any existing Cache-Control
	// header. The actual headers served in responses will not be altered.
	SignedURLCacheMaxAgeSec *float64 `json:"signedUrlCacheMaxAgeSec,omitempty" tf:"signed_url_cache_max_age_sec,omitempty"`
}

type CdnPolicyObservation struct {

	// Bypass the cache when the specified request headers are matched - e.g. Pragma or Authorization headers. Up to 5 headers can be specified. The cache is bypassed for all cdnPolicy.cacheMode settings.
	// Structure is documented below.
	BypassCacheOnRequestHeaders []BypassCacheOnRequestHeadersObservation `json:"bypassCacheOnRequestHeaders,omitempty" tf:"bypass_cache_on_request_headers,omitempty"`

	// The CacheKeyPolicy for this CdnPolicy.
	// Structure is documented below.
	CacheKeyPolicy []CacheKeyPolicyObservation `json:"cacheKeyPolicy,omitempty" tf:"cache_key_policy,omitempty"`

	// Specifies the cache setting for all responses from this backend.
	// The possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL and CACHE_ALL_STATIC
	// Possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL, CACHE_ALL_STATIC.
	CacheMode *string `json:"cacheMode,omitempty" tf:"cache_mode,omitempty"`

	// Specifies the maximum allowed TTL for cached content served by this origin.
	ClientTTL *float64 `json:"clientTtl,omitempty" tf:"client_ttl,omitempty"`

	// Specifies the default TTL for cached content served by this origin for responses
	// that do not have an existing valid TTL (max-age or s-max-age).
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// Specifies the maximum allowed TTL for cached content served by this origin.
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects.
	NegativeCaching *bool `json:"negativeCaching,omitempty" tf:"negative_caching,omitempty"`

	// Sets a cache TTL for the specified HTTP status code. negativeCaching must be enabled to configure negativeCachingPolicy.
	// Omitting the policy and leaving negativeCaching enabled will use Cloud CDN's default cache TTLs.
	// Structure is documented below.
	NegativeCachingPolicy []NegativeCachingPolicyObservation `json:"negativeCachingPolicy,omitempty" tf:"negative_caching_policy,omitempty"`

	// If true then Cloud CDN will combine multiple concurrent cache fill requests into a small number of requests to the origin.
	RequestCoalescing *bool `json:"requestCoalescing,omitempty" tf:"request_coalescing,omitempty"`

	// Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache.
	ServeWhileStale *float64 `json:"serveWhileStale,omitempty" tf:"serve_while_stale,omitempty"`

	// Maximum number of seconds the response to a signed URL request will
	// be considered fresh. After this time period,
	// the response will be revalidated before being served.
	// When serving responses to signed URL requests,
	// Cloud CDN will internally behave as though
	// all responses from this backend had a "Cache-Control: public,
	// max-age=[TTL]" header, regardless of any existing Cache-Control
	// header. The actual headers served in responses will not be altered.
	SignedURLCacheMaxAgeSec *float64 `json:"signedUrlCacheMaxAgeSec,omitempty" tf:"signed_url_cache_max_age_sec,omitempty"`
}

type CdnPolicyParameters struct {

	// Bypass the cache when the specified request headers are matched - e.g. Pragma or Authorization headers. Up to 5 headers can be specified. The cache is bypassed for all cdnPolicy.cacheMode settings.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	BypassCacheOnRequestHeaders []BypassCacheOnRequestHeadersParameters `json:"bypassCacheOnRequestHeaders,omitempty" tf:"bypass_cache_on_request_headers,omitempty"`

	// The CacheKeyPolicy for this CdnPolicy.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	CacheKeyPolicy []CacheKeyPolicyParameters `json:"cacheKeyPolicy,omitempty" tf:"cache_key_policy,omitempty"`

	// Specifies the cache setting for all responses from this backend.
	// The possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL and CACHE_ALL_STATIC
	// Possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL, CACHE_ALL_STATIC.
	// +kubebuilder:validation:Optional
	CacheMode *string `json:"cacheMode,omitempty" tf:"cache_mode,omitempty"`

	// Specifies the maximum allowed TTL for cached content served by this origin.
	// +kubebuilder:validation:Optional
	ClientTTL *float64 `json:"clientTtl,omitempty" tf:"client_ttl,omitempty"`

	// Specifies the default TTL for cached content served by this origin for responses
	// that do not have an existing valid TTL (max-age or s-max-age).
	// +kubebuilder:validation:Optional
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// Specifies the maximum allowed TTL for cached content served by this origin.
	// +kubebuilder:validation:Optional
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects.
	// +kubebuilder:validation:Optional
	NegativeCaching *bool `json:"negativeCaching,omitempty" tf:"negative_caching,omitempty"`

	// Sets a cache TTL for the specified HTTP status code. negativeCaching must be enabled to configure negativeCachingPolicy.
	// Omitting the policy and leaving negativeCaching enabled will use Cloud CDN's default cache TTLs.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	NegativeCachingPolicy []NegativeCachingPolicyParameters `json:"negativeCachingPolicy,omitempty" tf:"negative_caching_policy,omitempty"`

	// If true then Cloud CDN will combine multiple concurrent cache fill requests into a small number of requests to the origin.
	// +kubebuilder:validation:Optional
	RequestCoalescing *bool `json:"requestCoalescing,omitempty" tf:"request_coalescing,omitempty"`

	// Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache.
	// +kubebuilder:validation:Optional
	ServeWhileStale *float64 `json:"serveWhileStale,omitempty" tf:"serve_while_stale,omitempty"`

	// Maximum number of seconds the response to a signed URL request will
	// be considered fresh. After this time period,
	// the response will be revalidated before being served.
	// When serving responses to signed URL requests,
	// Cloud CDN will internally behave as though
	// all responses from this backend had a "Cache-Control: public,
	// max-age=[TTL]" header, regardless of any existing Cache-Control
	// header. The actual headers served in responses will not be altered.
	// +kubebuilder:validation:Optional
	SignedURLCacheMaxAgeSec *float64 `json:"signedUrlCacheMaxAgeSec,omitempty" tf:"signed_url_cache_max_age_sec,omitempty"`
}

type NegativeCachingPolicyInitParameters struct {

	// The HTTP status code to define a TTL against. Only HTTP status codes 300, 301, 308, 404, 405, 410, 421, 451 and 501
	// can be specified as values, and you cannot specify a status code more than once.
	Code *float64 `json:"code,omitempty" tf:"code,omitempty"`

	// The TTL (in seconds) for which to cache responses with the corresponding status code. The maximum allowed value is 1800s
	// (30 minutes), noting that infrequently accessed objects may be evicted from the cache before the defined TTL.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type NegativeCachingPolicyObservation struct {

	// The HTTP status code to define a TTL against. Only HTTP status codes 300, 301, 308, 404, 405, 410, 421, 451 and 501
	// can be specified as values, and you cannot specify a status code more than once.
	Code *float64 `json:"code,omitempty" tf:"code,omitempty"`

	// The TTL (in seconds) for which to cache responses with the corresponding status code. The maximum allowed value is 1800s
	// (30 minutes), noting that infrequently accessed objects may be evicted from the cache before the defined TTL.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type NegativeCachingPolicyParameters struct {

	// The HTTP status code to define a TTL against. Only HTTP status codes 300, 301, 308, 404, 405, 410, 421, 451 and 501
	// can be specified as values, and you cannot specify a status code more than once.
	// +kubebuilder:validation:Optional
	Code *float64 `json:"code,omitempty" tf:"code,omitempty"`

	// The TTL (in seconds) for which to cache responses with the corresponding status code. The maximum allowed value is 1800s
	// (30 minutes), noting that infrequently accessed objects may be evicted from the cache before the defined TTL.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

// BackendBucketSpec defines the desired state of BackendBucket
type BackendBucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackendBucketParameters `json:"forProvider"`
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
	InitProvider BackendBucketInitParameters `json:"initProvider,omitempty"`
}

// BackendBucketStatus defines the observed state of BackendBucket.
type BackendBucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackendBucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BackendBucket is the Schema for the BackendBuckets API. Backend buckets allow you to use Google Cloud Storage buckets with HTTP(S) load balancing.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type BackendBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackendBucketSpec   `json:"spec"`
	Status            BackendBucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackendBucketList contains a list of BackendBuckets
type BackendBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackendBucket `json:"items"`
}

// Repository type metadata.
var (
	BackendBucket_Kind             = "BackendBucket"
	BackendBucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackendBucket_Kind}.String()
	BackendBucket_KindAPIVersion   = BackendBucket_Kind + "." + CRDGroupVersion.String()
	BackendBucket_GroupVersionKind = CRDGroupVersion.WithKind(BackendBucket_Kind)
)

func init() {
	SchemeBuilder.Register(&BackendBucket{}, &BackendBucketList{})
}
