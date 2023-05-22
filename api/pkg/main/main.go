package main

import (
	"net/http"
	"os"
	"time"

	"github.com/Razikus/immuemail/pkg/vault"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var collectionName = os.Getenv("COLLECTION_NAME")
var ledger = os.Getenv("LEDGER")
var apikey = os.Getenv("API_KEY")
var vault_url = os.Getenv("VAULT_URL")
var certfile = os.Getenv("CERT_FILE")
var keyfile = os.Getenv("KEY_FILE")
var bindaddress = os.Getenv("BIND_ADDR")

func main() {

	vaultClient := vault.NewVaultClient(apikey, vault_url)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	r.POST("/form/:formid", func(c *gin.Context) {
		formid := c.Param("formid")
		var form map[string]interface{}
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid form",
			})
			return
		}

		added, err := vaultClient.AddToCollection(ledger, collectionName, formid, form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid form",
			})
			return
		}

		if !added {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Cannot add form to collection",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Message added to form: " + formid,
		})
	})
	if certfile != "" || keyfile != "" {
		r.RunTLS(bindaddress, certfile, keyfile)

	} else {
		r.Run(bindaddress)
	}
}
