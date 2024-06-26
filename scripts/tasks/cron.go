package tasks

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"time"

	"github.com/robfig/cron"
)

// 定时任务
var Cron *cron.Cron

// Run运行任务
func Run(job func() error) {
	//开始时间
	from := time.Now().UnixNano()
	err := job()
	//结束时间
	to := time.Now().UnixNano()
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	if err != nil {
		fmt.Printf("%s err:%dms\n", jobName, (to-from)/int64(time.Millisecond))
	} else {
		fmt.Printf("%s success:%dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}

func CronJob() {
	if Cron == nil {
		Cron = cron.New()
	}
	Cron.AddFunc("0 0 0 * * *", func() { Run(common) })
	Cron.Start()
	log.Println("执行定时任务")
}
