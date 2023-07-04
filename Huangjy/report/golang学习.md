## Golangѧϰ��¼

### Go�������ڷ���

> &emsp;&emsp;reflect��ʵ��������ʱ���䣬�����������������͵Ķ��󡣵����÷����þ�̬����interface{}����һ��ֵ��ͨ������TypeOf��ȡ�䶯̬������Ϣ���ú�������һ��Type����ֵ������ValueOf��������һ��Value����ֵ����ֵ��������ʱ�����ݡ�Zero����һ��Type���Ͳ���������һ�������������ֵ��Value����ֵ��
> &emsp;&emsp;ʹ��` import "reflect" `���������

�������������ķ��䣿

���赱ǰ��һ��struct������Ҫ��tag���ֶ����������һ��` map[string]interface{} `
```Go
type S struct {
    Name string `field:"name"`
    ID uint64 `field:"id"`
}
```
��������һ��GetMap����������һ���ṹ��interface{}������map
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
���Եõ����
> ```Go
> {Name:Astruct ID:20220114514}
> name: Astruct(string)
> id: 20220114514(uint64)
> ```

### Go�Ĳ�����ȫ����

����������δ���

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

һ��������x������100000�����100000�����Ӧ��Ϊ0��Ȼ�����н���ǲ�ȷ���ģ����Ǿ��Բ���0��������Ϊ����������Э����ͬʱ����һ��������Ϊ���ܹ�����������Ľ������Ҫ��ӻ�����Mutex����x���Ӻͼ���֮ǰ��x������������Ϻ�����ʼ�ձ�����ͬһ��ʱ��ֻ��һ��Э�̶�д���������

���һ��Mutex `var lock sync.Mutex`��forѭ���е�����Ϊ
``` Go
lock.Lock()
x += 1 // x -= 1
lock.Unlock()
```

���õ����0