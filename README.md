
# Installing

## Easy install
```
curl  https://srossross.github.io/template/get.sh | bash
```

## Github Releases

Go to https://github.com/srossross/template/releases and get the latest release for your system


# Examples

## Very simple example

### values.yaml
```yaml
# File: values.yaml
Image: library/postgres
```

### template.tpl
```yaml
# File: template.tpl
The docker image we should use is "{{ .Values.Image }}"
```

### shell

```sh
# Command line
$ template render -f values.yaml template.tpl
The docker image we should use is "library/postgres"
```
