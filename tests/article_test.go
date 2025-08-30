//go:build all
// +build all

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"aset-app/controllers"
	"aset-app/db"
	"aset-app/forms"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var auth = new(controllers.AuthController)

// TokenAuthMiddleware ...
// JWT Authentication middleware attached to each request that needs to be authenticated to validate the access_token in the header
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.TokenValid(c)
		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	// Set GIN to release mode to suppress debug output during tests
	gin.SetMode(gin.ReleaseMode)

	// Create a new engine without default middleware to reduce noise
	r := gin.New()

	// Add only essential middleware
	r.Use(gin.Recovery())

	// Custom form validator
	binding.Validator = new(forms.DefaultValidator)

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		user := new(controllers.UserController)

		v1.POST("/user/login", user.Login)
		v1.POST("/user/register", user.Register)
		v1.GET("/user/logout", user.Logout)

		/*** START AUTH ***/
		auth := new(controllers.AuthController)

		v1.POST("/token/refresh", auth.Refresh)

		/*** START Article ***/
		article := new(controllers.ArticleController)

		v1.POST("/article", TokenAuthMiddleware(), article.Create)
		v1.GET("/articles", TokenAuthMiddleware(), article.All)
		v1.GET("/article/:id", TokenAuthMiddleware(), article.One)
		v1.PUT("/article/:id", TokenAuthMiddleware(), article.Update)
		v1.DELETE("/article/:id", TokenAuthMiddleware(), article.Delete)
	}

	return r
}

var loginCookie string

var testEmail = "test-gin-boilerplate@test.com"
var testPassword = "123456"

var accessToken string
var refreshToken string

var articleID int

/**
* TestArticleAPI
* Tests the complete Article API lifecycle with visual checklist
*
* Must pass for Article functionality to work
* Outputs a checklist with ✅/❌ indicators for easy visual scanning
 */
