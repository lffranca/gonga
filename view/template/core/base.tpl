{{ define "base" }}
    <!doctype html>
    <html lang="en">
    <head>
        {{ template "head" . }}
    </head>
    <body>
    {{ template "header" . }}
    {{ template "content" . }}
    {{ template "footer" . }}
    {{ template "scripts" . }}
    </body>
    </html>
{{end}}
