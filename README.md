# Features
- Collect new Rick and Morty cards.
- List all your cards login with your Google account.

# Tech overview
- Go backend.
- React frontend.
- CockroachDB as database.
- Microservices.
- All dockerised.
- Communication with GRPC and/or HTTP
- Google Sign In.
- CI/CD with Github Actions.
- Test coverage (both backend and front).
- Google cloud deployment (GKE).
- Kubernetes on live environment.
- Run locally with docker compose.

# How to run the project locally

- Make sure you have Docker and docker-compose installed.
- Run `make cert` to create self-signed certificates for the services.
- Duplicate `auth.env.example` and rename it to `auth.env`. Inside that file you will have to replace the XXXX to your google project client id and client id in order to make Google oAuth2 work. See https://developers.google.com/adwords/api/docs/guides/authentication#create_a_client_id_and_client_secret
- In your google project authorized redirect URI's, add https://localhost/callback
- Run `make start` to build and start all services.
- Go to https://localhost
- Enjoy
