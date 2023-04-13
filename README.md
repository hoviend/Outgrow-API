# Outgrow-API
outgrow project backend service

### Pre-Commit

Please install `pre-commit` to conform code standard.

### Local Run

Please create and migrate DB for local with `make db.create && make migrate.apply`.

Then setup the environment variable and run the app from `/cmd/runserver`.

## Architecture

### Config

Update the `Config` struct on `/config/config.go` to read config from env variable.

### Routing

The app will listen on port `PORT` (from `config.Config.Port`). Then
the request path will be routed by [gorilla/mux](https://github.com/gorilla/mux) thru routes defined in
`handler.NewRouter` (`/handler/router.go`).