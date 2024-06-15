package auth

import (
	"context"
	firebase "firebase.google.com/go"
	"log"

	"firebase.google.com/go/auth"
)

var FirebaseAuth *auth.Client

func InitFirebase() {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	FirebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
}
