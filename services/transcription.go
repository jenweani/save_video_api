package services

import (
	// "encoding/json"
	// "fmt"
	// "log"
	// "os"
	// "path/filepath"

	// "github.com/openai/openai-go/v2"
	// "github.com/streadway/amqp"
)

type TranscriptionRequest struct {
    VideoURL string `json:"video_url"`
}

type TranscriptionResponse struct {
    Transcription string `json:"transcription"`
}

func GetVidTranscription(filename string){
	// rabbitMQURL := os.Getenv("RABBITMQ_URL")
    // queueName := "transcription_queue"

    // // Initialize RabbitMQ connection
    // conn, err := amqp.Dial(rabbitMQURL)
    // if err != nil {
    //     log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    // }
    // defer conn.Close()

    // ch, err := conn.Channel()
    // if err != nil {
    //     log.Fatalf("Failed to open a channel: %v", err)
    // }
    // defer ch.Close()

    // // Declare the queue
    // _, err = ch.QueueDeclare(
    //     queueName, // Queue name
    //     false,    // Durable
    //     false,    // Delete when unused
    //     false,    // Exclusive
    //     false,    // No-wait
    //     nil,      // Arguments
    // )
    // if err != nil {
    //     log.Fatalf("Failed to declare a queue: %v", err)
    // }

    // apiKey := os.Getenv("OPENAI_API_KEY")
    // client := openai.NewClient(apiKey)

    // // Consume messages from the queue
    // msgs, err := ch.Consume(
    //     queueName, // Queue name
    //     "",        // Consumer
    //     true,      // Auto-Ack
    //     false,     // Exclusive
    //     false,     // No-local
    //     false,     // No-Wait
    //     nil,       // Args
    // )
    // if err != nil {
    //     log.Fatalf("Failed to register a consumer: %v", err)
    // }

    // log.Printf("Waiting for messages. To exit, press CTRL+C")

    // for msg := range msgs {
    //     var request TranscriptionRequest
    //     if err := json.Unmarshal(msg.Body, &request); err != nil {
    //         log.Printf("Failed to decode message body: %v", err)
    //         continue
    //     }

    //     // Perform the transcription using OpenAI Whisper API
    //     response, err := client.Transcriptions.Create(&openai.TranscriptionRequest{
    //         AudioURL: request.VideoURL,
    //         Model:    "whisper-large",
    //     })
    //     if err != nil {
    //         log.Printf("Failed to transcribe audio: %v", err)
    //         continue
    //     }

    //     transcription := response.Transcription.Text

    //     // Save the transcription to a file on the server's disk
	// 	savePath := filepath.Join("./transcripts", fmt.Sprintf("%s.txt", filename))
    //     file, err := os.Create(savePath)
    //     if err != nil {
    //         log.Printf("Failed to create transcription file: %v", err)
    //         continue
    //     }
    //     defer file.Close()

    //     _, err = file.WriteString(transcription)
    //     if err != nil {
    //         log.Printf("Failed to write transcription to file: %v", err)
    //         continue
    //     }
	// }
}