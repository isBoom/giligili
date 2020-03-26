package task

import (
	"fmt"
	"github.com/robfig/cron"
	"reflect"
	"runtime"
	"time"
)

var Cron *cron.Cron

func Run(job func() error) {
	from := time.Now().UnixNano()
	err := job()
	to := time.Now().UnixNano()
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	if err != nil {
		fmt.Printf("%s error: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	} else {
		fmt.Printf("%s success: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}
func CronJob() {
	if Cron == nil {
		Cron = cron.New()
	}
	Cron.AddFunc("0 0 0 * * *", func() {
		Run(RestarDailyRank)
	})

	Cron.Start()
	fmt.Println("Cron start...")
}
