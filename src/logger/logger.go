package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

const logFileName = "log.txt"

type LogMessage struct {
	Method string
	Host string
	UrlPath string
	Query string
	Result string
}



func (m *LogMessage) String() string {
	logMessage := fmt.Sprintf("method: %v, host: %v, path: %v, query: %v => %s",
		m.Method,
		m.Host,
		m.UrlPath,
		m.Query,
		m.Result,
	)

	return logMessage

}



func (m LogMessage) WriteToLog(resultText string) {
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if len(resultText) == 0 {
		m.Result = "success"
	} else {
		m.Result = resultText
	}


	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)

	log.Println(m.String())
}
