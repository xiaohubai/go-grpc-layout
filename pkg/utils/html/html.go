package html

import (
	"bytes"
	"errors"
	"html/template"
)

var emailMaps = map[string]string{
	//告警html
	"warn": `
		<html>

		<body>
			<p>● DateTime: <td>{{.DateTime}}</td>
			<p>● TraceID: <td>{{.TraceID}}</td>
			<p>● Error: <td>{{.Error}}</td>
		</body>

		</html>
	`,
	//
}

// EmailHTMLByText 发送邮件的html格式化
func EmailHTMLByText(topic string, value interface{}) (string, error) {
	text, ok := emailMaps[topic]
	if !ok {
		return "", errors.New("topic not found")
	}
	tmpl, err := template.New("emailTemp").Parse(text)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, value)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
