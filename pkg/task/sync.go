package task

import (
	"os"
	"os/signal"
	"runtime"
	"time"
)

//同步执行任务
type RunnerAsync struct {
	//操作系统的信号检测
	interrupt chan os.Signal
	//记录执行完成的状态
	complete chan error
	//超时检测
	timeout <-chan time.Time
	//保存所有要执行的任务,顺序执行
	tasks []func(id int)
}

//new一个RunnerAsync对象
func NewRunnerAsync(d time.Duration) *RunnerAsync {
	return &RunnerAsync{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//添加一个任务
func (r *RunnerAsync) Add(tasks ...func(id int)) {
	r.tasks = append(r.tasks, tasks...)
}

//启动RunnerAsync，监听错误信息
func (r *RunnerAsync) Start() error {
	//开启多核
	runtime.GOMAXPROCS(runtime.NumCPU())
	//接收操作系统信号
	signal.Notify(r.interrupt, os.Interrupt)
	//执行任务
	go func() {
		r.complete <- r.Run()
	}()

	select {
	//返回执行结果
	case err := <-r.complete:
		return err
		//超时返回
	case <-r.timeout:
		return ErrTimeout
	}
}

//顺序执行所有的任务
func (r *RunnerAsync) Run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		//执行任务
		task(id)
	}
	return nil
}

//判断是否接收到操作系统中断信号
func (r *RunnerAsync) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		//停止接收别的信号
		signal.Stop(r.interrupt)
		return true
		//正常执行
	default:
		return false
	}
}
