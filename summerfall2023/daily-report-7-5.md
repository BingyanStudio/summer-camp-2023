# 7-5

## go

### 数组
```go
	var a[2] string
```
### 切片


```go
func main(){
	prinmes := [6]int {1,2,3,4,5}

	var s[]int = prinmes[2:5]

	fmt.Println(s)
}
```
output:[3,4,5]  

左闭右开  

正确的：
```go
	names := [4]string{
		"mollu","molly","johu","john"}
```
```go
	names := [4]string{
		"mollu","molly","johu","john",}
```
```go
	names := [4]string{
		"mollu","molly","johu","john",
        }
```
错误的：
```go
	names := [4]string{
		"mollu","molly","johu","john"
        }
```

切片下界的默认值为 0，上界则是该切片的长度。  

```go
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

```
output:[3,5,7]
       [3,5]
       [5]

---

*Print**与**Println**和**Printf的区别*

Printf是Go语言中的一个函数，用于格式化输出。它允许你使用占位符指定输出的格式，并将对应的值插入到格式化字符串中。  



```go
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Print(s)

	s = s[:2]
	fmt.Print(s)

	s = s[1:]
	fmt.Print(s)
}
```
output:[3,5,7][3,5][5]

-----
*print与fwt.Print*

在 Go 中，`print` 是一个内置函数，用于将参数打印到标准输出，但它不提供格式化功能。它接受任意数量的参数，并将它们转换为字符串后直接打印。

与之相反，`fmt.Print` 是 `fmt` 包中的函数，提供了更灵活的格式化打印功能。它接受类似 C 语言的格式字符串，并根据格式字符串将参数格式化后打印。

因此，在代码中，使用 `print(words)` 不会产生编译错误，因为 `print` 是一个有效的函数。然而，它只会简单地将 `words` 打印到标准输出，而不进行格式化。

如果想要更好的格式化输出，包括换行符和其他格式选项，应该使用 `fmt.Print` 或 `fmt.Println` 函数。

----

切片

len(s): 切片的长度就是它所包含的元素个数。

cap(s): 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

注意：  
切片的零值是 nil。  
nil 切片的长度和容量为 0 且没有底层数组。

make(a[],len,cap)创建元素为零的切片

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```
警告：
redundant type from array, slice, or map composite literal  
在Go语言中，当你使用数组、切片或映射的复合字面量时，如果你已经声明了该类型，那么在复合字面量中不需要再指定类型。  

```go
board := [][]string{
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
}
```
可以将类型声明简化为：

```go
board := [][]string{
	{"_", "_", "_"},
	{"_", "_", "_"},
	{"_", "_", "_"},
}
```
向切片追加元素

s = append(s, 2, 3, 4)
当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。

### range

for 循环的 range 形式可遍历切片或映射。

```go
for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
```
忽略索引或值
```go
for i, _ := range pow
for _, value := range pow
```
```go
for i := range pow
for value := range pow
```

### map 映射

1.声明

```go
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
```
```go
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

```

2.用make构造map
在映射 m 中插入或修改元素：

m[key] = elem
获取元素：

elem = m[key]
删除元素：

delete(m, key)
通过双赋值检测某个键是否存在：

elem, ok = m[key]

```go
v, ok := m["Answer"]
fmt.Println("The value:", v, "Present?", ok)
```

这里使用了一个特殊的 Go 语法，叫作**多返回值**。在这个例子中，`v, ok := m["Answer"]` 这一行代码尝试从 map `m` 中获取键为 "Answer" 的值。

- 如果键存在，那么 `v` 将会被赋值为该键对应的值，而 `ok` 的值将会是 `true`。
- 如果键不存在，那么 `v` 将会被赋值为 map 值类型的零值，而 `ok` 的值将会是 `false`。

### 一个小问题

fmt.Println 默认会打印切片的指针地址。

如果你想要打印切片中的元素，可以使用循环来逐个打印单词，或者使用 strings.Join 函数将单词连接成一个字符串后进行打印。

```go
for _, word := range words {
		fmt.Println(word)
		wordCount[word]++
	}
```
```go
fmt.Println(strings.Join(words, "\n")) // 使用换行符连接单词并打印
```

### 闭包

（联想Python中的闭包  
闭包可以访问和修改外部函数的变量
）  

----
而这在go中是不能实现的  
假如外层f有两个同级的g和h闭包函数，和在f中定义的变量a,实际上 g，h分别获得了a的拷贝，当g中修改了a的值，不会影响到h中a的值

go:使用副本  
Python：直接引用