func TestArticleAPI(t *testing.T) {
	// Initialize checklist for visual test tracking
	checklist := NewTestChecklist()

	// Print test header with clear separation
	t.Log("\n" + strings.Repeat("═", 80))
	t.Logf("🚀 ARTICLE API TEST SUITE")
	t.Logf("👤 Test User: %s", testEmail)
	t.Log(strings.Repeat("═", 80))

	// Test 1: Database Initialization
	t.Run("Database Setup", func(t *testing.T) {
		testStart := time.Now()

		//Load the .env file
		err := godotenv.Load("../.env")
		if err != nil {
			t.Fatal("Error loading .env file:", err)
		}

		db.Init()
		db.InitRedis(1)

		dbSuccess := true
		testDuration := time.Since(testStart)
		checklist.AddResult("Database Initialization", dbSuccess, testDuration)

		t.Log("✅ Database initialized successfully")
	})

	// Test 2: User Registration
	t.Run("User Registration", func(t *testing.T) {
		testStart := time.Now()
		testRouter := SetupRouter()

		var registerForm forms.RegisterForm
		registerForm.Name = "testing"
		registerForm.Email = testEmail
		registerForm.Password = testPassword

		data, _ := json.Marshal(registerForm)
		req, err := http.NewRequest("POST", "/v1/user/register", bytes.NewBufferString(string(data)))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			t.Fatal("Error creating request:", err)
		}

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		registerSuccess := assert.Equal(t, http.StatusOK, resp.Code, "User registration should succeed")
		testDuration := time.Since(testStart)
		checklist.AddResult("User Registration", registerSuccess, testDuration)

		if registerSuccess {
			t.Log("✅ User registration successful")
		} else {
			t.Logf("❌ User registration failed with status: %d", resp.Code)
		}
	})

	// Test 3: Invalid Email Registration
	t.Run("Invalid Email Registration", func(t *testing.T) {
		testStart := time.Now()
		testRouter := SetupRouter()

		var registerForm forms.RegisterForm
		registerForm.Name = "testing"
		registerForm.Email = "invalid@email"
		registerForm.Password = testPassword

		data, _ := json.Marshal(registerForm)
		req, err := http.NewRequest("POST", "/v1/user/register", bytes.NewBufferString(string(data)))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			t.Fatal("Error creating request:", err)
		}

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		invalidEmailSuccess := assert.Equal(t, http.StatusNotAcceptable, resp.Code, "Invalid email should be rejected")
		testDuration := time.Since(testStart)
		checklist.AddResult("Invalid Email Registration", invalidEmailSuccess, testDuration)

		if invalidEmailSuccess {
			t.Log("✅ Invalid email properly rejected")
		}
	})

	// Test 4: User Login
	t.Run("User Login", func(t *testing.T) {
		testStart := time.Now()
		testRouter := SetupRouter()

		var loginForm forms.LoginForm
		loginForm.Email = testEmail
		loginForm.Password = testPassword

		data, _ := json.Marshal(loginForm)
		req, err := http.NewRequest("POST", "/v1/user/login", bytes.NewBufferString(string(data)))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			t.Fatal("Error creating request:", err)
		}

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("Error reading response:", err)
		}

		var res struct {
			Message string `json:"message"`
			User    struct {
				CreatedAt int64  `json:"created_at"`
				Email     string `json:"email"`
				ID        int64  `json:"id"`
				Name      string `json:"name"`
				UpdatedAt int64  `json:"updated_at"`
			} `json:"user"`
			Token struct {
				AccessToken  string `json:"access_token"`
				RefreshToken string `json:"refresh_token"`
			} `json:"token"`
		}
		json.Unmarshal(body, &res)

		accessToken = res.Token.AccessToken
		refreshToken = res.Token.RefreshToken

		loginSuccess := assert.Equal(t, http.StatusOK, resp.Code, "User login should succeed")
		testDuration := time.Since(testStart)
		checklist.AddResult("User Login", loginSuccess, testDuration)

		if loginSuccess {
			t.Log("✅ User login successful")
			t.Logf("🔑 Access token obtained (length: %d)", len(accessToken))
		} else {
			t.Logf("❌ User login failed with status: %d", resp.Code)
		}
	})

	// Test 5: Invalid Login
	t.Run("Invalid Login", func(t *testing.T) {
		testStart := time.Now()
		testRouter := SetupRouter()

		var loginForm forms.LoginForm
		loginForm.Email = "wrong@email.com"
		loginForm.Password = testPassword

		data, _ := json.Marshal(loginForm)
		req, err := http.NewRequest("POST", "/v1/user/login", bytes.NewBufferString(string(data)))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			t.Fatal("Error creating request:", err)
		}

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		invalidLoginSuccess := assert.Equal(t, http.StatusNotAcceptable, resp.Code, "Invalid login should be rejected")
		testDuration := time.Since(testStart)
		checklist.AddResult("Invalid Login", invalidLoginSuccess, testDuration)

		if invalidLoginSuccess {
			t.Log("✅ Invalid login properly rejected")
		}
	})

	// Test 6: Article Creation
	t.Run("Article Creation", func(t *testing.T) {
		testStart := time.Now()
		testRouter := SetupRouter()

		var form forms.CreateArticleForm
		form.Title = "Testing article title"
		form.Content = "Testing article content"

		data, _ := json.Marshal(form)
		req, err := http.NewRequest("POST", "/v1/article", bytes.NewBufferString(string(data)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

		if err != nil {
			t.Fatal("Error creating request:", err)
		}

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("Error reading response:", err)
		}

		var res struct {
			Status int
			ID     int
		}
		json.Unmarshal(body, &res)
		articleID = res.ID

		createSuccess := assert.Equal(t, http.StatusOK, resp.Code, "Article creation should succeed")
		testDuration := time.Since(testStart)
		checklist.AddResult("Article Creation", createSuccess, testDuration)

		if createSuccess {
			t.Log("✅ Article creation successful")
			t.Logf("📝 Article created with ID: %d", articleID)
		} else {
			t.Logf("❌ Article creation failed with status: %d", resp.Code)
		}
	})

	// Test 7: Invalid Article Creation
	t.Run("Invalid Article Creation", func(t *testing.T) {
		testStart := time.Now()
		testRouter := SetupRouter()

		var form forms.CreateArticleForm
		form.Title = "Testing article title" // Missing content

		data, _ := json.Marshal(form)
		req, err := http.NewRequest("POST", "/v1/article", bytes.NewBufferString(string(data)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

		if err != nil {
			t.Fatal("Error creating request:", err)
		}

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		invalidCreateSuccess := assert.Equal(t, http.StatusNotAcceptable, resp.Code, "Invalid article should be rejected")
		testDuration := time.Since(testStart)
		checklist.AddResult("Invalid Article Creation", invalidCreateSuccess, testDuration)

		if invalidCreateSuccess {
			t.Log("✅ Invalid article properly rejected")
		}
	})

	// Test 8: Get Article
	t.Run("Get Article", func(t *testing.T) {
		testStart := time.Now()
		testRouter := SetupRouter()

		req, err := http.NewRequest("GET", fmt.Sprintf("/v1/article/%d", articleID), nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

		if err != nil {
			t.Fatal("Error creating request:", err)
		}

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		getSuccess := assert.Equal(t, http.StatusOK, resp.Code, "Get article should succeed")
		testDuration := time.Since(testStart)
		checklist.AddResult("Get Article", getSuccess, testDuration)

		if getSuccess {
			t.Log("✅ Get article successful")
		} else {
			t.Logf("❌ Get article failed with status: %d", resp.Code)
		}
	})

	// Test 9: Get Invalid Article
	t.Run("Get Invalid Article", func(t *testing.T) {
		testStart := time.Now()
		testRouter := SetupRouter()

		req, err := http.NewRequest("GET", "/v1/article/invalid", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

		if err != nil {
			t.Fatal("Error creating request:", err)
		}

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		invalidGetSuccess := assert.Equal(t, http.StatusNotFound, resp.Code, "Get invalid article should return 404")
		testDuration := time.Since(testStart)
		checklist.AddResult("Get Invalid Article", invalidGetSuccess, testDuration)

		if invalidGetSuccess {
			t.Log("✅ Invalid article properly returns 404")
		}
	})

	// Test 10: Get Article Not Logged In
	t.Run("Get Article Not Logged In", func(t *testing.T) {
		testStart := time.Now()
		testRouter := SetupRouter()

		req, err := http.NewRequest("GET", fmt.Sprintf("/v1/article/%d", articleID), nil)

		if err != nil {
			t.Fatal("Error creating request:", err)
		}

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		unauthorizedSuccess := assert.Equal(t, http.StatusUnauthorized, resp.Code, "Unauthorized access should be rejected")
		testDuration := time.Since(testStart)
		checklist.AddResult("Get Article Unauthorized", unauthorizedSuccess, testDuration)

		if unauthorizedSuccess {
			t.Log("✅ Unauthorized access properly rejected")
		}
	})

	// Test 11: Update Article
	t.Run("Update Article", func(t *testing.T) {
		testStart := time.Now()
		testRouter := SetupRouter()

		var form forms.CreateArticleForm
		form.Title = "Testing new article title"
		form.Content = "Testing new article content"

		data, _ := json.Marshal(form)
		url := fmt.Sprintf("/v1/article/%d", articleID)
		req, err := http.NewRequest("PUT", url, bytes.NewBufferString(string(data)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

		if err != nil {
			t.Fatal("Error creating request:", err)
		}

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		updateSuccess := assert.Equal(t, http.StatusOK, resp.Code, "Article update should succeed")
		testDuration := time.Since(testStart)
		checklist.AddResult("Update Article", updateSuccess, testDuration)

		if updateSuccess {
			t.Log("✅ Article update successful")
		} else {
			t.Logf("❌ Article update failed with status: %d", resp.Code)
		}
	})

	// Test 12: Delete Article
	t.Run("Delete Article", func(t *testing.T) {
		testStart := time.Now()
		testRouter := SetupRouter()

		url := fmt.Sprintf("/v1/article/%d", articleID)
		req, err := http.NewRequest("DELETE", url, nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

		if err != nil {
			t.Fatal("Error creating request:", err)
		}

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		deleteSuccess := assert.Equal(t, http.StatusOK, resp.Code, "Article deletion should succeed")
		testDuration := time.Since(testStart)
		checklist.AddResult("Delete Article", deleteSuccess, testDuration)

		if deleteSuccess {
			t.Log("✅ Article deletion successful")
		} else {
			t.Logf("❌ Article deletion failed with status: %d", resp.Code)
		}
	})

	// Test 13: Refresh Token
	t.Run("Refresh Token", func(t *testing.T) {
		testStart := time.Now()
		testRouter := SetupRouter()

		var tokenForm forms.Token
		tokenForm.RefreshToken = refreshToken

		data, _ := json.Marshal(tokenForm)
		req, err := http.NewRequest("POST", "/v1/token/refresh", bytes.NewBufferString(string(data)))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			t.Fatal("Error creating request:", err)
		}

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		refreshSuccess := assert.Equal(t, http.StatusOK, resp.Code, "Token refresh should succeed")
		testDuration := time.Since(testStart)
		checklist.AddResult("Refresh Token", refreshSuccess, testDuration)

		if refreshSuccess {
			t.Log("✅ Token refresh successful")
		} else {
			t.Logf("❌ Token refresh failed with status: %d", resp.Code)
		}
	})

	// Test 14: Invalid Refresh Token
	t.Run("Invalid Refresh Token", func(t *testing.T) {
		testStart := time.Now()
		testRouter := SetupRouter()

		var tokenForm forms.Token
		tokenForm.RefreshToken = "invalid_token"

		data, _ := json.Marshal(tokenForm)
		req, err := http.NewRequest("POST", "/v1/token/refresh", bytes.NewBufferString(string(data)))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			t.Fatal("Error creating request:", err)
		}

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		invalidRefreshSuccess := assert.Equal(t, http.StatusUnauthorized, resp.Code, "Invalid token refresh should be rejected")
		testDuration := time.Since(testStart)
		checklist.AddResult("Invalid Refresh Token", invalidRefreshSuccess, testDuration)

		if invalidRefreshSuccess {
			t.Log("✅ Invalid token refresh properly rejected")
		}
	})

	// Test 15: User Logout
	t.Run("User Logout", func(t *testing.T) {
		testStart := time.Now()
		testRouter := SetupRouter()

		req, err := http.NewRequest("GET", "/v1/user/logout", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

		if err != nil {
			t.Fatal("Error creating request:", err)
		}

		resp := httptest.NewRecorder()
		testRouter.ServeHTTP(resp, req)

		logoutSuccess := assert.Equal(t, http.StatusOK, resp.Code, "User logout should succeed")
		testDuration := time.Since(testStart)
		checklist.AddResult("User Logout", logoutSuccess, testDuration)

		if logoutSuccess {
			t.Log("✅ User logout successful")
		} else {
			t.Logf("❌ User logout failed with status: %d", resp.Code)
		}
	})

	// Test 16: Cleanup
	t.Run("Cleanup", func(t *testing.T) {
		testStart := time.Now()

		_, err := db.GetDB().Exec("DELETE FROM public.user WHERE email=$1", testEmail)
		cleanupSuccess := assert.NoError(t, err, "Database cleanup should succeed")
		testDuration := time.Since(testStart)
		checklist.AddResult("Database Cleanup", cleanupSuccess, testDuration)

		if cleanupSuccess {
			t.Log("✅ Database cleanup successful")
		} else {
			t.Logf("❌ Database cleanup failed: %v", err)
		}
	})

	// Print the final checklist
	checklist.PrintChecklist(t)

	t.Logf("\n🎯 Article API Test Complete - Checklist Above Shows Results")
}
