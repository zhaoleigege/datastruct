1. 创建bridge网络
    ```shell
    docker network create nats
    ```
2. 使用网络
    ```shell
    docker run --rm --network nats --network-alias nats-server -d --name nats -p 4222:4222 -p 6222:6222 -p 8222:8222 nats
    ```
   
    ```shell
    docker run --network nats --rm nats-consumer
    ```