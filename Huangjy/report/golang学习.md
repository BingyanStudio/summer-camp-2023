### Golangѧϰ��¼

#### Go�������ڷ���

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