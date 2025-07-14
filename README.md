# kevin

## Introdução

Kevin é uma ferramenta dinâmica de criação de comandos que permite definir e executar comandos com base em um arquivo de configuração. Ela foi projetada para simplificar o processo de execução de scripts e comandos complexos.

[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=realfabecker_kevin&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=realfabecker_kevin)

## Features

- Criação dinâmica de comandos
- Configuração via arquivo `kevin.yml`
- Suporte a flags e argumentos
- Integração fácil com shell scripts

## Instalação

Você pode instalar ou atualizar o Kevin usando seu script de instalação:

```
curl -so- https://raw.githubusercontent.com/realfabecker/kevin/master/install.sh | bash
```

## Usando

Aqui um exemplo de como definir e usar um comando com o Kevin no seu arquivo kevin.yml:

```yaml
commands:
  - name: "pg-restore"
    short: "Restaura um banco de dados a partir de um arquivo de backup"
    flags:
      - name: "database"
        required: true
      - name: "backup"
        required: true
    cmd: |
      pg_restore -h localhost -p 5432 -U postgres \
            --database {{ .GetFlag "database" }} --backup  {{ .GetFlag "backup" }}
```

O arquivo de configuração kevin.yml pode ser armazenado globalmente no diretório do usuário, ou por criando um arquivo no mesmo diretório em que o comando kevin é invocado.n command.

Com o arquivo pronto, será possível chamar o comando personalizado da seguinte forma:

```bash
kevin pg-restore --database app --backup ./backup.sql
```

## Contribuindo

Você gostaria de contribuir para o projeto? Confira nosso [guia][link-contrib] para saber como contribuir.

## Versionamento

O versionamento desse projeto é baseado no [SemVer](https://semver.org/). Verifique as [tags do projeto][link-tags] para informações sobre as versões disponíveis.

## Licença

Este projeto está sob a licença MIT. Confira os detalhes no arquivo [Lincença][link-license]

[link-tags]: https://github.com/realfabecker/kevin/tags
[link-license]: https://github.com/realfabecker/.github/blob/main/.github/LICENSE.md
[link-contrib]: https://github.com/realfabecker/.github/blob/main/.github/CONTRIBUTING.md
