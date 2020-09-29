package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
	"timewise/db"
	"timewise/handler"
	log "timewise/log"
	"timewise/model"
	"timewise/repository"
	"timewise/router"
)

func init(){
	os.Setenv("APP_NAME", "timewise")
	log.InitLogger(false)
}

func main() {

	var cfg model.Config
	loadConfig(&cfg)
	setEnv(&cfg)

	var sql = new(db.SQL)
	sql.Connect(cfg)
	defer sql.Close()

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	//setup validator
	//structValidator := validator.NewStrcutValidator()
	//structValidator.ReGisterValidate()
	//e.Validator = structValidator

	//setup router + handler
	userHandler := handler.UserHandler{
		UserRepo: repository.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupAdminRouter()
	// Tạo một Firewall
	// chỉ cho phép ip của máy tính admin truy cập vào port 8000
	e.Logger.Fatal(e.Start(fmt.Sprintf(":8000")))
	//e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}

/**
set env variable
*/
func setEnv(cfg *model.Config) {
	jwtExpires := strconv.Itoa(cfg.Server.JwtExpires)
	os.Setenv("JwtExpires", jwtExpires)
}

/**
Load config
*/
func loadConfig(cfg *model.Config) {
	f, err := os.Open("../../env.dev.yml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
}
