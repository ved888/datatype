
# 🚫 Deadlocks in Go – Full Guide

Deadlocks occur when goroutines wait on each other indefinitely, causing the program to stop making progress.

---

## 🔒 All Possible Deadlock Conditions in Go

### 1. Sending on an Unbuffered Channel With No Receiver

```go
func main() {
    ch := make(chan int)
    ch <- 10 // blocks forever (no goroutine is receiving)
}
```

---

### 2. Receiving from an Unbuffered Channel With No Sender

```go
func main() {
    ch := make(chan int)
    val := <-ch // blocks forever (no goroutine is sending)
    fmt.Println(val)
}
```

---

### 3. Buffered Channel: Send When Full or Receive When Empty

```go
func main() {
    ch := make(chan int, 1)
    ch <- 1
    ch <- 2 // deadlock: buffer full, no receiver
}
```

```go
func main() {
    ch := make(chan int, 1)
    fmt.Println(<-ch) // deadlock: buffer empty, no sender
}
```

---

### 4. Goroutines Waiting on Each Other

```go
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        ch1 <- 1
        <-ch2
    }()

    ch2 <- 2
    <-ch1
}
```

---

### 5. Range on a Channel That’s Never Closed

```go
func main() {
    ch := make(chan int)

    go func() {
        ch <- 1
        ch <- 2
        // channel not closed
    }()

    for val := range ch {
        fmt.Println(val) // blocks forever after last value
    }
}
```

✅ Fix: `close(ch)` after sending.

---

### 6. Select Without Default and No Ready Case

```go
func main() {
    ch := make(chan int)

    select {
    case <-ch:
        fmt.Println("Received")
    }
}
```

---

### 7. WaitGroup Never Done

```go
func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    // wg.Done() never called
    wg.Wait() // waits forever
}
```

---

### 8. Missing Goroutine for Channel Operation

```go
func main() {
    ch := make(chan int)
    ch <- 42 // blocks; no goroutine to handle it
}
```

---

### 9. Lock Without Unlock

```go
var mu sync.Mutex

func main() {
    mu.Lock()
    mu.Lock() // deadlock: trying to lock again without unlocking
}
```

---

### 10. Circular Locking Between Goroutines

```go
var mu1, mu2 sync.Mutex

func main() {
    go func() {
        mu1.Lock()
        defer mu1.Unlock()
        time.Sleep(1 * time.Second)
        mu2.Lock()
        defer mu2.Unlock()
    }()

    mu2.Lock()
    time.Sleep(1 * time.Second)
    mu1.Lock()
    mu2.Unlock()
    mu1.Unlock()
}
```

---

## 🧠 Summary Table

| Cause                             | Description |
|----------------------------------|-------------|
| Unbuffered send w/o receiver     | Blocks forever |
| Unbuffered receive w/o sender    | Blocks forever |
| Buffered channel full/empty misuse | No goroutine to read/write |
| Cyclic wait between goroutines   | Mutual blocking |
| Range on channel not closed      | Blocks after all values |
| Select with no valid case        | Blocks forever |
| WaitGroup never `.Done()`        | `.Wait()` blocks forever |
| No goroutine for channel op      | Stuck on send/receive |
| Mutex not unlocked               | Lock blocks forever |
| Circular locking (mutexes)       | Classic deadlock |

---

**Avoiding Deadlocks:** Always ensure there is a matching sender/receiver, all goroutines are started correctly, and channels are properly closed when used with `range`.