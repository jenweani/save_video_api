package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/sashabaranov/go-openai"
	"github.com/streadway/amqp"
)

type TranscriptRequest struct {
    VideoURL string `json:"video_url"`
}

type TranscriptResponse struct {
    Transcription string `json:"transcription"`
}

func GetVidTranscription(filename string){
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
    queueName := "transcription_queue"

    // Initialize RabbitMQ connection
    conn, err := amqp.Dial(rabbitMQURL)
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }
    defer ch.Close()

    // Declare the queue
    _, err = ch.QueueDeclare(
        queueName, // Queue name
        false,    // Durable
        false,    // Delete when unused
        false,    // Exclusive
        false,    // No-wait
        nil,      // Arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare a queue: %v", err)
    }

    // Initialise Openai connection
    apiKey := os.Getenv("OPENAI_API_KEY")
    client := openai.NewClient(apiKey)
    ctx := context.Background()
    filePath := filepath.Join("./uploads", filename)
    req := openai.AudioRequest{
        Model: openai.Whisper1,
        FilePath: filePath,
    }

    // Consume messages from the queue
    msgs, err := ch.Consume(
        queueName, // Queue name
        "",        // Consumer
        true,      // Auto-Ack
        false,     // Exclusive
        false,     // No-local
        false,     // No-Wait
        nil,       // Args
    )
    if err != nil {
        log.Fatalf("Failed to register a consumer: %v", err)
    }

    log.Printf("Waiting for messages. To exit, press CTRL+C")

    for msg := range msgs {
        var request TranscriptRequest
        if err := json.Unmarshal(msg.Body, &request); err != nil {
            log.Printf("Failed to decode message body: %v", err)
            continue
        }

        // Perform the transcription using OpenAI Whisper API
        response, err := client.CreateTranscription(ctx, req)
        if err != nil {
            log.Printf("Failed to transcribe video: %v", err)
            continue
        }

        transcription := response.Text

        // Save the transcription to a file on the server's disk
		savePath := filepath.Join("./transcripts", fmt.Sprintf("%s.txt", filename))
        file, err := os.Create(savePath)
        if err != nil {
            log.Printf("Failed to create transcription file: %v", err)
            continue
        }
        file.Close()

        _, err = file.WriteString(transcription)
        if err != nil {
            log.Printf("Failed to write transcription to file: %v", err)
            continue
        }
	}
}