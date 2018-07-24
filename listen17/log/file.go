package log

import "fmt"

type Filelog struct {

}

func NewFileLog(file string) *Filelog{
	return &Filelog{}
}

func (f *Filelog) LogDebug(msg string) {
	fmt.Println(msg)
}
func (f *Filelog) LogWarn(msg string) {
	fmt.Println(msg)
}