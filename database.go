package nemail

import (
	"context"
	"time"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
)

type DB struct {
	*dgo.Dgraph
	cfg *DBConfig
}

type DBConfig struct {
	Host              string
	ConnectionTimeout string
	DgraphCloud       bool
	Cloud             struct {
		Endpoint string
		APIKey   string
	}
}

func NewDBConnection(cfg *DBConfig) (_ *DB, err error) {
	client := &grpc.ClientConn{}
	if cfg.DgraphCloud {
		client, err = dgo.DialSlashEndpoint(cfg.Cloud.Endpoint, cfg.Cloud.APIKey)
		if err != nil {
			return nil, err
		}
	} else {
		timeoutDuration, err := time.ParseDuration(cfg.ConnectionTimeout)
		if err != nil {
			return nil, err
		}
		connectionContext, _ := context.WithTimeout(context.Background(), timeoutDuration)
		client, err = grpc.DialContext(connectionContext, cfg.Host)
		if err != nil {
			return nil, err
		}
	}
	dgraph := dgo.NewDgraphClient(api.NewDgraphClient(client))
	return &DB{dgraph, cfg}, nil
}
