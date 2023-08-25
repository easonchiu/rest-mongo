
<!-- # rest-mongo -->

<!-- ## 接口配置 -->

```yml
collections:
  users:
    name: {type: string, pattern: /[0-9]/, index: 1, unique: true, required: true}
    age: {type: int}
    data: {type: json}
    $$indexes:
      - [id_1, name_-1]

api:
  - /users
    collection: users
    method: GET
    first: true
    searchable: [name, age]
    retain: [name]

  - /users/:id
    collection: users
    method: GET
    retain: [name]
  
  - /users
    collection: users
    method: POST
  
  - /users
    collection: users
    method: PATCH
    modifiable: [name]

  - /users
    collection: users
    method: PUT

  - /users
    collection: users
    method: DELETE
    many: true
    condition: [age]

  - /users/:id
    collection: users
    method: DELETE
```

### collections

用于建立 collection，其下面的 `users` 说明要建立一个名为 `users` 的集合
`users` 内部除 `$$indexes` 用于保留字段，其他均为定义数据的格式：
 - type: 数据类型，支持 `int` `string` `bool` `float` `array` `json` 这几类格式，其中 `array` 与 `json` 不做内部结构校验
 - pattern: 正则表达式，在创建或更新数据时，会对目标字段做校验
 - index: 该字段的单索引，请用 `1` 或者 `-1`
 - unique: 唯一索引，配置为 `true` 之后，重复创建会报错
 - required: 必填字段，如果简单的非空即可，可以使用该功能，若有严格的校验，建议使用 `pattern`
 - $$indexes: 保留字段，用于创建该集合的联合索引

### api

用于创建具体的 api 接口，其中 method 仅支持 `GET` `POST` `PATCH` `PUT` `DELETE`，method 决定了该接口的功能

**method 定义**
- GET 获取数据
- POST 创建数据
- PATCH 局部更新
- PUT 全量更新（注意全量更新意味着，除 id 以外，新的数据会直接替换原有数据）
- DELETE 删除数据

