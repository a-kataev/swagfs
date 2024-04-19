# swagfs

## Swagger-UI

Update dist files in `files/dist/*`

```bash
./swagger-ui-dist.sh 5.16.2
```

## Example

Use `echo`

```golang
package main

import (
	"log"

	"github.com/a-kataev/swagfs"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg := &swagfs.Config{}
	cfg.AddURL("https://petstore.swagger.io/v2/swagger.json", "petstore")
	cfg.SetLayout("StandaloneLayout")

	e.StaticFS("/*", swagfs.FS(cfg))
	
	if err := e.Start("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}
```
