# LANGUAGE GUIDE
为了快速上手的学习 GO 这门语言，我决定通过和 Swift 对比。

## The Basics

### Constants and Variables

```go
const PI = 3.1415926
var age int
age := 18
```

### Type Annotations
```go
var msg string
msg = "hello world"
var hour, min, second int
```

### Naming Constants and Variables
```go
// todo
```

### Comments
```go
// single line
/*
    mutil lines
*/
```

### Semicolons

### Intergers
[整型范围](https://golang.org/ref/spec#Numeric_types)
```go
import "math" // 定义在 math 这个包里
const MaxUInt32 = math.MaxUint32
const MinInt32 = math.MinInt32
```
  
### Type Safety and Type Inference
```go
var age int
age = 18.5 // 报错: constant 18.5 truncated to integer
```

### Numeric Literals
```go

```

