/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AwsCloudMapObservation struct {
}

type AwsCloudMapParameters struct {

	// String map that contains attributes with values that you can use to filter instances by any custom attribute that you specified when you registered the instance. Only instances that match all of the specified key/value pairs will be returned.
	// +kubebuilder:validation:Optional
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// Name of the AWS Cloud Map namespace to use.
	// Use the aws_service_discovery_http_namespace resource to configure a Cloud Map namespace. Must be between 1 and 1024 characters in length.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/servicediscovery/v1beta1.HTTPNamespace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// Reference to a HTTPNamespace in servicediscovery to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameRef *v1.Reference `json:"namespaceNameRef,omitempty" tf:"-"`

	// Selector for a HTTPNamespace in servicediscovery to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameSelector *v1.Selector `json:"namespaceNameSelector,omitempty" tf:"-"`

	// attribute of the dns object to hostname.
	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`
}

type BackendDefaultsClientPolicyObservation struct {
}

type BackendDefaultsClientPolicyParameters struct {

	// Transport Layer Security (TLS) client policy.
	// +kubebuilder:validation:Optional
	TLS []BackendDefaultsClientPolicyTLSParameters `json:"tls,omitempty" tf:"tls,omitempty"`
}

type BackendDefaultsClientPolicyTLSCertificateObservation struct {
}

type BackendDefaultsClientPolicyTLSCertificateParameters struct {

	// Local file certificate.
	// +kubebuilder:validation:Optional
	File []ClientPolicyTLSCertificateFileParameters `json:"file,omitempty" tf:"file,omitempty"`

	// A Secret Discovery Service certificate.
	// +kubebuilder:validation:Optional
	Sds []ClientPolicyTLSCertificateSdsParameters `json:"sds,omitempty" tf:"sds,omitempty"`
}

type BackendDefaultsClientPolicyTLSObservation struct {
}

type BackendDefaultsClientPolicyTLSParameters struct {

	// Virtual node's client's Transport Layer Security (TLS) certificate.
	// +kubebuilder:validation:Optional
	Certificate []BackendDefaultsClientPolicyTLSCertificateParameters `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Whether the policy is enforced. Default is true.
	// +kubebuilder:validation:Optional
	Enforce *bool `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// One or more ports that the policy is enforced for.
	// +kubebuilder:validation:Optional
	Ports []*float64 `json:"ports,omitempty" tf:"ports,omitempty"`

	// TLS validation context.
	// +kubebuilder:validation:Required
	Validation []BackendDefaultsClientPolicyTLSValidationParameters `json:"validation" tf:"validation,omitempty"`
}

type BackendDefaultsClientPolicyTLSValidationObservation struct {
}

type BackendDefaultsClientPolicyTLSValidationParameters struct {

	// SANs for a TLS validation context.
	// +kubebuilder:validation:Optional
	SubjectAlternativeNames []ClientPolicyTLSValidationSubjectAlternativeNamesParameters `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// TLS validation context trust.
	// +kubebuilder:validation:Required
	Trust []ClientPolicyTLSValidationTrustParameters `json:"trust" tf:"trust,omitempty"`
}

type BackendObservation struct {
}

type BackendParameters struct {

	// Virtual service to use as a backend for a virtual node.
	// +kubebuilder:validation:Required
	VirtualService []BackendVirtualServiceParameters `json:"virtualService" tf:"virtual_service,omitempty"`
}

type BackendVirtualServiceObservation struct {
}

type BackendVirtualServiceParameters struct {

	// Client policy for the backend.
	// +kubebuilder:validation:Optional
	ClientPolicy []VirtualServiceClientPolicyParameters `json:"clientPolicy,omitempty" tf:"client_policy,omitempty"`

	// Name of the virtual service that is acting as a virtual node backend. Must be between 1 and 255 characters in length.
	// +kubebuilder:validation:Required
	VirtualServiceName *string `json:"virtualServiceName" tf:"virtual_service_name,omitempty"`
}

type BaseEjectionDurationObservation struct {
}

