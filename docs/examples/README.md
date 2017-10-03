
# Examples

## Very simple example

### values.yaml

{% capture fileContent %}
  {% include simple/values.yaml %}
{% endcapture %}

{% highlight yaml %}
{{ fileContent }}
{% endhighlight %}

### template.tpl

{% capture fileContent %}
  {% include simple/template.tpl %}
{% endcapture %}

{% highlight yaml %}
{{ fileContent }}
{% endhighlight %}

### shell

{% capture fileContent %}
  {% include simple/simple.sh %}
{% endcapture %}

{% highlight shell %}
{{ fileContent }}
{% endhighlight %}

### Output:

```
The docker image we should use is "library/postgres"
```
