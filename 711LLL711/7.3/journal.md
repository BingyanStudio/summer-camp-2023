# 🚀7.3 task

## 💡Things I Learned  
- markdown(already learned)
- git    
   *  fork -> clone -> edit -> pull request workflow  
- go  
   * download and set the environment  
   * "hello world" project  
   * about mod     
      what is mod?   
      It is recommended to create a module for a project that you intend to share or that has external dependencies. A module provides a way to define and manage dependencies, allowing others to easily build and run your project. To create a module, you can use the go mod init command as mentioned in the previous response.

## 🖥️Golang   
### packages   
1. Every Go program is made up of packages.    
2. import:factored or multipy import  statements   
   ~~~go
   package main

   import (
      "fmt"
      "math"
   )
   ~~~
3. export   
In Go, a name is exported if it begins with a capital letter.   
When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.   
### functions   
1. argument  
   ~~~go
   func add(x int, y int) int {
      return x + y
   } 
   ~~~
2. return result   
   A function can return any number of results.
   ~~~go
   func swap(x, y string) (string, string) {
      return y, x
   }
   ~~~
   named return values   
   Go's return values may be named. If so, they are treated as variables defined at the top of the function.  
3. Function values   
   Functions are values too. They can be passed around just like other values.    
   Function values may be used as function arguments and return values.   
4. Function closures     
   ~~~go
   package main
   import "fmt"
   // fibonacci is a function that returns
   // a function that returns an int.
   func fibonacci() func() int {
      prev, curr := 0, 1 // Initialize the previous and current Fibonacci numbers

      return func() int {
         result := prev
         prev, curr = curr, prev+curr // Calculate the next Fibonacci number
         return result
      }
   }

   func main() {
      f := fibonacci()
      for i := 0; i < 10; i++ {
         fmt.Println(f())
      }
   }
   ~~~
### variables
1. declare   
   ~~~go
   var c, python, java bool
   ~~~
2. Variables with initializers   
   ~~~go 
   var i, j int = 1, 2
   ~~~
3. Short variable declarations   
   usually in function :=  
   ~~~go
   func main() {
	var i, j int = 1, 2
	k := 3
   }
   ~~~
4. basic types   
   ~~~go
   var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
   )
   ~~~
5. zero values   
   Variables declared without an explicit initial value are given their zero value.  
   0 ,false ,""  
6. type convesion  
   The expression T(v) converts the value v to the type T.    
7. Type inference    
   When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.  
8. constants  
   ~~~go
   const Pi = 3.14
   ~~~
9. Numeric Constants   
   high-precision values.  


### Flow control statement   
1. for loop
   ~~~go
   for i := 0; i < 10; i++ {
		sum += i
	}
   ~~~
   without the init and post statements ,"for" is the while in C  
   ~~~go
   for i < 10 {
		sum += i
	}
   ~~~
   Forever
   ~~~go
   for {
	}
   ~~~
2. if else
   ~~~go
   if x < 0 {
      return sqrt(-x) + "i"
   }
   ~~~
   Like for, the if statement can start with a short statement to execute before the condition.   
   Variables declared by the statement are only in scope until the end of the if.   
   ~~~go
      ~~~go
   if v:=2 ;x < 0 {
      return sqrt(-x) + "i"
   }
   ~~~ 
3. switch  
   - the break statement is provided automatically in Go.    
      Go's switch cases need not be constants, and the values involved need not be integers.  
   - Switch without a condition is the same as switch true.  
   ~~~go
   t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
   ~~~
4. defer
   - A defer statement defers the execution of a function until the surrounding function returns.    
   ~~~go
   func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
   }
   ~~~

   - Last-In-First-Out (LIFO) order   
   - The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.  
### more types
1. pointers   
   ~~~go
   var p *int
   i:= 1
   p = &i
   fmt.Println(*p) 
   ~~~
2. structs   
   ~~~go
   type Vertex struct {
	X int
	Y int
   }
   v := Vertex{1, 2}
   v.X = 4
   fmt.Println(v.X)
   p = &v
   fmt.Println((*p).X) #p.X is also OK.
   ~~~
3. arrays   
   ~~~go
   var a [10]int
   ~~~
4. slices  
   ~~~go
   primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
   ~~~
   - Changing the elements of a slice modifies the corresponding elements of its underlying array.  
   ~~~go
   []bool{true, true, false}
   ~~~
   - create an arr and a slice of it
   - omit low or high part of a slice
   - Slice length and capacity
      The length of a slice is the number of elements it contains. -- len(slice)     
      The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.  -- cap(slice)    
   - Creating a slice with make
      slice := make([]Type, length, capacity)   
   - Appending to a slice  
      slice = append(slice, element1, element2, ...)  
      or slice1 = append(slice1 ,slice2 ,...)
   - range
      The range form of the for loop iterates over a slice or map.  

      When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.  
      ~~~go
      var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
      for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	   }
      ~~~
      omit index or element  
      ~~~go
      for i, _ := range pow
      for _, value := range pow
      ~~~   
5. maps   
   ~~~go
   var m map[KeyType]ValueType   
   m := make(map[KeyType]ValueType)
   ~~~
   Mutating Maps
   ~~~go
   Insert or update an element in map m
   m[key] = elem

   Retrieve an element:
   elem = m[key]

   Delete an element:
   delete(m, key)

   Test that a key is present with a two-value assignment:
   elem, ok = m[key]
   ~~~ 

## ❓things that  I am confused about  
   The concept of function closure  
