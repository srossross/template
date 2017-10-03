
# Examples

## Very simple example

### values.yaml

{% highlight yaml %}
{% raw %}
# File: values.yaml
Image: library/postgres
{% endraw %}
{% endhighlight %}

### template.tpl

{% highlight yaml %}
{% raw %}
# File: template.tpl
Your username is {{ default "<unknown>" .Env.USER }}
The docker image we should use is "{{ .Values.Image }}"
{% endraw %}
{% endhighlight %}

### shell

{% highlight shell %}
{% raw %}
template render -f values.yaml template.tpl
{% endraw %}
{% endhighlight %}

### Output:

```
The docker image we should use is "library/postgres"
```
