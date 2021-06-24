package main

import (
	"context"
	"faves4/graph/generated"
	"faves4/graph/resolvers"
	"faves4/infrastructure/auth"
	"faves4/infrastructure/database/conf"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{DB: client}}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	usePlayGround := true

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Use(auth.Middleware(ctx, client))

	router.Group(func(r chi.Router) {
		r.Handle("/query", srv)
		if usePlayGround {
			r.Handle("/", playground.Handler("faves4 ground", "/query"))
		}
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