type BaseEjectionDurationParameters struct {

	// Unit of time. Valid values: ms, s.
	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`

	// Number of time units. Minimum value of 0.
	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type ClientPolicyTLSCertificateFileObservation struct {
}

type ClientPolicyTLSCertificateFileParameters struct {

	// Certificate chain for the certificate.
	// +kubebuilder:validation:Required
	CertificateChain *string `json:"certificateChain" tf:"certificate_chain,omitempty"`

	// Private key for a certificate stored on the file system of the mesh endpoint that the proxy is running on.
	// +kubebuilder:validation:Required
	PrivateKey *string `json:"privateKey" tf:"private_key,omitempty"`
}

type ClientPolicyTLSCertificateObservation struct {
}

type ClientPolicyTLSCertificateParameters struct {

	// Local file certificate.
	// +kubebuilder:validation:Optional
	File []TLSCertificateFileParameters `json:"file,omitempty" tf:"file,omitempty"`

	// A Secret Discovery Service certificate.
	// +kubebuilder:validation:Optional
	Sds []TLSCertificateSdsParameters `json:"sds,omitempty" tf:"sds,omitempty"`
}

type ClientPolicyTLSCertificateSdsObservation struct {
}

type ClientPolicyTLSCertificateSdsParameters struct {

	// Name of the secret secret requested from the Secret Discovery Service provider representing Transport Layer Security (TLS) materials like a certificate or certificate chain.
	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type ClientPolicyTLSObservation struct {
}

type ClientPolicyTLSParameters struct {

	// Virtual node's client's Transport Layer Security (TLS) certificate.
	// +kubebuilder:validation:Optional
	Certificate []ClientPolicyTLSCertificateParameters `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Whether the policy is enforced. Default is true.
	// +kubebuilder:validation:Optional
	Enforce *bool `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// One or more ports that the policy is enforced for.
	// +kubebuilder:validation:Optional
	Ports []*float64 `json:"ports,omitempty" tf:"ports,omitempty"`

	// TLS validation context.
	// +kubebuilder:validation:Required
	Validation []ClientPolicyTLSValidationParameters `json:"validation" tf:"validation,omitempty"`
}

type ClientPolicyTLSValidationObservation struct {
}

type ClientPolicyTLSValidationParameters struct {

	// SANs for a TLS validation context.
	// +kubebuilder:validation:Optional
	SubjectAlternativeNames []TLSValidationSubjectAlternativeNamesParameters `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// TLS validation context trust.
	// +kubebuilder:validation:Required
	Trust []TLSValidationTrustParameters `json:"trust" tf:"trust,omitempty"`
}

type ClientPolicyTLSValidationSubjectAlternativeNamesMatchObservation struct {
}

type ClientPolicyTLSValidationSubjectAlternativeNamesMatchParameters struct {

	// Values sent must match the specified values exactly.
	// +kubebuilder:validation:Required
	Exact []*string `json:"exact" tf:"exact,omitempty"`
}

type ClientPolicyTLSValidationSubjectAlternativeNamesObservation struct {
}

type ClientPolicyTLSValidationSubjectAlternativeNamesParameters struct {

	// Criteria for determining a SAN's match.
	// +kubebuilder:validation:Required
	Match []ClientPolicyTLSValidationSubjectAlternativeNamesMatchParameters `json:"match" tf:"match,omitempty"`
}

type ClientPolicyTLSValidationTrustFileObservation struct {
}

type ClientPolicyTLSValidationTrustFileParameters struct {

	// Certificate chain for the certificate.
	// +kubebuilder:validation:Required
	CertificateChain *string `json:"certificateChain" tf:"certificate_chain,omitempty"`
}

type ClientPolicyTLSValidationTrustObservation struct {
}

