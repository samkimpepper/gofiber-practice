package main

import (
	"context"
	"go-note/module/auth"
	"log"

	"go-note/ent"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {
	app := fiber.New(*getConfig())
	app.Use(cors.New(*getCorsConfig()))

	db := dbConnect()

	auth.Routes(app.Group("/auth"), db)

	app.Listen(":3000")
}

func getConfig() *fiber.Config {
	return &fiber.Config{
		Prefork: false,
		AppName: "GoNote",
		Views:   getViewHandler(),
	}
}

func getViewHandler() *html.Engine {
	handler := html.New("./views", ".html")
	return handler
}

func dbConnect() *ent.Client {
	client, err := ent.Open("mysql", "vegielcl:akvkenqn12@tcp(turi-db.crnshyl5ky2g.us-east-1.rds.amazonaws.com:3306)/gonote?parseTime=True")
	if err != nil {
		log.Fatalf("failed connection database: %v", err)
	}

	//defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema: %v", err)
	}

	return client
}

func getCorsConfig() *cors.Config {
	return &cors.Config{
		AllowCredentials: true,
	}
}
