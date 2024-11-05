package serve

import (
	"context"
	log2 "github.com/ipfs/go-log/v2"
)

var log = log2.Logger("serve")

type Server struct {
}

func (s Server) Start(ctx context.Context) error {
	log.Info("Started server")

	select {} // temp wait
	return nil
}
