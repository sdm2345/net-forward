

net-forward
软件层面流量转发

go build . -o net-forward 
./net-forward tcp/0.0.0.0/7780/tcp/www.baidu.com/80

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o  bin/net-forward.linux  .
 
