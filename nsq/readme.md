### nsq使用

1. nsq lookup的配置
    ```shell
    docker run --rm --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq /nsqlookupd
    ```
   
2. 获取nsq lookup的容器内ip地址
    ```shell
    docker inspect -f '{{.Name}} - {{.NetworkSettings.IPAddress }}' $(docker ps -aq)
    ```  
   假设ip地址为: 172.17.0.3 
3. nsqd的配置
    ```sehll
   # --broadcast-address为nsqd自己的ip地址
    docker run --rm --name nsqd -p 4150:4150 -p 4151:4151 \
    nsqio/nsq /nsqd \
    --broadcast-address=172.17.0.4 \
    --lookupd-tcp-address=172.17.0.3:4160
    ```
   
4. 参考
   ```http
    https://www.liwenzhou.com/posts/Go/go_nsq/
   ```