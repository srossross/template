
# Examples

## Very simple example

### values.yaml

{% capture simpleValues %}
  {% include simple/values.yaml %}
{% endcapture %}

{% highlight yaml %}
{{ fileContent }}
{% endhighlight %}

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
