package data
var index = `

<!DOCTYPE html>
<html lang="fr">
<head>
	<meta charset="utf-8">
	<title>Functions et Types</title>
</head>
<body>
	<h1>Functions et Types</h1>
	<h2>Fonctions</h2>
	<ul id="list">
		{{ range . }}
		<li class="type_{{.Type}}">

			<h3>{{ .Name }}</h3>
			<div class="type">{{- .Type -}}</div>
			<div class="fileRef">
				{{- .FileName }}:{{ .LineNum -}}
			</div>
			{{ range .Comment -}}
				<p class="comment">{{.}}</p>
			{{- end -}}
		</li>
		{{- end }}

	</ul>
</body>
</html>
`
