

## Usage examples

No dependencies:

```console
$ helm-graph render --name "common" -r https://charts.bitnami.com/bitnami -v 1.8.0 -o yaml
name: common
version: 1.8.0
repo: https://charts.bitnami.com/bitnami
deps: []
```

Nested dependencies:

```console
$ helm-graph render --name "kafka" -r https://charts.bitnami.com/bitnami -v 14.0.5 -o yaml
name: kafka
version: 14.0.5
repo: https://charts.bitnami.com/bitnami
deps:
- name: common
  version: 1.8.0
  repo: https://charts.bitnami.com/bitnami
  deps: []
- name: zookeeper
  version: 7.4.1
  repo: https://charts.bitnami.com/bitnami
  deps:
  - name: common
    version: 1.8.0
    repo: https://charts.bitnami.com/bitnami
    deps: []
```

Pretty print json output with jq

```console
$ helm-graph render --name "kafka" -r https://charts.bitnami.com/bitnami -v 14.0.5 -o json | jq .
{
  "Name": "kafka",
  "Version": "14.0.5",
  "Repo": "https://charts.bitnami.com/bitnami",
  "Deps": [
    {
      "Name": "common",
      "Version": "1.8.0",
      "Repo": "https://charts.bitnami.com/bitnami",
      "Deps": null
    },
    {
      "Name": "zookeeper",
      "Version": "7.4.1",
      "Repo": "https://charts.bitnami.com/bitnami",
      "Deps": [
        {
          "Name": "common",
          "Version": "1.8.0",
          "Repo": "https://charts.bitnami.com/bitnami",
          "Deps": null
        }
      ]
    }
  ]
}
```

Flat json output with jq:

```console
$ helm-graph render --name "kafka" -r https://charts.bitnami.com/bitnami -v 14.0.5 -o json |  jq '[leaf_paths as $path | {"key": $path | join("."), "value": getpath($path)}] | from_entries'
{
  "Name": "kafka",
  "Version": "14.0.5",
  "Repo": "https://charts.bitnami.com/bitnami",
  "Deps.0.Name": "common",
  "Deps.0.Version": "1.8.0",
  "Deps.0.Repo": "https://charts.bitnami.com/bitnami",
  "Deps.1.Name": "zookeeper",
  "Deps.1.Version": "7.4.1",
  "Deps.1.Repo": "https://charts.bitnami.com/bitnami",
  "Deps.1.Deps.0.Name": "common",
  "Deps.1.Deps.0.Version": "1.8.0",
  "Deps.1.Deps.0.Repo": "https://charts.bitnami.com/bitnami"
}
```

All the charts as objects in a flat array, removing duplicated:

```console
$ helm-graph render --name "airflow" -r https://charts.bitnami.com/bitnami -v 10.3.1 -o json |  jq '[.. | objects | select(has("Deps")) | {Name: .Name, Version: .Version, Repo: .Repo}] |unique'                                                                                                     
[
  {
    "Name": "airflow",
    "Version": "10.3.1",
    "Repo": "https://charts.bitnami.com/bitnami"
  },
  {
    "Name": "common",
    "Version": "1.7.1",
    "Repo": "https://charts.bitnami.com/bitnami"
  },
  {
    "Name": "common",
    "Version": "1.8.0",
    "Repo": "https://charts.bitnami.com/bitnami"
  },
  {
    "Name": "postgresql",
    "Version": "10.9.3",
    "Repo": "https://charts.bitnami.com/bitnami"
  },
  {
    "Name": "redis",
    "Version": "14.8.11",
    "Repo": "https://charts.bitnami.com/bitnami"
  }
]
```