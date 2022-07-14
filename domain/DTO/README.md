# DTO

---

<p>Expected to contain the application Data Transfer Objects.</p>
<p>e.g.</p>

```go
type CreateUserDTO struct {
	Age       int    `json:"age,omitempty"`
	City      string `json:"city,omitempty" `
	Country   string `json:"country,omitempty" `
	Email     string `json:"email"  `
	FirstName string `json:"firstName" `
	LastName  string `json:"lastName"  `
	Pronouns  string `json:"pronouns,omitempty"`
}
```