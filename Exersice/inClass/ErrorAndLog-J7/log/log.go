package log

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"ErrorAndLog/richError"
	"ErrorAndLog/simpleError"
)

type Log struct{
	Errors []richError.RichError
}

// func (l Log) Print(){
// 	for _, v := range l.Errors{
// 		fmt.Println(v)
// 	}
// }

func (l Log) Save() error {
	var fileHandler *os.File

	if f, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666); err != nil{
		return fmt.Errorf("error occured: %w", err)
	} else {
		fileHandler = f
		defer fileHandler.Close()
	}

	data, _ := json.Marshal(l.Errors)
	fileHandler.Write(data)

	return nil
}

func (l *Log) Add(r error) {
	if err, ok := r.(richError.RichError); ok{
		l.Errors = append(l.Errors, err)
	} else if err, ok := r.(simpleError.SimpleError); ok{
		l.Errors = append(l.Errors, richError.RichError{
										Message: err.Message,
										Operation: err.Operation,
										MetaData: nil,
										Time: time.Now(),
		})
	} else {
		l.Errors = append(l.Errors, richError.RichError{
										Message: r.Error(),
										Operation: "unkown",
										MetaData: nil,
										Time: time.Now(),
		})
	}
	
}
