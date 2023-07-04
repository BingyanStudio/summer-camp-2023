### Golang学习记录

#### Go的运行期反射

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