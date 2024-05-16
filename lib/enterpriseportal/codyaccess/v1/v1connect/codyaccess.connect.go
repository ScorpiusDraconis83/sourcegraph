// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: codyaccess.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/sourcegraph/sourcegraph/lib/enterpriseportal/codyaccess/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// CodyAccessServiceName is the fully-qualified name of the CodyAccessService service.
	CodyAccessServiceName = "enterpriseportal.codyaccess.v1.CodyAccessService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CodyAccessServiceGetCodyGatewayAccessProcedure is the fully-qualified name of the
	// CodyAccessService's GetCodyGatewayAccess RPC.
	CodyAccessServiceGetCodyGatewayAccessProcedure = "/enterpriseportal.codyaccess.v1.CodyAccessService/GetCodyGatewayAccess"
	// CodyAccessServiceListCodyGatewayAccessesProcedure is the fully-qualified name of the
	// CodyAccessService's ListCodyGatewayAccesses RPC.
	CodyAccessServiceListCodyGatewayAccessesProcedure = "/enterpriseportal.codyaccess.v1.CodyAccessService/ListCodyGatewayAccesses"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	codyAccessServiceServiceDescriptor                       = v1.File_codyaccess_proto.Services().ByName("CodyAccessService")
	codyAccessServiceGetCodyGatewayAccessMethodDescriptor    = codyAccessServiceServiceDescriptor.Methods().ByName("GetCodyGatewayAccess")
	codyAccessServiceListCodyGatewayAccessesMethodDescriptor = codyAccessServiceServiceDescriptor.Methods().ByName("ListCodyGatewayAccesses")
)

// CodyAccessServiceClient is a client for the enterpriseportal.codyaccess.v1.CodyAccessService
// service.
type CodyAccessServiceClient interface {
	// Retrieve Cody Gateway access granted to an Enterprise subscription.
	GetCodyGatewayAccess(context.Context, *connect.Request[v1.GetCodyGatewayAccessRequest]) (*connect.Response[v1.GetCodyGatewayAccessResponse], error)
	// List all Cody Gateway accesses granted to any Enterprise subscription.
	ListCodyGatewayAccesses(context.Context, *connect.Request[v1.ListCodyGatewayAccessesRequest]) (*connect.Response[v1.ListCodyGatewayAccessesResponse], error)
}

// NewCodyAccessServiceClient constructs a client for the
// enterpriseportal.codyaccess.v1.CodyAccessService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCodyAccessServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) CodyAccessServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &codyAccessServiceClient{
		getCodyGatewayAccess: connect.NewClient[v1.GetCodyGatewayAccessRequest, v1.GetCodyGatewayAccessResponse](
			httpClient,
			baseURL+CodyAccessServiceGetCodyGatewayAccessProcedure,
			connect.WithSchema(codyAccessServiceGetCodyGatewayAccessMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		listCodyGatewayAccesses: connect.NewClient[v1.ListCodyGatewayAccessesRequest, v1.ListCodyGatewayAccessesResponse](
			httpClient,
			baseURL+CodyAccessServiceListCodyGatewayAccessesProcedure,
			connect.WithSchema(codyAccessServiceListCodyGatewayAccessesMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// codyAccessServiceClient implements CodyAccessServiceClient.
type codyAccessServiceClient struct {
	getCodyGatewayAccess    *connect.Client[v1.GetCodyGatewayAccessRequest, v1.GetCodyGatewayAccessResponse]
	listCodyGatewayAccesses *connect.Client[v1.ListCodyGatewayAccessesRequest, v1.ListCodyGatewayAccessesResponse]
}

// GetCodyGatewayAccess calls enterpriseportal.codyaccess.v1.CodyAccessService.GetCodyGatewayAccess.
func (c *codyAccessServiceClient) GetCodyGatewayAccess(ctx context.Context, req *connect.Request[v1.GetCodyGatewayAccessRequest]) (*connect.Response[v1.GetCodyGatewayAccessResponse], error) {
	return c.getCodyGatewayAccess.CallUnary(ctx, req)
}

// ListCodyGatewayAccesses calls
// enterpriseportal.codyaccess.v1.CodyAccessService.ListCodyGatewayAccesses.
func (c *codyAccessServiceClient) ListCodyGatewayAccesses(ctx context.Context, req *connect.Request[v1.ListCodyGatewayAccessesRequest]) (*connect.Response[v1.ListCodyGatewayAccessesResponse], error) {
	return c.listCodyGatewayAccesses.CallUnary(ctx, req)
}

// CodyAccessServiceHandler is an implementation of the
// enterpriseportal.codyaccess.v1.CodyAccessService service.
type CodyAccessServiceHandler interface {
	// Retrieve Cody Gateway access granted to an Enterprise subscription.
	GetCodyGatewayAccess(context.Context, *connect.Request[v1.GetCodyGatewayAccessRequest]) (*connect.Response[v1.GetCodyGatewayAccessResponse], error)
	// List all Cody Gateway accesses granted to any Enterprise subscription.
	ListCodyGatewayAccesses(context.Context, *connect.Request[v1.ListCodyGatewayAccessesRequest]) (*connect.Response[v1.ListCodyGatewayAccessesResponse], error)
}

// NewCodyAccessServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCodyAccessServiceHandler(svc CodyAccessServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	codyAccessServiceGetCodyGatewayAccessHandler := connect.NewUnaryHandler(
		CodyAccessServiceGetCodyGatewayAccessProcedure,
		svc.GetCodyGatewayAccess,
		connect.WithSchema(codyAccessServiceGetCodyGatewayAccessMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	codyAccessServiceListCodyGatewayAccessesHandler := connect.NewUnaryHandler(
		CodyAccessServiceListCodyGatewayAccessesProcedure,
		svc.ListCodyGatewayAccesses,
		connect.WithSchema(codyAccessServiceListCodyGatewayAccessesMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/enterpriseportal.codyaccess.v1.CodyAccessService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CodyAccessServiceGetCodyGatewayAccessProcedure:
			codyAccessServiceGetCodyGatewayAccessHandler.ServeHTTP(w, r)
		case CodyAccessServiceListCodyGatewayAccessesProcedure:
			codyAccessServiceListCodyGatewayAccessesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCodyAccessServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCodyAccessServiceHandler struct{}

func (UnimplementedCodyAccessServiceHandler) GetCodyGatewayAccess(context.Context, *connect.Request[v1.GetCodyGatewayAccessRequest]) (*connect.Response[v1.GetCodyGatewayAccessResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("enterpriseportal.codyaccess.v1.CodyAccessService.GetCodyGatewayAccess is not implemented"))
}

func (UnimplementedCodyAccessServiceHandler) ListCodyGatewayAccesses(context.Context, *connect.Request[v1.ListCodyGatewayAccessesRequest]) (*connect.Response[v1.ListCodyGatewayAccessesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("enterpriseportal.codyaccess.v1.CodyAccessService.ListCodyGatewayAccesses is not implemented"))
}