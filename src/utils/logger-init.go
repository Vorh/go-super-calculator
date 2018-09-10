package utils

import (
	"src/github.com/op/go-logging"
)


var log = logging.MustGetLogger("test")


var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func InitLogger()  {
	logging.SetFormatter( format)
}

func GetLogger() *logging.Logger {
	return log
}
