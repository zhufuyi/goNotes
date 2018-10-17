package main

import (
	"fmt"
	"sanwenyu/liveshow/apps/riskControl/stat"
	"sanwenyu/liveshow/logger"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	NextRunSeconds = 86400 // 一天的秒，可以改为
)

// 任务处理函数类型
type TaskFunc func() error

type Tasker struct {
	Status        int       // 0表示任务还没执行，1表示任务执行成功，-1或其他值表示任务执行失败
	IsRepeat      bool      // true表示每天重复(默认)，false表示不重复
	TaskTimeStamp int64     // 执行任务的时间戳
	TaskName      string    // 任务名称
	F             TaskFunc  // 任务处理
	StopChan      chan bool // 退出任务通道
}

// 新建任务
func NewTask(dt string, f TaskFunc, isRepeat bool) (*Tasker, error) {
	task := &Tasker{
		IsRepeat: isRepeat,
		StopChan: make(chan bool, 1),
		F:        f,
	}

	timeStamp, err := getTimeStamp(dt)
	if err != nil {
		return nil, err
	}
	task.TaskTimeStamp = timeStamp
	return task, nil
}

// 参数输入格式："2016-01-01 01:01:01"、"01:01:01"默认为当天日期、"2016-01-01"默认为凌晨零点
func getTimeStamp(dt string) (int64, error) {
	dtLen := len(dt)
	var timeStamp int64
	if dtLen >= 5 && dtLen <= 8 {
		// 只有时间没有日期格式
		tm := strings.Split(dt, ":")
		if len(tm) != 3 {
			return 0, errors.New("input time format error, eg: 01:01:01")
		}
		hour, _ := strconv.Atoi(tm[0])
		minute, _ := strconv.Atoi(tm[1])
		second, _ := strconv.Atoi(tm[2])
		now := time.Now()
		timeStamp = time.Date(now.Year(), now.Month(), now.Day(), hour, minute, second, 0, time.Local).Unix()
	} else if dtLen == 10 {
		// 只有日期格式
		tm := strings.Split(dt, "-")
		if len(tm) != 3 {
			return 0, errors.New("input time format error, eg: 2016-01-01")
		}
		year, _ := strconv.Atoi(tm[0])
		month, _ := strconv.Atoi(tm[1])
		day, _ := strconv.Atoi(tm[2])
		timeStamp = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Unix()
	} else if dtLen >= 15 && dtLen <= 19 {
		// 有日期和时间格式
		tm, err := time.Parse("2006-01-02 15:04:05", dt)
		if err != nil {
			return 0, errors.New("input time format error, eg: 2016-01-01 01:01:01")
		}
		timeStamp = tm.Unix()
	} else {
		return 0, errors.New("input time format error, eg: 2016-01-01 01:01:01")
	}
	return timeStamp, nil
}

// 开始任务
func (t *Tasker) Start() {
	tick := time.Tick(time.Second) // 1秒定时器
	tickCount := time.Now().Unix() // 秒计算器
	// 判断闹钟是否过时，过时则改为明天的闹钟
	if tickCount > t.TaskTimeStamp {
		t.TaskTimeStamp += stat.DaySeconds
	}
	stat.NextTaskStamp(t.TaskTimeStamp) // 下一个执行任务时间点

	logger.Info("Timing ......, waiting for performing the task at ", time.Unix(t.TaskTimeStamp, 0).Format("2006-01-02 15:04:05"))

LOOP:
	for {
		select {
		case <-tick:
			tickCount++
			// 判断执行任务时间是否已经来到
			if tickCount >= t.TaskTimeStamp {
				err := t.F()
				// 记录执行状态
				if err != nil {
					//					println(err.Error())	// 记录执行失败原因
					t.Status = -1
				} else {
					t.Status = 1
				}
				println(t.getString())

				// 判断是否每天重复自行
				if t.IsRepeat {
					t.TaskTimeStamp += NextRunSeconds // 下一个执行的任务时间点
					logger.Info("Timing ......, waiting for performing the task at ", time.Unix(t.TaskTimeStamp, 0).Format("2006-01-02 15:04:05"))
				} else {
					logger.Info("alarm exit.")
					break LOOP
				}
			}

		case <-t.StopChan:
			break LOOP
		}
	}
}

// 获取任务状态的字符串形式
func (t *Tasker) getString() string {
	dt := time.Unix(t.TaskTimeStamp, 0).Format("2006-01-02 15:04:05")

	st := ""
	switch t.Status {
	case 0:
		st = "waiting for execution"
	case 1:
		st = "success"
	default:
		st = "failed"
	}

	repeat := ""
	if t.IsRepeat {
		repeat = "yes"
	} else {
		repeat = "no"
	}

	return fmt.Sprintf("execute time: %s,  status: %s,  repeat: %s", dt, st, repeat)
}

// 停止任务
func (t *Tasker) Stop() {
	t.StopChan <- true
}

// 处理的任务
func handleTast1() error {
	println("(tast1) output time now =", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

// 处理的任务
func handleTast2() error {
	println("(tast2) output time now =", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

func main() {
	task1, err := NewTask("21:37:00", handleTast1, false)
	if err != nil {
		println(err.Error())
		return
	}
	go task1.Start()

	task2, err := NewTask("21:37:30", handleTast2, true)
	if err != nil {
		println(err.Error())
		return
	}
	go task2.Start()

	for {
		time.Sleep(3600)
	}
}
