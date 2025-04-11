## Golang Interview Questions and Answers for 2 Years Experience

---

### 📘 Basics & Syntax

**1. What are the basic data types in Go?**
- `int`, `float64`, `string`, `bool`, `byte`, `rune`, `complex64`, `complex128`.
- These types cover numbers, characters, booleans, and complex numbers. `byte` is alias for `uint8`, and `rune` is alias for `int32`.

**2. What is the difference between `var x int` and `x := 10`?**
- `var x int` declares a variable named `x` of type `int`, but doesn't assign a value.
- `x := 10` is a short declaration that both declares and initializes `x` to 10 with type inferred.

**3. What is the zero value of a variable in Go?**
- A variable declared without explicit initialization gets a default "zero" value.
  - `0` for `int`, `0.0` for `float`, `""` for `string`, `false` for `bool`, `nil` for pointers, slices, maps, channels, interfaces.

**4. What are short variable declarations? Where can you use them?**
- Syntax: `x := 5`. This shorthand can only be used inside functions.

**5. What is a pointer in Go? How do you use it?**
- A pointer stores the memory address of a variable.
```go
var x = 10
var p *int = &x
fmt.Println(*p) // prints 10
```
- `&x` gives the address, `*p` dereferences it.

**6. What are slices and how are they different from arrays?**
- Arrays are fixed-size and defined as `[n]type`.
- Slices are dynamic and reference an underlying array: `[]type`.
- Slices can grow/shrink and are used more frequently.

**7. What is the difference between `make()` and `new()`?**
- `make()` creates slices, maps, and channels, initializing internal data structures.
- `new()` allocates memory for a type and returns a pointer to it.

---

### 🧵 Concurrency

**8. What is a goroutine?**
- A goroutine is a lightweight thread of execution managed by the Go runtime. Use `go` keyword.
```go
go fmt.Println("Running in a goroutine")
```

**9. What is a channel and how does it work?**
- Channels are used to communicate between goroutines. Data can be sent with `ch <- val` and received with `val := <-ch`.

**10. What is the difference between buffered and unbuffered channels?**
- Unbuffered: sends block until receiver is ready.
- Buffered: can store limited elements without blocking until the buffer is full.

**11. What is a deadlock and how does it occur in Go?**
- A deadlock happens when goroutines wait on each other forever, like sending to a channel without a receiver.

**12. How does the `select` statement work in Go?**
- `select` waits on multiple channel operations. It picks one case that’s ready or blocks if none are ready.

**13. What is the use of `sync.WaitGroup`?**
- It waits for a collection of goroutines to finish.
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
  defer wg.Done()
  // work
}()
wg.Wait()
```

**14. What is a race condition and how can you detect it in Go?**
- When two goroutines access a shared resource concurrently and at least one write occurs.
- Detect using: `go run -race your_file.go`

---

### 🏗️ Structs & Interfaces

**15. What are structs in Go?**
- A collection of fields. Custom data types for grouping data.
```go
type Person struct {
  Name string
  Age  int
}
```

**16. How do you implement an interface in Go?**
- Any type that implements all the methods of an interface implicitly satisfies it.

**17. What is the difference between an interface and a struct?**
- Structs hold data; interfaces define behavior via method sets.

**18. What is type assertion in Go?**
- Extracts the underlying value of an interface.
```go
var i interface{} = "hello"
s := i.(string)
```

**19. What is type embedding in Go?**
- Composition of types: a struct can include another struct.
```go
type Address struct { City string }
type Employee struct {
  Name string
  Address // embedded
}
```

---

### 🔧 Error Handling & Packages

**20. How do you handle errors in Go?**
```go
val, err := someFunc()
if err != nil {
    log.Fatal(err)
}
```

**21. What is the difference between `panic`, `recover`, and `defer`?**
- `panic`: aborts the program.
- `recover`: regains control after a panic inside a deferred function.
- `defer`: schedules a function call to run after the current function completes.

**22. How do you create and use packages in Go?**
- Group files with same package name. Import with `import "path/to/package"`

**23. What is the purpose of the `init()` function?**
- Automatically runs before `main()` to initialize package variables or setup.

---

### 🌐 Web & APIs

**24. How do you build a simple HTTP server in Go?**
```go
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
```

**25. What packages do you use to make HTTP requests in Go?**
- `net/http`
- Use `http.NewRequest()`, `http.Get()`, and `http.Client.Do()`

**26. How do you handle JSON in Go?**
- Use `encoding/json`.
```go
json.Marshal() // encode struct to JSON
json.Unmarshal() // decode JSON to struct
```

**27. How do you manage routes in a Go web application?**
- Use `http.HandleFunc` or frameworks like `mux` or `chi` for better routing control.

---

### 🧪 Testing

**28. How do you write unit tests in Go?**
- Create a file ending in `_test.go`, use `testing` package.
```go
func TestAdd(t *testing.T) {
  got := Add(2, 3)
  want := 5
  if got != want {
    t.Errorf("got %d, want %d", got, want)
  }
}
```

**29. What is the use of the `testing` package?**
- To write, run, and organize unit tests and benchmarks.

**30. How do you mock a function or API call in Go?**
- Define interfaces for the external service and inject fake/mock implementations in tests.

---

Let me know if you'd like 30 more intermediate/advanced questions or want to export this file!

