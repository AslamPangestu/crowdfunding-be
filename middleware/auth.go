package middleware

import (
	"crowdfunding/config"
	"crowdfunding/helper"
	"crowdfunding/services"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware : Middleware for auth
func AuthMiddleware(authService config.AuthService, userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			errResponse := helper.ResponseHandler("Unauthorize", http.StatusUnauthorized, "failed", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}
		//Get plain token
		var parseToken string
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) == 2 {
			parseToken = splitToken[1]
		}

		//Validate token
		token, err := authService.ValidateToken(parseToken)
		if err != nil {
			errResponse := helper.ResponseHandler("Unauthorize", http.StatusUnauthorized, "failed", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !(ok || token.Valid) {
			errResponse := helper.ResponseHandler("Unauthorize", http.StatusUnauthorized, "failed", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}
		//Find user
		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			errResponse := helper.ResponseHandler("Unauthorize", http.StatusUnauthorized, "failed", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}
		c.Set("currentUser", user)
	}
}
