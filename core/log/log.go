package log

import (
	"fmt"
	"os"
	"time"
)

var log_path string
var log_stat bool
var silent bool

func Set_config(new_log string, new_log_stat bool, new_silent bool) {
	log_path = new_log
	log_stat = new_log_stat
	silent = new_silent

	// check if log file exsits if not create one
	if _, err := os.Stat(log_path); err != nil {
		file, err := os.Create(log_path)
		if err != nil {
			fmt.Println("[error] couldn't create log file due to " + err.Error())
			return
		}
		defer file.Close()
	}
}

func Error(msg string) {
	if log_stat {
		f, err := os.OpenFile(log_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("[error] can't open log file due to " + err.Error())
			return
		}
		defer f.Close()

		date := time.Now().Format("2006-01-02 15:04:05")
		_, err = f.WriteString("ERROR [" + date + "] ~ " + msg + "\n")

		if !silent {
			fmt.Println(msg)
		}

		if err != nil {
			fmt.Println("[error] can't write log file due to " + err.Error())
		}
	}
}

func Info(msg string) {
	if log_stat {
		f, err := os.OpenFile(log_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("[error] can't open log file due to " + err.Error())
			return
		}
		defer f.Close()

		date := time.Now().Format("2006-01-02 15:04:05")
		_, err = f.WriteString("INFO [" + date + "] ~ " + msg + "\n")

		if !silent {
			fmt.Println(msg)
		}

		if err != nil {
			fmt.Println("[error] can't write log file due to " + err.Error())
		}
	}
}

func Warning(msg string) {
	if log_stat {
		f, err := os.OpenFile(log_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("[error] can't open log file due to " + err.Error())
			return
		}
		defer f.Close()

		date := time.Now().Format("2006-01-02 15:04:05")
		_, err = f.WriteString("WARNING [" + date + "] ~ " + msg + "\n")

		if !silent {
			fmt.Println(msg)
		}

		if err != nil {
			fmt.Println("[error] can't write log file due to " + err.Error())
		}
	}
}
