# Repositories Interfaces

---

<p>Expected to contain the Repositories Interfaces. The Implementations are generally located under <b>infra/db/repositories</b</p>
<p>e.g.</p>

```go
type IUserRepository interface {
    Create(dto *user.CreateUserDTO) models.User
    GetAll() ([]models.User, error)
    GetById(id uint) (models.User, error)
    Delete(id uint) error
    Update(id uint, dto *user.UpdateUserDTO) (models.User, error)
}
```