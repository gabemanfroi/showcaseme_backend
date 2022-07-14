# Services Interfaces

---

<p>Expected to contain the Services Interfaces. The Implementations are generally located under <b>domain/services</b</p>
<p>e.g.</p>

```go
type IUserService interface {
    Create(dto *user.CreateUserDTO) models.User
    GetAll() ([]models.User, error)
    GetById(id uint) (models.User, error)
    Delete(id uint) error
    Update(id uint, dto *user.UpdateUserDTO) (models.User, error)
}
```