net-forward
应用层流量转发
暂时 仅仅支持 tcp 协议

```bash
 
go get -u https://github.com/sdm2345/net-forward

 
# 测试
./net-forward -l tcp://0.0.0.0:7780 -f tcp://www.baidu.com:80 \
  -l  tcp://0.0.0.0:7781 -f tcp://www.baidu.com:80 
 
# 手工编译
git clone https://github.com/sdm2345/net-forward
# mac 编译 linux 版本
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o  net-forward.linux .


# 本机 启动 redis
测试
go build . 
./net-forward -l tcp://0.0.0.0:80 -f tcp://0.0.0.0:6379

直接访问:
redis-benchmark -p 6379 GET
====== GET ======
  100000 requests completed in 13.30 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1

代理访问:
╰─$ redis-benchmark -p 80 GET                                                                                                                130 ↵
====== GET ======
  100000 requests completed in 13.90 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1

```