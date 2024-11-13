// Code generated by go-swagger; DO NOT EDIT.

package internal_operations_pvm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new internal operations pvm instances API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new internal operations pvm instances API client with basic auth credentials.
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

// New creates a new internal operations pvm instances API client with a bearer token for authentication.
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
Client for internal operations pvm instances API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	InternalV1OperationsPvminstancesDelete(params *InternalV1OperationsPvminstancesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1OperationsPvminstancesDeleteNoContent, error)

	InternalV1OperationsPvminstancesPost(params *InternalV1OperationsPvminstancesPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1OperationsPvminstancesPostCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
InternalV1OperationsPvminstancesDelete deletes a p VM instance c r n
*/
func (a *Client) InternalV1OperationsPvminstancesDelete(params *InternalV1OperationsPvminstancesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1OperationsPvminstancesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalV1OperationsPvminstancesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "internal.v1.operations.pvminstances.delete",
		Method:             "DELETE",
		PathPattern:        "/internal/v1/operations/pvm-instances/{resource_crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InternalV1OperationsPvminstancesDeleteReader{formats: a.formats},
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
	success, ok := result.(*InternalV1OperationsPvminstancesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for internal.v1.operations.pvminstances.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InternalV1OperationsPvminstancesPost creates a c r n for a p VM instance
*/
func (a *Client) InternalV1OperationsPvminstancesPost(params *InternalV1OperationsPvminstancesPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1OperationsPvminstancesPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalV1OperationsPvminstancesPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "internal.v1.operations.pvminstances.post",
		Method:             "POST",
		PathPattern:        "/internal/v1/operations/pvm-instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InternalV1OperationsPvminstancesPostReader{formats: a.formats},
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
	success, ok := result.(*InternalV1OperationsPvminstancesPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for internal.v1.operations.pvminstances.post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
