package azure

import "strings"

// OutboundType is a strategy for how egress from cluster is achieved.
// +kubebuilder:validation:Enum="";Loadbalancer;UserDefinedRouting
type OutboundType string

const (
	// LoadbalancerOutboundType uses Standard loadbalancer for egress from the cluster.
	// see https://docs.microsoft.com/en-us/azure/load-balancer/load-balancer-outbound-connections#lb
	LoadbalancerOutboundType OutboundType = "Loadbalancer"

	// UserDefinedRoutingOutboundType uses user defined routing for egress from the cluster.
	// see https://docs.microsoft.com/en-us/azure/virtual-network/virtual-networks-udr-overview
	UserDefinedRoutingOutboundType OutboundType = "UserDefinedRouting"
)

// Platform stores all the global configuration that all machinesets
// use.
type Platform struct {
	// Region specifies the Azure region where the cluster will be created.
	Region string `json:"region"`

	// BaseDomainResourceGroupName specifies the resource group where the Azure DNS zone for the base domain is found.
	BaseDomainResourceGroupName string `json:"baseDomainResourceGroupName,omitempty"`

	// DefaultMachinePlatform is the default configuration used when
	// installing on Azure for machine pools which do not define their own
	// platform configuration.
	// +optional
	DefaultMachinePlatform *MachinePool `json:"defaultMachinePlatform,omitempty"`

	// NetworkResourceGroupName specifies the network resource group that contains an existing VNet
	//
	// +optional
	NetworkResourceGroupName string `json:"networkResourceGroupName,omitempty"`

	// VirtualNetwork specifies the name of an existing VNet for the installer to use
	//
	// +optional
	VirtualNetwork string `json:"virtualNetwork,omitempty"`

	// ControlPlaneSubnet specifies an existing subnet for use by the control plane nodes
	//
	// +optional
	ControlPlaneSubnet string `json:"controlPlaneSubnet,omitempty"`

	// ComputeSubnet specifies an existing subnet for use by compute nodes
	//
	// +optional
	ComputeSubnet string `json:"computeSubnet,omitempty"`

	// cloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK
	// with the appropriate Azure API endpoints.
	// If empty, the value is equal to "AzurePublicCloud".
	// +optional
	CloudName CloudEnvironment `json:"cloudName,omitempty"`

	// OutboundType is a strategy for how egress from cluster is achieved. When not specified default is "Loadbalancer".
	//
	// +kubebuilder:default=Loadbalancer
	// +optional
	OutboundType OutboundType `json:"outboundType"`
}

// CloudEnvironment is the name of the Azure cloud environment
// +kubebuilder:validation:Enum="";AzurePublicCloud;AzureUSGovernmentCloud;AzureChinaCloud;AzureGermanCloud
type CloudEnvironment string

const (
	// PublicCloud is the general-purpose, public Azure cloud environment.
	PublicCloud CloudEnvironment = "AzurePublicCloud"

	// USGovernmentCloud is the Azure cloud environment for the US government.
	USGovernmentCloud CloudEnvironment = "AzureUSGovernmentCloud"

	// ChinaCloud is the Azure cloud environment used in China.
	ChinaCloud CloudEnvironment = "AzureChinaCloud"

	// GermanCloud is the Azure cloud environment used in Germany.
	GermanCloud CloudEnvironment = "AzureGermanCloud"
)

// Name returns name that Azure uses for the cloud environment.
// See https://github.com/Azure/go-autorest/blob/ec5f4903f77ed9927ac95b19ab8e44ada64c1356/autorest/azure/environments.go#L13
func (e CloudEnvironment) Name() string {
	return string(e)
}

//SetBaseDomain parses the baseDomainID and sets the related fields on azure.Platform
func (p *Platform) SetBaseDomain(baseDomainID string) error {
	parts := strings.Split(baseDomainID, "/")
	p.BaseDomainResourceGroupName = parts[4]
	return nil
}
