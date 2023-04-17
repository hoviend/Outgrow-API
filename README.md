# Outgrow-API
outgrow project backend service

<!-- ### Pre-Commit

Please install `pre-commit` to conform code standard. -->

### Local Run

Please create and migrate DB for local with `make db.create && make migrate.apply`.

Then setup the environment variable and run the app from `/cmd/runserver`.

## Architecture

### Config

Update the `Config` struct on `/config/config.go` to read config from env variable.

### Routing

The app will listen on port `PORT` (from `config.Config.Port`). Then
the request path will be routed by [gofiber](https://github.com/gofiber/fiber) thru routes defined in
`handler.SetupRoutes` (`/handler/router.go`).

### Middleware

Middleware is used to guard or modifying request or response before or after
accessing handler. For example all `/users` prefix use middleware to authenticate
request by bearer token.

```go
// file: /handler/router.go
package handler

func (h *RouteHandler) SetupRoutes(app *fiber.App, cfg *config.Config) {
	...

	userAPI := v1.Group("/users", middleware.GoogleIDTokenMiddleware())
	userAPI.Get("/:id", h.UserHandler.GetUser)

	...
}
```

### Handler

Each requests will be handled by handler defined in `/handler`. 

### Request

The `dto` models is stored in `/dto` with suffix `Payload` for request body.
The request query wil use suffix `Param` to distinguish them.

```go
// file: dto/payloads.go
package dto

type SomethingPayload {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type SomethingParam {
	Search string `json:"search"`
}
```

### Response

The `dto` models is stored in `/dto` wit suffix `Response` for respnose body.

```go
// file: dto/replies.go
package dto

type SomethingResponse struct {
	GUID uuid.UUID `json:"guid"`
}
```

### Validation

Validation request body or other struct may using `validator.Validate` from
`go-playground/validator/v10` package.

```go
// file: dto/payloads.go
package dto

type SomethingPayload {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
```

### Error Handling

- TBD

## Database

### Models

All database entitiy models or domain stored in `/ent`. Please refer to this [entgo](https://entgo.io/).

### Migration

Migration steps defined in `migrations`. Run `atlas` tools to do migration (refer to [this](https://atlasgo.io/))

To generate migration file with auto sequence number, use this command: `make migrate.new file=file_name`, or it can generated manually by this command `make migrate.manual.new file=file_name`.

Example below:

```bash
make migrate.new file=create_users_table`
```