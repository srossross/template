# Installing

## Easy install

```
curl  https://srossross.github.io/template/get.sh | sh
```

This will install a single executable file `template` into `/usr/local/bin`

Optionally you can set the version os and arch and install prefix

```
export TEMPLATE_VERSION=v1.0.0 TEMPLATE_ARCH=arm64 TEMPLATE_OS=linux TEMPLATE_INSTALL_PREFIX=/usr/local/bin
curl srossross.github.io/template/get.sh | bash
```

## Inside a Docker container

```
FROM ...

RUN curl srossross.github.io/template/get.sh | sh
```

# Getting Started

Check out the [examples](examples) directory for a list of examples.  
The template command exposes the [go template language](https://golang.org/pkg/text/template/#hdr-Actions)
And adds functions from [sprig](http://masterminds.github.io/sprig/)

This leans heavily from the templates commands of Kubernetes Helm https://docs.helm.sh

## My first template

```yaml
# template.tpl
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Values.Name }}
data:
  myvalue: "Hello World"
```

```
$ template render --set Name=SomeConfig - <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Values.Name }}
data:
  myvalue: "Hello World"
EOF

apiVersion: v1
kind: ConfigMap
metadata:
  name: SomeConfig
data:
  myvalue: "Hello World"

```

# Values Files

* A values file if passed into template with the `-f` flag (`template render -f myvals.yaml ./mytemplate.tpl`)
* Individual parameters passed with `--set` (such as `template render --set foo=bar ./mytemplate.tpl`)

Each `-f` can be overridden by more user-supplied values files, which can in turn be overridden by `--set` parameters.

Values files are plain YAML files. Let’s edit values.yaml and then edit our ConfigMap template.

# Built-in Objects

TODO

# Template Functions and Pipelines

TODO

# Flow Control

TODO

# Named Templates

TODO
