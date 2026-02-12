# net/http包详解

## 1.启动实例

## 2.服务端详解

    1.Serve

在构造handler时，如果没有显性声明，就使用默认全局路由(DefaultServeMux)
    type Serve struct{
        Addr string            //地址
        Handler Handler//路由处理器
    }

## 3.客户端详解


