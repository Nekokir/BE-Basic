package main

import (
	"net/http"
	"github.com/BurntSushi/toml"
	"strconv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/julienschmidt/httprouter"
	"crypto/md5"
	"encoding/hex"
)

var mux = httprouter.New()
var db *gorm.DB

func main() {

	_, err := toml.DecodeFile("config.toml", &GlobCfg)
	if err != nil {
		panic(err)
	}

	db, err = gorm.Open("sqlite3", "bt1.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Answer{}, &User{})

	mux.POST("/quiz", SignUp)
	mux.PUT("/quiz", EditName)
	mux.DELETE("/quiz", SubmitQuiz)
	mux.POST("/question/:qid", AnswerQuestion)
	mux.PUT("/question/:qid", EditAnswer)
	mux.GET("/question/:qid", ShowAnswer)
	mux.GET("/", ShowNames)
	mux.POST("/", ShowDetails)

	http.ListenAndServe(":"+strconv.FormatInt(GlobCfg.PORT, 10), mux)
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}