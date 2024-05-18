package main

import (
	"fmt"
	"github.com/SawitProRecruitment/UserService/router"
	service2 "github.com/SawitProRecruitment/UserService/service"
	"github.com/joho/godotenv"
	"os"

	"github.com/SawitProRecruitment/UserService/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	newServer()
}

func newServer() {
	e := echo.New()

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	dbDsn := os.Getenv("DATABASE_URL")
	fmt.Println("db: ", dbDsn)
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})

	var service = service2.NewService(repo)

	router.NewRouter(e, service)

	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":8080"))
}
