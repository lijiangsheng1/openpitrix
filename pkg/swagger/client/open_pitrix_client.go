// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/pkg/swagger/client/app_runtimes"
	"openpitrix.io/openpitrix/pkg/swagger/client/apps"
	"openpitrix.io/openpitrix/pkg/swagger/client/clusters"
	"openpitrix.io/openpitrix/pkg/swagger/client/repos"
)

// Default open pitrix HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new open pitrix HTTP client.
func NewHTTPClient(formats strfmt.Registry) *OpenPitrix {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new open pitrix HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *OpenPitrix {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new open pitrix client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *OpenPitrix {
	cli := new(OpenPitrix)
	cli.Transport = transport

	cli.AppRuntimes = app_runtimes.New(transport, formats)

	cli.Apps = apps.New(transport, formats)

	cli.Clusters = clusters.New(transport, formats)

	cli.Repos = repos.New(transport, formats)

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

// OpenPitrix is a client for open pitrix
type OpenPitrix struct {
	AppRuntimes *app_runtimes.Client

	Apps *apps.Client

	Clusters *clusters.Client

	Repos *repos.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *OpenPitrix) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.AppRuntimes.SetTransport(transport)

	c.Apps.SetTransport(transport)

	c.Clusters.SetTransport(transport)

	c.Repos.SetTransport(transport)

}
