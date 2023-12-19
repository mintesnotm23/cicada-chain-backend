package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jamyMarkos/backend/controllers"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/api/getVerificationCode", controllers.VerifyEmailHandler)
    r.POST("/api/verifyEmail", controllers.VerifyEmail)
}
