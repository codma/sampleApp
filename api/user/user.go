package user

import "github.com/gofiber/fiber/v2"

func ApplyRoutes(r fiber.Router) {
	apis := r.Group("user")
	{
		apis.Post("", CreateUser)
		//apis.Put("", UpdateUser)
		//apis.Delete("", DeleteUser)
	}
}
