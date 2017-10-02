
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
Your username is {{ default "<unknown>" .Env.USER }}
The docker image we should use is "{{ .Values.Image }}"
```

### shell

```sh
# Command line
$ template render -f values.yaml template.tpl
The docker image we should use is "library/postgres"
```
