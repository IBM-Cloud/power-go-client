// Code generated by go-swagger; DO NOT EDIT.

package internal_operations_network_security_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new internal operations network security groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new internal operations network security groups API client with basic auth credentials.
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

// New creates a new internal operations network security groups API client with a bearer token for authentication.
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
Client for internal operations network security groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	InternalV1OperationsNetworksecuritygroupsDelete(params *InternalV1OperationsNetworksecuritygroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1OperationsNetworksecuritygroupsDeleteNoContent, error)

	InternalV1OperationsNetworksecuritygroupsPost(params *InternalV1OperationsNetworksecuritygroupsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1OperationsNetworksecuritygroupsPostCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
InternalV1OperationsNetworksecuritygroupsDelete deletes a network security group c r n
*/
func (a *Client) InternalV1OperationsNetworksecuritygroupsDelete(params *InternalV1OperationsNetworksecuritygroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1OperationsNetworksecuritygroupsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalV1OperationsNetworksecuritygroupsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "internal.v1.operations.networksecuritygroups.delete",
		Method:             "DELETE",
		PathPattern:        "/internal/v1/operations/network-security-groups/{resource_crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InternalV1OperationsNetworksecuritygroupsDeleteReader{formats: a.formats},
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
	success, ok := result.(*InternalV1OperationsNetworksecuritygroupsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for internal.v1.operations.networksecuritygroups.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InternalV1OperationsNetworksecuritygroupsPost creates a c r n for a network security group
*/
func (a *Client) InternalV1OperationsNetworksecuritygroupsPost(params *InternalV1OperationsNetworksecuritygroupsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1OperationsNetworksecuritygroupsPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalV1OperationsNetworksecuritygroupsPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "internal.v1.operations.networksecuritygroups.post",
		Method:             "POST",
		PathPattern:        "/internal/v1/operations/network-security-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InternalV1OperationsNetworksecuritygroupsPostReader{formats: a.formats},
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
	success, ok := result.(*InternalV1OperationsNetworksecuritygroupsPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for internal.v1.operations.networksecuritygroups.post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}