package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const ALIAS = "minio"

type (
	// Config for the plugin.
	Config struct {
		Actions    []string
		IsQuiet    bool
		IsNoColor  bool
		IsDebug    bool
		IsJSON     bool
		IsInsecure bool
		URL        string
		AccessKey  string
		SecretKey  string

		// rm flag
		IsForce      bool
		IsRecursive  bool
		IsDangerous  bool
		IsFake       bool
		IsIncomplete bool
		EncryptKey   string
		OlderThan    int
		NewerThan    int
		Path         string
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
		p.addConfigCommand(),
	}
	// Add commands listed from Actions
	for _, action := range p.Config.Actions {
		switch action {
		case "rm":
			commands = append(commands, p.rmCommand())
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

func (p *Plugin) rmCommand() *exec.Cmd {
	args := []string{
		"rm",
	}

	if p.Config.IsRecursive {
		args = append(args, "--recursive")
	}

	if p.Config.IsIncomplete {
		args = append(args, "--incomplete")
	}

	if p.Config.IsFake {
		args = append(args, "--fake")
	}

	if p.Config.IsDangerous {
		args = append(args, "--dangerous")
	}

	if p.Config.IsForce {
		args = append(args, "--force")
	}

	if p.Config.OlderThan != 0 {
		args = append(args, "--older-than", fmt.Sprintf("%d", p.Config.OlderThan))
	}

	if p.Config.NewerThan != 0 {
		args = append(args, "--newer-than", fmt.Sprintf("%d", p.Config.NewerThan))
	}

	if p.Config.EncryptKey != "" {
		args = append(args, "--encrypt-key", `"`+p.Config.EncryptKey+`"`)
	}

	args = append(args, ALIAS+"/"+strings.TrimLeft(p.Config.Path, "/"))

	return exec.Command(
		"mc",
		args...,
	)
}

func (p *Plugin) addConfigCommand() *exec.Cmd {
	args := []string{
		"config",
		"host",
		"add",
		ALIAS,
		p.Config.URL,
		p.Config.AccessKey,
		p.Config.SecretKey,
	}

	return exec.Command(
		"mc",
		args...,
	)
}

func (p *Plugin) versionCommand() *exec.Cmd {
	args := []string{
		"version",
	}

	return exec.Command(
		"mc",
		args...,
	)
}

func trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}
