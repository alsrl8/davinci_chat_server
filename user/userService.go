package user

import (
	"context"
	"davinci-chat/auth"
	"davinci-chat/storage"
	auth2 "firebase.google.com/go/auth"
	"sync"
)

type Service interface {
	AddUser(ctx context.Context, user *User, email string) error
	GetUser(ctx context.Context, email string) (*User, error)
	Close() error
}

type userService struct {
	client storage.FirestoreClient
}

func (us *userService) AddUser(ctx context.Context, user *User, email string) error {
	doc := us.client.Collection("users").Doc(email)
	userPublic := UserPublic{
		Name: user.Name,
	}
	if _, err := doc.Set(ctx, userPublic); err != nil {
		return err
	}

	create := &(auth2.UserToCreate{})
	userToCreate := create.Email(email).Password(user.Password)
	_, err := auth.FirebaseAuth.CreateUser(context.Background(), userToCreate)
	if err != nil {
		return err
	}

	return nil
}

func (us *userService) GetUser(ctx context.Context, email string) (*User, error) {
	doc, err := us.client.Collection("users").Doc(email).Get(ctx)
	if err != nil {
		return nil, err
	}

	var user User
	err = doc.DataTo(&user)
	return &user, err
}

func (us *userService) Close() error {
	return us.client.Close()
}

var (
	userServiceInstance Service
	userServiceOnce     sync.Once
)

func NewUserService(ctx context.Context) Service {
	userServiceOnce.Do(func() {
		client := storage.GetFirestoreClient(ctx)
		userServiceInstance = &userService{client: client}
	})
	return userServiceInstance
}
