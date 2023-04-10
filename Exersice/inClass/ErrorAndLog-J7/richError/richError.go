package richError

import "time"
//import "fmt"

//import "fmt"

type RichError struct{
	Message string
	Operation string
	MetaData map[string]string
	Time time.Time
}

func (r RichError) Error() string{
	return r.Message
}

// func (r RichError) String() string{
// 	return fmt.Sprintf("{Message is: %s, Operation is: %s, MetaData is: %+v\n",
// 			r.Message, r.Operation, r.MetaData)
// }

