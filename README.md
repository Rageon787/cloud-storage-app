# Cloud Service

This is a lightweight cloud storage service — a RESTful API written in Golang using:
	•	Gin-Gonic, An HTTP web framework
	•	SQLite, Database for storing user accounts 
	•	Azure Storage Accounts for file storage in the cloud 

It includes support for file upload, deletion, retrieval, and user authentication — all with zero-knowledge encryption! 

## Features 
	🔐 User registration and authentication (with hashed passwords)
  •	🗃️ File upload, retrieval, and deletion
	•	🧾 SQLite for local account management and file metadata
	•	☁️ Azure Blob Storage integration
	•	🔒 Zero-knowledge encryption — files are encrypted client-side
	•	📦 RESTful API design following best practices

 ## API Endpoints
Authentication
	•	POST /register
Register a new user with email and password.
	•	POST /login
Authenticate and receive a JWT token.


📁 File Operations
	•	POST /upload
Upload a file. Requires authentication.
	•	GET /files/:id
Retrieve a file by ID. Requires authentication.
	•	DELETE /files/:id
Delete a file by ID. Requires authentication.
	•	GET /files
List all uploaded files for the authenticated user.
