1. 环境启动
    ```shell
    docker run -d --name jaeger -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 14250:14250 -p 9411:9411 jaegertracing/all-in-one
    ```
   
2. 概念理解
    * trace: 调用链，其中包含了多个span
    * span: 跨度，计量的最小单位，每个跨度都有开始时间和截止时间

3. jaeger组件
    * agent
       是一个在UDP端口监听span数据的网络守护进程，并将数据批量发送给collector。agent将client library和Collector解耦，
       为client library屏蔽了路由和发现Collector的细节。  
       
       | Port  | Protocol | Function                                                     |
       | :---- | :------- | :----------------------------------------------------------- |
       | 6831  | UDP      | accept [jaeger.thrift](https://github.com/jaegertracing/jaeger-idl/blob/master/thrift/jaeger.thrift) in `compact` Thrift protocol used by most current Jaeger clients |
       | 6832  | UDP      | accept [jaeger.thrift](https://github.com/jaegertracing/jaeger-idl/blob/master/thrift/jaeger.thrift) in `binary` Thrift protocol used by Node.js Jaeger client (because [thriftrw](https://www.npmjs.com/package/thriftrw) npm package does not support `compact` protocol) |
       | 5778  | HTTP     | serve configs, sampling strategies                           |
       | 5775  | UDP      | accept [zipkin.thrift](https://github.com/jaegertracing/jaeger-idl/blob/master/thrift/zipkincore.thrift) in `compact` Thrift protocol (deprecated; only used by very old Jaeger clients, circa 2016) |
       | 14271 | HTTP     | admin port: health check at `/` and metrics at `/metrics`    |
       
       单独启动agent
       
       ```shell
       docker run --rm 
         -p5775:5775/udp \
         -p6831:6831/udp \
         -p6832:6832/udp \
         -p5778:5778/tcp \
         jaegertracing/jaeger-agent --collector.host-port=<JAEGER_COLLECTOR_HOST>:14267
       ```
       
    * collector
       接收agent端发送过来的数据，然后将数据写入后端存储。后端存储是一个可插拔的组件。
       
       | Port  | Protocol | Function                                                     |
       | :---- | :------- | :----------------------------------------------------------- |
       | 14267 | TChannel | used by **jaeger-agent** to send spans in jaeger.thrift format |
       | 14250 | gRPC     | used by **jaeger-agent** to send spans in model.proto format |
       | 14268 | HTTP     | can accept spans directly from clients in jaeger.thrift format over binary thrift protocol |
       | 9411  | HTTP     | can accept Zipkin spans in Thrift, JSON and Proto (disabled by default) |
       | 14269 | HTTP     | admin port: health check at `/` and metrics at `/metrics`    |
       
    * query
       接收查询请求，从后端存储系统检索trace并通过UI展示出来
       
    * 日志服务
        collector将接收到的span数据持久化到日志服务中，query会从日志服务中检索数据
    
4. 参考
   
    * [分布式链路追踪(Tracing)系统 – Jaeger在Golang中的使用](https://www.lizenghai.com/archives/6130.html)