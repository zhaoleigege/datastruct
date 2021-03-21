### 数据结构

1. Git项目的创建

   ```shell
   echo "### 数据结构" >> README.md
   git init
   git add README.md
   git commit -m "datastruct初始化"
   git remote add origin git@github.com:zhaoleigege/datastruct.git
   git push -u origin master
   ```

2. golang项目的初始化

   ```shell
   go mod init github.com/zhaoleigege/datastruct
   ```

   IDEA设置支持`vgo`

   `Preferences...` -> `GO` -> `Go Modules (vgo)` -> `Enable Go Modules (vgo) integration` 打上勾

