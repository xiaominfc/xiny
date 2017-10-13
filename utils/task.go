package utils

import(
    "time"
)

type Task struct{
    ticker *time.Ticker
    work func()
    waitTime time.Duration
    running bool
}


func (this *Task) Run() {
    defer func(){this.running = false}()
    for _ = range this.ticker.C {
        this.work()

    }
}

func (this *Task) Start() {
    if this.ticker == nil {
        this.ticker = time.NewTicker(this.waitTime)
    }
    if this.running {
        return
    }
    this.running = true
    go this.Run()
}

func AddTask(waitTime time.Duration, work func()) *Task {
    task := &Task{work:work,waitTime:waitTime,running:false}
    task.Start()
    return task
}


