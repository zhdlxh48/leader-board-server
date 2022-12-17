package database

import (
	"fmt"
	"log"
	"os"

	"github.com/zhdlxh48/leader-board-server/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func getDsn() string {
	dbId := os.Getenv("DB_ID")
	dbPwd := os.Getenv("DB_PWD")

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		dbUrl = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "test"
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbId, dbPwd, dbUrl, dbPort, dbName)
}

func Initialize() {
	conn, err := gorm.Open(mysql.Open(getDsn()), &gorm.Config{})
	if err != nil {
		log.Panic(err.Error())
	}
	conn.AutoMigrate(&model.GormScore{})

	DBConn = conn
}

func InsertBoardData(score *model.Score) (*model.GormScore, error) {
	insert := &model.GormScore{Score: *score}
	tx := DBConn.Create(&insert)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	return insert, nil
}

func SelectBoardData(gameTitle string, gameType string, count int) (*[]model.GormScore, error) {
	scores := new([]model.GormScore)
	tx := DBConn.Model(&model.GormScore{}).Where("game_title = ?", gameTitle).Where("game_type = ?", gameType).Order("user_score desc").Order("created_at").Limit(count).Find(scores)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	return scores, nil
}

func SelectRanks(gameTitle string, gameType string, count int) (*[]model.Rank, error) {
	ranks := new([]model.Rank)
	tx := DBConn.Raw("SELECT user_name, user_score, RANK() OVER (ORDER BY user_score DESC) as ranking FROM test_leader_board WHERE game_title = ? AND game_type = ? LIMIT ?", gameTitle, gameType, count).Scan(ranks)
	log.Println(ranks)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	return ranks, nil
}

func SelectUserRank(gameTitle string, gameType string, userName string) (*model.Rank, error) {
	rank := new(model.Rank)
	tx := DBConn.Raw("SELECT user_name, user_score, ranking FROM (SELECT user_name, user_score, RANK() OVER (ORDER BY user_score DESC) as ranking FROM test_leader_board WHERE game_title = ? AND game_type = ?) as CNT WHERE user_name = ? LIMIT 1", gameTitle, gameType, userName).Scan(rank)
	log.Println(rank)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	return rank, nil
}

func DatabaseVersion() string {
	ver := new(string)
	DBConn.Raw("SELECT VERSION()").Scan(ver)
	return *ver
}
