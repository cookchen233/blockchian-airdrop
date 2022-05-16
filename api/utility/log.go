package utility

//日志的各种策略

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//日志滚动存储策略hook
func RotateLogHook(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) *lfshook.LfsHook {
	baseLogPath := path.Join(logPath, logFileName)

	writer, err := rotatelogs.New(
		baseLogPath+".%Y-%m-%d",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	err_writer, err := rotatelogs.New(
		baseLogPath+".%Y-%m-%d",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	return lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: err_writer,
		log.FatalLevel: err_writer,
		log.PanicLevel: err_writer,
	}, &LogFormatter{})

}


//针对不同错误级别设置不同记录内容
type LogFormatter struct{

}

//格式策略接口方法
func (formatter *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	msg := fmt.Sprintf("[%s] [%s] %s\n", time.Now().Local().Format("2006-01-02 15:04:05"), strings.ToUpper(entry.Level.String()), entry.Message)
	if entry.Level <= log.ErrorLevel{
		msg = fmt.Sprintf("[%s] [%s] %s\n%s\n", time.Now().Local().Format("2006-01-02 15:04:05"), strings.ToUpper(entry.Level.String()), entry.Message, entry.Data["stack"], )
	}

	return []byte(msg), nil
}


// 日志邮件 hook
type MailHook struct {
	User string
	Pass string
	Host string
	Port string
	Receivers []string
}

// 触发执行接口方法
func (hook *MailHook) Fire(entry *log.Entry) error {
	subject := "录音文件转文本数据发生错误"
	body := fmt.Sprintf("<h2>%s</h2><p>%s<p>", entry.Message, entry.Data["stack"])
	arr := strings.Split(body, "\n")
	body = strings.Join(arr, "</p><p>")
	return hook.Send(hook.Receivers, subject, body)
}

// 触发级别接口方法
func (hook *MailHook) Levels() []log.Level {
	return []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
	}
}

func (hook *MailHook) Send(mailTo []string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From",  m.FormatAddress(hook.User, "Golang App Error")) //这种方式可以添加别名，即“XX官方”
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用m.SetHeader("From",mailConn["user"])
	//m.SetHeader("From", mailConn["user"])
	reg1 := regexp.MustCompile(`(.*?)<(.*?)>`)
	var to []string
	for _, v := range(mailTo){
		res := reg1.FindAllStringSubmatch(v, -1)
		if len(res) > 0{
			to = append(to, m.FormatAddress(res[0][2], res[0][1]))
		}
	}
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	port, _ := strconv.Atoi(hook.Port)
	d := gomail.NewDialer(hook.Host, port, hook.User, hook.Pass)
	err := d.DialAndSend(m)
	return err

}