# Template

Why did I create this?:

 * I am extremely happy with [Helm](https://docs.helm.sh) templating mechanism. But helm does not do simple templates.
 * Can accept values from the environment e.g. `{{ .Env.ENV_VAR }}`
 * Can accept values from a json or yaml file
 * Can override values on the command line.
 * Single easy to install executable. Great for docker containers or CI builds.
 * Powerful [sprig](https://godoc.org/github.com/Masterminds/sprig) template function library

# Similar Templating tools

* [envsubst](https://www.gnu.org/software/gettext/manual/html_node/envsubst-Invocation.html)
* [dockerize](https://github.com/jwilder/dockerize)
* [gomplate](https://github.com/hairyhenderson/gomplate)
* [kontemplate](https://github.com/tazjin/kontemplate)

# Installing

## Easy install

This command will work on osx, linux-32bit and linux-64bit

```
curl https://srossross.github.io/template/get.sh | sh
```

It will install a single executable file `template` into `/usr/local/bin`

Optionally you can set the version os and arch and install prefix

```sh
export  TEMPLATE_VERSION=v1.0.0 \
        TEMPLATE_ARCH=arm64 \
        TEMPLATE_OS=linux \
        TEMPLATE_INSTALL_PREFIX=/usr/local/bin

curl srossross.github.io/template/get.sh | bash
```

## Inside a Docker container

```
FROM ...
ENV TEMPLATE_VERSION v1.0.0
RUN curl srossross.github.io/template/get.sh | sh
```

## OSX Homebrew

The `curl` method above works, but you can also install this via Homebrew on a mac

```
brew tap srossross/template https://github.com/srossross/template
brew install template
```

# Getting Started

Check out the [examples](examples) directory for a list of examples.  
The template command exposes the [go template language](https://golang.org/pkg/text/template/#hdr-Actions)
And adds functions from [sprig](http://masterminds.github.io/sprig/)

This leans heavily from the templates commands of Kubernetes Helm https://docs.helm.sh

## My first template


```yaml
# mytemplate.tpl
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Values.Name }}
data:
  myvalue: "Hello World"
```

You can run this example like this:

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
