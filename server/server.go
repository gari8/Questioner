package server

import (
	"context"
	"faves4/graph/generated"
	"faves4/graph/resolvers"
	"faves4/infrastructure/auth"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"net/http"
)

func (g *graphQLServer)Serve(ctx context.Context, usePlayGround bool) (chi.Router, error){
	router := chi.NewRouter()

	acceptOrigins := []string{
		"http://localhost:3000",
		"https://faves4.com",
		"https://www.faves4.com",
		"https://faves4.vercel.app",
	}

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   acceptOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		Debug:            true,
	}).Handler)

	router.Use(auth.Middleware(ctx, g.Resolver.Repository))

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: g.Resolver}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})


	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	router.Group(func(r chi.Router) {
		r.Handle("/query", srv)
		if usePlayGround {
			r.Handle("/", playground.Handler("faves4 ground", "/query"))
		}
	})

	return router, nil
}

type graphQLServer struct {
	Resolver *resolvers.Resolver
}

type GraphQLServer interface {
	Serve(ctx context.Context, usePlayground bool) (chi.Router, error)
}

func NewGraphQLServer(resolver *resolvers.Resolver) GraphQLServer {
	return &graphQLServer{
		Resolver: resolver,
	}
}