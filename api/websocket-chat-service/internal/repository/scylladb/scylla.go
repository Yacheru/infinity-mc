package scylladb

import (
	"context"
	"os"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/migrate"

	"websocket-chat-service/init/config"
	"websocket-chat-service/init/logger"
	"websocket-chat-service/pkg/constants"
)

func NewScyllaSession(ctx context.Context, cfg *config.Config) (*gocqlx.Session, error) {
	cluster := gocql.NewCluster(cfg.ScyllaHosts...)

	cluster.Keyspace = cfg.ScyllaKeyspace
	cluster.Consistency = gocql.One
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())

	ses, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		logger.Error(err.Error(), constants.ScyllaCategory)
		return nil, err
	}

	if err := migrate.FromFS(ctx, ses, os.DirFS("./migrations")); err != nil {
		logger.Error(err.Error(), constants.ScyllaCategory)
		return nil, err
	}

	return &ses, nil
}
