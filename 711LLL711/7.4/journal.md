# 🚀7.4 task

## 💡Things I Learned  

- question solved-- closure
    1. 打包带走。   
        内层函数g把外层f的局部变量a装好，  
        然后外层的f把g送走，  
        这样g就带着a去任何地方享用而不用受f限制了。  
        内层函数使用了外层函数的变量，最后外层函数返回内层函数 ，虽然外层函数已经结束执行了，但是调用内层函数还是可以使用到这些变量   
    2. 闭包的作用：封装私有变量、实现回调函数（实现异步）、延迟计算   
    3. 内存泄漏问题  
        由于闭包会将它的外部函数的作用域也保存在内存中，因此会比其他函数更占用内存，这样的话，如果过度使用闭包，就会有内存泄露的威胁。  
    4. 变量拷贝   
        假如外层f有两个同级的g和h闭包函数，和在f中定义的变量a,实际上 g，h分别获得了a的拷贝，当g中修改了a的值，不会影响到h中a的值  
    5. 实践--用闭包实现斐波那契数列  
        思路：每次输出需要用到前两个数据，构建内层函数闭包，把他们放在外层函数中访问。
        [fibonacci](closure.go)  
- go
## 🖥️Go 

### methods  

1. definition  
    A method is a function with a special receiver argument.  
    ~~~go
    package main

    import (
        "fmt"
        "math"
    )

    type Vertex struct {
        X, Y float64
    }

    func (v Vertex) Abs() float64 {
        return math.Sqrt(v.X*v.X + v.Y*v.Y)
    }

    func main() {
        v := Vertex{3, 4}
        fmt.Println(v.Abs())
    }
    ~~~
    当然，也可以写成普通的函数的形式。
    ~~~go
    func Abs(v Vertex) float64{
        return math.Sqrt(v.X*v.X + v.Y*v.Y)
    }
    func main() {
        v := Vertex{3, 4}
        fmt.Println(Abs(v))
    }
    ~~~
    note:You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).  
2. Pointer receivers
    Methods with pointer receivers can modify the value to which the receiver points  
    相当于函数中传的是指针，可以修改真实值，而只传值的话之能修改副本的值  
    使用pointer receivers时，即使调用方法时使用的是值a而非指针&a，依然不会报错。因为go解释器自动解释为&a。  
    但对于函数而言，如果出现相同的情况就会报错。  
    同样 ，如果方法声明的是值，如果传进去的是指针也不会报错。而函数会报错  

    why use pointer redceivers?
    - modify the value 
    - avoid copying a huge struct or other variables.(more efficient)

### interfaces  

1. definition  
    An interface type is defined as a set of method signatures.(no complemention)   
    ~~~go
    type InterfaceName interface {
    Method1()
    Method2()
    // ...
    }
    ~~~
2. 类比python中的鸭子模型(类、实例、多态)  
    在go中，当不同的结构有相同的方法名称和签名时，这些相同的方法可以组成一个接口，任何声明为这个接口的变量都可以调用这些方法   
    go：
    ~~~go
    package main
    import "fmt"
    // Shape is an interface that defines a common method.
    type Shape interface {
        Area() float64
    }

    // Circle is a struct representing a circle.
    type Circle struct {
        Radius float64
    }

    // Area calculates the area of a Circle.
    func (c Circle) Area() float64 {
        return 3.14 * c.Radius * c.Radius
    }

    // Rectangle is a struct representing a rectangle.
    type Rectangle struct {
        Width  float64
        Height float64
    }

    // Area calculates the area of a Rectangle.
    func (r Rectangle) Area() float64 {
        return r.Width * r.Height
    }

    func main() {
        circle := Circle{Radius: 5}
        rectangle := Rectangle{Width: 3, Height: 4}

        shapes := []Shape{circle, rectangle}

        for _, shape := range shapes {
            fmt.Println(shape.Area())
        }
    }
    ~~~
    python:  
    ~~~python
    class Shape:
    def area(self):
        pass

    class Circle:
        def __init__(self, radius):
            self.radius = radius

        def area(self):
            return 3.14 * self.radius * self.radius

    class Rectangle:
        def __init__(self, width, height):
            self.width = width
            self.height = height

        def area(self):
            return self.width * self.height

    circle = Circle(radius=5)
    rectangle = Rectangle(width=3, height=4)

    shapes = [circle, rectangle]

    for shape in shapes:
        print(shape.area())
    ~~~
3. interface values  
    interface values can be thought of as a tuple of a value and a concrete type:(value ,type)   
    接口值可以被指定为接口中不同的value和type  
4. Interface values with nil underlying values    
    If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.  
5. Nil interface values  
    A nil interface value holds neither value nor concrete type.  
    会有runtime error报错，因为无法调用到具体的方法  
6. The empty interface
    interface{} 
    内部没有任何函数签名  
7. Type assertions  
    A type assertion provides access to an interface value's underlying concrete value.
    ~~~go
    t,ok := i.(T)
    ~~~
    the interface value i holds the concrete type T and assigns the underlying T value to the variable t.   
    If i holds a T, then t will be the underlying value and ok will be true.   
    如果是 `t:=i.(T)` ，如果i接口中没有T类型，程序会报错panic  
8. Type switches
    ~~~go
    switch value := interfaceVariable.(type) {
    case Type1:
        // code to be executed if the interfaceVariable holds Type1
    case Type2:
        // code to be executed if the interfaceVariable holds Type2
    // additional cases for other types
    default:
        // code to be executed if the interfaceVariable holds a type other than Type1 or Type2
    }
    ~~~

9. Stringers-- a commonly used interface  
    By implementing the String() method, you can control how your type is displayed when formatted as a string   
    当调用fmt时会自动寻找它的方法的定义，通过自定义String方法，可以调整输出的格式  
10. error  
    The error type is a built-in interface similar to fmt.Stringer  
    As with fmt.Stringer, the fmt package looks for the error interface when printing values.     
    通过自定义error方法，可以调整错误信息打印的格式    
    ~~~go
    type error interface {
        Error() string
    }
    ~~~

    A nil error denotes success; a non-nil error denotes failure.   

    ~~~go
    package main

    import (
        "fmt"
        "math"
    )

    type ErrNegativeSqrt float64

    func (e ErrNegativeSqrt) Error() string {
        return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
    }

    func Sqrt(x float64) (float64, error) {
        if x < 0 {
            return 0, ErrNegativeSqrt(x)
        }
        return math.Sqrt(x), nil
    }

    func main() {
        fmt.Println(Sqrt(2))
        fmt.Println(Sqrt(-2))
    }
    ~~~

11. Readers 
    func (T) Read(b []byte) (n int, err error)  
    n: the number of bytes read  
    error: an error value  
    ~~~go
    func main() {
	reader := strings.NewReader("Hello, World!")
	buffer := make([]byte, 5)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break // End of file
		}
		fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])
		buffer = make([]byte, 5) // Reset the buffer
	}
    }
    ~~~
12. Images
    ~~~go
    type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
    }
    ~~~

### generics  

1. Type parameters  
    comparable ,any ...
    e.g.   
    ~~~go
    func Index[T comparable](s []T, x T) int {
        for i, v := range s {
            // v and x are type T, which has the comparable
            // constraint, so we can use == here.
            if v == x {
                return i
            }
        }
        return -1
    }
    ~~~

### concurrency  
1. Goroutines  
    A goroutine is a lightweight thread managed by the Go runtime.  
     use  `go f(x, y, z)`  
2. Channels  

    ~~~go
    ch <- v    // Send v to channel ch.
    v := <-ch  // Receive from ch, and
           // assign value to v.
    ~~~       
