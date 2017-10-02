
# Installing

```
curl srossross.github.io/template/get | bash
```

```
ENV TEMPLATE_VERSION v0.0.3-3
wget https://github.com/srossross/template/releases/download/${TEMPLATE_VERSION}/template-linux-amd64.tgz && tar -xf template-linux-amd64.tgz && chmod +x ./template-linux-amd64 && mv template-linux-amd64 /usr/bin/template
```

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
