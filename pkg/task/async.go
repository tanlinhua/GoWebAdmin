package task

import (
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

// 异步执行任务
type Runner struct {
	//操作系统的信号检测
	interrupt chan os.Signal

	//记录执行完成的状态
	complete chan error

	//超时检测
	timeout <-chan time.Time

	//保存所有要执行的任务,顺序执行
	tasks []func(id int) error

	waitGroup sync.WaitGroup

	lock sync.Mutex

	errs []error
}

//new一个Runner对象
func NewRunner(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
		waitGroup: sync.WaitGroup{},
		lock:      sync.Mutex{},
	}
}

//添加一个任务
func (r *Runner) Add(tasks ...func(id int) error) {
	r.tasks = append(r.tasks, tasks...)
}

//启动Runner，监听错误信息
func (r *Runner) Start() error {
	//开启多核心
	runtime.GOMAXPROCS(runtime.NumCPU())
	//接收操作系统信号
	signal.Notify(r.interrupt, os.Interrupt)
	//并发执行任务
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

//异步执行所有的任务
func (r *Runner) Run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		gTask := task // fix

		r.waitGroup.Add(1)
		go func(id int) {
			r.lock.Lock()

			//执行任务
			err := gTask(id)
			//加锁保存到结果集中
			r.errs = append(r.errs, err)

			r.lock.Unlock()
			r.waitGroup.Done()
		}(id)
	}
	r.waitGroup.Wait()

	return nil
}

//判断是否接收到操作系统中断信号
func (r *Runner) gotInterrupt() bool {
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

//获取执行完的error
func (r *Runner) GetErrs() []error {
	return r.errs
}
