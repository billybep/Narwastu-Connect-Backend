package auth

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var FirebaseAuth *auth.Client

func InitFirebase() {
	ctx := context.Background()
	var app *firebase.App
	var err error

	// Cek env variable dulu (untuk Railway)
	if jsonCred := os.Getenv("FIREBASE_SERVICE_ACCOUNT_JSON"); jsonCred != "" {
		opt := option.WithCredentialsJSON([]byte(jsonCred))
		app, err = firebase.NewApp(ctx, nil, opt)
		if err != nil {
			panic(fmt.Sprintf("failed to init firebase app from env: %v", err))
		}
	} else {
		// fallback ke file lokal
		credFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
		if credFile == "" {
			credFile = "config/firebase_service_account.json"
		}
		opt := option.WithCredentialsFile(credFile)
		app, err = firebase.NewApp(ctx, nil, opt)
		if err != nil {
			panic(fmt.Sprintf("failed to init firebase app from file: %v", err))
		}
	}

	client, err := app.Auth(ctx)
	if err != nil {
		panic(fmt.Sprintf("failed to init firebase auth: %v", err))
	}

	FirebaseAuth = client
	fmt.Println("🔥 Firebase Auth initialized")
}
