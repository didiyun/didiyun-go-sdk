package common

import (
	"crypto/tls"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

func DialTCP(addr string) (*grpc.ClientConn, error) {
	perRPC := oauth.NewOauthAccess(fetchToken())
	return grpc.Dial(addr, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})), grpc.WithPerRPCCredentials(perRPC))
}

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: Token,
		TokenType:   "bearer",
	}
}
