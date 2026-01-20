package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"aset-app/controllers"
	"aset-app/db"
	"aset-app/forms"
	"aset-app/helpers"
	"aset-app/models"
	"aset-app/services"

	"github.com/gin-contrib/gzip"
	uuid "github.com/google/uuid"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CORSMiddleware ...
// CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

// RequestIDMiddleware ...
// Generate a unique ID and attach it to each request for future reference or use
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}

var auth = new(controllers.AuthController)

// TokenAuthMiddleware ...
// JWT Authentication middleware attached to each request that needs to be authenitcated to validate the access_token in the header
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.TokenValid(c)
		c.Next()
	}
}

func main() {
	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Start the default gin server
	r := gin.Default()

	//Custom form validator
	binding.Validator = new(forms.DefaultValidator)

	r.Use(CORSMiddleware())
	r.Use(RequestIDMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(helpers.LogAPIErrorMiddleware())

	//Start PostgreSQL database
	//Example: db.GetDB() - More info in the models folder
	db.Init()

	//Start Redis on database 1 - it's used to store the JWT but you can use it for anythig else
	//Example: db.GetRedis().Set(KEY, VALUE, at.Sub(now)).Err()
	db.InitRedis(1)

	services.InitializeMinioClient()

	// Start background tax status update scheduler (runs every 6 hours)
	go func() {
		ticker := time.NewTicker(6 * time.Hour)
		defer ticker.Stop()

		// Run immediately on startup
		log.Println("Running initial tax status update...")
		if err := models.UpdateTaxStatuses(); err != nil {
			log.Printf("Error updating tax statuses on startup: %v", err)
		} else {
			log.Println("Tax status update completed successfully")
		}

		// Then run every 6 hours
		for range ticker.C {
			log.Println("Running scheduled tax status update...")
			if err := models.UpdateTaxStatuses(); err != nil {
				log.Printf("Error updating tax statuses: %v", err)
			} else {
				log.Println("Tax status update completed successfully")
			}
		}
	}()

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		user := new(controllers.UserController)

		v1.POST("/user/login", user.Login)
		v1.POST("/user/register", user.Register)
		v1.GET("/user/logout", user.Logout)

		/*** START AUTH ***/
		auth := new(controllers.AuthController)

		//Refresh the token when needed to generate new access_token and refresh_token for the user
		v1.POST("/token/refresh", auth.Refresh)

		/*** START Article ***/
		article := new(controllers.ArticleController)

		v1.POST("/article", TokenAuthMiddleware(), article.Create)
		v1.GET("/articles", TokenAuthMiddleware(), article.All)
		v1.GET("/articles/:format", TokenAuthMiddleware(), article.All)
		v1.GET("/article/:id", TokenAuthMiddleware(), article.One)
		v1.GET("/article/:id/:format", TokenAuthMiddleware(), article.One)
		v1.PUT("/article/:id", TokenAuthMiddleware(), article.Update)
		v1.DELETE("/article/:id", TokenAuthMiddleware(), article.Delete)

		/*** START VehicleAsset ***/
		vehicleAsset := new(controllers.VehicleAssetController)

		// Vehicle asset routes
		v1.POST("/vehicle-asset", TokenAuthMiddleware(), vehicleAsset.Create)
		v1.PUT("/vehicle-asset/:id/upload-document", TokenAuthMiddleware(), vehicleAsset.UploadDocument)
		v1.PATCH("/vehicle-asset/:id/verify-document", TokenAuthMiddleware(), vehicleAsset.VerifyDocument)
		v1.DELETE("/vehicle-asset/:id/remove-document", TokenAuthMiddleware(), vehicleAsset.RemoveDocument)
		v1.PUT("/vehicle-asset/:id/update-payment", TokenAuthMiddleware(), vehicleAsset.UpdatePayment)
		v1.PUT("/vehicle-asset/:id/update-recommendation", TokenAuthMiddleware(), vehicleAsset.UpdateVehicleRecommendation)
		v1.PUT("/vehicle-asset/:id", TokenAuthMiddleware(), vehicleAsset.Update)
		v1.DELETE("/vehicle-asset/:id", TokenAuthMiddleware(), vehicleAsset.Delete)

		// Vehicle reference routes
		v1.GET("/vehicle/wheels", TokenAuthMiddleware(), vehicleAsset.GetWheels)
		v1.GET("/vehicle/models", TokenAuthMiddleware(), vehicleAsset.GetModels)
		v1.GET("/vehicle/colors", TokenAuthMiddleware(), vehicleAsset.GetColors)
		v1.GET("/vehicle/fuels", TokenAuthMiddleware(), vehicleAsset.GetFuels)
		v1.GET("/vehicle/companies", TokenAuthMiddleware(), vehicleAsset.GetCompanies)
		v1.GET("/vehicle/owning-types", TokenAuthMiddleware(), vehicleAsset.GetOwningTypes)
		v1.GET("/vehicle/brands", TokenAuthMiddleware(), vehicleAsset.GetBrands)
		v1.GET("/vehicle-assets", TokenAuthMiddleware(), vehicleAsset.GetAll)
		v1.GET("/vehicle-assets/:format", TokenAuthMiddleware(), vehicleAsset.GetAll)
		v1.GET("/vehicle-asset/:id", TokenAuthMiddleware(), vehicleAsset.Get)
		v1.GET("/vehicle-asset/:id/:format", TokenAuthMiddleware(), vehicleAsset.Get)
		v1.GET("/vehicle-asset/license/:plate", TokenAuthMiddleware(), vehicleAsset.GetByLicensePlate)
		v1.GET("/vehicle-asset/company/:companyId", TokenAuthMiddleware(), vehicleAsset.GetByCompany)
		v1.GET("/vehicle-asset/company/:companyId/summary", TokenAuthMiddleware(), vehicleAsset.GetCompanySummary)
		v1.GET("/vehicle-assets/summary", TokenAuthMiddleware(), vehicleAsset.GetSummary)
		v1.POST("/vehicle-assets/update-tax-statuses", TokenAuthMiddleware(), vehicleAsset.UpdateTaxStatuses)

		v1.GET("/signed-url/:objectName", TokenAuthMiddleware(), func(c *gin.Context) {
			objectName := c.Param("objectName")
			signedURL, err := services.GenerateSignedURL(objectName)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"signed_url": signedURL})
		})

		/*** START APP VERSION ***/
		app := new(controllers.AppController)
		v1.GET("/version", TokenAuthMiddleware(), app.Version)
		v1.GET("/error-logs", TokenAuthMiddleware(), app.ErrorLogs)
	}

	r.LoadHTMLGlob("./public/html/*")

	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	port := os.Getenv("PORT")

	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	if os.Getenv("SSL") == "TRUE" {

		//Generated using sh generate-certificate.sh
		SSLKeys := &struct {
			CERT string
			KEY  string
		}{
			CERT: "./cert/myCA.cer",
			KEY:  "./cert/myCA.key",
		}

		r.RunTLS(":"+port, SSLKeys.CERT, SSLKeys.KEY)
	} else {
		r.Run(":" + port)
	}

}
