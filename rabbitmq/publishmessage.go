package rabbitmq

import (
	"encoding/json"
	"fmt"
	"payment/util"

	"github.com/streadway/amqp"
)

//function to publish the message
func PublishMessage(conn *amqp.Connection) {

	//now creating the channel for publishing our msg
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	//declaring a queue for our RMQ
	queue, err := ch.QueueDeclare(
		"PaymentQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(queue)

	//getting the msg to publish
	data := util.GetMessages()
	fmt.Println(len(data))

	//publishing msgs over loop
	for _, value := range data {
		fmt.Println(value)
		// marshelling the data before publishing it
		json_data2, err := json.Marshal(value)

		//now publishing our queue
		err = ch.Publish(
			"",
			"PaymentQueue",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(json_data2),
			},
		)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
	fmt.Println("Successfully Published the message")

}
