package credentials

import (
	"crypto/x509"
	"encoding/json"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/xds/bootstrap"
)

var _ bootstrap.Credentials = (*transportCredsBuilder)(nil)

// RegisterCredentials registers the TLS credentials used for connecting to the xds management server.
// cp and serverNameOverride are used to create the TLS credentials.
func RegisterCredentials(cp *x509.CertPool, serverNameOverride string) {
	builder := &transportCredsBuilder{
		cp:                 cp,
		serverNameOverride: serverNameOverride,
	}

	bootstrap.RegisterCredentials(builder)
}

type transportCredsBuilder struct {
	cp                 *x509.CertPool
	serverNameOverride string
}

func (t *transportCredsBuilder) Build(_ json.RawMessage) (credentials.Bundle, func(), error) {
	return &tlsCredentalsBundle{
		cp:                 t.cp,
		serverNameOverride: t.serverNameOverride,
	}, func() {}, nil
}

func (t *transportCredsBuilder) Name() string {
	return "tls"
}
