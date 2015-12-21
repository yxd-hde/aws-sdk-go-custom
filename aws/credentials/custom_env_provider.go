package credentials

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var (
	ErrAccessKeyIDNotFound = func(name string) awserr.Error {
		return awserr.New("EnvAccessKeyNotFound",
			fmt.Sprintf("%s not found in environment", name), nil)
	}

	ErrSecretAccessKeyNotFound = func(name string) awserr.Error {
		return awserr.New("EnvSecretNotFound",
			fmt.Sprintf("%s not found in environment", name), nil)
	}
)

type CustomEnvProvider struct {
	retrieved bool

	AccessKeyIDName     string
	SecretAccessKeyName string
	SessionTokenName    string
}

// Retrieve retrieves the keys from the environment.
func (e *CustomEnvProvider) Retrieve() (credentials.Value, error) {
	e.retrieved = false

	id := os.Getenv(e.AccessKeyIDName)
	if id == "" {
		return credentials.Value{},
			ErrAccessKeyIDNotFound(e.AccessKeyIDName)
	}

	secret := os.Getenv(e.SecretAccessKeyName)
	if secret == "" {
		return credentials.Value{},
			ErrSecretAccessKeyNotFound(e.SecretAccessKeyName)
	}

	e.retrieved = true
	return credentials.Value{
		AccessKeyID:     id,
		SecretAccessKey: secret,
		SessionToken:    os.Getenv(e.SessionTokenName),
	}, nil
}

// IsExpired returns if the credentials have been retrieved.
func (e *CustomEnvProvider) IsExpired() bool {
	return !e.retrieved
}
