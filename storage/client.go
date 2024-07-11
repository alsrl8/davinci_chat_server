package storage

import (
	"cloud.google.com/go/firestore"
)

type FirestoreClient interface {
	Collection(name string) *firestore.CollectionRef
	Doc(path string) *firestore.DocumentRef
	Close() error
}

type firestoreClientWrapper struct {
	client *firestore.Client
}
