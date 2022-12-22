# grpc-go-xds-tls-credentials

`grpc-go-xds-tls-credentials` is a Go library that provides a TLS credentials used for making a grpc-go client able to connect to a xDS control plane 🚀

## Usage

Just import this package like below:

```go
import _ "github.com/kauche/grpc-go-xds-tls-credentials"
```

This registers a new type of `channel_creds` named `tls` used for connecting to the xDS management server.

Then, make your grpc-go client use this `tls` channel_creds like below:

```sh
$ GRPC_XDS_BOOTSTRAP_CONFIG='{
  "xds_servers": [{
    "server_uri": "...",
    "channel_creds": [{"type": "tls"}],
    "server_features": ["xds_v3"]
  }],
  "node": {"id": "..."}
}' go run ./path/to/your/app
```
