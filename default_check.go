package check

import (
	"log"
)

type Logger interface {
	Printf(format string, v ...interface{})
	Print(v ...interface{})
	Println(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
}

var default_log Logger = log.Default()
var default_logger = func() *Checker {
	c := New(nil)
	c.defaultskip = 2
	return c
}()

// Check 只输出日志. 不panic
func Check(err error) bool {
	return default_logger.Check(err)
}

// CheckSkip // CheckReport 自动找寻到需要定位的函数reportfunc内部位置, 后可以忽略afterMatchSkip层堆栈. 递归函数可用
func CheckReport(err error, reportfunc string, afterMatchSkip int) bool {
	return default_logger.CheckReport(err, reportfunc, afterMatchSkip)
}

// CheckSkip 忽略多少层堆栈 只输出日志
func CheckSkip(err error, skip int) bool {
	return default_logger.CheckSkip(err, skip)
}

// CheckPanic 自动panic
func CheckPanic(err error) {
	default_logger.CheckPanic(err)
}

// CheckPanicReport // CheckReport 自动找寻到需要定位的函数reportfunc内部位置, 后可以忽略afterMatchSkip层堆栈. 递归函数可用
func CheckPanicReport(err error, reportfunc string, afterMatchSkip int) {
	default_logger.CheckPanicReport(err, reportfunc, afterMatchSkip)
}

// CheckPanicSkip 忽略多少层堆栈. 自动panic
func CheckPanicSkip(err error, skip int) {
	default_logger.CheckPanicSkip(err, skip)
}
