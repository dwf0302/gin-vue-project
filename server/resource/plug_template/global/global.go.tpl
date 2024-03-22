package global

{{- if .HasGlobal }}

import "go-vue-project/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}