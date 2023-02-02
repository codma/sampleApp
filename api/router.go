package api

import (
	"main/api/app"
	"main/api/user"
	"main/constant"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func InitRouter() *fiber.App {
	readTimeout := viper.GetDuration(constant.SERVER_READ_TIMEOUT) * time.Second
	writeTimeout := viper.GetDuration(constant.SERVER_WRITE_TIMEOUT) * time.Second
	appName := constant.APP_NAME

	api := fiber.New(fiber.Config{
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		AppName:      appName,
	})
	api.Use(logger.New())
	api.Use(recover.New())

	apiGroup := api.Group("")
	{
		app.ApplyRoutes(apiGroup)
		user.ApplyRoutes(apiGroup)
	}

	return api

}
