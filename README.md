# gouno

gouno封装了unoconv的操作，提供restful接口供转化文件，并且在unoconv基础上，额外提供了:

+ 基于Docker，快速安装并运行
+ 并行执行转化任务
+ 设置转化超时时间
+ 限制转化文件的大小

## 项目结构

```
.
├── boot.sh              # Dockerfile入口脚本
├── build.sh             # 项目编译脚本
├── config.go            # 存储配置信息
├── Dockerfile
├── handler.go
├── main.go
├── README.md
├── supervisord.conf     # supervisor配置文件，管理unoconv及gouno服务
├── unoconv.go           # 主要工具
└── vendor               # govendor固化依赖
    │   └── ...
    └── vendor.json

```

## 部署
### 拉取代码
```   
git clone https://www.github.com/xiongsyao/gouno.git && cd gouno
```

### 编译   
当前项目采用`golang:1.8`镜像作为编译环境，编译脚本已经写好，为`build.sh`
```
chmod 777 build.sh
./build.sh
```   
成功后将在当前目录下生成名为`gouno`的可执行文件。
首次执行`build.sh`命令时，会先拉取`golang:1.8`镜像，所需时间较长。

### 生成镜像
当前目录下执行:
```
docker build -t gouno .
```

## 运行服务
在跑运行命令之前，需要先了解以下几点:
+ 附件转化依赖于libreoffice，**容器需要挂载宿主机的字体文件**，否则转化的文件中包含中文时会乱码。
+ 内部服务的端口为3000

启动容器:
```
docker run -d -p 127.0.0.1:3000:3000 -e TIMEOUT=30 -e WORKER=5 -e MAXSIZE=-1 -v /usr/share/fonts:/usr/share/fonts/extra --name=gouno gouno
```

## 使用

假如我们想把一个docx的文件转化成pdf，通过使用curl，示例如下：
```
curl --form file=@example.docx --form "compress=1" http://127.0.0.1:3000/unoconv/pdf > example.pdf
```
其中，
+ `compress`为可选参数，当`compress`为1且转化后的文件包含媒体文件(例如图片)时，会压缩媒体文件。
+ `/unoconv/{format-to-conver-to}` 可以接受的转化规则如下:

    |Input|Output|
    |-----|------|
    |html|pptx,docx,xlsx,pdf|
    |docx|pdf,html|
    |xlsx|docx,html,pdf|
    
    docx类型的转化为html效果最佳，docx类型的，依情况转化为pdf、html效果较好。
    
    更多转化规则，自行摸索。
