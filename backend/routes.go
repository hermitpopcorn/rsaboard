package main

import (
	cryptoRand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	mathRand "math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hermitpopcorn/rsaboard/internal/database"
)

type FindMessageRequest struct {
	Code string `uri:"code" binding:"required"`
}

type CreateMessageRequest struct {
	Message             string `json:"message" binding:"required,max=180"`
	Email               string `json:"email" binding:"required"`
	ShouldBurnInMinutes int    `json:"shouldBurnInMinutes"`
}

type ReadMessageRequest struct {
	PrivateKey []byte `json:"privateKey" binding:"required"`
}

func injectRoutes(router *gin.Engine, db database.Database) {
	router.Use(optionsHandlerMiddleware())

	router.GET("/messages/:code", func(context *gin.Context) {
		var request FindMessageRequest
		if err := context.ShouldBindUri(&request); err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		message, err := db.FindMessage(request.Code)
		if err != nil {
			context.JSON(err.Code(), gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": message})
	})

	router.POST("/messages", func(context *gin.Context) {
		var request CreateMessageRequest
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		code, err := generateUniqueCode(db)
		if err != nil {
			context.JSON(500, gin.H{"error": "Failed generating unique code."})
			return
		}

		privateKey, encryptedText, err := encryptText(request.Message)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed encrypting text."})
			return
		}

		message, dbErr := db.CreateMessage(code, encryptedText, request.Email, request.ShouldBurnInMinutes)
		if dbErr != nil {
			context.JSON(dbErr.Code(), gin.H{"error": "Failed saving message."})
			return
		}

		privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
		context.JSON(http.StatusCreated, gin.H{"code": message.Code, "privateKey": privateKeyBytes})
	})

	router.POST("/messages/:code", func(context *gin.Context) {
		var uriRequest FindMessageRequest
		if err := context.ShouldBindUri(&uriRequest); err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
		var bodyRequest ReadMessageRequest
		if err := context.ShouldBindJSON(&bodyRequest); err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		message, dbErr := db.FindMessage(uriRequest.Code)
		if dbErr != nil {
			context.JSON(dbErr.Code(), gin.H{"error": "Could not find message with requested code."})
			return
		}

		privateKey, err := x509.ParsePKCS1PrivateKey(bodyRequest.PrivateKey)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Failed parsing private key."})
			return
		}
		decryptedText, err := decryptText(message.EncryptedText, privateKey)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Failed decrypting message with provided private key."})
			return
		}

		accessLog, dbErr := db.CreateAccessLog(&message, context.ClientIP())
		if dbErr != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Failed marking message (A)."})
			return
		}
		if message.ShouldBurn && message.DeleteAt == nil {
			burnAfterMinutes := message.ShouldBurnInMinutes * int64(time.Minute)
			deleteAt := accessLog.AccessedAt.Add(time.Duration(burnAfterMinutes))

			dbErr := db.SetMessageDeleteAt(&message, deleteAt)
			if dbErr != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": "Failed marking message (B)."})
				return
			}
		}

		context.JSON(http.StatusOK, gin.H{"message": decryptedText})
	})
}

func generateUniqueCode(db database.Database) (string, error) {
	unique := false
	code := ""
	for !unique {
		code = generateCode()
		codeAlreadyExists, err := db.CheckCodeExists(code)
		if err != nil {
			return "", err
		}

		if !codeAlreadyExists {
			break
		}
	}

	return code, nil
}

func generateCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 4)
	for i := range code {
		code[i] = charset[mathRand.Intn(len(charset))]
	}

	return string(code)
}

func encryptText(text string) (*rsa.PrivateKey, string, error) {
	key, err := rsa.GenerateKey(cryptoRand.Reader, 2048)
	if err != nil {
		return nil, "", err
	}

	encrypted, err := rsa.EncryptOAEP(sha256.New(), cryptoRand.Reader, &key.PublicKey, []byte(text), nil)
	if err != nil {
		return nil, "", err
	}

	return key, string(encrypted), nil
}

func decryptText(text string, privateKey *rsa.PrivateKey) (string, error) {
	decrypted, err := rsa.DecryptOAEP(sha256.New(), nil, privateKey, []byte(text), nil)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

func optionsHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
