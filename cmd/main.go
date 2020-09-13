package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v2"
	"os"
	"timewise/db"
	"timewise/handler"
	"timewise/model"
	"timewise/repository"
	"timewise/router"
)

func main(){
	var cfg model.Config
	loadConfig(&cfg)

	var sql = new (db.SQL)
	sql.Connect(cfg)
	defer  sql.Close()

	e :=echo.New()
	//e.Use(middleware.CORS())
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"localhost:8080"},
	//	AllowHeaders: []string{echo.HeaderOrigin,echo.HeaderContentType, echo.HeaderAccept},
	//}))

	//setup validator
	//structValidator := validator.NewStrcutValidator()
	//structValidator.ReGisterValidate()
	//e.Validator = structValidator

	//setup router + handler
	userHandler := handler.UserHandler{
		UserRepo: repository.NewUserRepo(sql),
	}

	api:= router.API{
		Echo: e,
		UserHandler: userHandler,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}
func loadConfig(cfg *model.Config)  {
	f, err := os.Open("../env.dev.yml")
	if(err !=nil){
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
}