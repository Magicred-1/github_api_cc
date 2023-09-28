package main

import (
	"github_api/routers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setUpRoutes(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New())
	app.Static("/", "./public")
	app.Get("/api/repos/:username", routers.GetGHAllUserRepos)
	app.Get("/api/orgs/:orgname", routers.GetGHAllOrgRepo)
	app.Get("/api/repos/:username/:reponame", routers.GetGHUserRepo)
	app.Get("/api/repos/:username/:reponame/download", routers.DownloadRepoSource)
}

func main() {
	app := fiber.New(fiber.Config{})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Static("/", "./public")

	setUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
