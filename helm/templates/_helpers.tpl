{{/* The name of the application this chart installs */}}
{{- define "ack-documentdb-controller.app.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "ack-documentdb-controller.app.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/* The name and version as used by the chart label */}}
{{- define "ack-documentdb-controller.chart.name-version" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/* The name of the service account to use */}}
{{- define "ack-documentdb-controller.service-account.name" -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}

{{- define "ack-documentdb-controller.watch-namespace" -}}
{{- if eq .Values.installScope "namespace" -}}
{{ .Values.watchNamespace | default .Release.Namespace }}
{{- end -}}
{{- end -}}

{{/* The mount path for the shared credentials file */}}
{{- define "ack-documentdb-controller.aws.credentials.secret_mount_path" -}}
{{- "/var/run/secrets/aws" -}}
{{- end -}}

{{/* The path the shared credentials file is mounted */}}
{{- define "ack-documentdb-controller.aws.credentials.path" -}}
{{ $secret_mount_path := include "ack-documentdb-controller.aws.credentials.secret_mount_path" . }}
{{- printf "%s/%s" $secret_mount_path .Values.aws.credentials.secretKey -}}
{{- end -}}

{{/* The rules a of ClusterRole or Role */}}
{{- define "ack-documentdb-controller.rbac-rules" -}}
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - patch
  - watch
- apiGroups:
  - ""
  resources:
  - namespaces
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - get
  - list
  - patch
  - watch
- apiGroups:
  - documentdb.services.k8s.aws
  resources:
  - dbclusters
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - documentdb.services.k8s.aws
  resources:
  - dbclusters/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - documentdb.services.k8s.aws
  resources:
  - dbinstances
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - documentdb.services.k8s.aws
  resources:
  - dbinstances/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - documentdb.services.k8s.aws
  resources:
  - dbsubnetgroups
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - documentdb.services.k8s.aws
  resources:
  - dbsubnetgroups/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - ec2.services.k8s.aws
  resources:
  - securitygroups
  verbs:
  - get
  - list
- apiGroups:
  - ec2.services.k8s.aws
  resources:
  - securitygroups/status
  verbs:
  - get
  - list
- apiGroups:
  - ec2.services.k8s.aws
  resources:
  - subnets
  verbs:
  - get
  - list
- apiGroups:
  - ec2.services.k8s.aws
  resources:
  - subnets/status
  verbs:
  - get
  - list
- apiGroups:
  - kms.services.k8s.aws
  resources:
  - keys
  verbs:
  - get
  - list
- apiGroups:
  - kms.services.k8s.aws
  resources:
  - keys/status
  verbs:
  - get
  - list
- apiGroups:
  - services.k8s.aws
  resources:
  - adoptedresources
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - services.k8s.aws
  resources:
  - adoptedresources/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - services.k8s.aws
  resources:
  - fieldexports
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - services.k8s.aws
  resources:
  - fieldexports/status
  verbs:
  - get
  - patch
  - update
{{- end }}