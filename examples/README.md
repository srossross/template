
# Examples

## Very simple example

### values.yaml

{% capture simpleValues %}
  {% include simple/values.yaml %}
{% endcapture %}

{% highlight yaml %}
{{ simpleValues }}
{% endhighlight %}

### template.tpl

{% capture simpleValues %}
  {% include simple/values.yaml %}
{% endcapture %}

{% highlight yaml %}
{{ fileContent }}
{% endhighlight %}

### shell

```sh
# Command line
$ template render -f values.yaml template.tpl
The docker image we should use is "library/postgres"
```
