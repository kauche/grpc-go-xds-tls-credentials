package credentials

import (
	"github.com/kauche/grpc-go-xds-tls-credentials/credentials"
)

func init() {
	credentials.RegisterCredentials(nil, "")
}
