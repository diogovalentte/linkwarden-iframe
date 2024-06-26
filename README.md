# 2024-04-08: Moved to the project [Homarr iFrames](https://github.com/diogovalentte/homarr-iframes)

# Linkwarden iFrame

A very simple and limited API that gets your links from your [Linkwarden](https://github.com/linkwarden/linkwarden) instance and creates a nice HTML code to be used in an iFrame (designed to be used in [Homarr](https://github.com/ajnart/homarr))

This is how the iFrame is shown on the dashboard (on the right in the image below). It is made based on the Homarr app to show movies/show requests on Jellyseer/Overseer (on the left):

![image](https://github.com/diogovalentte/linkwarden-iframe/assets/49578155/90cddcc1-7fca-4ea2-9f0e-6418209ddb80)

# How to run:
To run, the application will need two environment variables. The way you will export them to the application depends on how you will run the application:
- `LINKWARDEN_ADDRESS`: your Linkwarden instance address, like https://sub.domain.com or http://192.168.1.45:8080
- `LINKWARDEN_TOKEN`: an access token used to access your Linkwarden instance API to get your links. You can get it by going to **Settings -> Access Tokens -> New Access Token**.

## Using Docker:

1. Run the latest version:

```sh
docker run --name linkwarden-iframe -p 8080:8080 -e LINKWARDEN_ADDRESS=https://sub.domain.com -e LINKWARDEN_TOKEN=linkwarden-token ghcr.io/diogovalentte/linkwarden-iframe:latest
```

## Using Docker Compose:

1. There is a `docker-compose.yml` file in this repository. Clone this repository to use this file or create one.
2. Create a .env file with your Linkwarden instance address and token. It should be like the `.env.example` file and be in the same directory as the `docker-compose.yml` file.
3. Start the container by running:
```sh
docker compose up
```

## Manually:

1. Install the dependencies:
```sh
go mod download
```
2. Export the environment variables `LINKWARDEN_ADDRESS` and `LINKWARDEN_TOKEN`.
3. Run:
```sh
go run main.go
```

# Simple docs:
- `/v1/health`
- `/v1/links`: returns all links of all collections in a JSON. Allow the following query arguments:
  - `limit` (optional): limit the number of links returned.
  - `collectionId` (optional): return all links of a specific collection. You can get the collection ID by going to the collection page. The ID should be on the URL. The ID of the default collection **Unorganized** is 1 because the URL is https://domain.com/collections/1.
- `/v1/links/iframe`: returns all links of all collections in an HTML document that can be used as an iFrame (designed to be used with [Homarr](https://github.com/ajnart/homarr)). Allow the following query arguments:
  - `limit`: same as above.
  - `collectionId`: same as above.
  - `theme` (optional): "light" or "dark". It's used to match the HTML returned with the Homarr theme. Defaults to "light".

# Adding to Homarr
1. In your Homarr dashboard, click on **Enter edit mode -> Add a tile -> Widgets -> iFrame**.
2. Click to edit the new iFrame widget.
3. Add the API URL, like `http://192.168.1.15:8080/v1/links/iframe?collectionId=1&limit=3&theme=dark`. Change the query arguments for your needs.

# Obs:
- Anyone who can access the API will be able to **see** all information about your links, including their collections and tags. You can add an authentication portal like [Authelia](https://github.com/authelia/authelia) or [Authentik](https://github.com/goauthentik/authentik) in front of the API to secure it, this is how I do it.
