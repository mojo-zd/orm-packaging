### 目的
本项目主要针对平时使用的xorm、beego orm做了简单封装
- 对复杂操作进行包装
- 简化代码复杂度，提高代码重用率
- 增加分页功能，更贴切实际项目使用

## [beego orm](#0)

### [Insert](#1)
### [Query](#2)
### [Query by page](#3)
### [Delete](#4)
### [Update](#5)
### [Update with map](#6)

## xorm

<h2 id="0"> beego orm </h2>
<h3 id="1">Insert</h3>

> 添加记录并返回持久化的记录

```
    Insert(v interface)
```

<h3 id="2">Query</h3>

> 根据指定字段进行查询

```
    Read(v interface, cols ...string)
```

<h3 id="3">Query By Page</h3>

> 分页查询

```
    // condition eg:map[string]interface{}{"age__gt":1, "name__icontains":"mojo", "sex":1} 查找age大于1 & name包含mojo & sex为1的记录
    QueryElement(page *pagination.Pagination, condition map[string]interface{}, table interface{})
```

<h3 id="4">Delete</h3>

> 删除指定的记录

```
    // 根据指定cols作为条件进行删除
    Delete(v interface, cols ...string)
```

<h3 id="5">Update</h3>

> 更新记录

```
    // 更新指定字段
    Update(v interface, cols ...string)
```

<h3 id="6">Update with map</h3>

> 以map作为对象来更新

```
    // condition 为条件  m为将要更新的值 table为指定表对应的实体类  调用该方法后持久化后的对象将会赋值给table
    UpdateWithMap(condition, m map[string]interface{}, table interface{})
```
