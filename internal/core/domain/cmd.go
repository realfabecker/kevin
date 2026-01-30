package domain

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Cmd struct {
	Name     string  `yaml:"name"`
	Short    string  `yaml:"short"`
	Cmd      string  `yaml:"cmd"`
	Flags    []Flag  `yaml:"flags"`
	Args     []Arg   `yaml:"args"`
	Lang     string  `yaml:"lang"`
	Ref      string  `yaml:"ref"`
	Commands []Cmd   `yaml:"commands"`
	Matrix   *Matrix `yaml:"matrix"`
	Workdir  string
}

type Matrix struct {
	Name []string `yaml:"name"`
}

type Flag struct {
	Name     string   `yaml:"name"`
	Short    string   `yaml:"short"`
	Value    string   `yaml:"value"`
	Usage    string   `yaml:"usage"`
	Required bool     `yaml:"required"`
	Default  string   `yaml:"default"`
	Enum     []string `yaml:"enum"`
}

type Arg struct {
	Name     string   `yaml:"name"`
	Value    string   `yaml:"value"`
	Required bool     `yaml:"required"`
	Default  string   `yaml:"default"`
	Enum     []string `yaml:"enum"`
}

func (c *Cmd) SetFlag(flag string, value string) error {
	for i, f := range c.Flags {
		if f.Name == flag {
			if len(f.Enum) > 0 && !slices.Contains(f.Enum, value) {
				return fmt.Errorf("value for flag %s not in the enumeration: %s", flag, strings.Join(f.Enum, ","))
			}
			c.Flags[i].Value = value
			break
		}
	}
	return nil
}

func (c *Cmd) GetFlag(flag string) string {
	for _, f := range c.Flags {
		if f.Name == flag {
			if f.Value == "" {
				return f.Default
			}
			return f.Value
		}
	}
	return ""
}

func (c *Cmd) SetArg(arg int, value string) error {
	for i, a := range c.Args {
		if i == arg {
			if len(a.Enum) > 0 && !slices.Contains(a.Enum, value) {
				return fmt.Errorf("value for arg %s not in the enumeration: %s", a.Name, strings.Join(a.Enum, ","))
			}
			c.Args[i].Value = value
			break
		}
	}
	return nil
}

func (c *Cmd) GetArg(arg string) string {
	for _, a := range c.Args {
		if a.Name == arg {
			if a.Value == "" {
				return a.Default
			}
			return a.Value
		}
	}
	return ""
}

func (c *Cmd) GetNofRequiredArgs() int {
	total := 0
	for _, a := range c.Args {
		if a.Required {
			total += 1
		}
	}
	return total
}

func (c *Cmd) GetEnv(env string) string {
	return os.Getenv(env)
}

func (c *Cmd) SetWd(wd string) {
	c.Workdir = wd
}

func (c *Cmd) GetWd() string {
	return c.Workdir
}
