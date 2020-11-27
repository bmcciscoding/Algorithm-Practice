# 字符串

首先，字符串看起来很简单，但是如果你想用得好，不仅要求需要理解它是如何工作的，也要区分字符串和byte、character、rune 的区别。区分 Unicode 和 UTF-8 的区别。区分 string 和 string literal 的区别。

多问自己：什么时候该用索引操作字符串？为什么得不到第n个字符？这两个问题会带你走向现代文本工作原理的诸多细节。

## 字符串是什么？

在 GO 里，字符串就是**只读字节切片**（read-only slice of bytes）。下文假设你是理解切片，并明白其是如何工作的。

```go
const str = `⌘` 
```
思考下，`len(str)` 输出多少？


```GO
const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" // fmt.Println => ��=� ⌘
```
假设我们写了上面这些字符串，是一些16进制的，打印出了乱码。

## 字符串常用方法

格式化字符串
```go
p1 := "i am p1"
str := fmt.Sprintf("foo: %s", bar)
```


**参考**
- [The Go Blog](https://blog.golang.org/string)