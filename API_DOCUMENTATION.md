# LKE Application API Documentation

## Overview

This document provides comprehensive documentation for the LKE (Lembar Kerja Evaluasi) Application API, an Indonesian Government Evaluation System. The API is organized by modules and includes detailed information about all endpoints, parameters, request/response schemas, authentication requirements, and usage examples.

**Base URL:** `http://localhost:8080/v1`
**Version:** 1.4.11.1
**Authentication:** Bearer Token (JWT)

## Table of Contents

1. [Authentication](#authentication)
2. [Articles](#articles)
3. [LKE Rekap](#lke-rekap)
4. [LKE Evaluasi](#lke-evaluasi)
5. [LKE Komponen](#lke-komponen)
6. [LKE Rekomendasi](#lke-rekomendasi)
7. [Application](#application)

## Authentication

All protected endpoints require a Bearer token in the Authorization header.

### POST /v1/user/login

**User Login**

Authenticate user and return access token.

**Parameters:**

- `loginForm` (body, required): Login credentials
  - `email` (string): User email
  - `password` (string, 3-50 chars): User password

**Responses:**

- `200`: Successfully logged in with user data and token
- `400`: Invalid request
- `401`: Invalid credentials

**Example:**

```bash
curl -X POST http://localhost:8080/v1/user/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'
```

### POST /v1/user/register

**User Registration**

Register a new user account.

**Parameters:**

- `registerForm` (body, required): User registration data
  - `email` (string): User email
  - `name` (string, 3-20 chars): Full name
  - `password` (string, 3-50 chars): User password

**Responses:**

- `201`: Successfully registered
- `400`: Invalid request or validation failed

### GET /v1/user/logout

**User Logout**

Logout user and invalidate access token.

**Parameters:**

- `Authorization` (header, required): Bearer token

**Responses:**

- `200`: Successfully logged out
- `401`: Unauthorized

### POST /v1/token/refresh

**Refresh Access Token**

Refresh access token using refresh token.

**Parameters:**

- `tokenForm` (body, required): Token refresh request
  - `refresh_token` (string, required): Refresh token

**Responses:**

- `200`: Successfully refreshed token
- `401`: Invalid or expired token

## Articles

Article management endpoints for authenticated users.

### POST /v1/article

**Create Article**

Create a new article.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `createForm` (body, required): Article data
  - `title` (string, 3-100 chars, required): Article title
  - `content` (string, 3-1000 chars, required): Article content

**Responses:**

- `200`: Article created
- `406`: Validation failed

### GET /v1/articles

**Get All Articles**

Get all articles for the authenticated user.

**Parameters:**

- `Authorization` (header, required): Bearer token

**Responses:**

- `200`: Articles list
- `406`: Could not get articles

### GET /v1/article/{id}

**Get Article by ID**

Get a specific article by ID.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `id` (path, required): Article ID (integer)

**Responses:**

- `200`: Article data
- `404`: Article not found or invalid parameter

### PUT /v1/article/{id}

**Update Article**

Update an existing article by ID.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `id` (path, required): Article ID (integer)
- `updateForm` (body, required): Updated article data
  - `title` (string, 3-100 chars, required): Article title
  - `content` (string, 3-1000 chars, required): Article content

**Responses:**

- `200`: Article updated
- `404`: Invalid parameter
- `406`: Validation failed or could not update

### DELETE /v1/article/{id}

**Delete Article**

Delete an article by ID.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `id` (path, required): Article ID (integer)

**Responses:**

- `200`: Article deleted
- `404`: Invalid parameter
- `406`: Could not delete

## LKE Rekap

LKE (Lembar Kerja Evaluasi) Rekap management endpoints.

### POST /v1/lke-rekap

**Create LKE Rekap**

Create a new LKE Rekap record.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `createForm` (body, required): LKE Rekap data
  - `id_opd` (integer, required): OPD ID
  - `tahun` (integer, required): Year
  - `kelengkapan` (number, required): Completeness score
  - `nilai_capaian` (number, required): Achievement score
  - `predikat_akhir` (string, max 50, required): Final predicate
  - `predikat` (string, max 50, required): Predicate
  - `kelengkapan_m` (number, required): Modified completeness score
  - `nilai_capaian_m` (number, required): Modified achievement score
  - `predikat_akhir_m` (string, max 50, required): Modified final predicate
  - `predikat_m` (string, max 50, required): Modified predicate
  - `status_evaluasi` (string, max 50, required): Evaluation status
  - `id_verifikator` (integer, required): Verifikator ID
  - `id_ketua` (integer, required): Ketua ID
  - `id_evaluator` (integer, required): Evaluator ID
  - `id_pengendali` (integer, required): Pengendali ID

**Responses:**

- `200`: LKE Rekap created
- `406`: Validation failed

### GET /v1/lke-rekaps

**Get All LKE Rekap**

Get all LKE Rekap records, optionally filtered by year.

**Parameters:**

- `Authorization` (header, required): Bearer token

**Responses:**

- `200`: LKE Rekap records list
- `400`: Invalid year parameter
- `406`: Could not get LKE Rekap records

### GET /v1/lke-rekaps/{tahun}

**Get LKE Rekap by Year**

Get LKE Rekap records filtered by year.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `tahun` (path): Year filter (defaults to current year - 1)

**Responses:**

- `200`: LKE Rekap records list
- `400`: Invalid year parameter
- `406`: Could not get LKE Rekap records

### GET /v1/lke-rekap/{id}

**Get LKE Rekap by ID**

Get a specific LKE Rekap record by ID.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `id` (path, required): LKE Rekap ID (integer)

**Responses:**

- `200`: LKE Rekap data
- `404`: LKE Rekap not found or invalid parameter

### PUT /v1/lke-rekap/{id}

**Update LKE Rekap**

Update an existing LKE Rekap record by ID.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `id` (path, required): LKE Rekap ID (integer)
- `updateForm` (body, required): Updated LKE Rekap data (all fields optional)

**Responses:**

- `200`: LKE Rekap updated
- `404`: Invalid parameter
- `406`: Validation failed or could not update

### DELETE /v1/lke-rekap/{id}

**Delete LKE Rekap**

Delete an LKE Rekap record by ID.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `id` (path, required): LKE Rekap ID (integer)

**Responses:**

- `200`: LKE Rekap deleted
- `404`: Invalid parameter
- `406`: Could not delete

### GET /v1/lke-rekap/opd/{id_opd}/tahun/{tahun}

**Get LKE Rekap by OPD and Year**

Get a LKE Rekap record by OPD ID and year, including evaluasi children.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `id_opd` (path, required): OPD ID (integer)
- `tahun` (path, required): Year (integer)

**Responses:**

- `200`: LKE Rekap data with evaluasi children
- `404`: LKE Rekap not found or invalid parameters

## LKE Evaluasi

LKE Evaluasi management endpoints.

### POST /v1/lke-evaluasi

**Create LKE Evaluasi**

Create a new LKE Evaluasi record or update existing one.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `createForm` (body, required): LKE Evaluasi data
  - `lke_rekap_id` (integer, required): LKE Rekap ID
  - `kode_evaluasi` (string, max 50, required): Evaluation code
  - `jawaban` (string, optional): Answer
  - `evaluasi` (string, optional): Evaluation
  - `berkas` (string, optional): File reference
  - `pranala` (string, optional): Link URL
  - `catatan` (string, optional): Notes
- `file` (formData, optional): File to upload (for PUT requests only)

**Responses:**

- `200`: LKE Evaluasi created or updated
- `400`: File required for PUT or invalid request
- `406`: Validation failed
- `500`: Failed to upload file or update rekap values

### PUT /v1/lke-evaluasi

**Update LKE Evaluasi**

Update an existing LKE Evaluasi record.

**Parameters:** Same as POST /v1/lke-evaluasi

### GET /v1/lke-evaluasis

**Get All LKE Evaluasi**

Get all LKE Evaluasi records for the authenticated user.

**Parameters:**

- `Authorization` (header, required): Bearer token

**Responses:**

- `200`: LKE Evaluasi records list
- `406`: Could not get LKE Evaluasi records

### GET /v1/lke-evaluasi/{id}

**Get LKE Evaluasi by ID**

Get a specific LKE Evaluasi record by ID.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `id` (path, required): LKE Evaluasi ID (integer)

**Responses:**

- `200`: LKE Evaluasi data
- `404`: LKE Evaluasi not found or invalid parameter

### PUT /v1/lke-evaluasi/{id}

**Update LKE Evaluasi by ID**

Update an existing LKE Evaluasi record by ID.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `id` (path, required): LKE Evaluasi ID (integer)
- `createForm` (body, required): Updated LKE Evaluasi data

**Responses:**

- `200`: LKE Evaluasi updated
- `404`: Invalid parameter
- `406`: Validation failed or could not update

### DELETE /v1/lke-evaluasi/{id}

**Delete LKE Evaluasi**

Delete an LKE Evaluasi record by ID.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `id` (path, required): LKE Evaluasi ID (integer)

**Responses:**

- `200`: LKE Evaluasi deleted
- `404`: Invalid parameter
- `406`: Could not delete

### GET /v1/lke-evaluasi/signed-url/{lke_rekap_id}/{kode_evaluasi}

**Get Signed URL for File**

Generate a signed URL to access a file associated with an LKE Evaluasi record.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `lke_rekap_id` (path, required): LKE Rekap ID (integer)
- `kode_evaluasi` (path, required): Kode Evaluasi (string)

**Responses:**

- `200`: Signed URL for file access
- `400`: Missing required parameters
- `404`: Record not found
- `500`: Failed to generate signed URL

### GET /v1/lke-evaluasi/sync-evaluasi

**Sync LKE Evaluasi**

Synchronize jawaban data into evaluasi records for the specified LKE Rekap ID.

**Parameters:**

- `Authorization` (header, required): Bearer token
- `lke_rekap_id` (query, required): LKE Rekap ID (integer)

**Responses:**

- `200`: Successfully synced evaluasi from jawaban
- `400`: Invalid lke_rekap_id parameter
- `500`: Failed to sync evaluasi or update rekap values

## LKE Komponen

LKE Komponen management endpoints.

### POST /v1/lke-komponen

**Create LKE Komponen**

Create a new LKE Komponen record.

**Parameters:**

- `createForm` (body, required): LKE Komponen data
  - `kode_evaluasi` (string, 3-50 chars, required): Evaluation code
  - `bobot` (number, required): Weight/score
  - `komponen` (string, 3-255 chars, required): Component name
  - `eviden` (string, 3-255 chars, required): Evidence
  - `level` (string, 3-50 chars, required): Level

**Responses:**

- `200`: LKE Komponen created
- `406`: Validation failed

### GET /v1/lke-komponens

**Get All LKE Komponen**

Get all LKE Komponen records.

**Responses:**

- `200`: LKE Komponen records list
- `406`: Could not get LKE Komponen records

### GET /v1/lke-komponen/{id}

**Get LKE Komponen by ID**

Get a specific LKE Komponen record by ID.

**Parameters:**

- `id` (path, required): LKE Komponen ID (integer)

**Responses:**

- `200`: LKE Komponen data
- `404`: LKE Komponen not found or invalid parameter

### PUT /v1/lke-komponen/{id}

**Update LKE Komponen**

Update an existing LKE Komponen record by ID.

**Parameters:**

- `id` (path, required): LKE Komponen ID (integer)
- `createForm` (body, required): Updated LKE Komponen data

**Responses:**

- `200`: LKE Komponen updated
- `404`: Invalid parameter
- `406`: Validation failed or could not update

### DELETE /v1/lke-komponen/{id}

**Delete LKE Komponen**

Delete an LKE Komponen record by ID.

**Parameters:**

- `id` (path, required): LKE Komponen ID (integer)

**Responses:**

- `200`: LKE Komponen deleted
- `404`: Invalid parameter
- `406`: Could not delete

## LKE Rekomendasi

LKE Rekomendasi management endpoints.

### POST /v1/lke-rekomendasi

**Create LKE Rekomendasi**

Create a new LKE Rekomendasi record or update existing one based on parent_id.

**Parameters:**

- `createForm` (body, required): LKE Rekomendasi data
  - `parent_id` (integer, required): Parent ID
  - `ta1a`, `ta1b`, `ta1c`, `ta2a`, `ta2b`, `ta2c` (string, optional): TA recommendations
  - `tb1`, `tb2`, `tb3` (string, optional): TB recommendations
  - `tc1`, `tc2`, `tc3` (string, optional): TC recommendations
  - `td1`, `td2`, `td3` (string, optional): TD recommendations

**Responses:**

- `200`: LKE Rekomendasi created or updated
- `406`: Validation failed

### GET /v1/lke-rekomendasis

**Get All LKE Rekomendasi**

Get all LKE Rekomendasi records.

**Responses:**

- `200`: LKE Rekomendasi records list
- `406`: Could not get LKE Rekomendasi records

### GET /v1/lke-rekomendasi/{id}

**Get LKE Rekomendasi by ID**

Get a specific LKE Rekomendasi record by ID.

**Parameters:**

- `id` (path, required): LKE Rekomendasi ID (integer)

**Responses:**

- `200`: LKE Rekomendasi data
- `404`: LKE Rekomendasi not found or invalid parameter

### PUT /v1/lke-rekomendasi/{id}

**Update LKE Rekomendasi**

Update an existing LKE Rekomendasi record by ID.

**Parameters:**

- `id` (path, required): LKE Rekomendasi ID (integer)
- `createForm` (body, required): Updated LKE Rekomendasi data

**Responses:**

- `200`: LKE Rekomendasi updated
- `404`: Invalid parameter
- `406`: Validation failed or could not update

### DELETE /v1/lke-rekomendasi/{id}

**Delete LKE Rekomendasi**

Delete an LKE Rekomendasi record by ID.

**Parameters:**

- `id` (path, required): LKE Rekomendasi ID (integer)

**Responses:**

- `200`: LKE Rekomendasi deleted
- `404`: Invalid parameter
- `406`: Could not delete

## Application

Application-level endpoints.

### GET /v1/version

**Get Application Version**

Get current application version and patch information.

**Responses:**

- `200`: Application version information

### GET /v1/error-logs

**Get Error Logs**

Retrieve recent application error logs for debugging.

**Parameters:**

- `Authorization` (header, required): Bearer token

**Responses:**

- `200`: Recent error logs
- `500`: Failed to get error logs

## Error Codes

- `200`: Success
- `201`: Created
- `400`: Bad Request
- `401`: Unauthorized
- `404`: Not Found
- `406`: Not Acceptable (Validation Error)
- `500`: Internal Server Error

## Notes

- All endpoints requiring authentication need a valid Bearer token in the Authorization header
- File uploads are supported for certain endpoints using multipart/form-data
- Response formats can be JSON, XML, or YAML depending on the endpoint and parameters
- Some endpoints support format specification via path parameters
- LKE endpoints are part of the Indonesian Government Evaluation System (Lembar Kerja Evaluasi)

## Validation Rules

- String lengths are validated with min/max constraints
- Required fields are marked as such in the schema definitions
- Email validation is applied where appropriate
- Numeric fields have appropriate type constraints

## Rate Limiting

Currently not implemented but may be added in future versions.

## Deprecations

No endpoints are currently deprecated in this version.
