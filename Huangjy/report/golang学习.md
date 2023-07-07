## Golang学习记录

### Go的运行期反射

> &emsp;&emsp;reflect包实现了运行时反射，允许程序操作任意类型的对象。典型用法是用静态类型interface{}保存一个值，通过调用TypeOf获取其动态类型信息，该函数返回一个Type类型值。调用ValueOf函数返回一个Value类型值，该值代表运行时的数据。Zero接受一个Type类型参数并返回一个代表该类型零值的Value类型值。
> &emsp;&emsp;使用` import "reflect" `引入这个包

如何运用这个包的反射？

假设当前有一个struct，我们要将tag和字段提出来建立一个` map[string]interface{} `
```Go
type S struct {
    Name string `field:"name"`
    ID uint64 `field:"id"`
}
```
单独定义一个GetMap函数，接受一个结构的interface{}，返回map
```Go
func getMap(s interface{}) map[string]interface{} {
    var m map[string]interface{} = make(map[string]interface{})
    reType := reflect.TypeOf(s)
    reValue := reflect.ValueOf(s)
    for i, n := 0, reType.NumField(); i < n; i++ {
    	sField := reType.Field(i)
    	m[sField.Tag.Get("field")] = reValue.FieldByName(sField.Name).Interface()
    }
    return m
}
```
可以得到结果
> ```Go
> {Name:Astruct ID:20220114514}
> name: Astruct(string)
> id: 20220114514(uint64)
> ```

### Go的并发安全与锁

尝试运行这段代码

```Go
func main() {
	var (
		wg sync.WaitGroup
		x int
	)
	wg.Add(2)
	go func() {
		for i := 0; i < 100000; i++ {
			x += 1
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100000; i++ {
			x -= 1
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(x)
}
```

一般来看，x先增加100000后减少100000，结果应该为0，然后运行结果是不确定的，但是绝对不是0，这是因为两个并发的协程在同时操作一个变量，为了能够获得所期望的结果，需要添加互斥锁Mutex，在x增加和减少之前给x上锁，读完完毕后开锁，始终保持在同一个时间只有一个协程读写这个变量。

添加一个Mutex `var lock sync.Mutex`将for循环中的语句改为
``` Go
lock.Lock()
x += 1 // x -= 1
lock.Unlock()
```

最后得到结果为 0

### Go的运行期反射再使用

从mongodb中查找数据时，有时候会涉及从多个集合中查找不同数据，封装到一个函数中。

``` Go
func baseQuery(ctx context.Context, filter *bson.D, c *mongo.Collection, t reflect.Type, opts ...*options.FindOptions) ([]*interface{}, error) {
	cur, err := c.Find(ctx, filter, opts...)
	infos := []*interface{}{}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for cur.Next(context.TODO()) {
		p := reflect.New(t)
		if err = cur.Decode(p.Interface()); err != nil {
			log.Println(err)
			return infos, err
		}
		c := p.Elem().Interface()
		infos = append(infos, &c)
	}
	return infos, nil
}
```
在这里，函数接受一个`reflect.Type`的参数，通过`reflect.New`按照这个类型定义一个新变量`reflect.Value`，这个Value是指向类型与所传入结构相同类型的结构的指针，将这个Value传入Decode解码后，用Elem方法解指针，最后加入到切片中返回。