##proxy test:

```
curl -Lv  --proxy http://127.0.0.1:9999 https://www.jetbrains.com/intellij-repository/snapshots/com/jetbrains/gateway/BUILD/222.3739.24-CUSTOM-SNAPSHOT/BUILD-222.3739.24-CUSTOM-SNAPSHOT.txt

#Start as a server:
nohup ./easy_proxy -port 9999 2>&1 > proxy.log&
#Build for other OS:
GOOS=linux GOARCH=amd64 go build -o easy_proxy main.go
```

#正向代理服务

##简介
本服务可向外提供代理服务，客户端可配置该服务地址实现外网访问，本服务支持http形式的外部代理服务。

##用法
```sdf
#demo
./easy_proxy -port 9999
#as daemon servicesdf
./easy_proxy -port 9999 2>&1 > ./proxy.log &
```
##编译
### 准备
golang 1.16+
### 过程

```
#本地环境编译
go build -o easy_proxy main.go
#跨环境编译给Linux服务器使用
GOOS=linux GOARCH=amd64 go build -o easy_proxy main.go
```

### 测试

在DMZ服务器安装好后，可通过curl命令，测试是否连接成功。
```
```
或者直接在浏览器配置代理服务器，以谷歌浏览器为例：
浏览器--设置--高级--系统--打开您计算机的代理设置
打开手动代理设置，地址设置为部署正向代理服务的服务器地址，比如`172.22.22.100`，端口为启动服务的端口，比如`9999`
