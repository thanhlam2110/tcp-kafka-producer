package service

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/thanhlam/tcp-kafka-producer/model"
)

func PushKafkaMessage(c echo.Context) error {
	pushKafka := new(model.PushKafka)
	err := c.Bind(pushKafka)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, map[string]interface{}{"code": "6", "message": "Body is Invalid"})
	}
	message := pushKafka.Message
	topic := pushKafka.Topic
	status := TcpPushMessage(message, topic)
	if status != "200 OK" {
		return c.JSON(200, map[string]interface{}{"code": "1", "message": "ERROR PUSH MESSAGE TO KAFKA"})
	} else {
		return c.JSON(200, map[string]interface{}{"code": "0", "message": "SUCCESS PUSH MESSAGE TO KAFKA"})
	}
}
func TcpPushMessage(message, topic string) string {
	url := "http://18.139.3.218:19092/topics/" + topic + "/messages"
	var query = []byte(message)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(query))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//fmt.Println("response Status:", resp.Status)
	return resp.Status
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
}
