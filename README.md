# Live project on GKE

- Visit https://www.donot.cards and have fun!

# How to run the project locally

- Make sure you have Docker and docker-compose installed.
- Run `make cert` to create self-signed certificates for the services.
- Duplicate `auth.env.example` and rename it to `auth.env`. Inside that file you will have to replace the XXXX to your google project client id and client id in order to make Google oAuth2 work. See https://developers.google.com/adwords/api/docs/guides/authentication#create_a_client_id_and_client_secret
- In your google project authorized redirect URI's, add https://localhost/callback
- Run `make start` to build and start all services.
- Go to https://localhost
- Enjoy