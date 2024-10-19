package server

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strings"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/gin-gonic/gin"

	"github.com/BrandonC98/fortify/services/fortify/internal/database"
	"github.com/BrandonC98/fortify/services/fortify/internal/model"
)

func StartServer(config model.Config) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	var client standardHTTPClient
	r := database.NewSecretsRepository(config.DBHost, "fortify_db", config.DBUser, config.DBPassword)
	r.Setup()

	router.GET("/ping", pingHandler)
	router.GET("/", homeHandler)
	router.GET("/generate", generateHandler(fmt.Sprintf("%s/generate", config.GeneratorURL), &client, config))
	router.POST("/save", saveHandler(r, config))
	router.GET("/show", showHandler(r, config))

	err := router.Run(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func showHandler(r database.Repository, config model.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		creds := r.RetriveAllRecords()
		var sb strings.Builder
		for i := 0; i < len(creds); i++ {
			decryptedValue, err := helper.DecryptMessageWithPassword([]byte(config.Key), creds[i].Value)
			if err != nil {
				slog.Warn("Unable to decrypt key pair " + creds[i].Name + " value, continuing with value in current form")
				sb.WriteString(fmt.Sprint(creds[i].Name, ": ", creds[i].Value, "\n"))
				continue
			}

			sb.WriteString(fmt.Sprint(creds[i].Name, ": ", decryptedValue, "\n"))
		}

		c.String(http.StatusOK, sb.String())
	}
}

func saveHandler(r database.Repository, config model.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials model.Secret
		key := config.Key

		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		plaintext := &credentials.Value
		encryptedText, err := helper.EncryptMessageWithPassword([]byte(key), *plaintext)
		if err != nil {
			log.Fatal(err)
		}

		encryptedCredentials := model.Secret{
			Name:  credentials.Name,
			Value: encryptedText,
		}

		r.AddRecord(&encryptedCredentials)
		c.String(http.StatusOK, "successful")
	}
}

func generateHandler(endpointURL string, client HTTPClient, config model.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := config.Key
		s := generate(endpointURL, client)

		s, err := helper.DecryptMessageWithPassword([]byte(key), s)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": s})
	}
}

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title": "main website",
	})
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
