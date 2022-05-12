{{ define "go-api.deployment" }}
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ print .name | quote }}
spec:
  selector:
    matchLabels:
      app: {{ .name | quote }}
  template:
    metadata:
      labels:
        app: {{ .name | quote }}
    spec:
      containers:
        - name: {{ .name | quote }}
          image: {{ print .image.name ":" .image.tag | quote }}
          ports:
            - containerPort: {{ .config.port }}
{{- end }}
