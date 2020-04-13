package ora2uml

import (
	"os"
	"text/template"
)

const (
	TemplPlantUML string = `@startuml sample

!define Table(name,desc) class name as " + "\"desc\"" + @" << (T,#FFAAAA) >>
	
!define primary_key(x) <b>x</b>
!define unique(x) <color:green>x</color>
!define not_null(x) <u>x</u>
	
hide methods
hide stereotypes

' Tables ...

{{ with .Tables }} 
{{ range . }}
Table({{ .TableName }} "{{ .TableName }}" {
}
{{ end }} 
{{ end }}

@enduml
`
)

func Export(model Model, templ string) {
	t, err := template.New("model").Parse(templ)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, model)
	if err != nil {
		panic(err)
	}
}
