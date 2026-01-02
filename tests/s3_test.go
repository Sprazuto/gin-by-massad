//go:build all
// +build all

package tests

import (
	"context"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/stretchr/testify/assert"
)

/**
* TestS3Connection
* Tests S3/MinIO connection and basic operations with visual checklist
*
* Must pass for S3 functionality to work
* Outputs a checklist with ✅/❌ indicators for easy visual scanning
 */
func TestS3Connection(t *testing.T) {
	// Initialize checklist for visual test tracking
	checklist := NewTestChecklist()
	startTime := time.Now()

	// Print test header with clear separation
	t.Log("\n" + strings.Repeat("═", 80))
	t.Logf("🚀 S3 CONNECTION TEST SUITE")
	t.Log(strings.Repeat("═", 80))

	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file:", err)
	}
	checklist.AddResult("Environment Setup", err == nil, time.Since(startTime))

	// Get configuration from environment
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	bucketName := os.Getenv("MINIO_BUCKET_NAME")

	// Verify required environment variables
	envCheck := assert.NotEmpty(t, endpoint, "MINIO_ENDPOINT environment variable is required") &&
		assert.NotEmpty(t, accessKeyID, "MINIO_ACCESS_KEY environment variable is required") &&
		assert.NotEmpty(t, secretAccessKey, "MINIO_SECRET_KEY environment variable is required") &&
		assert.NotEmpty(t, bucketName, "MINIO_BUCKET_NAME environment variable is required")

	checklist.AddResult("Environment Variables Validation", envCheck, 0)

	t.Logf("🚀 Testing S3 connection to endpoint: %s", endpoint)
	t.Logf("🪣 Using bucket: %s", bucketName)

	// Test 1: Basic DNS resolution
	t.Run("DNS Resolution", func(t *testing.T) {
		testStart := time.Now()
		host, _, err := net.SplitHostPort(endpoint + ":443")
		if err != nil {
			// If no port specified, use the endpoint as host
			host = endpoint
		}

		ips, err := net.LookupIP(host)
		dnsSuccess := assert.NoError(t, err, "DNS resolution should succeed") &&
			assert.NotEmpty(t, ips, "Should have at least one IP address")

		testDuration := time.Since(testStart)
		checklist.AddResult("DNS Resolution", dnsSuccess, testDuration)

		if dnsSuccess {
			t.Logf("✅ DNS resolution successful for %s:", host)
			for _, ip := range ips {
				t.Logf("   📡 %s", ip.String())
			}
		}
	})

	// Test 2: TCP connection test
	t.Run("TCP Connection", func(t *testing.T) {
		testStart := time.Now()
		host, port, _ := net.SplitHostPort(endpoint + ":443")
		if host == "" {
			host = endpoint
			port = "443"
		}

		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 10*time.Second)
		tcpSuccess := err == nil
		testDuration := time.Since(testStart)
		checklist.AddResult("TCP Connection", tcpSuccess, testDuration)

		if err != nil {
			t.Fatalf("❌ TCP connection failed: %v", err)
		}
		defer conn.Close()

		t.Logf("✅ TCP connection successful to %s:%s", host, port)
	})

	// Test 3: HTTP/S connection test
	t.Run("HTTP/S Connection", func(t *testing.T) {
		testStart := time.Now()
		host, port, _ := net.SplitHostPort(endpoint + ":443")
		if host == "" {
			host = endpoint
			port = "443"
		}

		client := &http.Client{
			Timeout: 15 * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   10 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 5 * time.Second,
				ResponseHeaderTimeout: 15 * time.Second,
			},
		}

		url := "https://" + net.JoinHostPort(host, port)
		resp, err := client.Get(url)
		httpSuccess := err == nil && resp != nil
		testDuration := time.Since(testStart)
		checklist.AddResult("HTTP/S Connection", httpSuccess, testDuration)

		if err != nil {
			t.Fatalf("❌ HTTP/S connection failed: %v", err)
		}
		defer resp.Body.Close()

		t.Logf("✅ HTTP/S connection successful - Status: %s", resp.Status)
	})

	// Test 4: MinIO client initialization
	t.Run("MinIO Client Initialization", func(t *testing.T) {
		testStart := time.Now()
		// Create custom transport with timeouts
		customTransport := &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          10,
			IdleConnTimeout:       30 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 5 * time.Second,
			ResponseHeaderTimeout: 30 * time.Second,
		}

		minioClient, err := minio.New(endpoint, &minio.Options{
			Creds:     credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure:    true,
			Transport: customTransport,
			Region:    "us-east-1",
		})
		clientSuccess := assert.NoError(t, err, "MinIO client initialization should succeed") &&
			assert.NotNil(t, minioClient, "MinIO client should not be nil")
		testDuration := time.Since(testStart)
		checklist.AddResult("MinIO Client Initialization", clientSuccess, testDuration)

		if clientSuccess {
			t.Log("✅ MinIO client initialized successfully")
		}
	})

	// Test 5: MinIO bucket operations
	t.Run("MinIO Bucket Operations", func(t *testing.T) {
		// Create custom transport with timeouts
		customTransport := &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          10,
			IdleConnTimeout:       30 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 5 * time.Second,
			ResponseHeaderTimeout: 30 * time.Second,
		}

		minioClient, err := minio.New(endpoint, &minio.Options{
			Creds:     credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure:    true,
			Transport: customTransport,
			Region:    "us-east-1",
		})
		assert.NoError(t, err, "MinIO client initialization should succeed")

		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Test listing buckets
		t.Run("ListBuckets", func(t *testing.T) {
			testStart := time.Now()
			buckets, err := minioClient.ListBuckets(ctx)
			listSuccess := assert.NoError(t, err, "ListBuckets should succeed") &&
				assert.NotNil(t, buckets, "Buckets list should not be nil")
			testDuration := time.Since(testStart)
			checklist.AddResult("List Buckets", listSuccess, testDuration)

			if listSuccess {
				t.Logf("✅ ListBuckets successful - Found %d buckets:", len(buckets))
				for _, bucket := range buckets {
					t.Logf("   🪣 %s (created: %s)", bucket.Name, bucket.CreationDate)
				}
			}
		})

		// Test bucket existence
		t.Run("BucketExists", func(t *testing.T) {
			testStart := time.Now()
			exists, err := minioClient.BucketExists(ctx, bucketName)
			existsSuccess := assert.NoError(t, err, "BucketExists check should succeed")
			testDuration := time.Since(testStart)
			checklist.AddResult("Bucket Existence Check", existsSuccess, testDuration)

			if existsSuccess {
				if exists {
					t.Logf("✅ Bucket '%s' exists", bucketName)
				} else {
					t.Logf("ℹ️  Bucket '%s' does not exist", bucketName)
				}
			}
		})

		// Test bucket creation (if it doesn't exist)
		t.Run("MakeBucket", func(t *testing.T) {
			testStart := time.Now()
			exists, err := minioClient.BucketExists(ctx, bucketName)
			existsCheck := assert.NoError(t, err, "BucketExists check should succeed")

			var createSuccess bool
			if existsCheck && !exists {
				err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
				createSuccess = assert.NoError(t, err, "MakeBucket should succeed")
				if createSuccess {
					t.Logf("✅ Bucket '%s' created successfully", bucketName)
				}
			} else if existsCheck && exists {
				createSuccess = true
				t.Logf("ℹ️  Bucket '%s' already exists, skipping creation", bucketName)
			} else {
				createSuccess = false
			}

			testDuration := time.Since(testStart)
			checklist.AddResult("Create Bucket (if needed)", createSuccess, testDuration)
		})
	})

	// Test 6: Simple upload/download test
	t.Run("File Upload/Download", func(t *testing.T) {
		// Create custom transport with timeouts
		customTransport := &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          10,
			IdleConnTimeout:       30 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 5 * time.Second,
			ResponseHeaderTimeout: 30 * time.Second,
		}

		minioClient, err := minio.New(endpoint, &minio.Options{
			Creds:     credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure:    true,
			Transport: customTransport,
			Region:    "us-east-1",
		})
		assert.NoError(t, err, "MinIO client initialization should succeed")

		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		testObjectName := "test-connection.txt"
		testContent := "This is a test file to verify S3 connectivity"

		// Test upload
		t.Run("PutObject", func(t *testing.T) {
			testStart := time.Now()
			_, err = minioClient.PutObject(ctx, bucketName, testObjectName,
				strings.NewReader(testContent), int64(len(testContent)),
				minio.PutObjectOptions{
					ContentType: "text/plain",
				})
			uploadSuccess := assert.NoError(t, err, "PutObject should succeed")
			testDuration := time.Since(testStart)
			checklist.AddResult("File Upload", uploadSuccess, testDuration)

			if uploadSuccess {
				t.Log("✅ PutObject successful")
			}
		})

		// Test download
		t.Run("GetObject", func(t *testing.T) {
			testStart := time.Now()
			obj, err := minioClient.GetObject(ctx, bucketName, testObjectName, minio.GetObjectOptions{})
			getSuccess := assert.NoError(t, err, "GetObject should succeed")

			var downloadSuccess bool
			if getSuccess {
				defer obj.Close()

				// Read the content
				buffer := new(strings.Builder)
				_, err = io.Copy(buffer, obj)
				readSuccess := assert.NoError(t, err, "Reading object should succeed")

				if readSuccess {
					downloadedContent := buffer.String()
					contentMatch := assert.Equal(t, testContent, downloadedContent, "Downloaded content should match uploaded content")
					downloadSuccess = contentMatch
					if contentMatch {
						t.Log("✅ Content verification successful")
					}
				} else {
					downloadSuccess = false
				}
			} else {
				downloadSuccess = false
			}

			testDuration := time.Since(testStart)
			checklist.AddResult("File Download & Verify", downloadSuccess, testDuration)
		})

		// Clean up - delete test object
		t.Run("RemoveObject", func(t *testing.T) {
			testStart := time.Now()
			err := minioClient.RemoveObject(ctx, bucketName, testObjectName, minio.RemoveObjectOptions{})
			removeSuccess := assert.NoError(t, err, "RemoveObject should succeed")
			testDuration := time.Since(testStart)
			checklist.AddResult("File Cleanup", removeSuccess, testDuration)

			if removeSuccess {
				t.Log("✅ RemoveObject successful")
			}
		})
	})

	// Print the final checklist
	checklist.PrintChecklist(t)

	t.Logf("\n🎯 S3 Connection Test Complete - Checklist Above Shows Results")
}