# Controllers Interfaces

---

<p>Expected to contain the Controllers Interfaces. The Implementations are generally located under <b>application/controllers</b</p>
<p>e.g.</p>

```go
type IUserController interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}
```