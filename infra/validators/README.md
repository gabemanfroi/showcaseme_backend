# Validators

---

<p>Expected to have the application validation schemas (used to validate request bodies).</p>
<p>e.g.</p>

```go
type CreateUserValidator struct {
	FirstName string `validate:"required,min=3"`
	LastName  string `validate:"required,min=3"`
	Email     string `validate:"required,email"`
	Age       uint8  `validate:"omitempty,gte=0,lte=130"`
	City      string `validate:"omitempty,min=3"`
	Country   string `validate:"omitempty,min=3"`
	Pronouns  string `validate:"omitempty,min=3"`
	Password  string `validate:"required"`
}
```
