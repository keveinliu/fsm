package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MeshConfig is the type used to represent the mesh configuration.
// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:deprecatedversion
type MeshConfig struct {
	// Object's type metadata.
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Object's metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Spec is the MeshConfig specification.
	// +optional
	Spec MeshConfigSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

// MeshConfigSpec is the spec for FSM's configuration.
type MeshConfigSpec struct {
	// ClusterSetSpec defines the configurations of cluster.
	ClusterSet ClusterSetSpec `json:"clusterSet,omitempty"`

	// Sidecar defines the configurations of the proxy sidecar in a mesh.
	Sidecar SidecarSpec `json:"sidecar,omitempty"`

	// RepoServer defines the configurations of pipy repo server.
	RepoServer RepoServerSpec `json:"repoServer,omitempty"`

	// Traffic defines the traffic management configurations for a mesh instance.
	Traffic TrafficSpec `json:"traffic,omitempty"`

	// Observalility defines the observability configurations for a mesh instance.
	Observability ObservabilitySpec `json:"observability,omitempty"`

	// Certificate defines the certificate management configurations for a mesh instance.
	Certificate CertificateSpec `json:"certificate,omitempty"`

	// FeatureFlags defines the feature flags for a mesh instance.
	FeatureFlags FeatureFlags `json:"featureFlags,omitempty"`

	// PluginChains defines the default plugin chains.
	PluginChains PluginChainsSpec `json:"pluginChains,omitempty"`
}

// SidecarSpec is the type used to represent the specifications for the proxy sidecar.
type SidecarSpec struct {
	// EnablePrivilegedInitContainer defines a boolean indicating whether the init container for a meshed pod should run as privileged.
	EnablePrivilegedInitContainer bool `json:"enablePrivilegedInitContainer"`

	// LogLevel defines the logging level for the sidecar's logs. Non developers should generally never set this value. In production environments the LogLevel should be set to error.
	LogLevel string `json:"logLevel,omitempty"`

	// SidecarClass defines the container provider used for the proxy sidecar.
	SidecarClass string `json:"sidecarClass,omitempty"`

	// SidecarImage defines the container image used for the proxy sidecar.
	SidecarImage string `json:"sidecarImage,omitempty"`

	// SidecarDisabledMTLS defines whether mTLS is disabled.
	SidecarDisabledMTLS bool `json:"sidecarDisabledMTLS"`

	// InitContainerImage defines the container image used for the init container injected to meshed pods.
	InitContainerImage string `json:"initContainerImage,omitempty"`

	// SidecarDrivers defines the sidecar supported.
	SidecarDrivers []SidecarDriverSpec `json:"sidecarDrivers,omitempty"`

	// MaxDataPlaneConnections defines the maximum allowed data plane connections from a proxy sidecar to the FSM controller.
	MaxDataPlaneConnections int `json:"maxDataPlaneConnections,omitempty"`

	// ConfigResyncInterval defines the resync interval for regular proxy broadcast updates.
	ConfigResyncInterval string `json:"configResyncInterval,omitempty"`

	// SidecarTimeout defines the connect/idle/read/write timeout.
	SidecarTimeout int `json:"sidecarTimeout,omitempty"`

	// Resources defines the compute resources for the sidecar.
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
}

// TrafficSpec is the type used to represent FSM's traffic management configuration.
type TrafficSpec struct {
	// InterceptionMode defines a string indicating which traffic interception mode is used.
	InterceptionMode string `json:"interceptionMode"`

	// EnableEgress defines a boolean indicating if mesh-wide Egress is enabled.
	EnableEgress bool `json:"enableEgress"`

	// OutboundIPRangeExclusionList defines a global list of IP address ranges to exclude from outbound traffic interception by the sidecar proxy.
	OutboundIPRangeExclusionList []string `json:"outboundIPRangeExclusionList"`

	// OutboundPortExclusionList defines a global list of ports to exclude from outbound traffic interception by the sidecar proxy.
	OutboundPortExclusionList []int `json:"outboundPortExclusionList"`

	// InboundPortExclusionList defines a global list of ports to exclude from inbound traffic interception by the sidecar proxy.
	InboundPortExclusionList []int `json:"inboundPortExclusionList"`

	// EnablePermissiveTrafficPolicyMode defines a boolean indicating if permissive traffic policy mode is enabled mesh-wide.
	EnablePermissiveTrafficPolicyMode bool `json:"enablePermissiveTrafficPolicyMode"`

	// ServiceAccessMode defines a string indicating service access mode.
	ServiceAccessMode string `json:"serviceAccessMode"`

	// InboundExternalAuthorization defines a ruleset that, if enabled, will configure a remote external authorization endpoint
	// for all inbound and ingress traffic in the mesh.
	InboundExternalAuthorization ExternalAuthzSpec `json:"inboundExternalAuthorization,omitempty"`

	// HTTP1PerRequestLoadBalancing defines a boolean indicating if load balancing based on request is enabled for http1.
	HTTP1PerRequestLoadBalancing bool `json:"http1PerRequestLoadBalancing"`

	// HTTP1PerRequestLoadBalancing defines a boolean indicating if load balancing based on request is enabled for http2.
	HTTP2PerRequestLoadBalancing bool `json:"http2PerRequestLoadBalancing"`
}

// ObservabilitySpec is the type to represent FSM's observability configurations.
type ObservabilitySpec struct {
	// FSMLogLevel defines the log level for FSM control plane logs.
	FSMLogLevel string `json:"fsmLogLevel,omitempty"`

	// Tracing defines FSM's tracing configuration.
	Tracing TracingSpec `json:"tracing,omitempty"`

	// RemoteLogging defines FSM's remot logging configuration.
	RemoteLogging RemoteLoggingSpec `json:"remoteLogging,omitempty"`
}

// TracingSpec is the type to represent FSM's tracing configuration.
type TracingSpec struct {
	// Enable defines a boolean indicating if the sidecars are enabled for tracing.
	Enable bool `json:"enable"`

	// Port defines the tracing collector's port.
	Port uint16 `json:"port,omitempty"`

	// Address defines the tracing collectio's hostname.
	Address string `json:"address,omitempty"`

	// Endpoint defines the API endpoint for tracing requests sent to the collector.
	Endpoint string `json:"endpoint,omitempty"`

	// SampledFraction defines the sampled fraction.
	SampledFraction *float32 `json:"sampledFraction,omitempty"`
}

// RemoteLoggingSpec is the type to represent FSM's remote logging configuration.
type RemoteLoggingSpec struct {
	// Enable defines a boolean indicating if the sidecars are enabled for remote logging.
	Enable bool `json:"enable"`

	// Level defines the remote logging's level.
	Level uint16 `json:"level,omitempty"`

	// Port defines the remote loggings port.
	Port uint16 `json:"port,omitempty"`

	// Address defines the remote logging's hostname.
	Address string `json:"address,omitempty"`

	// Endpoint defines the API endpoint for remote logging requests sent to the collector.
	Endpoint string `json:"endpoint,omitempty"`

	// Authorization defines the access entity that allows to authorize someone in remote logging service.
	Authorization string `json:"authorization,omitempty"`

	// SampledFraction defines the sampled fraction.
	SampledFraction *float32 `json:"sampledFraction,omitempty"`
}

// ExternalAuthzSpec is a type to represent external authorization configuration.
type ExternalAuthzSpec struct {
	// Enable defines a boolean indicating if the external authorization policy is to be enabled.
	Enable bool `json:"enable"`

	// Address defines the remote address of the external authorization endpoint.
	Address string `json:"address,omitempty"`

	// Port defines the destination port of the remote external authorization endpoint.
	Port uint16 `json:"port,omitempty"`

	// StatPrefix defines a prefix for the stats sink for this external authorization policy.
	StatPrefix string `json:"statPrefix,omitempty"`

	// Timeout defines the timeout in which a response from the external authorization endpoint.
	// is expected to execute.
	Timeout string `json:"timeout,omitempty"`

	// FailureModeAllow defines a boolean indicating if traffic should be allowed on a failure to get a
	// response against the external authorization endpoint.
	FailureModeAllow bool `json:"failureModeAllow"`
}

// CertificateSpec is the type to reperesent FSM's certificate management configuration.
type CertificateSpec struct {
	// ServiceCertValidityDuration defines the service certificate validity duration.
	ServiceCertValidityDuration string `json:"serviceCertValidityDuration,omitempty"`

	// CertKeyBitSize defines the certicate key bit size.
	CertKeyBitSize int `json:"certKeyBitSize,omitempty"`

	// IngressGateway defines the certificate specification for an ingress gateway.
	// +optional
	IngressGateway *IngressGatewayCertSpec `json:"ingressGateway,omitempty"`
}

// IngressGatewayCertSpec is the type to represent the certificate specification for an ingress gateway.
type IngressGatewayCertSpec struct {
	// SubjectAltNames defines the Subject Alternative Names (domain names and IP addresses) secured by the certificate.
	SubjectAltNames []string `json:"subjectAltNames"`

	// ValidityDuration defines the validity duration of the certificate.
	ValidityDuration string `json:"validityDuration"`

	// Secret defines the secret in which the certificate is stored.
	Secret corev1.SecretReference `json:"secret"`
}

// MeshConfigList lists the MeshConfig objects.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MeshConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MeshConfig `json:"items"`
}

