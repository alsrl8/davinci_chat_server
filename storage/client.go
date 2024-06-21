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

func (c *firestoreClientWrapper) Collection(name string) *firestore.CollectionRef {
	return c.client.Collection(name)
}

func (c *firestoreClientWrapper) Doc(path string) *firestore.DocumentRef {
	return c.client.Doc(path)
}

func (c *firestoreClientWrapper) Close() error {
	return c.client.Close()
}
