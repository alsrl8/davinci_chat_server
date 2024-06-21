package storage

import (
	"context"
	firebase "firebase.google.com/go"
	"log"
	"sync"
)

var (
	firestoreClient FirestoreClient
	once            sync.Once
)

func GetFirestoreClient(ctx context.Context) FirestoreClient {
	once.Do(func() {
		app, err := firebase.NewApp(ctx, nil)
		if err != nil {
			log.Fatalf("error initializing firebase app: %v\n", err)
		}

		client, err := app.Firestore(ctx)
		if err != nil {
			log.Fatalf("error getting Firestore client: %v\n", err)
		}

		firestoreClient = &firestoreClientWrapper{client: client}
	})
	return firestoreClient
}