type ClientPolicyTLSValidationTrustParameters struct {

	// TLS validation context trust for an AWS Certificate Manager (ACM) certificate.
	// +kubebuilder:validation:Optional
	Acm []ValidationTrustAcmParameters `json:"acm,omitempty" tf:"acm,omitempty"`

	// Local file certificate.
	// +kubebuilder:validation:Optional
	File []ClientPolicyTLSValidationTrustFileParameters `json:"file,omitempty" tf:"file,omitempty"`

	// A Secret Discovery Service certificate.
	// +kubebuilder:validation:Optional
	Sds []ClientPolicyTLSValidationTrustSdsParameters `json:"sds,omitempty" tf:"sds,omitempty"`
}

type ClientPolicyTLSValidationTrustSdsObservation struct {
}

type ClientPolicyTLSValidationTrustSdsParameters struct {

	// Name of the secret secret requested from the Secret Discovery Service provider representing Transport Layer Security (TLS) materials like a certificate or certificate chain.
	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type ConnectionPoolGRPCObservation struct {
}

type ConnectionPoolGRPCParameters struct {

	// Maximum number of inflight requests Envoy can concurrently support across hosts in upstream cluster. Minimum value of 1.
	// +kubebuilder:validation:Required
	MaxRequests *float64 `json:"maxRequests" tf:"max_requests,omitempty"`
}

type ConnectionPoolHTTPObservation struct {
}

type ConnectionPoolHTTPParameters struct {

	// Maximum number of outbound TCP connections Envoy can establish concurrently with all hosts in upstream cluster. Minimum value of 1.
	// +kubebuilder:validation:Required
	MaxConnections *float64 `json:"maxConnections" tf:"max_connections,omitempty"`

	// Number of overflowing requests after max_connections Envoy will queue to upstream cluster. Minimum value of 1.
	// +kubebuilder:validation:Optional
	MaxPendingRequests *float64 `json:"maxPendingRequests,omitempty" tf:"max_pending_requests,omitempty"`
}

type ConnectionPoolHttp2Observation struct {
}

type ConnectionPoolHttp2Parameters struct {

	// Maximum number of inflight requests Envoy can concurrently support across hosts in upstream cluster. Minimum value of 1.
	// +kubebuilder:validation:Required
	MaxRequests *float64 `json:"maxRequests" tf:"max_requests,omitempty"`
}

type DNSObservation struct {
}

type DNSParameters struct {

	// DNS host name for your virtual node.
	// +kubebuilder:validation:Required
	Hostname *string `json:"hostname" tf:"hostname,omitempty"`
}

type GRPCIdleObservation struct {
}

type GRPCIdleParameters struct {

	// Unit of time. Valid values: ms, s.
	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`

	// Number of time units. Minimum value of 0.
	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type GRPCPerRequestObservation struct {
}

type GRPCPerRequestParameters struct {

	// Unit of time. Valid values: ms, s.
	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`

	// Number of time units. Minimum value of 0.
	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type HTTPIdleObservation struct {
}

type HTTPIdleParameters struct {

	// Unit of time. Valid values: ms, s.
	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`

	// Number of time units. Minimum value of 0.
	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type HTTPPerRequestObservation struct {
}

type HTTPPerRequestParameters struct {

	// Unit of time. Valid values: ms, s.
	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`

	// Number of time units. Minimum value of 0.
	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type Http2IdleObservation struct {
}

type Http2IdleParameters struct {

	// Unit of time. Valid values: ms, s.
	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`

	// Number of time units. Minimum value of 0.
	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type Http2PerRequestObservation struct {
}

type Http2PerRequestParameters struct {

	// Unit of time. Valid values: ms, s.
	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`

	// Number of time units. Minimum value of 0.
	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type IntervalObservation struct {
}

type IntervalParameters struct {

	// Unit of time. Valid values: ms, s.
	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`

	// Number of time units. Minimum value of 0.
	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type ListenerConnectionPoolObservation struct {
}

type ListenerConnectionPoolParameters struct {

	// Connection pool information for gRPC listeners.
	// +kubebuilder:validation:Optional
	GRPC []ConnectionPoolGRPCParameters `json:"grpc,omitempty" tf:"grpc,omitempty"`

	// Connection pool information for HTTP listeners.
	// +kubebuilder:validation:Optional
	HTTP []ConnectionPoolHTTPParameters `json:"http,omitempty" tf:"http,omitempty"`

	// Connection pool information for HTTP2 listeners.
	// +kubebuilder:validation:Optional
	Http2 []ConnectionPoolHttp2Parameters `json:"http2,omitempty" tf:"http2,omitempty"`

	// Connection pool information for TCP listeners.
	// +kubebuilder:validation:Optional
	TCP []TCPParameters `json:"tcp,omitempty" tf:"tcp,omitempty"`
}

type ListenerHealthCheckObservation struct {
}

type ListenerHealthCheckParameters struct {

	// Number of consecutive successful health checks that must occur before declaring listener healthy.
	// +kubebuilder:validation:Required
	HealthyThreshold *float64 `json:"healthyThreshold" tf:"healthy_threshold,omitempty"`

	// Time period in milliseconds between each health check execution.
	// +kubebuilder:validation:Required
	IntervalMillis *float64 `json:"intervalMillis" tf:"interval_millis,omitempty"`

	// File path to write access logs to. You can use /dev/stdout to send access logs to standard out. Must be between 1 and 255 characters in length.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Port used for the port mapping.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol used for the port mapping. Valid values are http, http2, tcp and grpc.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Amount of time to wait when receiving a response from the health check, in milliseconds.
	// +kubebuilder:validation:Required
	TimeoutMillis *float64 `json:"timeoutMillis" tf:"timeout_millis,omitempty"`

	// Number of consecutive failed health checks that must occur before declaring a virtual node unhealthy.
	// +kubebuilder:validation:Required
	UnhealthyThreshold *float64 `json:"unhealthyThreshold" tf:"unhealthy_threshold,omitempty"`
}

type ListenerPortMappingObservation struct {
}

type ListenerPortMappingParameters struct {

	// Port used for the port mapping.
	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// Protocol used for the port mapping. Valid values are http, http2, tcp and grpc.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type ListenerTLSCertificateFileObservation struct {
}

type ListenerTLSCertificateFileParameters struct {

	// Certificate chain for the certificate.
	// +kubebuilder:validation:Required
	CertificateChain *string `json:"certificateChain" tf:"certificate_chain,omitempty"`

	// Private key for a certificate stored on the file system of the mesh endpoint that the proxy is running on.
	// +kubebuilder:validation:Required
	PrivateKey *string `json:"privateKey" tf:"private_key,omitempty"`
}

type ListenerTLSCertificateObservation struct {
}

type ListenerTLSCertificateParameters struct {

	// TLS validation context trust for an AWS Certificate Manager (ACM) certificate.
	// +kubebuilder:validation:Optional
	Acm []TLSCertificateAcmParameters `json:"acm,omitempty" tf:"acm,omitempty"`

	// Local file certificate.
	// +kubebuilder:validation:Optional
	File []ListenerTLSCertificateFileParameters `json:"file,omitempty" tf:"file,omitempty"`

	// A Secret Discovery Service certificate.
	// +kubebuilder:validation:Optional
	Sds []ListenerTLSCertificateSdsParameters `json:"sds,omitempty" tf:"sds,omitempty"`
}

type ListenerTLSCertificateSdsObservation struct {
}

type ListenerTLSCertificateSdsParameters struct {

	// Name of the secret secret requested from the Secret Discovery Service provider representing Transport Layer Security (TLS) materials like a certificate or certificate chain.
	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type ListenerTLSValidationObservation struct {
}

type ListenerTLSValidationParameters struct {

	// SANs for a TLS validation context.
	// +kubebuilder:validation:Optional
	SubjectAlternativeNames []ListenerTLSValidationSubjectAlternativeNamesParameters `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// TLS validation context trust.
	// +kubebuilder:validation:Required
	Trust []ListenerTLSValidationTrustParameters `json:"trust" tf:"trust,omitempty"`
}

type ListenerTLSValidationSubjectAlternativeNamesMatchObservation struct {
}

type ListenerTLSValidationSubjectAlternativeNamesMatchParameters struct {

	// Values sent must match the specified values exactly.
	// +kubebuilder:validation:Required
	Exact []*string `json:"exact" tf:"exact,omitempty"`
}

type ListenerTLSValidationSubjectAlternativeNamesObservation struct {
}

type ListenerTLSValidationSubjectAlternativeNamesParameters struct {

	// Criteria for determining a SAN's match.
	// +kubebuilder:validation:Required
	Match []ListenerTLSValidationSubjectAlternativeNamesMatchParameters `json:"match" tf:"match,omitempty"`
}

type ListenerTLSValidationTrustFileObservation struct {
}

type ListenerTLSValidationTrustFileParameters struct {

	// Certificate chain for the certificate.
	// +kubebuilder:validation:Required
	CertificateChain *string `json:"certificateChain" tf:"certificate_chain,omitempty"`
}

type ListenerTLSValidationTrustObservation struct {
}

type ListenerTLSValidationTrustParameters struct {

	// Local file certificate.
	// +kubebuilder:validation:Optional
	File []ListenerTLSValidationTrustFileParameters `json:"file,omitempty" tf:"file,omitempty"`

	// A Secret Discovery Service certificate.
	// +kubebuilder:validation:Optional
	Sds []ListenerTLSValidationTrustSdsParameters `json:"sds,omitempty" tf:"sds,omitempty"`
}

type ListenerTLSValidationTrustSdsObservation struct {
}

type ListenerTLSValidationTrustSdsParameters struct {

	// Name of the secret secret requested from the Secret Discovery Service provider representing Transport Layer Security (TLS) materials like a certificate or certificate chain.
	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type ListenerTimeoutObservation struct {
}

type ListenerTimeoutParameters struct {

	// Connection pool information for gRPC listeners.
	// +kubebuilder:validation:Optional
	GRPC []TimeoutGRPCParameters `json:"grpc,omitempty" tf:"grpc,omitempty"`

	// Connection pool information for HTTP listeners.
	// +kubebuilder:validation:Optional
	HTTP []TimeoutHTTPParameters `json:"http,omitempty" tf:"http,omitempty"`

	// Connection pool information for HTTP2 listeners.
	// +kubebuilder:validation:Optional
	Http2 []TimeoutHttp2Parameters `json:"http2,omitempty" tf:"http2,omitempty"`

	// Connection pool information for TCP listeners.
	// +kubebuilder:validation:Optional
	TCP []TimeoutTCPParameters `json:"tcp,omitempty" tf:"tcp,omitempty"`
}

type LoggingAccessLogFileObservation struct {
}

type LoggingAccessLogFileParameters struct {

	// File path to write access logs to. You can use /dev/stdout to send access logs to standard out. Must be between 1 and 255 characters in length.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type LoggingAccessLogObservation struct {
}

type LoggingAccessLogParameters struct {

	// Local file certificate.
	// +kubebuilder:validation:Optional
	File []LoggingAccessLogFileParameters `json:"file,omitempty" tf:"file,omitempty"`
}

type OutlierDetectionObservation struct {
}

type OutlierDetectionParameters struct {

	// Base amount of time for which a host is ejected.
	// +kubebuilder:validation:Required
	BaseEjectionDuration []BaseEjectionDurationParameters `json:"baseEjectionDuration" tf:"base_ejection_duration,omitempty"`

	// Time interval between ejection sweep analysis.
	// +kubebuilder:validation:Required
	Interval []IntervalParameters `json:"interval" tf:"interval,omitempty"`

	// Maximum percentage of hosts in load balancing pool for upstream service that can be ejected. Will eject at least one host regardless of the value.
	// Minimum value of 0. Maximum value of 100.
	// +kubebuilder:validation:Required
	MaxEjectionPercent *float64 `json:"maxEjectionPercent" tf:"max_ejection_percent,omitempty"`

	// Number of consecutive 5xx errors required for ejection. Minimum value of 1.
	// +kubebuilder:validation:Required
	MaxServerErrors *float64 `json:"maxServerErrors" tf:"max_server_errors,omitempty"`
}

type ServiceDiscoveryObservation struct {
}

type ServiceDiscoveryParameters struct {

	// Any AWS Cloud Map information for the virtual node.
	// +kubebuilder:validation:Optional
	AwsCloudMap []AwsCloudMapParameters `json:"awsCloudMap,omitempty" tf:"aws_cloud_map,omitempty"`

	// DNS service name for the virtual node.
	// +kubebuilder:validation:Optional
	DNS []DNSParameters `json:"dns,omitempty" tf:"dns,omitempty"`
}

type SpecBackendDefaultsObservation struct {
}

type SpecBackendDefaultsParameters struct {

	// Client policy for the backend.
	// +kubebuilder:validation:Optional
	ClientPolicy []BackendDefaultsClientPolicyParameters `json:"clientPolicy,omitempty" tf:"client_policy,omitempty"`
}

type SpecListenerObservation struct {
}

type SpecListenerParameters struct {

	// Connection pool information for the listener.
	// +kubebuilder:validation:Optional
	ConnectionPool []ListenerConnectionPoolParameters `json:"connectionPool,omitempty" tf:"connection_pool,omitempty"`

	// Health check information for the listener.
	// +kubebuilder:validation:Optional
	HealthCheck []ListenerHealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// Outlier detection information for the listener.
	// +kubebuilder:validation:Optional
	OutlierDetection []OutlierDetectionParameters `json:"outlierDetection,omitempty" tf:"outlier_detection,omitempty"`

	// Port mapping information for the listener.
	// +kubebuilder:validation:Required
	PortMapping []ListenerPortMappingParameters `json:"portMapping" tf:"port_mapping,omitempty"`

	// Transport Layer Security (TLS) client policy.
	// +kubebuilder:validation:Optional
	TLS []SpecListenerTLSParameters `json:"tls,omitempty" tf:"tls,omitempty"`

	// Timeouts for different protocols.
	// +kubebuilder:validation:Optional
	Timeout []ListenerTimeoutParameters `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type SpecListenerTLSObservation struct {
}

type SpecListenerTLSParameters struct {

	// Virtual node's client's Transport Layer Security (TLS) certificate.
	// +kubebuilder:validation:Required
	Certificate []ListenerTLSCertificateParameters `json:"certificate" tf:"certificate,omitempty"`

	// Listener's TLS mode. Valid values: DISABLED, PERMISSIVE, STRICT.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// TLS validation context.
	// +kubebuilder:validation:Optional
	Validation []ListenerTLSValidationParameters `json:"validation,omitempty" tf:"validation,omitempty"`
}

type SpecLoggingObservation struct {
}

type SpecLoggingParameters struct {

	// Access log configuration for a virtual node.
	// +kubebuilder:validation:Optional
	AccessLog []LoggingAccessLogParameters `json:"accessLog,omitempty" tf:"access_log,omitempty"`
}

type TCPIdleObservation struct {
}

type TCPIdleParameters struct {

	// Unit of time. Valid values: ms, s.
	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`

	// Number of time units. Minimum value of 0.
	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type TCPObservation struct {
}

type TCPParameters struct {

	// Maximum number of outbound TCP connections Envoy can establish concurrently with all hosts in upstream cluster. Minimum value of 1.
	// +kubebuilder:validation:Required
	MaxConnections *float64 `json:"maxConnections" tf:"max_connections,omitempty"`
}

type TLSCertificateAcmObservation struct {
}

type TLSCertificateAcmParameters struct {

	// ARN for the certificate.
	// +kubebuilder:validation:Required
	CertificateArn *string `json:"certificateArn" tf:"certificate_arn,omitempty"`
}

type TLSCertificateFileObservation struct {
}

type TLSCertificateFileParameters struct {

	// Certificate chain for the certificate.
	// +kubebuilder:validation:Required
	CertificateChain *string `json:"certificateChain" tf:"certificate_chain,omitempty"`

	// Private key for a certificate stored on the file system of the mesh endpoint that the proxy is running on.
	// +kubebuilder:validation:Required
	PrivateKey *string `json:"privateKey" tf:"private_key,omitempty"`
}

type TLSCertificateSdsObservation struct {
}

type TLSCertificateSdsParameters struct {

	// Name of the secret secret requested from the Secret Discovery Service provider representing Transport Layer Security (TLS) materials like a certificate or certificate chain.
	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type TLSValidationSubjectAlternativeNamesMatchObservation struct {
}

type TLSValidationSubjectAlternativeNamesMatchParameters struct {

	// Values sent must match the specified values exactly.
	// +kubebuilder:validation:Required
	Exact []*string `json:"exact" tf:"exact,omitempty"`
}

type TLSValidationSubjectAlternativeNamesObservation struct {
}

type TLSValidationSubjectAlternativeNamesParameters struct {

	// Criteria for determining a SAN's match.
	// +kubebuilder:validation:Required
	Match []TLSValidationSubjectAlternativeNamesMatchParameters `json:"match" tf:"match,omitempty"`
}

type TLSValidationTrustFileObservation struct {
}

type TLSValidationTrustFileParameters struct {

	// Certificate chain for the certificate.
	// +kubebuilder:validation:Required
	CertificateChain *string `json:"certificateChain" tf:"certificate_chain,omitempty"`
}

type TLSValidationTrustObservation struct {
}

type TLSValidationTrustParameters struct {

	// TLS validation context trust for an AWS Certificate Manager (ACM) certificate.
	// +kubebuilder:validation:Optional
	Acm []TrustAcmParameters `json:"acm,omitempty" tf:"acm,omitempty"`

	// Local file certificate.
	// +kubebuilder:validation:Optional
	File []TLSValidationTrustFileParameters `json:"file,omitempty" tf:"file,omitempty"`

	// A Secret Discovery Service certificate.
	// +kubebuilder:validation:Optional
	Sds []TLSValidationTrustSdsParameters `json:"sds,omitempty" tf:"sds,omitempty"`
}

type TLSValidationTrustSdsObservation struct {
}

type TLSValidationTrustSdsParameters struct {

	// Name of the secret secret requested from the Secret Discovery Service provider representing Transport Layer Security (TLS) materials like a certificate or certificate chain.
	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type TimeoutGRPCObservation struct {
}

type TimeoutGRPCParameters struct {

	// Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
	// +kubebuilder:validation:Optional
	Idle []GRPCIdleParameters `json:"idle,omitempty" tf:"idle,omitempty"`

	// Per request timeout.
	// +kubebuilder:validation:Optional
	PerRequest []GRPCPerRequestParameters `json:"perRequest,omitempty" tf:"per_request,omitempty"`
}

type TimeoutHTTPObservation struct {
}

type TimeoutHTTPParameters struct {

	// Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
	// +kubebuilder:validation:Optional
	Idle []HTTPIdleParameters `json:"idle,omitempty" tf:"idle,omitempty"`

	// Per request timeout.
	// +kubebuilder:validation:Optional
	PerRequest []HTTPPerRequestParameters `json:"perRequest,omitempty" tf:"per_request,omitempty"`
}

type TimeoutHttp2Observation struct {
}

type TimeoutHttp2Parameters struct {

	// Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
	// +kubebuilder:validation:Optional
	Idle []Http2IdleParameters `json:"idle,omitempty" tf:"idle,omitempty"`

	// Per request timeout.
	// +kubebuilder:validation:Optional
	PerRequest []Http2PerRequestParameters `json:"perRequest,omitempty" tf:"per_request,omitempty"`
}

type TimeoutTCPObservation struct {
}

type TimeoutTCPParameters struct {

	// Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
	// +kubebuilder:validation:Optional
	Idle []TCPIdleParameters `json:"idle,omitempty" tf:"idle,omitempty"`
}

type TrustAcmObservation struct {
}

type TrustAcmParameters struct {

	// One or more ACM ARNs.
	// +kubebuilder:validation:Required
	CertificateAuthorityArns []*string `json:"certificateAuthorityArns" tf:"certificate_authority_arns,omitempty"`
}

type ValidationTrustAcmObservation struct {
}

type ValidationTrustAcmParameters struct {

	// One or more ACM ARNs.
	// +kubebuilder:validation:Required
	CertificateAuthorityArns []*string `json:"certificateAuthorityArns" tf:"certificate_authority_arns,omitempty"`
}

type VirtualNodeObservation struct {

	// ARN of the virtual node.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Creation date of the virtual node.
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	// ID of the virtual node.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Last update date of the virtual node.
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date,omitempty"`

	// Resource owner's AWS account ID.
	ResourceOwner *string `json:"resourceOwner,omitempty" tf:"resource_owner,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VirtualNodeParameters struct {

	// Name of the service mesh in which to create the virtual node. Must be between 1 and 255 characters in length.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appmesh/v1beta1.Mesh
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	MeshName *string `json:"meshName,omitempty" tf:"mesh_name,omitempty"`

	// Reference to a Mesh in appmesh to populate meshName.
	// +kubebuilder:validation:Optional
	MeshNameRef *v1.Reference `json:"meshNameRef,omitempty" tf:"-"`

	// Selector for a Mesh in appmesh to populate meshName.
	// +kubebuilder:validation:Optional
	MeshNameSelector *v1.Selector `json:"meshNameSelector,omitempty" tf:"-"`

	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	// +kubebuilder:validation:Optional
	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner,omitempty"`

	// Name to use for the virtual node. Must be between 1 and 255 characters in length.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Virtual node specification to apply.
	// +kubebuilder:validation:Required
	Spec []VirtualNodeSpecParameters `json:"spec" tf:"spec,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VirtualNodeSpecObservation struct {
}

type VirtualNodeSpecParameters struct {

	// Backends to which the virtual node is expected to send outbound traffic.
	// +kubebuilder:validation:Optional
	Backend []BackendParameters `json:"backend,omitempty" tf:"backend,omitempty"`

	// Defaults for backends.
	// +kubebuilder:validation:Optional
	BackendDefaults []SpecBackendDefaultsParameters `json:"backendDefaults,omitempty" tf:"backend_defaults,omitempty"`

	// Listeners from which the virtual node is expected to receive inbound traffic.
	// +kubebuilder:validation:Optional
	Listener []SpecListenerParameters `json:"listener,omitempty" tf:"listener,omitempty"`

	// Inbound and outbound access logging information for the virtual node.
	// +kubebuilder:validation:Optional
	Logging []SpecLoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// Service discovery information for the virtual node.
	// +kubebuilder:validation:Optional
	ServiceDiscovery []ServiceDiscoveryParameters `json:"serviceDiscovery,omitempty" tf:"service_discovery,omitempty"`
}

type VirtualServiceClientPolicyObservation struct {
}

type VirtualServiceClientPolicyParameters struct {

	// Transport Layer Security (TLS) client policy.
	// +kubebuilder:validation:Optional
	TLS []ClientPolicyTLSParameters `json:"tls,omitempty" tf:"tls,omitempty"`
}

// VirtualNodeSpec defines the desired state of VirtualNode
type VirtualNodeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualNodeParameters `json:"forProvider"`
}

// VirtualNodeStatus defines the observed state of VirtualNode.
type VirtualNodeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualNodeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNode is the Schema for the VirtualNodes API. Provides an AWS App Mesh virtual node resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VirtualNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNodeSpec   `json:"spec"`
	Status            VirtualNodeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNodeList contains a list of VirtualNodes
type VirtualNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNode `json:"items"`
}

// Repository type metadata.
var (
	VirtualNode_Kind             = "VirtualNode"
	VirtualNode_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualNode_Kind}.String()
	VirtualNode_KindAPIVersion   = VirtualNode_Kind + "." + CRDGroupVersion.String()
	VirtualNode_GroupVersionKind = CRDGroupVersion.WithKind(VirtualNode_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualNode{}, &VirtualNodeList{})
}
