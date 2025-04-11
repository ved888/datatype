# 1. What is a Channel in Go?

A **channel** in Go is a powerful feature used for **communication between goroutines**. It allows goroutines to **safely share data** without using explicit locks or mutexes.

## 📌 Key Points

- Channels are typed conduits — they allow sending and receiving values of a specific type.
- Channels ensure **synchronization** between sender and receiver.
- They help implement **concurrent and parallel** logic in a clean and readable way.

## 🔧 Syntax


ch := make(chan int) // creating an unbuffered channel of type int


# 4. What is the Difference Between a Buffered and an Unbuffered Channel?

In Go, channels are used for communication between goroutines. Channels can be either **buffered** or **unbuffered**, and the difference lies in how they handle synchronization and data flow.

---

## 🔄 Unbuffered Channels

An **unbuffered channel** has zero capacity. It requires both the sender and receiver to be ready at the same time. The send operation will block until another goroutine is ready to receive the data.

### ➕ Characteristics:
- **Synchronous communication**
- Sender waits until receiver is ready
- Ensures **tight synchronization** between goroutines

### 🔧 Example:
```go
ch := make(chan int) // unbuffered

go func() {
    ch <- 10 // blocks until receiver is ready
}()

val := <-ch // receives the value
fmt.Println(val)
```

---

## 🧺 Buffered Channels

A **buffered channel** has a defined capacity. The sender can send data to the channel without waiting, as long as the buffer is not full. The receiver can retrieve data when it is ready.

### ➕ Characteristics:
- **Asynchronous communication**
- Sender only blocks when buffer is full
- Receiver only blocks when buffer is empty
- Helps **decouple** goroutines

### 🔧 Example:
```go
ch := make(chan int, 2) // buffered with capacity 2

ch <- 1 // doesn't block
ch <- 2 // doesn't block
// ch <- 3 // would block if uncommented (buffer full)

fmt.Println(<-ch) // prints 1
fmt.Println(<-ch) // prints 2
```

---

## 🔍 Conceptual Difference

| Aspect                     | Unbuffered Channel                          | Buffered Channel                          |
|---------------------------|---------------------------------------------|-------------------------------------------|
| Capacity                  | 0 (no buffer)                               | >0 (fixed buffer size)                    |
| Communication Style       | Synchronous                                 | Asynchronous                              |
| Send Blocks When          | No receiver is ready                        | Buffer is full                            |
| Receive Blocks When       | No value is available                       | Buffer is empty                           |
| Synchronization Use Case  | Tight goroutine coordination                | Looser communication, task queuing        |
| Example Use               | Task signaling, step-by-step execution      | Work queues, producer-consumer patterns   |

---

## 🧠 Summary
- Use **unbuffered channels** when you want goroutines to strictly synchronize.
- Use **buffered channels** when you want to decouple send and receive operations, or implement task queues.

' 
# 🚫 Channel Deadlocks in Go – Key Conditions

Deadlocks in Go related to **channels** usually occur due to improper send/receive synchronization between goroutines.

---

## 🔄 Common Channel Deadlock Scenarios

### 1. Sending on an Unbuffered Channel Without a Receiver

```go
func main() {
    ch := make(chan int)
    ch <- 10 // ❌ Blocks forever (no goroutine receiving)
}
```

---

### 2. Receiving from an Unbuffered Channel Without a Sender

```go
func main() {
    ch := make(chan int)
    val := <-ch // ❌ Blocks forever (no goroutine sending)
}
```

---

### 3. Sending on a Buffered Channel That Is Full

```go
func main() {
    ch := make(chan int, 1)
    ch <- 1
    ch <- 2 // ❌ Deadlock: buffer is full, no receiver
}
```

---

### 4. Receiving from a Buffered Channel That Is Empty

```go
func main() {
    ch := make(chan int, 1)
    <-ch // ❌ Deadlock: nothing in buffer, no sender
}
```

---

### 5. Using `range` on a Channel That Is Never Closed

```go
func main() {
    ch := make(chan int)
    go func() {
        ch <- 1
        ch <- 2
    }()

    for val := range ch { // ❌ Deadlock after last value (channel not closed)
        fmt.Println(val)
    }
}
```

---

### 6. `select` with No Ready Case and No `default`

```go
func main() {
    ch := make(chan int)
    select {
    case <-ch: // ❌ Deadlock: no sender
    }
}
```

---

## ✅ Best Practices to Avoid Channel Deadlocks

- Always pair `send` and `receive` operations correctly.
- Use goroutines for async operations.
- Close channels when done (especially with `range`).
- Use buffered channels carefully (watch for full/empty states).
- Use `select` with `default` if there's no guarantee of readiness.


12. What happens if you receive data from a closed channel?

Receiving from a closed channel in Go is completely safe and does not cause a panic.

✅ Behavior:

You will receive the zero value of the channel's type.

The comma-ok idiom can be used to check if the channel is closed.

📌 Example:

`ch := make(chan int)
close(ch)

val, ok := <-ch
fmt.Println(val, ok) // Output: 0 false`

🔍 Explanation:

val receives the zero value of type int, which is 0.

ok will be false, indicating the channel is closed.

🧠 Notes:

This mechanism allows graceful shutdown of goroutines listening on channels.

Always use the ok flag if you're not sure whether a channel is closed.

❗ Contrast With Sending:

Unlike receiving, sending to a closed channel causes a panic.

``ch := make(chan int)
close(ch)
ch <- 10 // panic: send on closed channel``

### ❓ 21. What is the default buffer size of an unbuffered channel?
** ✅ The default buffer size of an unbuffered channel is 0.**

### ❓ 38. How do you drain a buffered channel before closing it?

✅ **Draining a buffered channel** means receiving and processing all values in the channel **before closing or exiting**, to avoid data loss or deadlocks.

---

### ✅ Correct Way to Drain a Buffered Channel

You typically **don’t drain before closing**, but rather:

- **Close the channel first**, then
- Use a `for-range` loop to drain it.

```go
func drainBufferedChannel(ch chan int) {
    close(ch) // Close the channel to stop sending
    for val := range ch {
        fmt.Println("Drained:", val)
    }
}
```

This ensures all data in the buffered channel is read and handled properly after it has been closed.

```go
func main() {
    ch := make(chan int, 5)

    // Fill the buffered channel
    for i := 0; i < 5; i++ {
        ch <- i
    }

    drainBufferedChannel(ch)
}
```