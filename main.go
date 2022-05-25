package main

import (
	"context"
	_ "embed"

	"github.com/qingwave/opa-gin-authz/server"

	"github.com/open-policy-agent/opa/rego"
	"go.uber.org/zap"
)

var (
	//go:embed authz/authz.rego
	opaConf string
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	query, err := rego.New(rego.Query("data.authz.allow"), rego.Module("authz.repo", opaConf)).PrepareForEval(context.TODO())
	if err != nil {
		logger.Fatal("failed to create rego query", zap.Error(err))
	}

	s := server.New(&query, logger)
	s.Run()
}
