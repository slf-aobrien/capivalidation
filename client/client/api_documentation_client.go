// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"client/client/admin"
	"client/client/agency"
	"client/client/agent"
	"client/client/basic_error_controller"
	"client/client/benefit"
	"client/client/bill"
	"client/client/c_b_s"
	"client/client/cache"
	"client/client/dependent"
	"client/client/error"
	"client/client/focus"
	"client/client/member"
	"client/client/organization"
	"client/client/person"
	"client/client/policy"
	"client/client/policy_bill_group"
	"client/client/sunbeam"
	"client/client/system"
)

// Default API documentation HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "dev-corp-wk-ca1-k8s.sunlifecorp.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/sit-compass-integration-ns"
	//DefaultBasePath string = "/dev-compass-integration-ns"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new API documentation HTTP client.
func NewHTTPClient(formats strfmt.Registry) *APIDocumentation {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new API documentation HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *APIDocumentation {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new API documentation client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *APIDocumentation {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(APIDocumentation)
	cli.Transport = transport
	cli.Admin = admin.New(transport, formats)
	cli.Agency = agency.New(transport, formats)
	cli.Agent = agent.New(transport, formats)
	cli.BasicErrorController = basic_error_controller.New(transport, formats)
	cli.Benefit = benefit.New(transport, formats)
	cli.Bill = bill.New(transport, formats)
	cli.Cbs = c_b_s.New(transport, formats)
	cli.Cache = cache.New(transport, formats)
	cli.Dependent = dependent.New(transport, formats)
	cli.Error = error.New(transport, formats)
	cli.Focus = focus.New(transport, formats)
	cli.Member = member.New(transport, formats)
	cli.Organization = organization.New(transport, formats)
	cli.Person = person.New(transport, formats)
	cli.Policy = policy.New(transport, formats)
	cli.PolicyBillGroup = policy_bill_group.New(transport, formats)
	cli.Sunbeam = sunbeam.New(transport, formats)
	cli.System = system.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// APIDocumentation is a client for API documentation
type APIDocumentation struct {
	Admin admin.ClientService

	Agency agency.ClientService

	Agent agent.ClientService

	BasicErrorController basic_error_controller.ClientService

	Benefit benefit.ClientService

	Bill bill.ClientService

	Cbs c_b_s.ClientService

	Cache cache.ClientService

	Dependent dependent.ClientService

	Error error.ClientService

	Focus focus.ClientService

	Member member.ClientService

	Organization organization.ClientService

	Person person.ClientService

	Policy policy.ClientService

	PolicyBillGroup policy_bill_group.ClientService

	Sunbeam sunbeam.ClientService

	System system.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *APIDocumentation) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Admin.SetTransport(transport)
	c.Agency.SetTransport(transport)
	c.Agent.SetTransport(transport)
	c.BasicErrorController.SetTransport(transport)
	c.Benefit.SetTransport(transport)
	c.Bill.SetTransport(transport)
	c.Cbs.SetTransport(transport)
	c.Cache.SetTransport(transport)
	c.Dependent.SetTransport(transport)
	c.Error.SetTransport(transport)
	c.Focus.SetTransport(transport)
	c.Member.SetTransport(transport)
	c.Organization.SetTransport(transport)
	c.Person.SetTransport(transport)
	c.Policy.SetTransport(transport)
	c.PolicyBillGroup.SetTransport(transport)
	c.Sunbeam.SetTransport(transport)
	c.System.SetTransport(transport)
}
