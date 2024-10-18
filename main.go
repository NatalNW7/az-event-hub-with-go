package main

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
	"github.com/joho/godotenv"
)

func main() {
    _ = godotenv.Load()

    credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Error to geenerate credential: %v", err)
	}

    namespace := "<YOUR-NAMESPACE>.servicebus.windows.net"
    eventHub := "<YOUR-EVENTHUB>"
    
    // its possible to use connection string
    // conStr := "Endpoint=sb://<YOUR-NAMESPACE>.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=XXXXXXXXXX"
    // producerClient, err := azeventhubs.NewProducerClientFromConnectionString(conStr, eventHub, nil)
    producerClient, err := azeventhubs.NewProducerClient(namespace, eventHub, credential, nil)
	if err != nil {
		log.Fatalf("Error to create ProducerClient: %v", err)
	}

    defer producerClient.Close(context.TODO())

    events := createEventsForSample()

    newBatchOptions := &azeventhubs.EventDataBatchOptions{}

    batch, err := producerClient.NewEventDataBatch(context.TODO(), newBatchOptions)
    if err != nil {
        log.Fatalf("Error to add events: %v", err)
    }

    for i := 0; i < len(events); i++ {
        _ = batch.AddEventData(events[i], nil)
    }

    err = producerClient.SendEventDataBatch(context.TODO(), batch, nil)
    if err != nil {
        log.Fatalf("Error to send events: %v", err)
    }

    log.Println("Event Sent")
}

func createEventsForSample() []*azeventhubs.EventData {
    return []*azeventhubs.EventData{
        {
            Body: []byte("{ \"message\": \"Hello, World\" }"),
        },
        {
            Body: []byte("{ \"message\": \"Hello, EventHub\" }"),
        },
    }
}