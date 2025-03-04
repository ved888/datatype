### 4. What will happen if you exceed the maximum value of an int in Go?
In Go, if you exceed the maximum value of an int type, the value will wrap around (also known as overflow). This means the value will cycle back to the minimum value of that type and continue counting from there.

### 5. How does Go handle integer overflow?
Go does not panic or raise an error when an integer overflow happens. Instead, it performs modular arithmetic. The value wraps around within the range of the integer type.

👉 The range of the integer type determines how Go will wrap the number:

`int8: -128 to 127
int16: -32,768 to 32,767
int32: -2,147,483,648 to 2,147,483,647
int64: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807`

### 7. What is the size of an int on a 32-bit system?
On a 32-bit system, the size of an int in Go is 4 bytes (32 bits), which means it can store values from `-2,147,483,648 to 2,147,483,647.`

### 8. What is the size of an int on a 64-bit system?
On a 64-bit system, the size of an int in Go is 8 bytes (64 bits), which means it can store values from `-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807.`

### 16. Can you assign an int8 value to an int variable? Why or why not?
❌ No, Go does not allow direct assignment between different integer types without conversion.

`var a int8 = 100`
`var b int = a // Error`

You need to explicitly convert the type:

`var b int = int(a) // ✅ Works`

### 19. What happens if you try to divide an integer by zero in Go?
It will cause a runtime panic.

Example:

`fmt.Println(10 / 0) // panic: runtime error: integer divide by zero`

### 20. How do you convert a string to an integer in Go?
You can use the strconv.Atoi() function.