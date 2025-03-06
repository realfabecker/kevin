# kevin

## Introduction

Kevin is a dynamic command creation tool that allows you to define and execute commands based on a configuration file.
It is designed to simplify the process of running complex scripts and commands.

[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=realfabecker_kevin&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=realfabecker_kevin)

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
  - name: "hello"
    short: "Say hello"
    cmd: |
      echo "Hello friend!"

  - name: "greeter"
    short: "greets the user"
    commands:
      - name: "morning"
        short: "Say good morning"
        cmd: |
          echo "Good morning"

      - name: "afternoon"
        short: "Say good good afternoon"
        cmd: |
          echo "Good afternoon"   
```

The kevin.yml configuration file can be stored globally in the user's home directory, or specifically by creating a file in the same directory as the invocation of the kevin command.

With the file ready, it will be possible to call a custom command as follows:

```bash
kevin hello
```

Or a custom sub command

```bash
kevin greeter morning
```


## Contributing

We welcome contributions! Please refer to the [contribution guide](./docs/CONTRIBUTING.md) for details on how to
contribute to the project.

## Versioning

This project uses [SemVer](https://semver.org/) for versioning. For all available versions, check the
[tags in this repository](https://github.com/realfabecker/kevin/tags).

## Licence

This project is licensed under the MIT License. See the [License](LICENSE.md) for more information.
