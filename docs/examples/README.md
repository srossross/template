
# Examples

## Very simple example

### values.yaml

{% highlight yaml %}
{% raw %}
# File: values.yaml
Thing: library/postgres
ThingName: image
{% endraw %}
{% endhighlight %}

### template.tpl

{% highlight yaml %}
{% raw %}
# File: template.tpl
Your username is {{ default "<unknown>" .Env.USER }}
The default {{ .Values.ThingName }} we should use is "{{ .Values.Thing }}"
{% endraw %}
{% endhighlight %}

### shell

{% highlight shell %}
{% raw %}
template render -f values.yaml template.tpl
{% endraw %}
{% endhighlight %}

### Output:

{% highlight shell %}
{% raw %}
Your username is alice
The default image we should use is "library/postgres"
{% endraw %}
{% endhighlight %}

## Overriding values with `--set`

### shell

{% highlight shell %}
{% raw %}
template render -f values.yaml --set Thing=library/mongodb template.tpl
{% endraw %}
{% endhighlight %}

### Output:

{% highlight shell %}
{% raw %}
Your username is alice
The default image we should use is "library/mongodb"
{% endraw %}
{% endhighlight %}
