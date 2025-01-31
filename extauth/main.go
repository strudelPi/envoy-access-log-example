package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"net"
	"strings"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	auth_pb "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"github.com/gogo/googleapis/google/rpc"
	typev3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
)

// struct with check method
type AuthServer struct{}

func (server *AuthServer) Check(
	ctx context.Context,
	request *auth_pb.CheckRequest,
) (*auth_pb.CheckResponse, error) {
	authHeader, ok := request.Attributes.Request.Http.Headers["authorization"]
	var splitToken []string
	if ok {
		splitToken = strings.Split(authHeader, "Bearer ")

		// Normally this is where you'd go check with the system that knows if it's a valid token.
		if len(splitToken) == 2 {
			token := splitToken[1]
			sha := sha256.New()
			sha.Write([]byte(token))
			tokenSha := base64.StdEncoding.EncodeToString(sha.Sum(nil))

			// valid tokens have exactly 3 characters. #secure.
			if len(token) == 3 {
				return &auth_pb.CheckResponse{
					Status: &status.Status{
						Code: int32(rpc.OK),
					},
					HttpResponse: &auth_pb.CheckResponse_OkResponse{
						OkResponse: &auth_pb.OkHttpResponse{
							Headers: []*corev3.HeaderValueOption{
								{
									Header: &corev3.HeaderValue{
										Key:   "x-ext-auth-ratelimit",
										Value: tokenSha,
									},
								},
							},
						},
					},
				}, nil
			}
		}
	}
	return &auth_pb.CheckResponse{
		Status: &status.Status{
			Code: int32(rpc.UNAUTHENTICATED),
		},
		HttpResponse: &auth_pb.CheckResponse_DeniedResponse{
			DeniedResponse: &auth_pb.DeniedHttpResponse{
				Status: &typev3.HttpStatus{
					Code: typev3.StatusCode_Unauthorized,
				},
				Body: "Need an Authorization Header with a 3 character bearer token! #secure",
			},
		},
	}, nil

}

func main() {
	// create a TCP listener on port 4000
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %s", lis.Addr())

	grpcServer := grpc.NewServer()
	authServer := &AuthServer{}
	auth_pb.RegisterAuthorizationServer(grpcServer, authServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
