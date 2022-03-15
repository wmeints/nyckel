package runtime

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/wmeints/nyckel/pkg/config"
)

type NyckelRuntime struct {
	Config *config.Configuration
}

// CreateOpaqueSecretFromData creates a new opaque secret file from a piece of data.
func (r *NyckelRuntime) CreateOpaqueSecretFromData(name string, key string, data string) error {
	secret := config.NewOpaqueSecret(name)
	r.Config.Secret = &secret

	return r.AddSecretFromData(key, data)
}

// CreateOpaqueSecretFromFile creates a new opaque secret file from a file.
func (r *NyckelRuntime) CreateOpaqueSecretFromFile(name string, key string, data string) error {
	secret := config.NewOpaqueSecret(name)
	r.Config.Secret = &secret

	return r.AddSecretFromFile(key, data)
}

// AddSecretFromData adds a secret to the configuration based on data provided on the command line.
func (r *NyckelRuntime) AddSecretFromData(key, data string) error {
	if _, ok := r.Config.Secret.Data[key]; ok {
		return fmt.Errorf("the key '%s' already exists in the secret", key)
	}

	secretValue := b64.StdEncoding.EncodeToString([]byte(data))
	r.Config.Secret.Data[key] = secretValue

	return nil
}

// AddSecretFromFile adds a secret to the configuration based on data provided from a file.
func (r *NyckelRuntime) AddSecretFromFile(key, file string) error {
	if _, ok := r.Config.Secret.Data[key]; ok {
		return fmt.Errorf("the key '%s' already exists in the secret", key)
	}

	data, err := ioutil.ReadFile(file)

	if err != nil {
		return err
	}

	secretValue := b64.StdEncoding.EncodeToString(data)

	r.Config.Secret.Data[key] = secretValue

	return nil
}

// RemoveSecret removes a secret from the configuration.
func (r *NyckelRuntime) RemoveSecret(key string) error {
	if _, ok := r.Config.Secret.Data[key]; !ok {
		return fmt.Errorf("the key '%s' does not exist in the secret", key)
	}

	delete(r.Config.Secret.Data, key)

	return nil
}

// Save stores the configuration used by the application.
func (r *NyckelRuntime) SaveConfiguration() error {
	return r.Config.Save()
}

func New(path string) (*NyckelRuntime, error) {
	config, err := config.Load(path)

	if err != nil {
		return nil, err
	}

	return &NyckelRuntime{Config: config}, nil
}
