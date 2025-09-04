package main

import (
	"context"
	"fmt"
	"os"

	"github.com/danhenke/env"
	infisical "github.com/infisical/go-sdk"
)

const CLIENT_ID_ENV = "CLIENT_ID"
const CLIENT_SECRET_ENV = "CLIENT_SECRET"
const SECRET_KEY_ENV = "SECRET_KEY"
const SECRET_PATH_ENV = "SECRET_PATH"
const SECRET_ENVIRONMENT_ENV = "SECRET_ENVIRONMENT"
const SECRET_PROJECT_ID_ENV = "SECRET_PROJECT_ID"

func main() {
	ctx := context.Background()

	client := infisical.NewInfisicalClient(ctx, infisical.Config{})

	_, err := client.Auth().UniversalAuthLogin(
		env.MustGet(CLIENT_ID_ENV),
		env.MustGet(CLIENT_SECRET_ENV),
	)

	if err != nil {
		fmt.Printf("Authentication failed: %v", err)
		os.Exit(1)
	}

	secret, err := client.Secrets().Retrieve(infisical.RetrieveSecretOptions{
		SecretKey:   env.MustGet(SECRET_KEY_ENV),
		Environment: env.MustGet(SECRET_ENVIRONMENT_ENV),
		ProjectID:   env.MustGet(SECRET_PROJECT_ID_ENV),
		SecretPath:  env.MustGet(SECRET_PATH_ENV),
	})

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n", secret.SecretValue)
}
