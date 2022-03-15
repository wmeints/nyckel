package main

import (
	"fmt"
	"regexp"

	"github.com/alecthomas/kong"
	"github.com/wmeints/nyckel/pkg/runtime"
)

const KeyPattern = `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{1,62}){1}(.[a-zA-Z0-9]{1}[a-zA-Z0-9_-]{1,62})*$`

type CreateCmd struct {
	Path      string `help:"The path to the secret file" name:"path" type:"path"`
	Name      string `help:"The name of the secret" name:"name" type:"string"`
	Key       string `help:"The key of the secret" name:"key" type:"string"`
	Data      string `help:"The unencoded secret value" optional:"" name:"data" type:"string"`
	InputFile string `help:"The input file containing the secret data" name:"input-file" optional:"" type:"path"`
}

func (cmd *CreateCmd) Run() error {
	app, err := runtime.New(cmd.Path)

	if err != nil {
		return err
	}

	if cmd.Name == "" {
		return fmt.Errorf("please provide a name for the secret")
	}

	if isNameValid, _ := regexp.Match(KeyPattern, []byte(cmd.Name)); !isNameValid {
		return fmt.Errorf("the provided name '%s' is not valid", cmd.Name)
	}

	if cmd.Key == "" {
		return fmt.Errorf("please provide a key for the secret")
	}

	if isKeyValid, _ := regexp.Match(KeyPattern, []byte(cmd.Key)); !isKeyValid {
		return fmt.Errorf("the provided key '%s' is not valid", cmd.Key)
	}

	if cmd.InputFile != "" {
		if err := app.AddSecretFromFile(cmd.Key, cmd.InputFile); err != nil {
			return err
		}

		return app.SaveConfiguration()
	}

	if cmd.Data != "" {
		err = app.CreateOpaqueSecretFromData(cmd.Name, cmd.Key, cmd.Data)

		if err != nil {
			return err
		}

		return app.SaveConfiguration()
	}

	return fmt.Errorf("please provide data using the --data argument or an input file using the --input-file argument")
}

type AddCmd struct {
	Path      string `help:"The path to the secret file" name:"path" type:"path"`
	Key       string `help:"The key of the secret" name:"key"`
	Data      string `help:"The unencoded secret value" optional:"" name:"data"`
	InputFile string `help:"The input file containing the secret data" name:"input-file" optional:"" type:"path"`
}

func (cmd *AddCmd) Run() error {
	app, err := runtime.New(cmd.Path)

	if err != nil {
		return err
	}

	if cmd.Key == "" {
		return fmt.Errorf("please provide a key for the secret")
	}

	if isKeyValid, _ := regexp.Match(KeyPattern, []byte(cmd.Key)); !isKeyValid {
		return fmt.Errorf("the provided key '%s' is not valid", cmd.Key)
	}

	if cmd.InputFile != "" {
		if err := app.AddSecretFromFile(cmd.Key, cmd.InputFile); err != nil {
			return err
		}

		return app.SaveConfiguration()
	}

	if cmd.Data != "" {
		err = app.AddSecretFromData(cmd.Key, cmd.Data)

		if err != nil {
			return err
		}

		return app.SaveConfiguration()
	}

	return fmt.Errorf("please provide data using the --data argument or an input file using the --input-file argument")
}

type RemoveCmd struct {
	Path string `help:"The path to the secret file" name:"path" type:"path"`
	Key  string `help:"The key of the secret" name:"key"`
}

func (cmd *RemoveCmd) Run() error {
	app, err := runtime.New(cmd.Path)

	if err != nil {
		return nil
	}

	if cmd.Key == "" {
		return fmt.Errorf("please provide a key for the secret")
	}

	if isKeyValid, _ := regexp.Match(KeyPattern, []byte(cmd.Key)); !isKeyValid {
		return fmt.Errorf("the provided key '%s' is not valid", cmd.Key)
	}

	app.RemoveSecret(cmd.Key)

	return nil
}

type UpdateCmd struct {
	Path      string `help:"The path to the secret file" name:"path" type:"path"`
	Key       string `help:"The key of the secret" name:"key"`
	Data      string `help:"The unencoded secret value" optional:"" name:"data"`
	InputFile string `help:"The input file containing the secret data" name:"input-file" optional:"" type:"path"`
}

func (cmd *UpdateCmd) Run() error {
	app, err := runtime.New(cmd.Path)

	if err != nil {
		return nil
	}

	if cmd.Key == "" {
		return fmt.Errorf("please provide a key for the secret")
	}

	if isKeyValid, _ := regexp.Match(KeyPattern, []byte(cmd.Key)); !isKeyValid {
		return fmt.Errorf("the provided key '%s' is not valid", cmd.Key)
	}

	if cmd.InputFile != "" {
		app.RemoveSecret(cmd.Key)
		app.AddSecretFromFile(cmd.Key, cmd.InputFile)

		return app.SaveConfiguration()
	}

	if cmd.Data != "" {
		app.RemoveSecret(cmd.Key)
		app.AddSecretFromData(cmd.Key, cmd.Data)

		return app.SaveConfiguration()
	}

	return fmt.Errorf("please provide data using the --data argument or an input file using the --input-file argument")
}

var cli struct {
	Create CreateCmd `cmd:"create" help:"Create a new opaque secret file"`
	Add    AddCmd    `cmd:"add" help:"Add a secret to an existing opaque secret file"`
	Remove RemoveCmd `cmd:"remove" help:"Remove a secret from an existing opaque secret file"`
	Update UpdateCmd `cmd:"update" help:"Update a secret in an existing opaque secret file"`
}

func main() {
	ctx := kong.Parse(&cli, kong.UsageOnError())
	err := ctx.Run()

	ctx.FatalIfErrorf(err)
}
