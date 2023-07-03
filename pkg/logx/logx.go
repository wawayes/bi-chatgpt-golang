package logx

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

// logrus自带的格式输出，传字符串格式时需要两个转义字符，且每次手动输入debug.Stack()加转义字符的转换不方便
// 目前只做了stdout，log-pilot会做收集，如需输入到日志文件中还需添加

// Info 常规信息的输出，程序运行正常
func Info(msg string) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("\"%s\" line=\"%d\" time=\"%s\" level=info msg=\"%s\"\n", file, line, time.Now().Format("2006-01-02 15:04:05.999999"), msg)
}

// Warning 警告信息的输出，重要，需要尽快去查看，但不需要立刻终止程序
func Warning(msg string) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("\"%s\" line=\"%d\" time=\"%s\" level=warning msg=\"%s\"\n", file, line, time.Now().Format("2006-01-02 15:04:05.999999"), msg)
}

// Error 发生重大错误，程序无法运行下去，会调用os.Exit()终止程序；
// 对于调用的第三方包，若希望进行异常recover，也在recover后进行调用，以确保打印信息后退出
func Error(msg string) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("\"%s\" line=\"%d\" time=\"%s\" level=error msg=\"%s\"\nstack=\"%s\"\n", file, line, time.Now().Format("2006-01-02 15:04:05.999999"), msg, debug.Stack())
	os.Exit(1)
}
