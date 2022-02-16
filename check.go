package check

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

type _caller struct {
	PC       uintptr
	Skip     int
	Line     int
	File     string
	FuncName string
}

type Checker struct {
	logger        Logger
	reportSkipMap sync.Map
	defaultskip   int
}

func New(logger Logger) *Checker {
	if logger == nil {
		return &Checker{logger: default_log, defaultskip: 1}
	}
	return &Checker{logger: logger, defaultskip: 1}
}

func (check *Checker) SetDefaultSkip(skip int) {
	check.defaultskip = skip
}

// CheckReport 自动找寻到需要定位的函数reportfunc内部位置, 后可以忽略afterMatchSkip层堆栈. 递归函数可用
func (check *Checker) CheckReport(err error, reportfunc string, afterMatchSkip int) bool {

	if err != nil {
		var caller *_caller
		v, ok := check.reportSkipMap.Load(reportfunc)
		if !ok {

			caller = &_caller{}
			for i := 0; i < 100; i++ {
				caller.PC, caller.File, caller.Line, _ = runtime.Caller(i)
				caller.FuncName = runtime.FuncForPC(caller.PC).Name()
				if regexp.MustCompile(reportfunc).MatchString(caller.FuncName[strings.LastIndexByte(caller.FuncName, '.')+1:]) {
					check.reportSkipMap.Store(reportfunc, caller)
					caller.Skip = i + afterMatchSkip
					if afterMatchSkip != 0 {
						caller.PC, caller.File, caller.Line, _ = runtime.Caller(caller.Skip)
						caller.FuncName = runtime.FuncForPC(caller.PC).Name()
					}
					break
				}
			}
		} else {
			caller = v.(*_caller)
		}

		check.logger.Println(fmt.Sprintf("%s:%d\n%s", caller.File, caller.Line, err))
		return true
	}
	return false
}

// Check 不主动panic
func (check *Checker) Check(err error) bool {
	if err != nil {
		_, file, line, _ := runtime.Caller(check.defaultskip)
		check.logger.Println(fmt.Sprintf("%s:%d\n%s", file, line, err))
		return true
	}
	return false
}

func (check *Checker) CheckSkip(err error, skip int) bool {
	if err != nil {
		_, file, line, _ := runtime.Caller(check.defaultskip + skip)
		check.logger.Fatalln(fmt.Sprintf("%s:%d\n%s", file, line, err))
		return true
	}
	return false
}

func (check *Checker) CheckPanic(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(check.defaultskip)
		check.logger.Fatalln(fmt.Sprintf("%s:%d\n%s", file, line, err))
	}
}

// CheckPanicReport 忽略多少层堆栈. 不主动panic
func (check *Checker) CheckPanicReport(err error, reportfunc string, afterMatchSkip int) {

	if err != nil {
		var caller *_caller
		v, ok := check.reportSkipMap.Load(reportfunc)
		if !ok {

			caller = &_caller{}
			for i := 0; i < 100; i++ {
				caller.PC, caller.File, caller.Line, _ = runtime.Caller(i)
				caller.FuncName = runtime.FuncForPC(caller.PC).Name()
				if regexp.MustCompile(reportfunc).MatchString(caller.FuncName[strings.LastIndexByte(caller.FuncName, '.')+1:]) {
					check.reportSkipMap.Store(reportfunc, caller)
					caller.Skip = i + afterMatchSkip
					if afterMatchSkip != 0 {
						caller.PC, caller.File, caller.Line, _ = runtime.Caller(caller.Skip)
						caller.FuncName = runtime.FuncForPC(caller.PC).Name()
					}
					break
				}
			}
		} else {
			caller = v.(*_caller)
		}

		check.logger.Fatalln(fmt.Sprintf("%s:%d\n%s", caller.File, caller.Line, err))
		return
	}
	return
}

// CheckSkip 忽略多少层堆栈

func (check *Checker) CheckPanicSkip(err error, skip int) {
	if err != nil {
		_, file, line, _ := runtime.Caller(check.defaultskip + skip)
		check.logger.Fatalln(fmt.Sprintf("%s:%d\n%s", file, line, err))
		return
	}
	return
}
