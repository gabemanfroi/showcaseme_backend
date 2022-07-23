# Models

---

<p>Expected to contain all the application domain models.</p>
<p>e.g.</p>

```go
type User struct {
	gorm.Model
	Age       int
	City      string
	Country   string
	Email     string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Pronouns  string
	Password  string `gorm:"not null"`
}

```