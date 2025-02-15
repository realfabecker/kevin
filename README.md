# kevin

## Introduction

Kevin is a dynamic command creation tool that allows you to define and execute commands based on a configuration file.
It is designed to simplify the process of running complex scripts and commands.

## Features

- Dynamic command creation
- Configuration via `kevin.yml`
- Support for flags and arguments
- Easy integration with shell scripts

## Installation

You can install or update Kevin using the installation script:

```bash
curl -so- https://raw.githubusercontent.com/realfabecker/kevin/master/install.sh | bash
```

## Usage

Here is an example of how to define and use a command with Kevin in the kevin.yml file:

```yaml
commands:
  - name: "ssh-key-gen"
    short: "generates a new ssh key"
    flags:
      - name: "key"
        short: "k"
        required: true
      - name: "comment"
        short: "c"
        default: "my-key"
    cmd: |
      key="$HOME/.ssh/{{ .GetFlag "key" }}.id_rsa"
      if [[ -f $key ]]; then
        echo "key $key already exists"
        exit 1
      fi;

      ssh-keygen -t rsa \
        -q \
        -f "$key" \
        -C {{ .GetFlag "comment"}} \
        -N ""

      ssh-add $key && cat "${key}.pub"      
```

The kevin.yml configuration file can be stored globally in the user's home directory, or specifically by creating a file in the same directory as the invocation of the kevin command.

With this file ready, it will be possible to call the custom command as follows:

```bash
kevin ssh-key-gen -k nuvem
```

## Contributing

We welcome contributions! Please refer to the [contribution guide](./docs/CONTRIBUTING.md) for details on how to
contribute to the project.

## Versioning

This project uses [SemVer](https://semver.org/) for versioning. For all available versions, check the
[tags in this repository](https://github.com/realfabecker/kevin/tags).

## Licence

This project is licensed under the MIT License. See the [License](LICENSE.md) for more information.
