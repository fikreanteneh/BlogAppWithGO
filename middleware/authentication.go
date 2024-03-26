package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
    "BlogApp/domain/model"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
            return
        }

        token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(secret), nil
        })

        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Error parsing token"})
            return
        }

        if !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }
        
        // Extract claims from token
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            return
        }

        // Set user ID and username in the Gin context
        UserID, ok := claims["id"].(string)
        Email, ok := claims["email"].(string)
        Username, ok := claims["username"].(string)
        Role, ok := claims["role"].(string)


        if !ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
            return
        }
        if !ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Username not found in token"})
            return
        }
        c.Set("AuthenticatedUser", &model.AuthenticatedUser{
            UserID: UserID,
            Email: Email,
            Username: Username,
            Role: Role,
        })

        c.Next()
    }
}
