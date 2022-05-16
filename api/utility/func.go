package utility

//一些工具杂项函数

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime/debug"
	"sync/atomic"
)

//列出目录内容
func ListDirName(path string) []string {
	var files []string
	filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
		files = append(files, file)
		return nil
	})
	return files

}

func Pr(args ...interface{}) {
	fmt.Println(args...)
}

//生成整型范围切片
func RangeInt(min, max, step int) []int {
	var a []int
	if step < 0{
		for i := min; i > max; i+= step {
			a = append(a, i)
		}
	}
	for i := min; i < max; i+= step {
		a = append(a, i)
	}
	return a
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//生成 md5字符串
func Md5Str(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}


//带数量限制的协程, 方便直接控制 go 方法执行的数量
type Golimit struct {
	Num int
	ch chan struct{}
}

func NewGolimit(num int) Golimit {
	return Golimit{
		Num: num,
		ch: make(chan struct{}, num),
	}
}

func (gl *Golimit) Run(fn interface{}, args ...interface{}) {
	gl.ch <- struct{}{}
	go func() {
		defer func(){<-gl.ch}()
		v := reflect.ValueOf(fn)
		rargs := make([]reflect.Value, len(args))
		for i, arg := range args {
			rargs[i] = reflect.ValueOf(arg)
		}
		v.Call(rargs)
	}()
}

//defer 的封装
func SafeDefer(params ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			msg := fmt.Sprintf("%+v", r)
			stack := string(debug.Stack())
			log.Error(fmt.Sprintf("recovery from panic:\n%s\n%s", msg, stack), true)
			return
		}
	}()

	r := recover()
	if r == nil {
		return
	}

	err := fmt.Errorf("%+v", r)
	stack := string(debug.Stack())
	log.Error(fmt.Sprintf("recovery from panic:\n%s\n%s", err.Error(), stack), true)

	if paramsLen := len(params); paramsLen > 0 {
		if reflect.TypeOf(params[0]).String()[0:4] != "func" {
			return
		}
		var args []reflect.Value
		if paramsLen > 1 {
			args = append(args, reflect.ValueOf(err))
			for _, v := range params[1:] {
				args = append(args, reflect.ValueOf(v))
			}
		}
		reflect.ValueOf(params[0]).Call(args)
	}
}

//安全的协程方法, 可对子协程的异常捕获
func deprecated_SafeGo(params ...interface{}) {
	if len(params) == 0 {
		return
	}

	pg := &panicGroup{panics: make(chan string, 1), dones: make(chan struct{}, 1)}
	defer pg.closeChan()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				pg.panics <- fmt.Sprintf("%+v\n%s", r, string(debug.Stack()))
				return
			}
			pg.dones <- struct{}{}
		}()
		var args []reflect.Value
		if len(params) > 1 {
			for _, v := range params[1:] {
				args = append(args, reflect.ValueOf(v))
			}
		}
		reflect.ValueOf(params[0]).Call(args)
	}()

	for {
		select {
		case <-pg.dones:
			return
		case p := <-pg.panics:
			panic(p)
		}
	}
}

type SafeGo interface {
	Go(...interface{}) *panicGroup
	Wait()
}

//安全的协程方法
//limit 并发协程数量限制
func NewSafeGo(limit int) SafeGo {
	p := &panicGroup{
		panics: make(chan string, 1),
		dones:  make(chan struct{}, limit),
		limit:  make(chan struct{}, limit),
	}
	p.Go()
	return p
}

type panicGroup struct {
	panics chan string // 协程 panic 通知通道
	dones  chan struct{}    // 协程完成通知通道
	jobN   int32       // 协程计数
	limit   chan struct{}       //限制器
}

//协程运行
//params 第一个为执行方法名, 后续为方法参数
func (g *panicGroup) Go(params ...interface{}) *panicGroup {
	if len(params) == 0 {
		params = []interface{}{func() {}}
	}
	atomic.AddInt32(&g.jobN, 1)
	go func() {
		defer func() {
			<- g.limit
			if r := recover(); r != nil {
				func(){
					defer func(){
						if r := recover(); r != nil {}
					}()
					g.panics <- fmt.Sprintf("%+v\n%s", r, string(debug.Stack()))
				}()
			}

			func(){
				defer func(){
					if r := recover(); r != nil {}
				}()
				g.dones <- struct{}{}
			}()
		}()
		g.limit <- struct{}{}
		var args []reflect.Value
		if len(params) > 1 {
			for _, v := range params[1:] {
				args = append(args, reflect.ValueOf(v))
			}
		}
		reflect.ValueOf(params[0]).Call(args)
	}()
	return g
}

//等待子协程执行完毕, 与WaitGroup类似
func (g *panicGroup) Wait() {
	defer g.closeChan()
	for {
		select {
		case <-g.dones:
			if atomic.AddInt32(&g.jobN, -1) == 0 {
				return
			}
		case p := <-g.panics:
			panic(p)
		}
	}
}

func (g *panicGroup) closeChan() {
	close(g.dones)
	close(g.panics)
}


func IsFile(filename string) bool {
	file, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !file.IsDir()
}

func IsDir(filename string) bool {
	file, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return file.IsDir()
}

//将字节大小转换为人类可读单位(M, G, T)
func ByteCountBinary(size int64) string {
	const unit int64 = 1024
	if size < unit {
		return fmt.Sprintf("%dB", size)
	}
	div, exp := unit, 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB", float64(size)/float64(div), "KMGTPE"[exp])
}

func FileSize(filename string) int64{
	file, err:=os.Stat(filename)
	if err != nil {
		return 0
	}
	return file.Size()
}

//仅提取文件名与其上层目录的拼接路径
func PartFilename(filename string) (string){
	return path.Join(path.Base(path.Dir(filename)), path.Base(filename))
}
