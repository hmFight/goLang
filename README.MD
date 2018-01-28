## note for go

目录说明

 ```
--practice
  	--ext    extra说明
  	--idgen  go练习项目：http id 生成器
 ```

### Id 生成器

#### Feature

*  Snowfalke ID ✔️
*  自增 Id 生成✔️
*  批量获取 
*  Id 重置✔️

#### Http Api

```shell
//从 snowflake id生成器获取id
curl -G [-d 'num=num'] http://host/id/snowflake   

//从 指定 key 的自增id生成器获取id,
//key - 如果不指定 key,则从默认 key 为 default
//num - 获取数量，默认数量为1
curl -G [-d 'key=val&num='] http://host/id/incr

//重置所有自增 id 生成器 当前 id 为 0
curl http://host/id/incr/resetall

//重置指定 key 自增 id 生成器 当前 id 为 0
curl -G [-d 'key=val'] http://host/id/incr/reset

//重置指定 key 自增 id 生成器 当前 id 为 数值 num ，待增加
curl -G [-d 'key=val&to=tar_id'] http://host/id/incr/reset

//待完善
//批量获取 id
curl -G [-d 'key=val&num=num'] http://host/id/incr
```

