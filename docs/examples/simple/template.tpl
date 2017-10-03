# File: template.tpl
Your username is {{ default "<unknown>" .Env.USER }}
The docker image we should use is "{{ .Values.Image }}"
