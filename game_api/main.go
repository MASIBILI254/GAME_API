package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

// create a game request struct

type GameRequest struct {
	PlayerID string `json:"player_id"`
	GameID   string `json:"game_id"`
}

// create a game response struct
type GameResponse struct {
	PlayerID string `json:"player_id"`
	GameID   string `json:"game_id"`
	Results  string `json:"results"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	app := fiber.New()

	//create server

	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{"message": "Game serve API running"})

	})

	// play game

	app.Post("/play", func(c *fiber.Ctx) error {
		var req GameRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON"})
		}
		if req.PlayerID == "" || req.GameID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "playerID or GameID are required"})

		}

		result := "loss"

		if rand.Intn(2) == 0 {
			result = "win"

		}

		resp := GameResponse{
			PlayerID: req.PlayerID,
			GameID:   req.GameID,
			Results:  result,
		}
		return c.JSON(resp)

	})
	log.Println("listening on port:5000")
	log.Fatal(app.Listen(":5000"))

}
