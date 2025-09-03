## Golang Interview Questions and Answers

### 1. What are the main differences between goroutines and threads? How does Go handle concurrency under the hood?

Goroutines are lightweight, user-space threads managed by the Go runtime. Unlike OS threads, they are cheaper in terms of memory (a few KB per goroutine) and scheduling. Go runtime uses a scheduler to multiplex thousands of goroutines onto a smaller number of OS threads. This concurrency model is efficient and simplifies writing scalable applications.

### 2. Explain how the Go scheduler works and how it relates to goroutines and system threads.

Go's scheduler uses a model called GPM:

* **G**: Goroutine (task to run)
* **P**: Processor (manages execution context)
* **M**: Machine (OS thread)
  It schedules goroutines (G) on logical processors (P), which run on OS threads (M). This model enables efficient concurrency with minimal thread switching overhead.

### 3. What is a race condition? How would you detect and prevent it in Go?

A race condition occurs when multiple goroutines access shared memory concurrently and at least one is writing. To detect it, use `go run -race`. To prevent it, use synchronization tools like `sync.Mutex`, `sync.RWMutex`, or channels to control access to shared data.

### 4. What are channels in Go? How do buffered and unbuffered channels differ?

Channels allow goroutines to communicate safely. An unbuffered channel blocks on send until receive is ready. A buffered channel allows a fixed number of values to be stored without blocking the sender.

### 5. Can you explain the purpose of the `select` statement in Go? Provide a practical use case.

`select` is used to wait on multiple channel operations. It blocks until one of the channels is ready. Use case: implementing timeout with `time.After()`.

### 6. How would you implement a worker pool pattern in Go? Why is it useful?

Create a fixed number of worker goroutines that read tasks from a channel and write results to another. It's useful to limit concurrency and reuse goroutines efficiently.

### 7. How does Go handle memory management and garbage collection?

Go has a concurrent garbage collector which automatically reclaims memory no longer in use. It's optimized for low latency and runs concurrently with the application.

### 8. Explain the `context` package in Go. How do you use it for timeouts and cancellations?

`context` is used to carry deadlines, cancellation signals, and other request-scoped values. Use `context.WithTimeout` or `context.WithCancel` to automatically stop operations after a time or on signal.

### 9. What are best practices to write concurrent, safe Go code in microservices?

* Avoid shared memory, prefer message passing.
* Use context for request lifecycle management.
* Use channels or sync primitives for synchronization.
* Test with `-race` flag.

### 10. Describe the difference between `sync.Mutex`, `sync.RWMutex`, and `sync.WaitGroup` with use cases.

* `Mutex`: mutual exclusion for one writer.
* `RWMutex`: allows multiple readers or one writer.
* `WaitGroup`: waits for a collection of goroutines to finish.

### 11. Which Go web frameworks or routers have you used (e.g., Gin, Echo)? Why did you choose them?

Gin is fast, minimal, and has middleware support. Echo is more feature-rich and suitable for larger apps. Choice depends on project complexity.

### 12. How do you structure a REST API project in Go for scalability and maintainability?

* Separate layers: handler, service, repository.
* Use dependency injection.
* Organize by feature or module.
* Use interfaces for testing.

### 13. How would you handle input validation and error handling in a Go REST API?

Use struct tags with libraries like `go-playground/validator`. For errors, return consistent structured JSON responses and appropriate HTTP status codes.

### 14. Describe how to implement middleware in a Go-based web service.

Middleware is a function that wraps a handler to perform tasks like logging, authentication, or rate limiting. Most routers like Gin and Echo support chaining middleware functions.

### 15. How do you manage versioning in REST APIs using Go?

Use URI-based versioning like `/v1/resource`. Separate handler logic by version. Maintain backward compatibility for existing versions.

### 16. Which Go libraries do you use for SQL database interaction? (e.g., `database/sql`, `gorm`, `sqlx`)

`database/sql` for low-level control. `gorm` for ORM-style mapping. `sqlx` for flexible mapping and better developer ergonomics.

### 17. How would you handle database transactions in Go?

Use `db.Begin()` to start a transaction, `tx.Commit()` to commit, and `tx.Rollback()` on error. Ensure `defer tx.Rollback()` is used to handle panics.

### 18. How do you prevent SQL injection attacks in Go?

Always use parameterized queries (`?` or named parameters) instead of string concatenation for queries.

### 19. What’s your approach to handling database migrations in a Go project?

Use tools like `golang-migrate/migrate` or `goose` to manage versioned migrations. Maintain migration history in a schema table.

### 20. How would you model relationships (e.g., one-to-many, many-to-many) using GORM in Go?

Use `has many`, `belongs to`, and `many2many` GORM tags on struct fields. GORM handles foreign keys and join tables.

### 21. How do you use AWS Lambda with Go? What are the trade-offs compared to a standard Go microservice?

Build your Go binary, package it, and deploy to Lambda using AWS CLI or SAM. Trade-offs: good for event-driven workloads, but limited in runtime control and cold start latency.

### 22. How do you structure and deploy a Go application in a Docker container?

Use a multi-stage Dockerfile. Compile binary in one stage, copy to a minimal base image (e.g., `scratch` or `alpine`). Expose ports and define entrypoint.

### 23. Explain a scenario where you containerized a Go application and deployed it on Kubernetes.

Build and push Docker image to registry. Define Kubernetes deployment and service YAML. Use `kubectl` or Helm to deploy.

### 24. What is your approach for logging and monitoring in a Go-based cloud-native service?

Use structured logging (e.g., logrus or zap). Emit logs in JSON. Integrate with monitoring tools like Prometheus and centralized log systems (e.g., ELK, Loki).

### 25. How do you optimize Go services for cold start performance in AWS Lambda or GCP Cloud Functions?

* Reduce binary size and dependencies
* Avoid global initializations
* Pre-warm with scheduled invocations
* Keep handlers lightweight
