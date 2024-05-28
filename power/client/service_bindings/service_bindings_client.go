// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new service bindings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new service bindings API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new service bindings API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for service bindings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ServiceBindingBinding(params *ServiceBindingBindingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingBindingOK, *ServiceBindingBindingCreated, *ServiceBindingBindingAccepted, error)

	ServiceBindingGet(params *ServiceBindingGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingGetOK, error)

	ServiceBindingLastOperationGet(params *ServiceBindingLastOperationGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingLastOperationGetOK, error)

	ServiceBindingUnbinding(params *ServiceBindingUnbindingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingUnbindingOK, *ServiceBindingUnbindingAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ServiceBindingBinding generations of a service binding
*/
func (a *Client) ServiceBindingBinding(params *ServiceBindingBindingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingBindingOK, *ServiceBindingBindingCreated, *ServiceBindingBindingAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingBindingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceBinding.binding",
		Method:             "PUT",
		PathPattern:        "/v2/service_instances/{instance_id}/service_bindings/{binding_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBindingBindingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *ServiceBindingBindingOK:
		return value, nil, nil, nil
	case *ServiceBindingBindingCreated:
		return nil, value, nil, nil
	case *ServiceBindingBindingAccepted:
		return nil, nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for service_bindings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ServiceBindingGet gets a service binding
*/
func (a *Client) ServiceBindingGet(params *ServiceBindingGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceBinding.get",
		Method:             "GET",
		PathPattern:        "/v2/service_instances/{instance_id}/service_bindings/{binding_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBindingGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServiceBindingGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serviceBinding.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ServiceBindingLastOperationGet lasts requested operation state for service binding
*/
func (a *Client) ServiceBindingLastOperationGet(params *ServiceBindingLastOperationGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingLastOperationGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingLastOperationGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceBinding.lastOperation.get",
		Method:             "GET",
		PathPattern:        "/v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBindingLastOperationGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServiceBindingLastOperationGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serviceBinding.lastOperation.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ServiceBindingUnbinding deprovisions of a service binding
*/
func (a *Client) ServiceBindingUnbinding(params *ServiceBindingUnbindingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingUnbindingOK, *ServiceBindingUnbindingAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingUnbindingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceBinding.unbinding",
		Method:             "DELETE",
		PathPattern:        "/v2/service_instances/{instance_id}/service_bindings/{binding_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBindingUnbindingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ServiceBindingUnbindingOK:
		return value, nil, nil
	case *ServiceBindingUnbindingAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for service_bindings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
