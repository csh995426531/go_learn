package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"os"
	"time"
)

func main(){
	logrus_test()
}

func logrus_test() {

	fmt.Printf("<<<<<<<<<logrus test>>>>>>>>>>>>>>\n")

	logrus.WithFields(logrus.Fields{
		"sb": "sbvalue",
	}).Info("A walrus appears")

	log1 := logrus.New()
	fmt.Printf("log1 level: %d\n", log1.Level)
	log1.Debug("log1 debug")
	log1.Debugf("log1 debug f, %d", 10)
	log1.Info("log1 info")
	log1.Warn("log1 warn")
	log1.Error("log1 error")
	// log1.Panic("log1 panic")

	log1.SetLevel(logrus.ErrorLevel)
	fmt.Printf("after set log1 level to errorlevel\n")
	log1.Debug("log1 debug")

	fmt.Printf("-------------test formater-------------\n")
	log1.SetLevel(logrus.DebugLevel)
	log1.Formatter = &logrus.TextFormatter{
		DisableColors:  true,
		FullTimestamp:  true,
		DisableSorting: true,
	}

	log1.Debug("log text formatter test")

	fmt.Printf("-----------json formatter-------------\n")
	log1.Formatter = &logrus.JSONFormatter{}
	log1.Debug("log json formatter test")

	fmt.Printf("-----------log to file test-----------\n")
	log2 := logrus.New()
	log2.SetLevel(logrus.DebugLevel)
	log2.Formatter = &logrus.TextFormatter{
		DisableColors:  true,
		FullTimestamp:  true,
		DisableSorting: true,
	}

	logger_name := "logrus"
	cur_time := time.Now()
	log_file_name := fmt.Sprintf("%s_%04d-%02d-%02d-%02d-%02d.txt",
		logger_name, cur_time.Year(), cur_time.Month(), cur_time.Day(), cur_time.Hour(), cur_time.Minute())
	log_file, err := os.OpenFile(log_file_name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeExclusive)
	if err != nil {
		fmt.Printf("try create logfile[%s] error[%s]\n", log_file_name, err.Error())
		return
	}

	defer log_file.Close()

	log2.SetOutput(log_file)

	for i := 0; i < 10; i++ {
		log2.Debugf("logrus to file test %d", i)
	}

}