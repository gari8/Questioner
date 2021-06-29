package main

import (
	"context"
	"faves4/graph/resolvers"
	"faves4/infrastructure/database/conf"
	"faves4/interactor"
	"faves4/server"
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	client, err := conf.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}
	if client == nil {
		panic(err)
	}
	defer func() {
		if client != nil {
			if err := client.Close(); err != nil {
				panic(err)
			}
		}
	}()

	ctx := context.Background()
	// 自動マイグレーションツールを実行して、すべてのスキーマリソースを作成します。
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	usePlayGround := true

	env := os.Getenv("DEPLOY_ENV")
	if env == "production" {
		usePlayGround = false
	}

	i := interactor.NewInteractor(client)
	r := i.NewRepository()

	resolver := resolvers.NewResolver(*r)

	s := server.NewGraphQLServer(resolver)

	router, err := s.Serve(ctx, usePlayGround)

	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
