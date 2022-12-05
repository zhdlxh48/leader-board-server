package routes

import (
	"log"
	"strconv"

	"github.com/zhdlxh48/leader-board-server/database"
	"github.com/zhdlxh48/leader-board-server/model"

	"github.com/gofiber/fiber/v2"
)

func SaveScore(ctx *fiber.Ctx) error {
	score := new(model.Score)
	err := ctx.BodyParser(score)
	if err != nil {
		log.Fatalln(err)
		return ctx.SendStatus(500)
	}

	if score.GameTitle == "" || score.GameType == "" {
		ctx.SendStatus(400)
		return ctx.JSON(fiber.Map{
			"message": "please insert title and type on request",
		})
	}

	if score.UserID == "" {
		score.UserID = "anonymous"
	}
	if score.UserName == "" {
		score.UserName = "익명"
	}

	insert, err := database.InsertBoardData(score)
	if err != nil {
		log.Fatalln(err)
		return ctx.SendStatus(500)
	}

	return ctx.JSON(fiber.Map{
		"score": insert,
	})
}

func GetScore(ctx *fiber.Ctx) error {
	gameTitle := ctx.Query("title")
	gameType := ctx.Query("type")
	count, err := strconv.Atoi(ctx.Query("count"))
	if err != nil {
		count = 10
	}

	if gameTitle == "" || gameType == "" {
		ctx.SendStatus(400)
		return ctx.JSON(fiber.Map{
			"message": "please insert title and type on request",
		})
	}

	scores, err := database.SelectBoardData(gameTitle, gameType, count)
	if err != nil {
		log.Fatalln(err)
		return ctx.SendStatus(500)
	}

	return ctx.JSON(fiber.Map{
		"scores": *scores,
	})
}
