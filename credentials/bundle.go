package credentials

import (
	"crypto/x509"

	"google.golang.org/grpc/credentials"
)

var _ credentials.Bundle = (*tlsCredentalsBundle)(nil)

type tlsCredentalsBundle struct {
	cp                 *x509.CertPool
	serverNameOverride string
}

func (b *tlsCredentalsBundle) TransportCredentials() credentials.TransportCredentials {
	return credentials.NewClientTLSFromCert(b.cp, b.serverNameOverride)
}

func (b *tlsCredentalsBundle) PerRPCCredentials() credentials.PerRPCCredentials {
	return nil
}

func (b *tlsCredentalsBundle) NewWithMode(mode string) (credentials.Bundle, error) {
	return &tlsCredentalsBundle{
		cp:                 b.cp,
		serverNameOverride: b.serverNameOverride,
	}, nil
}
