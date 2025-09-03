from pathlib import Path

# Prepare the markdown content
content = """# Advanced Golang Interview Questions and Answers

This document contains some of the most challenging questions and answers related to goroutines, wait groups, concurrency, mutexes, channels, select, and interfaces in Go.

---

## 🧵 Goroutine

1. **How does Go’s runtime manage preemption of long-running goroutines?**  
   Go uses cooperative preemption. It checks for safe points like function calls or specific instructions to yield execution, allowing other goroutines to run.

2. **What are the conditions under which goroutine leaks happen and how can you detect them in a production system?**  
   Leaks happen when goroutines wait on blocking calls (e.g., read from a channel) that never complete. Use pprof and runtime.Stack to detect them.

3. **Explain the differences between goroutines and OS threads in terms of stack growth and memory footprint.**  
   Goroutines start with a small stack (~2KB) that grows/shrinks dynamically. OS threads have fixed large stacks (~1MB).

4. **What does it mean that goroutines are “cooperatively scheduled”?**  
   The runtime scheduler relies on goroutines to yield execution voluntarily at safe points.

5. **How does `GOMAXPROCS` affect goroutine scheduling in a multicore environment?**  
   It sets the number of OS threads for parallel execution. More `GOMAXPROCS` means more parallelism.

6. **What tools or methods can you use to trace goroutine creation and lifecycle at runtime?**  
   Use `runtime/pprof`, `trace`, `debug.SetTraceback`, and tools like `go tool trace`.

7. **Can a goroutine block the entire scheduler? How?**  
   Yes, a tight loop without safe points can monopolize the CPU core, starving others.

8. **How do panic and recover behave when a panic occurs inside a goroutine?**  
   Recover must be called inside the same goroutine's defer. Otherwise, the panic crashes the program.

9. **What are the pitfalls of spawning goroutines inside a loop with closures, and how do you fix them?**  
   Closures capture loop variables. Use parameters to pass variables explicitly to the goroutine.

10. **How can you simulate goroutine starvation and how can it be avoided?**  
   Use long-running loops or blocking channels. Avoid by inserting preemption points or breaking long tasks.

---

## ⏳ WaitGroup

1. **What happens if `WaitGroup.Add()` is called after `Wait()` has started waiting?**  
   It’s safe if done before the counter hits zero, but dangerous if done after that.

2. **Can `WaitGroup` be reused? If yes, under what conditions is it safe?**  
   Yes, only after the counter is reset to zero.

3. **What are the consequences of calling `Done()` more times than `Add()`?**  
   Panic: `negative WaitGroup counter`.

4. **How would you design a wrapper over `WaitGroup` to include timeouts?**  
   Use a `select` with a timeout channel alongside `WaitGroup`.

5. **How does `WaitGroup` ensure memory visibility between goroutines?**  
   Internally uses atomic operations that include memory barriers.

6. **Can a race condition occur if multiple goroutines modify a shared `WaitGroup` incorrectly?**  
   Yes. `Add()` and `Done()` must be called safely.

7. **What happens internally when `Wait()` is called on a `WaitGroup` with zero counter?**  
   It returns immediately.

8. **Why is it not safe to copy a `WaitGroup`? What happens if you do?**  
   Copies the internal state, which corrupts wait behavior and causes panic.

9. **How can you ensure goroutines properly call `Done()` even when they encounter a panic?**  
   Use `defer wg.Done()` at the start of the goroutine.

10. **How does `WaitGroup` avoid busy waiting internally?**  
   It uses a semaphore and condition variables internally.

---

## ⚙️ Concurrency

1. **Compare Go’s concurrency model (CSP) with traditional thread-based models.**  
   Go uses goroutines + channels (CSP), unlike threads + locks. It simplifies reasoning and avoids shared state.

2. **How would you detect and debug race conditions in production Go code?**  
   Use `-race` flag, logging, or tools like Delve and pprof.

3. **What are the downsides of spawning too many goroutines concurrently?**  
   High memory, scheduler contention, runtime slowdowns.

4. **How would you build a rate limiter that controls concurrency using Go primitives?**  
   Use buffered channels as semaphores or time.Ticker for token buckets.

5. **What is the impact of goroutine scheduling delays on real-time systems?**  
   Increased latency, jitter, and deadline misses.

6. **How can you control concurrent access to a shared resource with minimal performance overhead?**  
   Use sync.Mutex, atomic ops, or channel-based synchronization.

7. **Design a concurrency-safe cache with TTL using only channels and goroutines.**  
   Maintain a goroutine to periodically clear expired keys, and serialize access via channels.

8. **What does it mean to have *structured concurrency* in Go and how would you implement it?**  
   Ensure all goroutines are tied to context. Use `context.Context` with `select`.

9. **How do you prevent zombie and orphan goroutines in a long-running service?**  
   Use contexts, timeouts, and centralized cancellation logic.

10. **Explain the concept of *context propagation* and its importance in concurrent code.**  
   Context carries cancellation, timeout, and metadata across API boundaries.

---

## 🔐 Mutex

1. **What is the difference between `sync.Mutex` and `sync.RWMutex` in terms of fairness and performance?**  
   `RWMutex` allows multiple readers. Writers block both. RWMutex may starve writers.

2. **How does Go’s `sync.Mutex` internally implement locking and avoid starvation?**  
   Uses a semaphore and spinning. New versions improve fairness with starvation detection.

3. **Under what circumstances can deadlocks still occur when using mutexes?**  
   Acquiring multiple locks in different orders across goroutines.

4. **How would you detect and fix a deadlock in a production system?**  
   Use `pprof` goroutine dumps and ordering analysis.

5. **How does priority inversion manifest with `sync.Mutex` and how can you mitigate it?**  
   Low-priority goroutine holds a lock. Starves high-priority. Mitigation: Avoid long locks.

6. **Can a panic occur while holding a mutex? What precautions must be taken?**  
   Yes. Use `defer mu.Unlock()` to ensure release.

7. **Is it safe to pass a `sync.Mutex` to multiple goroutines? Why or why not?**  
   Yes, as long as access to the mutex itself is not raced.

8. **How do you design a lock-free data structure in Go?**  
   Use atomic operations (`sync/atomic`) and CAS loops.

9. **What are the trade-offs of using a global vs. fine-grained mutex?**  
   Global mutexes are simpler but more contended. Fine-grained improves performance but increases complexity.

10. **Compare `sync.Mutex` with atomic operations in terms of performance and use cases.**  
   Atomics are faster for simple values. Mutexes support complex critical sections.

---

## 📦 Channel

1. **What’s the internal structure of a buffered vs unbuffered channel?**  
   Buffered channels have a ring buffer. Unbuffered block until both sender and receiver are ready.

2. **How does Go handle multiple senders and receivers on the same channel?**  
   FIFO queue internally, managed by the scheduler.

3. **What happens when a channel is closed and a value is sent on it?**  
   Panic: send on closed channel.

4. **What’s the difference between a nil channel and a closed channel?**  
   Nil blocks forever. Closed returns zero value on receive.

5. **How does Go prevent channel deadlocks and how do you identify one in logs?**  
   Compiler/runtime panic: "all goroutines are asleep - deadlock!"

6. **Can you implement a broadcast mechanism using only channels?**  
   Not directly. Requires fan-out with multiple channels and goroutines.

7. **How would you implement a priority queue using channels?**  
   Use separate channels per priority and select order.

8. **What are channel direction types and how do they help in safe concurrency?**  
   `<-chan`, `chan<-` restrict usage and clarify intent.

9. **Can you detect if a channel is full or empty? How?**  
   Not directly. Wrap with select + default for non-blocking ops.

10. **When does a channel block, and how does buffering size affect performance?**  
   Unbuffered: blocks immediately. Buffered: blocks when full/empty.

---

## 🔀 Select

1. **How does Go's `select` choose between multiple ready channels?**  
   Pseudo-randomly among ready cases for fairness.

2. **What is the role of `default` in a `select` block? What are its pros and cons?**  
   Makes select non-blocking. May cause busy loops.

3. **How would you implement timeouts and cancellation with `select`?**  
   Use `time.After` or `context.Done()`.

4. **What happens when multiple cases in a `select` are ready at once?**  
   One is chosen pseudo-randomly.

5. **How does `select` interact with `nil` channels?**  
   Ignores them (never ready).

6. **How would you simulate a fan-in/fan-out pattern using `select`?**  
   Fan-in: merge channels. Fan-out: range over shared channel.

7. **How does Go implement fairness in `select` internally?**  
   Randomized polling of cases.

8. **What happens when a `select` is blocked but one of the channels is closed?**  
   The receive case fires with zero value.

9. **How can you use `select` to implement a graceful shutdown?**  
   Use a `done` channel and select on it.

10. **How does a `select` block behave inside an infinite loop, and what are the risks?**  
   Risk: CPU burn if no sleep/default, or deadlock if all channels block.

---

## 🧩 Interface

1. **How are interfaces implemented under the hood in Go (type + data pair)?**  
   Stored as `itab` (type info) + pointer to value.

2. **What’s the difference between an interface with value `nil` and an interface that is `nil`?**  
   `(nil, type)` ≠ `nil`. Check with `if v == nil`.

3. **How does Go determine if a type satisfies an interface implicitly?**  
   By checking method sets during compilation.

4. **Can you assign an interface to itself? What does it mean semantically?**  
   Yes, but can cause hidden type conversion.

5. **How do type assertions differ from type switches? When would one fail silently?**  
   `val, ok := iface.(Type)` safe; panic without `ok`.

6. **How does the empty interface affect performance and type safety?**  
   Flexible but requires type assertion; slower.

7. **How do you implement polymorphism using interfaces in Go without inheritance?**  
   Define common methods and satisfy interface.

8. **What are the implications of method sets for interface satisfaction?**  
   Only exported or pointer methods count depending on receiver type.

9. **How would you mock an interface for unit testing with dependency injection?**  
   Use manually created mock structs or frameworks like gomock.

10. **How does Go handle interface values in memory (interface boxing/unboxing)?**  
   Value is boxed (wrapped), then unboxed via type assertion or reflection.
"""

# Save to markdown file
md_path = Path("/mnt/data/Advanced_Golang_Questions.md")
md_path.write_text(content)

md_path
