package richError

import "time"
import "fmt"


type RichError struct{
	Message string
	Operation string
	MetaData map[string]string
	Time time.Time
}

func (r RichError) String() string{
	return fmt.Sprintf("{Message is: %s, Operation is: %s, MetaData is: %+v}",
			r.Message, r.Operation, r.MetaData)
}

func (r RichError) Error() string{
	return r.Message
}


