package auth

import (
	proto "../../../proto/auth"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"time"
)

var RegisterAuthServiceServer = proto.RegisterAuthServiceServer
var expiresDuration time.Duration // Ablaufdatum des logins, wird in init fixiert

func init() {
	//Todo: eventuell über config einstellbar machen
	expiresDuration = time.Hour * time.Duration(24)
}

// Gibt den grpc ServiceServer zurück
func GetServiceServer() proto.AuthServiceServer {
	var s authServiceServer
	return &s
}

// authServiceServer is used to implement authServiceServer.
type authServiceServer struct {
}

// Override Funktion um nicht über die default auth-middleware
func (s *authServiceServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (authServiceServer) Login(ctx context.Context, req *proto.CredentialsRequest) (*empty.Empty, error) {
	_, err := login(req.Body.Username, req.Body.Password)
	if err == nil {
		// erfolg
		d := time.Now().Add(expiresDuration)
		//todo: Rollen aus user Table in token werfen
		jwtToken := createJWT()
		_ = grpc.SendHeader(ctx, metadata.Pairs("Set-Cookie", "Authorization=Bearer "+jwtToken+";expires="+d.UTC().String()))
	} else {
		// falscher login
		return nil, status.Errorf(codes.Unauthenticated, "login error %v", err)
	}
	return &empty.Empty{}, nil
}

func (s *authServiceServer) Logout(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	_ = grpc.SendHeader(ctx, metadata.Pairs("Set-Cookie", "Authorization=deleted; expires=Thu, 01 Jan 1970 00:00:00 GMT"))
	return &empty.Empty{}, nil
}
