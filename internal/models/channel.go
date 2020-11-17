package models

const (
	Success = iota
	Failure
)

type ChanResp struct {
	Status int
	ErrMsg string
	Data   interface{}
}
