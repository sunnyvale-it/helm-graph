

## Output examples

No dependencies:

```console
$ helm-graph render --name "common" -r https://charts.bitnami.com/bitnami -v 1.8.0 -o yaml
name: common
version: 1.8.0
repo: https://charts.bitnami.com/bitnami
deps: []
```

Flat dependencies:

```console
$ helm-graph render --name "grafana-tempo" -r https://charts.bitnami.com/bitnami -v 0.2.5 -o yaml
name: grafana-tempo
version: 0.2.5
repo: https://charts.bitnami.com/bitnami
deps:
- name: common
  version: 0.2.5
  repo: https://charts.bitnami.com/bitnami
  deps: []
- name: memcached
  version: 0.2.5
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