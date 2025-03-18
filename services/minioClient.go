package services

import (
	"context"
	"fmt"
	"log"
	"mime/multipart" // Added import for multipart
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

// InitializeMinioClient initializes the MinIO client
func InitializeMinioClient() {
	var err error
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := true // Set to true if using SSL

	log.Println("Initializing MinIO client...")
	log.Printf("Endpoint: %s\n", endpoint) // Added logging for the endpoint
	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func UploadFile(file *multipart.FileHeader) (string, error) {
	ctx := context.Background()
	bucketName := os.Getenv("MINIO_BUCKET_NAME") // Replace with your bucket name
	objectName := file.Filename

	// Open the file

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Check if MinioClient is initialized
	if MinioClient == nil {
		log.Println("MinIO client is not initialized")
		return "", fmt.Errorf("MinIO client is not initialized")
	}

	// Check if the bucket exists
	exists, err := MinioClient.BucketExists(ctx, bucketName)

	if err != nil {
		return "", err
	}

	// Create the bucket if it does not exist
	if !exists {
		err = MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Printf("Error creating bucket: %v\n", err)
			return "", err
		}
	}

	// Log the error if the upload fails
	if err != nil {
		log.Printf("Error uploading file: %v\n", err)
		return "", err
	}

	// Upload the file
	_, err = MinioClient.PutObject(ctx, bucketName, objectName, src, file.Size, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}

	// Return the file URL
	// fileURL := fmt.Sprintf("%s/%s/%s", MinioClient.EndpointURL(), bucketName, objectName)
	return objectName, nil
}

func DeleteObject(objectName string) error {
	ctx := context.Background()
	bucketName := os.Getenv("MINIO_BUCKET_NAME") // Replace with your bucket name

	// Remove the object from the bucket
	err := MinioClient.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		log.Printf("Error deleting object: %v\n", err)
		return err
	}

	return nil
}

// GenerateSignedURL generates a signed URL for a specified object in the MinIO bucket
func GenerateSignedURL(objectName string) (string, error) {
	ctx := context.Background()
	bucketName := os.Getenv("MINIO_BUCKET_NAME") // Replace with your bucket name

	// Generate a presigned URL for the object
	signedURL, err := MinioClient.PresignedGetObject(ctx, bucketName, objectName, 24*time.Hour, nil)
	if err != nil {
		log.Printf("Error generating signed URL: %v\n", err)
		return "", err
	}

	return signedURL.String(), nil
}
