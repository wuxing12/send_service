package send

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"textgrpc/configs"
	"time"

	"github.com/jordan-wright/email"
)

var (
	cstZone = time.FixedZone("CST", 8*3600) // 东八区
)

type SendService struct {
	UnimplementedSendServiceServer
}

func (this *SendService) Send(ctx context.Context, in *SendReq) (*SendRsp, error) {
	fmt.Printf("时间戳:%d，性能:%s，维度:%s，值:%f，告警类型: %s \n",
		in.Timestamp, in.Metric, in.Dimensions, in.Value, in.AlertType)
	// 根据告警类型发送相对应的邮件
	if in.AlertType == "SEVERE" && in.Metric == "cpu_rate" {
		SevereSendMail(in.Timestamp, in.Metric, in.Dimensions, in.Value, in.AlertType)
	}
	if in.AlertType == "FATAL" && in.Metric == "cpu_rate" {
		FatalSendMail(in.Timestamp, in.Metric, in.Dimensions, in.Value, in.AlertType)
	}
	if in.AlertType == "WARN" && in.Metric == "cpu_rate" {
		return &SendRsp{Code: 0, Msg: "cpu使用过高警告"}, nil
	}
	if in.AlertType == "SEVERE" && in.Metric == "mem_rate" {
		MemSevereSendMail(in.Timestamp, in.Metric, in.Dimensions, in.Value, in.AlertType)
	}
	if in.AlertType == "FATAL" && in.Metric == "mem_rate" {
		MemFatalSendMail(in.Timestamp, in.Metric, in.Dimensions, in.Value, in.AlertType)
	}
	if in.AlertType == "WARN" && in.Metric == "mem_rate" {
		return &SendRsp{Code: 0, Msg: "内存使用过高警告"}, nil
	}
	return &SendRsp{Code: 1, Msg: "Success"}, nil
}

func SevereSendMail(Time int64, metric string, dim map[string]string, value float64, Alter string) (feedback string) {
	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	ac := configs.GetConfig()

	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	// em.From = "CPU监控提醒 <462118329@qq.com>"
	em.From = fmt.Sprintf("CPU监控提醒 <%s>", ac.Sender)

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = ac.Recipients

	// 抄送
	// em.Cc = []string{Receiver}

	// 密送
	// em.Bcc = []string{Receiver}
	// 设置主题
	em.Subject = "您的CPU即将爆满了," + fmt.Sprintf("告警类型为 “%s” ", Alter)

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte("您的CPU阈值快满了，请尽快查看！！各项指标分别为：\n" +
		fmt.Sprintf("Timestamp: %s \n Metric: %s \n Value: %v \n Dimensions: %+v \n Altertype: %s \n",
			time.Unix(Time, 0).In(cstZone).Format("2006-01-02 15:04:05"), metric, value, dim, Alter))

	// em.Text = []byte("您的CPU阈值爆满了，请马上上线查看！！")
	// 设置服务器相关的配置
	err := em.Send(fmt.Sprintf("%s:%d", ac.Host, ac.Port),
		smtp.PlainAuth("", ac.Sender, ac.Password, ac.Host))

	if err != nil {
		log.Fatal(err)
	}
	log.Println("CPUSevere send successfully ... ")
	return "Send Completely"
}

func FatalSendMail(Time int64, metric string, dim map[string]string, value float64, Alter string) (feedback string) {
	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	ac := configs.GetConfig()

	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	// em.From = "CPU监控紧急通知！！！ <462118329@qq.com>"
	em.From = fmt.Sprintf("CPU监控紧急通知！！！ <%s>", ac.Sender)

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = ac.Recipients

	// 抄送
	// em.Cc = []string{Receiver}

	// 密送
	// em.Bcc = []string{Receiver}

	// 设置主题
	em.Subject = "急！！！您的CPU使用率过高," + fmt.Sprintf("告警类型为 “%s” ", Alter)

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte("您的CPU阈值爆满了，请马上上线修复！！！各项指标分别为：\n" +
		fmt.Sprintf("Timestamp: %s \n Metric: %s \n Value: %v \n Dimensions: %s \n Altertype: %s \n",
			time.Unix(Time, 0).In(cstZone).Format("2006-01-02 15:04:05"), metric, value, dim, Alter))

	// em.Text = []byte("您的CPU阈值爆满了，请请马上上线查看！！")
	// 设置服务器相关的配置
	err := em.Send(fmt.Sprintf("%s:%d", ac.Host, ac.Port),
		smtp.PlainAuth("", ac.Sender, ac.Password, ac.Host))

	if err != nil {
		log.Fatal(err)
	}
	log.Println("CPUFatal send successfully ... ")
	return "Send Completely"
}

func MemSevereSendMail(Time int64, metric string, dim map[string]string, value float64, Alter string) (feedback string) {
	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	ac := configs.GetConfig()

	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	// em.From = "内存监控提醒 <462118329@qq.com>"
	em.From = fmt.Sprintf("内存监控提醒 <%s>", ac.Sender)

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = ac.Recipients

	// 抄送
	// em.Cc = []string{Receiver}

	// 密送
	// em.Bcc = []string{Receiver}
	// 设置主题
	em.Subject = "您的内存即将爆满了," + fmt.Sprintf("告警类型为 “%s” ", Alter)

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte("您的内存阈值快满了，请尽快查看！！各项指标分别为：\n" +
		fmt.Sprintf("Timestamp: %s \n Metric: %s \n Value: %v \n Dimensions: %+v \n Altertype: %s \n",
			time.Unix(Time, 0).In(cstZone).Format("2006-01-02 15:04:05"), metric, value, dim, Alter))

	// em.Text = []byte("您的CPU阈值爆满了，请请马上上线查看！！")
	// 设置服务器相关的配置
	err := em.Send(fmt.Sprintf("%s:%d", ac.Host, ac.Port),
		smtp.PlainAuth("", ac.Sender, ac.Password, ac.Host))

	if err != nil {
		log.Fatal(err)
	}
	log.Println("MemSevere send successfully ... ")
	return "Send Completely"
}

func MemFatalSendMail(Time int64, metric string, dim map[string]string, value float64, Alter string) (feedback string) {
	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	ac := configs.GetConfig()

	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	// em.From = "内存监控紧急通知！！！ <462118329@qq.com>"
	em.From = fmt.Sprintf("内存监控紧急通知！！！ <%s>", ac.Sender)

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = ac.Recipients

	// 抄送
	// em.Cc = []string{Receiver}

	// 密送
	// em.Bcc = []string{Receiver}

	// 设置主题
	em.Subject = "急！！！您的内存使用率过高," + fmt.Sprintf("告警类型为 “%s” ", Alter)

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte("您的内存阈值爆满了，请马上上线修复！！！各项指标分别为：\n" +
		fmt.Sprintf("Timestamp: %s \n Metric: %s \n Value: %v \n Dimensions: %s \n Altertype: %s \n",
			time.Unix(Time, 0).In(cstZone).Format("2006-01-02 15:04:05"), metric, value, dim, Alter))

	// em.Text = []byte("您的CPU阈值爆满了，请请马上上线查看！！")
	// 设置服务器相关的配置
	err := em.Send(fmt.Sprintf("%s:%d", ac.Host, ac.Port),
		smtp.PlainAuth("", ac.Sender, ac.Password, ac.Host))

	if err != nil {
		log.Fatal(err)
	}
	log.Println("MemFatal send successfully ... ")
	return "Send Completely"
}
