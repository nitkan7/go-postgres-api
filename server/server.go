package server

import (
	db "api/DB"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type response struct {
	Id    int
	Value string
}

func ServerStart() {

	app := fiber.New()

	output := response{}
	output.Id, output.Value = db.DBGet()
	ouputJson := jsonMaker(output)
	fmt.Println(ouputJson)
	hello := ouputJson

	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send(hello)
	})

	app.Listen(":3000")
}

func jsonMaker(hello interface{}) []byte {

	json, err := json.Marshal(hello)

	if err != nil {
		panic(err)
	}
	return json
}
