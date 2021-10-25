# mongorm
> 基于 `go.mongodb.org/mongo-driver` 二次封装，可以进行链式操作的简单orm

## 支持操作
* 搜索
    - where  
    where 条件支持`and`和`or`，关键字分别为`where`和`orWhere`。
    通过`字段`，`操作符`，`值`进行搜索条件的输入。

    ```go
    where := mongorm.
        Where("create_time", ">", 1545623340).
        Where("_id", "==", mongorm.ObjectID("5c20575e716de1ba769cc295")).
        Result()
    ```
  
    - WhereOrResult  
    当出现两个条件出现`或`存在时，需要构建出两个`and`条件，然后进行`or`组合
    
    ```go
        whereOne := mongorm.
        	Where("create_time", ">", 1545623340)
        whereTwo := mongorm.
        	Where("_id", "==", mongorm.ObjectID("5c20575e716de1ba769cc295"))
        orWhere := mongorm.WhereOrResult(whereOne, whereTwo)
    ```
操作符号支持如下输入：`==`，`>`，`>=`，`<`，`<=`，`!=`，`in`，`IN`，`nin`，`NIN`，`not in`，`NOT IN`，`regex`，`REGEX`
    
    
`其他操作可以通过default选项进行输入，注意保证输入的正确性。`


* 属性设置

在查询时，需要分页，排序，强制命中索引时，可通过对应方法设置需要的内容。同样采用链式操作，进行组合设置。
    
```go
opt := mongorm.
    SetSelect("order_number,_id,create_time"). // 显示字段
    SetLimit(2).  // 限制条数
    SetOffset(2). // 偏移条数
    SetHint("_id"). // 索引名字
    SetSort("create_time", mongorm.DESC).  // 排序
    SetSort("update_time", mongorm.ASC). // 多条排序
    GetOneOption()
```
    
> 查询一条记录时，结尾使用`GetOneOption()`，查询多条数据时，结尾使用`GetOption()`


* 更新数据

可以同时进行`set` 和 `incr` 操作。有个彩蛋是，`Result()`方法支持传入字段命，用于设置为当前更新时间，比如传入`update_time`。

```go
where := mongorm.
    Where("_id", "==", mongorm.ObjectID("5c20575e716de1ba769cc295")).
    Result()
			
update := mongorm.
	Set("field_time",  time.Now().Unix()).
	Set("field_name",  "测试一下").
	Incr("field_status",  1).
	Incr("field_result",  1).
	Result()
ret := table.NewDemo().Update(where, update)
``` 

* 插入数据

插入一条数据

```go
where := mongorm.NewInsertOne().Value("field_supplier", field.Supplier).Result("update_time", "create_time")
```
