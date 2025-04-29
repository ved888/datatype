# SOLID Principles Explained in Very Easy Words

---

## 1. Single Responsibility Principle (SRP)

### 📌 Rule: One thing = one job only.

**✅ Do:** Write a function or struct that has only *one reason to change*.

**❌ Don't:** Mix different tasks in one place.

### ❌ Don’t do this:
```go
type User struct {
	Name string
	Email string
}

func (u *User) Register() {
	// Save user to database
	// Send welcome email
}
```

### ✅ Do this instead:
```go
type User struct {
	Name string
	Email string
}

func (u *User) Save() {
	// Save to database
}

func SendWelcomeEmail(u User) {
	// Send email
}
```

---

## 2. Open/Closed Principle (OCP)

### 📌 Rule: **Open for adding**, **closed for changing.**

**✅ Do:** Add new features using interfaces or functions.

**❌ Don't:** Keep editing old code again and again.

### ❌ Don’t do this:
```go
func GetDiscount(typ string, price float64) float64 {
	if typ == "new_year" {
		return price * 0.9
	} else if typ == "black_friday" {
		return price * 0.8
	}
	return price
}
```

### ✅ Do this:
```go
type Discount interface {
	Apply(price float64) float64
}

type NewYear struct{}
func (d NewYear) Apply(price float64) float64 { return price * 0.9 }

type BlackFriday struct{}
func (d BlackFriday) Apply(price float64) float64 { return price * 0.8 }

func GetFinalPrice(d Discount, price float64) float64 {
	return d.Apply(price)
}
```

---

## 3. Liskov Substitution Principle (LSP)

### 📌 Rule: **Substitute child for parent safely.**

**✅ Do:** Make sure new types behave like the original.

**❌ Don't:** Break things when swapping types.

### ❌ Don’t do this:
```go
type Bird interface {
	Fly()
}

type Ostrich struct{} // Can't fly
func (o Ostrich) Fly() {
	panic("I can't fly!")
}
```

### ✅ Do this:
```go
type Bird interface {
	Walk()
}

type Ostrich struct{}
func (o Ostrich) Walk() {
	fmt.Println("Ostrich walking")
}
```

---

## 4. Interface Segregation Principle (ISP)

### 📌 Rule: **Small interfaces** are better.

**✅ Do:** Make interfaces with **only the methods needed**.

**❌ Don't:** Force someone to implement unused stuff.

### ❌ Don’t do this:
```go
type Machine interface {
	Print()
	Scan()
	Fax()
}

type OldPrinter struct{}
func (o OldPrinter) Print() {}
func (o OldPrinter) Scan() { panic("Not supported") }
func (o OldPrinter) Fax() { panic("Not supported") }
```

### ✅ Do this:
```go
type Printer interface { Print() }
type Scanner interface { Scan() }

type OldPrinter struct{}
func (o OldPrinter) Print() {}
```

---

## 5. Dependency Inversion Principle (DIP)

### 📌 Rule: **Depend on interfaces**, not real objects.

**✅ Do:** Code against interfaces (abstractions).

**❌ Don't:** Tie your code to specific things.

### ❌ Don’t do this:
```go
type Database struct{}
func (d Database) Save(data string) {}

type App struct {
	db Database
}
```

### ✅ Do this:
```go
type Storage interface {
	Save(data string)
}

type Database struct{}
func (d Database) Save(data string) {}

type App struct {
	store Storage
}
```

---

# 🧠 Super Simple Summary

| Principle | DO ✅ | DON'T ❌ |
|-----------|---------|-----------|
| SRP | One thing per function | Mix multiple tasks |
| OCP | Add new logic safely | Keep editing old code |
| LSP | Subtypes behave the same | Break when swapping types |
| ISP | Small interfaces | Big forced interfaces |
| DIP | Use interfaces | Lock to concrete types |

---

