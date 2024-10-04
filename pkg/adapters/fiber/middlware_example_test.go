package fiber

import "github.com/gofiber/fiber/v2"

func Example() {
	app := fiber.New()
	app.Use(
		SentinelMiddleware(
			// customize resource extractor if required
			// method_path by default
			WithResourceExtractor(func(ctx *fiber.Ctx) string {
				headers := ctx.GetReqHeaders()
				ipValues, exists := headers["X-Real-IP"]
				var ip string
				if exists && len(ipValues) > 0 {
					ip = ipValues[0] // 获取第一个IP值
				} else {
					// 如果不存在或没有值，则可以设置默认值或者处理逻辑
					ip = "default-ip" // 示例中的默认值
				}
				return ip
			}),
			// customize block fallback if required
			// abort with status 429 by default
			WithBlockFallback(func(ctx *fiber.Ctx) error {
				return ctx.Status(400).JSON(struct {
					Error string `json:"error"`
					Code  int    `json:"code"`
				}{
					"too many request; the quota used up",
					10222,
				})
			})),
	)

	app.Get("/test", func(ctx *fiber.Ctx) error { return nil })
	_ = app.Listen(":8080")
}
