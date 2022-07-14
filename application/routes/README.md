# Routes

---

<p>Expected to contain the application routes</p>
<p>e.g.</p>

```go
func RegisterUserRoutes(router fiber.Router) {
	var controller controllers.IUserController

	utils.Check(container.Resolve(&controller), "Failed to create userController instance...")

	router.Post("/users", func(c *fiber.Ctx) error { return ValidateSchema(c, user.CreateUserValidator{}) }, controller.Create)
	router.Get("/users", controller.GetAll)
	router.Get("/users/:id", controller.GetById)
	router.Delete("/users/:id", controller.Delete)
	router.Patch("/users/:id", func(c *fiber.Ctx) error { return ValidateSchema(c, user.UpdateUserValidator{}) }, controller.Update)
}
```

*Notice that we use the ValidateSchema function to validate the requests bodies.