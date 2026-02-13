# net/http包详解

## 1.启动实例

## 2.服务端详解

#### （1）Server

Handler是Server中最核心的字段，实现了从请求路径path到具体处理函数 handler 的注册与映射

在构造Server对象时，如果没有显性声明Handler字段，则取net/http包下的单例对象DefaultServeMux(ServeMux类型)兜底。

```go

type Serve struct{
    Addr string     //地址
    Handler Handler//路由处理器
    ...
```

   }

#### （2）Handler

Handler是一个接口，声明ServerHTTP方法

该方法的作用，将接收的请求Request中的请求路径path映射到具体处理器函数handler，进行处理和响应

*Request接收客户端的请求参数，将处理结果写入ResponseWrite，再响应给客户端

**区分一下：Handler是路由处理器，handler是路径处理函数**

```go
type Handler interface{
    ServerHTTP(ResponseWriter, *Request)
}
```

####  （3）ServeMux

ServeMux是对Handler的具体实现，内部有一个map，维护从path到handler的映射关系

```go
type ServeMux struct{
  mu     sync.RWMutex
    tree   routingNode//存储路由树的根节点
    index  routingIndex//路由索引
    mux121 serveMux121 // used only when GODEBUG=httpmuxgo121=1
}
```

routingNode是一个自定义的树形结构节点，用于高效的存储和匹配URL路径。在1.23+版本引入更高效的基于树的路由算法，取代简单的 遍历切片 的方式。

routingIndex:

```go
type routingIndex struct {
    /*
    倒排索引
字段说明：
(1)routingIndexKey：对于路径/a/b/c,
"a"->key{0,"a"}
"b"->key{1,"b"}
"c"->key{2,"c"}，
(2)[]*pattern:存储所有在该位置有该值（例如{1,"b"}）的已注册路由模式
如果注册一下模式：/a/b,/a/b/c,/a/c,/b/a,/a/{x},那么键{1,"b"}包含/a/b,/a/b/c,b在第一段，不包含/a/c,/a/{x},/b/a
  */
    segments map[routingIndexKey][]*pattern
    //存储所有以多段通配符（multi-segment wildcard）结尾的路由模式。
    multis []*pattern
}
```

## 3.客户端详解


