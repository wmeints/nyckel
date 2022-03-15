package config

import (
	"io/ioutil"

	"github.com/goccy/go-yaml"
)

type OpaqueSecret struct {
	ApiVersion string               `yaml:"apiVersion"`
	Kind       string               `yaml:"kind"`
	Metadata   OpaqueSecretMetadata `yaml:"metadata"`
	Data       map[string]string    `yaml:"data"`
}

type OpaqueSecretMetadata struct {
	Name string `yaml:"name"`
}

type Configuration struct {
	Path   string
	Secret *OpaqueSecret
}

// Load loads an opaque secret from file for manipulation
func Load(path string) (*Configuration, error) {
	var secret OpaqueSecret

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &secret)

	if err != nil {
		return nil, err
	}

	if secret.Data == nil {
		secret.Data = make(map[string]string)
	}

	return &Configuration{Secret: &secret, Path: path}, nil
}

// Save stores the opaque secret to file
func (config *Configuration) Save() error {
	data, err := yaml.Marshal(config.Secret)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(config.Path, data, 0644)

	if err != nil {
		return err
	}

	return nil
}

// NewOpaqueSecret creates a new opaque secret
func NewOpaqueSecret(name string) OpaqueSecret {
	return OpaqueSecret{
		ApiVersion: "v1",
		Kind:       "Opaque",
		Metadata: OpaqueSecretMetadata{
			Name: name,
		},
		Data: make(map[string]string),
	}
}
