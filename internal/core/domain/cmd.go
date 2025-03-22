package domain

import "os"

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
}

type Matrix struct {
	Name []string `yaml:"name"`
}

type Flag struct {
	Name     string `yaml:"name"`
	Short    string `yaml:"short"`
	Value    string `yaml:"value"`
	Usage    string `yaml:"usage"`
	Required bool   `yaml:"required"`
	Default  string `yaml:"default"`
}

type Arg struct {
	Name     string `yaml:"name"`
	Value    string `yaml:"value"`
	Required bool   `yaml:"required"`
	Default  string `yaml:"default"`
}

func (c *Cmd) SetFlag(flag string, value string) {
	for i, f := range c.Flags {
		if f.Name == flag {
			c.Flags[i].Value = value
			break
		}
	}
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

func (c *Cmd) SetArg(arg string, value string) {
	for i, a := range c.Args {
		if a.Name == arg {
			c.Args[i].Value = value
			break
		}
	}
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
