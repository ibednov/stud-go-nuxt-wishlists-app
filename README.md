# go-nuxt-wishlists-app

Guide:

1. Change `.env` file
2. Create service: `make create-service`
3. Check if `go.mod` file is correct

    ``` go
    module wishlists

    go 1.23

    require (
    github.com/gin-gonic/gin v1.10.0
    github.com/joho/godotenv v1.5.1
    gorm.io/driver/mysql v1.5.7
    gorm.io/gorm v1.25.12
    )

    require (
    filippo.io/edwards25519 v1.1.0 // indirect
    github.com/bytedance/sonic v1.12.4 // indirect
    github.com/bytedance/sonic/loader v0.2.1 // indirect
    github.com/cloudwego/base64x v0.1.4 // indirect
    github.com/cloudwego/iasm v0.2.0 // indirect
    github.com/gabriel-vasile/mimetype v1.4.6 // indirect
    github.com/gin-contrib/sse v0.1.0 // indirect
    github.com/go-playground/locales v0.14.1 // indirect
    github.com/go-playground/universal-translator v0.18.1 // indirect
    github.com/go-playground/validator/v10 v10.22.1 // indirect
    github.com/go-sql-driver/mysql v1.8.1 // indirect
    github.com/goccy/go-json v0.10.3 // indirect
    github.com/jinzhu/inflection v1.0.0 // indirect
    github.com/jinzhu/now v1.1.5 // indirect
    github.com/json-iterator/go v1.1.12 // indirect
    github.com/klauspost/cpuid/v2 v2.2.9 // indirect
    github.com/leodido/go-urn v1.4.0 // indirect
    github.com/mattn/go-isatty v0.0.20 // indirect
    github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
    github.com/modern-go/reflect2 v1.0.2 // indirect
    github.com/pelletier/go-toml/v2 v2.2.3 // indirect
    github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
    github.com/ugorji/go/codec v1.2.12 // indirect
    golang.org/x/arch v0.12.0 // indirect
    golang.org/x/crypto v0.29.0 // indirect
    golang.org/x/net v0.31.0 // indirect
    golang.org/x/sys v0.27.0 // indirect
    golang.org/x/text v0.20.0 // indirect
    google.golang.org/protobuf v1.35.1 // indirect
    gopkg.in/yaml.v3 v3.0.1 // indirect
    )

    ```

4. Create Go Main file:`touch app/back/services/{service-name}/main.go`

    ``` go
    package main

    import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    )

    // Определение модели
    type YourModel struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"size:255"`
    // Добавьте другие поля по необходимости
    }

    func init() {
    if err := godotenv.Load(); err != nil {
    log.Fatal("Error loading .env file")
    }
    }

    func main() {
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
    log.Fatalf("failed to connect to database: %v", err)
    }

    // Используйте переменную db для выполнения операций с базой данных
    if err := db.AutoMigrate(&YourModel{}); err != nil {
    log.Fatalf("failed to migrate database: %v", err)
    }

    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
    c.String(200, "Hello World")
    })

    router.GET("/env", func(c *gin.Context) {
    c.String(200, "Hello %s", os.Getenv("NAME"))
    })

    router.Run(":8080")
    }

    ```

5. Run `make build` command
6. Run `make up` command
7. Check if service is running:
   1. Docker Desktop and check if container is running
   2. Check if `curl http://localhost:8080/` returns `Hello World`
   3. Check if `curl http://localhost:8080/env` returns `Hello {username}`

---

## Installing [Swagger](https://github.com/swaggo/gin-swagger)

<!-- 1. Install new package by running: `make update-dependencies`
    1. `github.com/swaggo/swag/cmd/swag`
    2. `github.com/swaggo/gin-swagger`
    3. `github.com/swaggo/files` -->
2. Run command to get into container: `make o-back`
3. And then into container run `go install github.com/swaggo/swag/cmd/swag@latest`
4. Run `make o-back` to go into container
5. Into container run `swag init`
6. Then you can see `docs` folder with `docs.go` file into your project directory
7. After this install this packages:
   1. `go get -u github.com/swaggo/gin-swagger`
   2. `go get -u github.com/swaggo/files`
8. Then add annotations to your routes in `routes.go` file
```
        // @Summary Get all wishlists
        // @Description Get a list of all wishlists
        // @Tags wishlists
        // @Accept json
        // @Produce json
        // @Success 200 {array} models.Wishlist
        // @Router /wishlists [get]
        func GetWishlists(c *gin.Context) {
            // Your code for getting wishlists
        }
```

1. Initialize swagger into your project
