package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type (
	// Config for the plugin.
	Config struct {
		Actions   []string
		Quiet     bool
		NoColor   bool
		Debug     bool
		JSON      bool
		Insecure  bool
		URL       string
		AccessKey string
		SecretKey string
	}

	// Plugin values
	Plugin struct {
		Config Config
	}
)

// Exec executes the plugin.
func (p *Plugin) Exec() error {
	if len(p.Config.Actions) == 0 {
		return errors.New("you must provide packer action")
	}

	commands := []*exec.Cmd{
		p.versionCommand(),
	}
	// Add commands listed from Actions
	for _, action := range p.Config.Actions {
		switch action {
		default:
			return fmt.Errorf("valid actions are: rm, cp. You provided %s", action)
		}
	}

	for _, cmd := range commands {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Env = os.Environ()

		trace(cmd)

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

func (p *Plugin) versionCommand() *exec.Cmd {
	args := []string{
		"--version",
	}

	return exec.Command(
		"mc",
		args...,
	)
}

func trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}
