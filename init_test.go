package credentials_test

import (
	"testing"

	"google.golang.org/grpc/xds/bootstrap"

	_ "github.com/kauche/grpc-go-xds-tls-credentials"
)

func TestInit(t *testing.T) {
	t.Parallel()

	c := bootstrap.GetChannelCredentials("tls")

	if c == nil {
		t.Errorf("expected the credentials to be registered")
	}
}
