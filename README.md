# Technical Documentation for the Backend of my Portfolio

## Introduction

The backend implements a portfolio management system using the `Gin` framework in Go. It includes features like 2FA authentication, user management, post handling, and security configurations.

---

## Project Structure

### Main Files and Directories

- **api/**: Contains controllers for handling API routes.
- **config/**: Manages application configuration.
- **middlewares/**: Includes middlewares (JWT authentication).
- **models/**: Defines data models.
- **services/**: Provides business logic and helper functions.
- **storage/**: Sets up the database connection.
- **main.go**: Application entry point.
- **.env**: Stores environment variables.

---

## File Descriptions

### `api/auth.go`

Defines endpoints related to authentication and user management:

#### Endpoints

1. **`POST /api/public/login`**: Initial authentication, sends 2FA code to the user’s email.
2. **`POST /api/public/verify_code`**: Verifies the 2FA code and generates a JWT token.
3. **`POST /api/public/register`**: Registers a new user.
4. **`POST /api/private/totp/setup`**: Sets up 2FA for an authenticated user.
5. **`POST /api/private/totp/verify`**: Validates a TOTP code.

---

### `auth/private.go`

Provides private endpoints for post management:

#### Endpoints

1. **`POST /api/private/posts`**: Creates a new post.

---

### `api/public.go`

Handles public endpoints related to posts:

#### Endpoints

1. **`GET /api/public/posts`**: Retrieves all public posts.

---

### `config/config.go`

Loads configurations from a `.env` file and initializes them:

#### Functions

- **`LoadConfig`**: Reads environment variables.
- **`parseAllowedEmails`**: Converts a list of allowed emails into a map.

---

### `middlewares/auth.go`

Implements the JWT authentication middleware:

#### Functions

- **`AuthenticateJWT`**: Validates the JWT token and adds claims to the context.

---

### `models/post.go`

Defines the `Post` model:

#### Attributes

- `ID`, `Image`, `Name`, `Description`, `Category`, `Tags`, `Media`, `StartDate`, `EndDate`, `Link`.

### `models/user.go`

Defines the `User` model:

#### Attributes

- `ID`, `Username`, `Password`, `Email`, `SecretTOTP`, `PendingCode`.

---

### `services/auth.go`

Provides services related to JWT and TOTP:

#### Functions

1. **`GenerateJWT`**: Generates a JWT token.
2. **`ValidateJWT`**: Validates a JWT token.
3. **`GenerateTOTP`**: Generates a URL and secret for TOTP.

---

### `services/email.go`

Handles email-related functionality:

#### Functions

1. **`GenerateEmailCode`**: Generates a random code.
2. **`SendEmail`**: Sends an email using SMTP.

---

### `services/posts.go`

Provides services to manage posts:

#### Functions

1. **`GetAllPosts`**: Retrieves all posts.
2. **`CreatePost`**: Creates a new post.

---

### `services/totp.go`

Provides services for TOTP code generation and validation:

#### Functions

1. **`GenerateTOTPSecret`**: Generates a new TOTP secret.
2. **`GenerateTOTPCode`**: Generates a TOTP code based on a secret.
3. **`ValidateTOTP`**: Validates a TOTP code.

---

### `services/users.go`

Handles user management:

#### Functions

1. **`CreateUser`**: Creates a user in the database.
2. **`GetUserByUsername`**: Retrieves a user by username.
3. **`HashPassword`**: Hashes a password.
4. **`CheckPassword`**: Compares plain text and hashed passwords.
5. **`SaveTOTPSecret`**, **`GetTOTPSecret`**: Manages TOTP secrets.

---

### `storage/database.go`

Configures the SQLite database connection and runs migrations.

#### Functions

1. **`InitDatabase`**: Initializes the database connection.
2. **`RunMigrations`**: Runs migrations for defined models.

---

### `.env`

Contains the following environment variables:

- `PORT`: Application port.
- `SECRET_KEY`: Secret key for JWT.
- `DATABASE`: Database path.
- `ISSUER_NAME`: JWT issuer name.
- `ALLOWED_EMAILS`: List of allowed email addresses.

---

### `main.go`

Entry point that initializes configuration, database, and routes:

1. **Public Routes**: Login, registration, 2FA verification, and post retrieval.
2. **Private Routes**: 2FA setup and post management.
3. **Middleware**: CORS and JWT authentication.

---

## Configuration

To run the project:

1. Configure variables in `.env`.
2. Initialize the database with `InitDatabase()`.
3. Start the server with `go run main.go`.

---

## Dependencies

- **`gin`**: HTTP framework for Go.
- **`gorm`**: ORM for database management.
- **`bcrypt`**: Password hashing.
- **`totp`**: TOTP generation and validation.
- **`godotenv`**: Environment variable management.

---

## Future Improvements

1. Implement unit tests for endpoints and services.
2. Add support for additional databases (e.g., PostgreSQL).
3. Enhance error handling for clearer messages.

---
