# Portfolio Backend

This project serves as the backend for the portfolio, designed to handle authentication (JWT + TOTP), post publishing, and user management using `Gin`, `SQLite`, and other modern tools from the Go ecosystem.

## Project Structure

```
portfolio-backend/
├── main.go        # Main entry point of the application
├── api/           # Public and private endpoints
│   ├── auth.go       # Authentication-related routes
│   ├── public.go     # Public routes
│   ├── private.go    # Private routes
├── models/        # Data models
│   ├── user.go       # User model
│   ├── post.go       # Post model
├── services/      # Business logic
│   ├── auth.go       # JWT and TOTP handling
│   ├── users.go      # User management
│   ├── posts.go      # Post management
├── storage/       # Persistence layer
│   ├── database.go   # Database initialization
├── middlewares/   # Middleware
│   ├── auth.go       # JWT validation middleware
├── config/        # Project configuration
│   ├── config.go    
├── .gitignore     # Ignored files and directories by Git
├── go.mod         # Dependencies
├── go.sum         # Dependency hashes
└── .env           # Environment variables (ignored by Git)
```

## Features

### Authentication
- **JWT:** Token generation and validation to secure private routes.
- **TOTP (Two-Factor Authentication):** Optional integration for an additional security layer.

### User Management
- Registration and authentication.
- Secure password storage (bcrypt).
- Storage of TOTP secrets.

### Post Management
- Creation and retrieval of posts.
- Use of JSON to store tags and associated media.

### Database
- **SQLite:** Lightweight embedded database.
- Automatic table creation on project initialization.

### Configuration
- Use of a `.env` file to define variables such as the execution port and secret key.

## Installation and Setup

### Prerequisites
- Go 1.18+
- SQLite3 installed (optional if the appropriate Go binary is available).

### Steps
1. Clone the repository:
   ```bash
   git clone <URL_REPO>
   cd portfolio-backend
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure environment variables in a `.env` file:
   ```env
   PORT=8080
   SECRET_KEY=my_very_secret_password
   DATABASE=./portfolio.db
   ISSUER_NAME=PortfolioBackend
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

5. The application will be available at `http://localhost:8080`.

## Key Routes

### Public
- `POST /api/public/login`: Authenticate and generate a JWT token.
- `POST /api/public/register`: Register a new user.

### Private (Protected by JWT)
- `POST /api/private/totp/setup`: Initial TOTP setup.
- `POST /api/private/totp/verify`: TOTP code verification.
- `POST /api/private/posts`: Create a new post.

## Ignored Files by Git

- `.env`: Contains sensitive environment variables.
- Compiled binaries and directories like `bin/` and `pkg/`.
- Temporary files from operating systems or IDEs, such as `.DS_Store`, `.vscode/`, `.idea/`.

## Testing

It is recommended to use tools like [Postman](https://www.postman.com/) or [Thunder Client](https://marketplace.visualstudio.com/items?itemName=rangav.vscode-thunder-client) to test the endpoints.

## Contributions

1. Fork the project.
2. Create a new branch for your changes:
   ```bash
   git checkout -b feature/new-feature
   ```
3. Make your changes and create a descriptive commit:
   ```bash
   git commit -m "Added new feature X"
   ```
4. Push the changes to your fork and open a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

---

Thank you for contributing and using this project!
