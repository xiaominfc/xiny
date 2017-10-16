package ylog

import "log"

func ILog(message ...interface{}) {
    log.Println(message)
}

func DLog(message ...interface{}) {
   log.Println(message)
}
