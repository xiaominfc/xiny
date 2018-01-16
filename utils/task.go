package utils

import(
    "time"
)

type Task struct{
    ticker *time.Ticker
    work func()
    waitTime time.Duration
    running bool
    paused bool
}


func (this *Task) Run() {
    defer func(){this.running = false}()
    for _ = range this.ticker.C {
        if !this.paused {
            this.work()    
        }
    }
}

func (this *Task) Start() {
    if this.ticker == nil {
        this.ticker = time.NewTicker(this.waitTime)
    }
    this.paused = false
    if this.running {
        return
    }
    this.running = true
    go this.Run()
}

func (this *Task) Pause() {
    this.paused = true;
}

func (this *Task) Stop() {
    this.ticker.Stop()
    this.running = false
    this.ticker = nil
    this.work = nil
}

func AddTask(waitTime time.Duration, work func()) *Task {
    task := &Task{work:work,waitTime:waitTime,running:false}
    task.Start()
    return task
}

