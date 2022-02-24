send文件夹： send.pb.go：proto文件生成的go文件 SendService.go：定义并实现发送服务，测试时可以更新const里面的变量切换接收者邮箱，根据告警参数发送不同的邮件。

server.go：main函数启动发送服务

Client文件夹： client.go：启动客户端传入数据测试

说明：首先先打开server.go启动服务，再启动客户端即可 接收者邮箱可以在SendService.go里面更换，发送者邮箱要更换的话首先需要更换协议，邮箱以及授权码
#更新
不用再代码里面更新授权码邮箱等信息，直接可以在配置文件config里面进行修改。
