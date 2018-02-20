# Swagger-doc middleware

Provide methods to expose swagger documentation endpoints to your API

## Installing

```bash
go get github.com/ricardo-ch/swagger-doc
```

## Getting Started

To add this in your API:

- Create an environment variable EXPOSE_DOC with boolean value 'true'. Without this, both endpoints will return 404 NotFound result

```bash
export EXPOSE_DOC=true
```

- Add import statement

```golang
import "github.com/ricardo-ch/swagger-doc"
```

- Add those 2 endpoints in your API.

```golang
// You can specify an URL where your swagger file is hosted, or if you are using the SwaggerHandler, just specify the endpoint(eg. /swagger)
router.Handle("/docs", middleware.DocumentationHandler("http://127.0.0.1:8081", "/docs")) // Add your custom port and endpoint name here
router.HandleFunc("/swagger", middleware.SwaggerHandler)
```

- /swagger will display your swagger.json if it exists, otherwise it will return 404
- /docs will render Redoc documentation for your swagger file