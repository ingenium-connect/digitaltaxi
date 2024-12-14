package rest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/common/helpers"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/utils"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/usecases/payperday"
)

// AuthenticationGinMiddleware is an authentication middleware for servers using Gin. It checks the user token and ensures
// that it is valid
func AuthenticationGinMiddleware(usecases payperday.PayPerDay) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) == 0 || len(fields) == 1 {
			ctx.Set("user_id", nil)
			ctx.Next()
		} else {
			if len(fields) != 0 && fields[0] == "Bearer" {
				access_token = fields[1]
			}

			sub, err := utils.ValidateToken(access_token, helpers.MustGetEnvVar("ACCESS_TOKEN_PUBLIC_KEY"))
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})

				return
			}

			str := fmt.Sprint(sub)
			inputArray := strings.Split(str, "++")

			user, err := usecases.GetUserProfileByUserID(ctx, inputArray[0])
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": fmt.Sprintf("Unable to retrieve user profile: %v", err),
				})

				return
			}

			ctx.Set("user_id", user.ID)
			ctx.Next()
		}
	}
}
