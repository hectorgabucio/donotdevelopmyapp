# Do Not Develop My App

A Rick and Morty card collection application built with microservices architecture. Collect and manage your favorite character cards with Google authentication.

## Features

- 🎴 Collect new Rick and Morty character cards
- 📋 List and manage all your collected cards
- 🔐 Secure authentication with Google Sign-In
- 🎨 Modern React-based user interface
- 🚀 Microservices architecture for scalability
- 🐳 Fully containerized with Docker
- ☸️ Kubernetes-ready for production deployment

## Tech Stack

### Backend
- **Language**: Go 1.17+
- **Architecture**: Microservices
- **Communication**: gRPC and HTTP
- **Database**: CockroachDB
- **Caching**: Redis
- **Authentication**: OAuth2 (Google)

### Frontend
- **Framework**: React 16.13+
- **UI Library**: Material-UI
- **HTTP Client**: Axios
- **Routing**: React Router

### Infrastructure
- **Containerization**: Docker & Docker Compose
- **Orchestration**: Kubernetes (GKE)
- **CI/CD**: GitHub Actions
- **TLS**: Self-signed certificates for local development

## Project Structure

```
donotdevelopmyapp/
├── cmd/                    # Application entry points
│   ├── auth/              # Authentication service
│   ├── backend/           # Main backend service
│   ├── characters/        # Character service
│   └── random-micro/      # Random service
├── internal/              # Internal packages
│   ├── auth/              # Auth protobuf definitions
│   ├── character/         # Character protobuf definitions
│   ├── random/            # Random protobuf definitions
│   ├── cipher/            # Encryption utilities
│   ├── config/            # Configuration management
│   ├── data/              # Data access layer
│   ├── jwt/               # JWT handling
│   ├── oauth/             # OAuth implementation
│   └── server/            # Server utilities
├── website/               # React frontend application
├── deployments/           # Docker Compose and deployment configs
├── k8s/                   # Kubernetes manifests
├── test/                  # Test utilities and mocks
└── tls/                   # TLS certificates
```

## Prerequisites

- Docker and Docker Compose
- Go 1.17+ (for local development)
- Node.js and npm (for frontend development)
- Google Cloud Project with OAuth2 credentials
- OpenSSL (for certificate generation)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/hectorgabucio/donotdevelopmyapp.git
cd donotdevelopmyapp
```

### 2. Set Up Google OAuth2

1. Create a Google Cloud Project or use an existing one
2. Enable the Google+ API
3. Create OAuth 2.0 credentials (Client ID and Client Secret)
4. Add `https://localhost/callback` to authorized redirect URIs
5. See [Google OAuth2 Documentation](https://developers.google.com/adwords/api/docs/guides/authentication#create_a_client_id_and_client_secret) for detailed instructions

### 3. Configure Environment Variables

```bash
cp auth.env.example auth.env
```

Edit `auth.env` and replace the placeholders with your Google OAuth2 credentials:
- Replace `XXXX` with your Google Client ID
- Replace `XXXX` with your Google Client Secret

### 4. Generate TLS Certificates

```bash
make cert
```

This creates self-signed certificates for secure gRPC communication between services.

### 5. Start the Application

```bash
make start
```

This will:
- Run code style checks
- Build all Docker images
- Start all services using Docker Compose

### 6. Access the Application

Open your browser and navigate to:
```
https://localhost
```

**Note**: You may need to accept the self-signed certificate warning in your browser.

## Development

### Running Tests

Run all tests (backend and frontend):
```bash
make test
```

Run backend tests only:
```bash
make test-back
```

Run frontend tests only:
```bash
make test-front
```

Generate test coverage:
```bash
make cov
```

### Code Style

Check code style compliance:
```bash
make check-style
```

This runs:
- `golangci-lint` for Go code
- ESLint for JavaScript/React code

### Generate Mocks

Generate mock files for testing:
```bash
make mocks
```

### Clean Up

Remove all containers and volumes:
```bash
make clean
```

## Available Make Targets

| Target | Description |
|--------|-------------|
| `make cert` | Generate TLS certificates |
| `make check-style` | Run linters on codebase |
| `make start` | Start local environment with Docker Compose |
| `make clean` | Remove containers and volumes |
| `make test` | Run all tests |
| `make test-back` | Run backend tests |
| `make test-front` | Run frontend tests |
| `make cov` | Generate test coverage |
| `make mocks` | Generate mock files |
| `make help` | Show all available targets |

## Deployment

### Kubernetes Deployment

The project includes Kubernetes manifests in the `k8s/` directory. Deploy to GKE or any Kubernetes cluster:

```bash
kubectl apply -k k8s/
```

Services are deployed in the following order:
1. Random service
2. Character service
3. Auth service
4. Backend service
5. Website

### CI/CD

The project uses GitHub Actions for continuous integration and deployment. The workflow includes:
- Code style checks
- Running tests
- Building Docker images
- Deploying to Google Kubernetes Engine (GKE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Rick and Morty API for character data
- Material-UI for React components
- The open-source community for amazing tools and libraries
