package app

import "github.com/gofiber/fiber/v2"

func ApplyRoutes(r fiber.Router) {
	apis := r.Group("apps")
	{
		//# 앱 라이프 사이클 이벤트 리스너 (앱 라이프사이클 이벤트 호출 HTTP Method 및 URL Path는 고정)
		apis.Post("/lifecycle/app-installed", AppInstalled())
		apis.Post("/lifecycle/app-deactivated", AppDeactived())
		apis.Post("/lifecycle/app-reactivated", AppReactive())
		apis.Post("/lifecycle/app-uninstalled", AppUninstalled())

		// 	# 앱에서 타 앱에서 발행하는 이벤트를 구독하기 위한 리스너 API (앱 개발자가 URL Path 정의)
	}
}
