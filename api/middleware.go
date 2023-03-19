package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"trackit/token"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// authMiddleware authorizes/validates a user
func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse("unauthorized", err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse("unauthorized", err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse("unauthorized", err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse("invalid access token", err))
			return
		}

		// store payload in the context
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

// getAuthorizationPayload retrieves the authorization payload from the context
func getAuthorizationPayload(ctx *gin.Context) (*token.Payload, error) {
	payload, ok := ctx.Get(authorizationPayloadKey)
	if !ok {
		return nil, errors.New("authorization payload not found")
	}

	return payload.(*token.Payload), nil
}

// setCorsHeaders sets the CORS headers
func setCorsHeaders(corsConfig *cors.Config) {
	corsConfig.AllowOrigins = []string{"https://localhost", "http://localhost", "http://localhost:3000", "https://localhost:3000"}

	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Requested-With", "Accept", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials", "Access-Control-Max-Age", "Access-Control-Expose-Headers", "Access-Control-Request-Headers", "Access-Control-Request-Method", "X-Forwarded-For", "X-Forwarded-Host", "X-Forwarded-Port", "X-Forwarded-Proto", "X-Real-Ip", "X-Request-Id", "X-Scheme", "X-Forwarded-Proto", "X-Forwarded-Protocol", "X-Forwarded-Ssl", "X-Url-Scheme", "X-Forwarded-Host", "X-Forwarded-Server", "X-Forwarded-For", "withCredentials"}

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS", "GET", "POST", "PUT", "DELETE", "PATCH")

	corsConfig.AllowCredentials = true
}
