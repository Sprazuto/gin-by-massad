package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/Massad/gin-boilerplate/controllers"
	"github.com/Massad/gin-boilerplate/db"
	"github.com/Massad/gin-boilerplate/forms"
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

	//Start PostgreSQL database
	//Example: db.GetDB() - More info in the models folder
	db.Init()

	//Start Redis on database 1 - it's used to store the JWT but you can use it for anythig else
	//Example: db.GetRedis().Set(KEY, VALUE, at.Sub(now)).Err()
	db.InitRedis(1)

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

		/*** START LkeRekap ***/
		lkeRekap := new(controllers.LkeRekapController)

		v1.POST("/lke-rekap", TokenAuthMiddleware(), lkeRekap.Create)
		v1.GET("/lke-rekaps", TokenAuthMiddleware(), lkeRekap.All)
		v1.GET("/lke-rekaps/:format", TokenAuthMiddleware(), lkeRekap.All)
		v1.GET("/lke-rekap/:id", TokenAuthMiddleware(), lkeRekap.One)
		v1.GET("/lke-rekap/:id/:format", TokenAuthMiddleware(), lkeRekap.One)
		v1.PUT("/lke-rekap/:id", TokenAuthMiddleware(), lkeRekap.Update)
		v1.DELETE("/lke-rekap/:id", TokenAuthMiddleware(), lkeRekap.Delete)
		v1.GET("/lke-rekap/opd/:id_opd/tahun/:tahun", TokenAuthMiddleware(), lkeRekap.GetByOPDAndTahun)

		/*** START LkeEvaluasi ***/
		lkeEvaluasi := new(controllers.LkeEvaluasiController)

		v1.POST("/lke-evaluasi", TokenAuthMiddleware(), lkeEvaluasi.Create)
		v1.GET("/lke-evaluasis", TokenAuthMiddleware(), lkeEvaluasi.All)
		v1.GET("/lke-evaluasis/:format", TokenAuthMiddleware(), lkeEvaluasi.All)
		v1.GET("/lke-evaluasi/:id", TokenAuthMiddleware(), lkeEvaluasi.One)
		v1.GET("/lke-evaluasi/:id/:format", TokenAuthMiddleware(), lkeEvaluasi.One)
		v1.PUT("/lke-evaluasi/:id", TokenAuthMiddleware(), lkeEvaluasi.Update)
		v1.DELETE("/lke-evaluasi/:id", TokenAuthMiddleware(), lkeEvaluasi.Delete)

		/*** START LkeKomponen ***/
		lkeKomponen := new(controllers.LkeKomponenController)

		v1.POST("/lke-komponen", TokenAuthMiddleware(), lkeKomponen.Create)
		v1.GET("/lke-komponens", TokenAuthMiddleware(), lkeKomponen.All)
		v1.GET("/lke-komponens/:format", TokenAuthMiddleware(), lkeKomponen.All)
		v1.GET("/lke-komponen/:id", TokenAuthMiddleware(), lkeKomponen.One)
		v1.GET("/lke-komponen/:id/:format", TokenAuthMiddleware(), lkeKomponen.One)
		v1.PUT("/lke-komponen/:id", TokenAuthMiddleware(), lkeKomponen.Update)
		v1.DELETE("/lke-komponen/:id", TokenAuthMiddleware(), lkeKomponen.Delete)
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
