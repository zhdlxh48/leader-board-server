package routes

import (
	"log"
	"strconv"

	"github.com/zhdlxh48/leader-board-server/database"
	"github.com/zhdlxh48/leader-board-server/model"

	"github.com/gofiber/fiber/v2"
)

func SaveScore(ctx *fiber.Ctx) error {
	payload := struct {
		GameTitle string `json:"game_title"`
		GameType  string `json:"game_type"`
		UserID    string `json:"user_id"`
		UserName  string `json:"user_name"`
		UserScore int64  `json:"user_score"`
	}{}
	if err := ctx.BodyParser(&payload); err != nil {
		log.Println(payload)
		log.Println("Body Parser:", err)
		return ctx.SendStatus(500)
	}

	log.Println(payload)

	if payload.GameTitle == "" || payload.GameType == "" {
		ctx.SendStatus(400)
		return ctx.JSON(fiber.Map{
			"message": "please insert title and type on request",
		})
	}

	if payload.UserID == "" {
		payload.UserID = "anonymous"
	}
	if payload.UserName == "" {
		payload.UserName = "익명"
	}

	data := &model.Score{GameTitle: payload.GameTitle, GameType: payload.GameType, UserID: payload.UserID, UserName: payload.UserName, UserScore: payload.UserScore}
	insert, err := database.InsertBoardData(data)
	if err != nil {
		log.Println("Database Error:", err)
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
		log.Println("Database Error:", err)
		return ctx.SendStatus(500)
	}

	return ctx.JSON(fiber.Map{
		"scores": *scores,
	})
}

func GetRanks(ctx *fiber.Ctx) error {
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

	ranks, err := database.SelectRanks(gameTitle, gameType, count)
	if err != nil {
		log.Println("Database Error:", err)
		return ctx.SendStatus(500)
	}

	return ctx.JSON(fiber.Map{
		"ranks": *ranks,
	})
}

func GetUserRank(ctx *fiber.Ctx) error {
	gameTitle := ctx.Query("title")
	gameType := ctx.Query("type")
	userName := ctx.Params("user")

	if gameTitle == "" || gameType == "" || userName == "" {
		ctx.SendStatus(400)
		return ctx.JSON(fiber.Map{
			"message": "please insert title, type and user on request",
		})
	}

	rank, err := database.SelectUserRank(gameTitle, gameType, userName)
	if err != nil {
		log.Println("Database Error:", err)
		return ctx.SendStatus(500)
	}

	return ctx.JSON(fiber.Map{
		"rank": *rank,
	})
}
