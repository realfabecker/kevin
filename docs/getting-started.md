# Iniciando com o Kevin

Bem vindo ao Kevin!

Esse tutorial vai te guiar no básico  da instalação e criação de sua primeira ferramenta de linha de comando com o kevin.

## Paço 1: Instalando o Kevin

Primeiro vamos instalar e configurar o Kevin no seu ambiente:

```bash
curl -so- https://raw.githubusercontent.com/realfabecker/kevin/master/install.sh | bash
```

Com o comando acima estamos baixando e executando o script de instalação.

Caso preferir você pode obter a [ultima versão](https://github.com/realfabecker/kevin/releases/tag/0.15.0) do aplicativo diretamente pelo github no github.

## Paço 2: Crie a sua Aplicação

Vamos nessa etapa criar a definição da nossa aplicação que será gerada pelo kevin.

```yml
commands:
  - name: "soma"
    short: "Realiza a soma entre dois números"
    args:
      - name: "num1"
        required: true
      - name: "num2"
        required: true
    cmd: |
      echo $(({{.GetArg "num1"}} + {{.GetArg "num2"}}))
```

O conteúdo acima deverá ser salvo em um arquivo de nome `kevin.yml`.

## Paço 3: Testando a Aplicação

No diretório em que o arquivo anterior foi salvo podemos agora testar a nossa aplicação:

```bash
$ kevin

Usage:
  kevin [command]

Available Commands:
  soma        Realiza a soma entre dois números

Flags:
      --dry-run   run in dry run mode
  -h, --help      help for kevin

Use "kevin [command] --help" for more information about a command.
```

A invocação do comando `kevin` sem nenhum argumento exibe o menu de ajuda com os comandos disponíveis.

> Nesse é possível observar o comando `soma` que foi especificado no `kevin.yml` bem como a sua descrição de uso.

```bash
$ kevin soma 1 2
3
```

!!! note
Perfeito, com essa simples configuração já temos um utilitário em funcionamento.