// FeatureFlags is a type to represent FSM's feature flags.
type FeatureFlags struct {
	// EnableEgressPolicy defines if FSM's Egress policy is enabled.
	EnableEgressPolicy bool `json:"enableEgressPolicy"`

	// EnableSnapshotCacheMode defines if XDS server starts with snapshot cache.
	EnableSnapshotCacheMode bool `json:"enableSnapshotCacheMode"`

	//EnableAsyncProxyServiceMapping defines if FSM will map proxies to services asynchronously.
	EnableAsyncProxyServiceMapping bool `json:"enableAsyncProxyServiceMapping"`

	// EnableIngressBackendPolicy defines if FSM will use the IngressBackend API to allow ingress traffic to
	// service mesh backends.
	EnableIngressBackendPolicy bool `json:"enableIngressBackendPolicy"`

	// EnableAccessControlPolicy defines if FSM will use the AccessControl API to allow access control traffic to
	// service mesh backends.
	EnableAccessControlPolicy bool `json:"enableAccessControlPolicy"`

	// EnableAccessCertPolicy defines if FSM can issue certificates for external services..
	EnableAccessCertPolicy bool `json:"enableAccessCertPolicy"`

	// EnableSidecarActiveHealthChecks defines if FSM will sidecar active health
	// checks between services allowed to communicate.
	EnableSidecarActiveHealthChecks bool `json:"enableSidecarActiveHealthChecks"`

	// EnableRetryPolicy defines if retry policy is enabled.
	EnableRetryPolicy bool `json:"enableRetryPolicy"`

	// EnablePluginPolicy defines if plugin policy is enabled.
	EnablePluginPolicy bool `json:"enablePluginPolicy"`

	// EnableAutoDefaultRoute defines if auto default route is enabled.
	EnableAutoDefaultRoute bool `json:"enableAutoDefaultRoute"`
}

