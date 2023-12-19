package controllers

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/jamyMarkos/backend/models"
	"github.com/jamyMarkos/backend/utils"
	"time"
	"math/rand"
	"context"
	"strconv"
	"regexp"
	"go.mongodb.org/mongo-driver/bson"
)


func VerifyEmailHandler(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}
	
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := request.Email
	
	if !isValidEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}
	

	verificationCode := generateVerificationCode()
	err := utils.SendVerificationEmail(email, verificationCode)
	
	if(err != nil){
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification email"})
		return
	}

	verification := models.EmailVerification{
        Email:            email,
        VerificationCode: verificationCode,
        CreatedAt:        time.Now(),
		ExpiresAt:        time.Now().Add(1 * time.Hour),
    }
	collection := models.GetClient().Database("email-verification").Collection("emailVerifications")
	 _, err =collection.InsertOne(context.Background(), verification)
	log.Println(err)


	c.JSON(http.StatusOK, gin.H{"message": "Verification code sent successfully"})
}

func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(900000) + 100000)
}


func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@gmail\.com$`)
	return emailRegex.MatchString(email)
}

func VerifyEmail(c *gin.Context){
		var request struct {
			Email string `json:"email"`
			Code string `json:"code"`
		}
		
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		 collection := models.GetClient().Database("email-verification").Collection("emailVerifications")
		var verification models.EmailVerification
		err := collection.FindOne(context.Background(), bson.M{"email": request.Email}).Decode(&verification)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find verification information"})
			return
		}

		if verification.ExpiresAt.Before(time.Now()) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Verification code expired"})
			return
		}

		if verification.VerificationCode != request.Code {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
	}
	
