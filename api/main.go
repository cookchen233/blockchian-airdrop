package main


import (
	"blockchain-deal-hunter/api/router"
	logrus_stack "github.com/Gurpartap/logrus-stack"
	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"blockchain-deal-hunter/api/utility"
)

var ca *cache.Cache
var ne = cache.NoExpiration
var loc, _ = time.LoadLocation("Asia/Shanghai")
var time_a time.Time

func init() {
	godotenv.Load()

	if os.Getenv("is_debug") == "1" {
		log.SetLevel(log.DebugLevel)
	}else{
		log.SetLevel(log.InfoLevel)
	}

	callerLevels := []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
	}
	stackLevels := []log.Level{log.PanicLevel, log.FatalLevel, log.ErrorLevel}

	log.AddHook(logrus_stack.NewHook(callerLevels, stackLevels))

	log.AddHook(utility.RotateLogHook("log", "stdout.log", 7*24*time.Hour, 24*time.Hour))

	if os.Getenv("environment") == "pro" {
		log.AddHook(&utility.MailHook{
			os.Getenv("mail_user"),
			os.Getenv("mail_pass"),
			os.Getenv("mail_host"),
			os.Getenv("mail_port"),
			strings.Split(os.Getenv("err_receivers"), ";"),
		})
	}

}

func main() {
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()

}