# Controllers

---

<p>Expected to contain the implementations for the Controllers Interfaces located in <b>domain/interfaces/controllers</b></p>
<p>e.g</p>

```go

type UserController struct {
	service services.IUserService
}

func CreateUserController() *UserController { return &UserController{service: getUserService()} }

func (controller UserController) Create(c *fiber.Ctx) error {
	var dto user.CreateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	return c.Status(200).JSON(controller.service.Create(&dto))
}

func (controller UserController) GetAll(c *fiber.Ctx) error {
	users, err := controller.service.GetAll()
	if err != nil {
		return err
	}
	utils.Check(json.NewEncoder(c).Encode(&users), "failed to encode users")
	return c.Status(200).JSON(users)
}

func (controller UserController) GetById(c *fiber.Ctx) error {
	u, err := controller.service.GetById(c.Params("id"))
	if err != nil {
		return err
	}
	utils.Check(json.NewEncoder(c).Encode(&u), "failed to encode user")
	return c.Status(200).JSON(u)
}

func (controller UserController) Delete(c *fiber.Ctx) error {
	err := controller.service.Delete(c.Params("id"))
	if err != nil {
		return err
	}

	return c.Status(200).JSON("user deleted")
}

func (controller UserController) Update(c *fiber.Ctx) error {
	var dto user.UpdateUserDTO

	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	updatedUser, err := controller.service.Update(c.Params("id"), &dto)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(updatedUser)
}

func getUserService() services.IUserService {
	var injector services.IUserService
	utils.Check(container.Resolve(&injector), "Error while retrieving UserService instance ")
	return injector
}


```