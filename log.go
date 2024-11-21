package mi_biblioteca

import (
	"flag"
	"fmt"
	"log"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

var logFile *string

func InitializeLog() {
	logFile = flag.String("logFile", "automation-paymet-", "tracer log")
	if *logFile != "" {
		// open output file
		date := time.Now().Format("2006-01-02")
		elLog, err := rotatelogs.New(*logFile + date + ".log")
		if err != nil {
			log.Fatalf("Could not open Log file:%v with err:%v", *logFile, err)
		}
		fmt.Println("Log File with format:" + *logFile + "YYYY-MM-DD.log")
		log.SetOutput(elLog)
		//defer elLog.Close()  algo
	}
	log.Printf("\n\nTime now for the program is %v\n", time.Now())
}
