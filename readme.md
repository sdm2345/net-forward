

net-forward
软件层面流量转发
git clone https://github.com/sdm2345/net-forward

cd net-forward;
go build ./ -o net-forward

测试 
./net-forward tcp/0.0.0.0/7780/tcp/www.baidu.com/80

mac 编译 linux 版本
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o  net-forward.linux  .
 
