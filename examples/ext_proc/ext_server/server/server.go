package server

import (
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/service/ext_proc/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
)

type ExternalProcessingServer struct {
	v3.UnimplementedExternalProcessorServer
}

func NewExternalProcessingServer() *ExternalProcessingServer {
	return &ExternalProcessingServer{}
}

// Process envoy External Processing filter implementation
// https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/ext_proc_filter
// add grey headers to affect route
func (s *ExternalProcessingServer) Process(srv v3.ExternalProcessor_ProcessServer) error {
	ctx := srv.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive stream request: %v", err)
		}
		resp := &v3.ProcessingResponse{}
		switch req.Request.(type) {
		case *v3.ProcessingRequest_RequestHeaders:
			var retHeaders []*corev3.HeaderValueOption

			isGrey := false

			headers := req.GetRequestHeaders().Headers.Headers
			for _, header := range headers {
				if header.Key == "grey" {
					isGrey = true
					break
				}
			}

			if isGrey {
				retHeaders = append(retHeaders, &corev3.HeaderValueOption{
					Header: &corev3.HeaderValue{
						Key:   "grey-version",
						Value: "localhost:8080=localhost-feature-test:8081",
					},
					Append: wrapperspb.Bool(false),
				})

				retHeaders = append(retHeaders, &corev3.HeaderValueOption{
					Header: &corev3.HeaderValue{
						Key:   "grey-version-localhost:8080",
						Value: "grey-version-localhost-feature-test:8081",
					},
					Append: wrapperspb.Bool(false),
				})
			}

			resp = &v3.ProcessingResponse{
				Response: &v3.ProcessingResponse_RequestHeaders{
					RequestHeaders: &v3.HeadersResponse{
						Response: &v3.CommonResponse{
							ClearRouteCache: true, // must be true to affect route
							HeaderMutation: &v3.HeaderMutation{
								SetHeaders: retHeaders,
							},
						},
					},
				},
			}
		default:
		}
		if err := srv.Send(resp); err != nil {
			return status.Errorf(codes.Unknown, "cannot send stream response: %v", err)
		}
	}
}
