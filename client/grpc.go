/**
 * @note
 * 用于dicloud用户系统
 *
 * @author wuwenyu
 * @date   2020/7/1
 */
package client

import (
	"crypto/tls"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

func GetGRPCClient(token, serverAddr string) (*grpc.ClientConn, error) {
	perRPC := oauth.NewOauthAccess(&oauth2.Token{
		AccessToken: token,
		TokenType:   "bearer",
	})
	return grpc.Dial(serverAddr,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithPerRPCCredentials(perRPC))
}
