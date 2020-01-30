// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ory/kratos-client-go/client/admin"
	"github.com/ory/kratos-client-go/client/common"
	"github.com/ory/kratos-client-go/client/health"
	"github.com/ory/kratos-client-go/client/public"
	"github.com/ory/kratos-client-go/client/version"
)

// Default ory kratos HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new ory kratos HTTP client.
func NewHTTPClient(formats strfmt.Registry) *OryKratos {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new ory kratos HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *OryKratos {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new ory kratos client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *OryKratos {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(OryKratos)
	cli.Transport = transport

	cli.Admin = admin.New(transport, formats)

	cli.Common = common.New(transport, formats)

	cli.Health = health.New(transport, formats)

	cli.Public = public.New(transport, formats)

	cli.Version = version.New(transport, formats)

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

// OryKratos is a client for ory kratos
type OryKratos struct {
	Admin *admin.Client

	Common *common.Client

	Health *health.Client

	Public *public.Client

	Version *version.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *OryKratos) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Admin.SetTransport(transport)

	c.Common.SetTransport(transport)

	c.Health.SetTransport(transport)

	c.Public.SetTransport(transport)

	c.Version.SetTransport(transport)

}