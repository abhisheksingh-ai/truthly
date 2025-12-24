package middleware

import (
	"net/http"
	"strings"
	"truthly/internals/util/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authToken *auth.AuthToken) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//1. Authorization Header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
			return
		}

		//2. Extract Bearer token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization format",
			})
			return
		}

		//3. verify token
		claims, err := authToken.VerifyJwtToken(
			token,
			ctx.Request.Context(),
		)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		ctx.Set("userId", claims.UserId)

		ctx.Next()
	}
}
