package model

type PushKafka struct {
	Message string `json:"message"`
	Topic   string `json:"topic"`
}