// SidecarDriverSpec is the type to represent FSM's sidecar driver define.
type SidecarDriverSpec struct {
	// SidecarName defines the name of the sidecar driver.
	SidecarName string `json:"sidecarName,omitempty"`

	// SidecarImage defines the container image used for the proxy sidecar.
	SidecarImage string `json:"sidecarImage,omitempty"`

	// InitContainerImage defines the container image used for the init container injected to meshed pods.
	InitContainerImage string `json:"initContainerImage,omitempty"`

	// ProxyServerPort is the port on which the Discovery Service listens for new connections from Sidecars
	ProxyServerPort uint32 `json:"proxyServerPort"`

	// SidecarDisabledMTLS defines if mTLS are disabled.
	SidecarDisabledMTLS bool `json:"sidecarDisabledMTLS"`
}

// RepoServerSpec is the type to represent repo server.
type RepoServerSpec struct {
	// IPAddr of the pipy repo server
	IPAddr string `json:"ipaddr"`

	// Codebase is the folder used by fsmController
	Codebase string `json:"codebase"`
}

// ClusterPropertySpec is the type to represent cluster property.
type ClusterPropertySpec struct {
	// Name defines the name of cluster property.
	Name string `json:"name"`

	// Value defines the name of cluster property.
	Value string `json:"value"`
}

// ClusterSetSpec is the type to represent cluster set.
type ClusterSetSpec struct {
	// Properties defines properties for cluster.
	Properties []ClusterPropertySpec `json:"properties"`
}

// PluginChainsSpec is the type to represent plugin chains.
type PluginChainsSpec struct {
	// InboundTCPChains defines inbound tcp chains
	InboundTCPChains []*PluginChainSpec `json:"inbound-tcp"`

	// InboundHTTPChains defines inbound http chains
	InboundHTTPChains []*PluginChainSpec `json:"inbound-http"`

	// OutboundTCPChains defines outbound tcp chains
	OutboundTCPChains []*PluginChainSpec `json:"outbound-tcp"`

	// OutboundHTTPChains defines outbound http chains
	OutboundHTTPChains []*PluginChainSpec `json:"outbound-http"`
}

// PluginChainSpec is the type to represent plugin chain.
type PluginChainSpec struct {
	// Plugin defines the name of plugin
	Plugin string `json:"plugin"`

	// Priority defines the priority of plugin
	Priority float32 `json:"priority"`

	// Disable defines the visibility of plugin
	Disable bool `json:"disable"`
}
