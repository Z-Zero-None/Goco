package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

type Level int8                    //日志等级
type Fields map[string]interface{} //日志的公共字段
//定义日志等级
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

//日志等级字符串化
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

//日志标准化
type Logger struct {
	newLogger *log.Logger
	ctx       context.Context //上下文
	fields    Fields          //日志公共字段
	callers   []string        //调用栈信息
}

//初始化日志对象
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

//克隆
func (l *Logger) clone() *Logger {
	//复制属性
	nl := *l
	return &nl
}

//设置日志公共字段
func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}
	for k, v := range f {
		ll.fields[k] = v
	}
	return ll
}

//设置日志上下文属性
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

//设置当前某一层的调用栈信息
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	//实参skip为上溯的栈帧数，0表示Caller的调用者
	//Caller报告当前go程调用栈所执行的函数的文件和行号信息。
	//函数的返回值为调用栈标识符、文件名、该调用在文件中的行号。如果无法获得信息，ok会被设为false。
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		//FuncForPC返回一个表示调用栈标识符pc对应的调用栈的*Func
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s:%d %s", file, line, f.Name())}
	}
	return ll
}

//设置当前整个调用栈的信息
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	var callers []string
	pcs := make([]uintptr, maxCallerDepth)
	//函数把当前go程调用栈上的调用栈标识符填入切片pcs中，返回写入到pc中的项数。
	//实参skip为开始在pcs中记录之前所要跳过的栈帧数，0表示Callers自身的调用栈，1表示Callers所在的调用栈。返回写入p的项数。
	depth := runtime.Callers(minCallerDepth, pcs)

	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}

//添加链路追踪
func (l *Logger) WithTrace() *Logger {
	ginCtx, ok := l.ctx.(*gin.Context)
	if ok {
		//将traceId和spanId添加到日志flied中
		return l.WithFields(Fields{
			"trace_id": ginCtx.MustGet("X-Trace-ID"),
			"span_id":  ginCtx.MustGet("X-Span-ID"),
		})
	}
	return l
}

//编写日志内容格式化
func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	//创建返回信息数据集合
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	//UnixNano将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位纳秒）。
	//如果纳秒为单位的unix时间超出了int64能表示的范围，结果是未定义的。注意这就意味着Time零值调用UnixNano方法的话，结果是未定义的。
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}

//日志输出的动作
func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JSONFormat(level, message))
	context := string(body)
	switch level {
	case LevelDebug:
		l.newLogger.Print(context)
	case LevelInfo:
		l.newLogger.Print(context)
	case LevelWarn:
		l.newLogger.Print(context)
	case LevelError:
		l.newLogger.Print(context)
	case LevelFatal:
		l.newLogger.Fatal(context)
	case LevelPanic:
		l.newLogger.Panic(context)
	}
}

//针对各个等级进行输出格式划分
func (l *Logger) Panic(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelPanic, fmt.Sprint(v...))
}
func (l *Logger) PanicF(format string, ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelPanic, fmt.Sprintf(format, v...))
}
func (l *Logger) Debug(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelDebug, fmt.Sprint(v...))
}
func (l *Logger) DebugF(format string, ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelDebug, fmt.Sprintf(format, v...))
}
func (l *Logger) Fatal(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelFatal, fmt.Sprint(v...))
}
func (l *Logger) FatalF(format string, ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelFatal, fmt.Sprintf(format, v...))
}
func (l *Logger) Error(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelError, fmt.Sprint(v...))
}
func (l *Logger) ErrorF(format string, ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelError, fmt.Sprintf(format, v...))
}
func (l *Logger) Warn(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelWarn, fmt.Sprint(v...))
}
func (l *Logger) WarnF(format string, ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelWarn, fmt.Sprintf(format, v...))
}
func (l *Logger) Info(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelInfo, fmt.Sprint(v...))
}
func (l *Logger) InfoF(format string, ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(LevelInfo, fmt.Sprintf(format, v...))
}
