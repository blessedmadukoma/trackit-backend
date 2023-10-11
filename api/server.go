package api

import (
	"fmt"
	"net/http"
	db "trackit/db/sqlc"
	"trackit/token"
	"trackit/util"

	"github.com/gin-gonic/gin"
)

// Server struct serves HTTP requests for our banking service
type Server struct {
	config     util.Config
	tokenMaker token.Maker
	store      *db.Store
	router     *gin.Engine
}

func healthy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Healthy")
	return
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store *db.Store) (*Server, error) {

	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}

	gin.SetMode(config.GinMode)

	router := gin.Default()

	// corsConfig := cors.Default()
	// router.Use(corsConfig)

	// corsConfig := cors.DefaultConfig()
	// setCorsHeaders(&corsConfig)

	// router.Use(cors.New(corsConfig))

	router.Use(CORS())
	router.Use(server.rateLimit())

	// do not trust all proxies
	// router.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
	router.SetTrustedProxies(nil)
	router.TrustedPlatform = gin.PlatformCloudflare

	Routes(router, server)

	server.router = router

	return server, nil
}

// StartServer runs the HTTP server on a specific address
func (srv *Server) StartServer(address string) error {
	fmt.Printf("Server starting on address: %s\n", address)
	return srv.router.Run(fmt.Sprintf(":%s", address))
}
