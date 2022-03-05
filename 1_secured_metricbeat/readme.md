# Elastic Beat Workshop #1 – secured Metricbeat

You can find here the code in full length for the workshop [Elastic Beat Workshop #1 – secured metricbeat](https://cdax.ch/2022/03/05/elasticsearch-beat-workshop-1-secured-metricbeat/)

## Install metricbeat

```
sudo apt install metricbeat=8.0.0
```

## Prerequisites

### /etc/hosts

```
192.168.1.68   srvelk8      srvelk8.local.ch
```

### elasticsearch.yml

```
cluster.name: cluster_3
node.name: srvelk8
path.data: /var/lib/elasticsearch
path.logs: /var/log/elasticsearch
network.host: srvelk8.local.ch
http.port: 9200

xpack.security.enabled: true
xpack.security.enrollment.enabled: true

xpack.security.http.ssl:
  enabled: true
  keystore.path: certs/http.p12

xpack.security.transport.ssl:
  enabled: true
  verification_mode: certificate
  keystore.path: certs/transport.p12
  truststore.path: certs/transport.p12

cluster.initial_master_nodes: ["srvelk8"]
http.host: [_local_, _site_]
```

### get the ssl fingerprint 

#### from kibana.yml

```
cat /etc/kibana/kibana.yml |grep fingerprint
```

### with OpenSSL

```
openssl x509 -fingerprint -sha256 -in /etc/elasticsearch/certs/http_ca.crt | \
  grep Fingerprint|cut -d  '=' -f2| tr -d ':'|tr A-Z a-z

a2929842b8920e5c0ebd91ea157c159f16b62df7e3b6998de93ad56ff2693b59
```
## Installation

```
curl -L -O https://artifacts.elastic.co/downloads/beats/metricbeat/metricbeat-8.0.1-amd64.deb
sudo dpkg -i metricbeat-8.0.1-amd64.deb
apt-mark hold metricbeat
```

## create the keystore


```
metricbeat keystore create
metricbeat keystore add MB_USER
metricbeat keystore add MB_PW
```

## enable modules

```
metricbeat modules enable elasticsearch-xpack
metricbeat modules enable kibana-xpack
metricbeat modules enable beat-xpack
```

## Configuration

### metricbeat.yml

```
metricbeat.config.modules:
  path: ${path.config}/modules.d/*.yml
  reload.enabled: false

setup.template.settings:
  index.number_of_shards: 2
  index.codec: best_compression

processors:
  - add_host_metadata: ~

monitoring.enabled: true

setup.ilm.overwrite: true

setup.kibana:
  host: "http://srvelk8.local.ch:5601"
  username: "${MB_USER}"
  password: "${MB_PW}"
  #api_key: "TxDLW38BJmYfkUzINFEF:0foLANnlS-qMNgx_jEXhGw"
  ssl:
    certificate_authorities: /etc/elasticsearch/certs/http_ca.crt
    enabled: true
    ca_trusted_fingerprint: "a2929842b8920e5c0ebd91ea157c159f16b62df7e3b6998de93ad56ff2693b59"

output.elasticsearch:
  hosts: ["https://192.168.1.68:9200"]
  protocol: "https"
  username: "${MB_USER}"
  password: "${MB_PW}"
  #api_key: "TxDLW38BJmYfkUzINFEF:0foLANnlS-qMNgx_jEXhGw"
  ssl:
    #certificate_authorities: /etc/elasticsearch/certs/http_ca.crt
    enabled: true
    ca_trusted_fingerprint: "a2929842b8920e5c0ebd91ea157c159f16b62df7e3b6998de93ad56ff2693b59"
```

### module configuration

#### elasticsearch-xpack.yml
```
- module: elasticsearch
  xpack.enabled: true
  period: 10s
  hosts: ["https://192.168.1.68:9200"]
  protocol: "https"
  username: "${MB_USER}"
  password: "${MB_PW}"
  #api_key: "TxDLW38BJmYfkUzINFEF:0foLANnlS-qMNgx_jEXhGw"
  ssl:
    #certificate_authorities: /etc/elasticsearch/certs/http_ca.crt
    enabled: true
    ca_trusted_fingerprint: "a2929842b8920e5c0ebd91ea157c159f16b62df7e3b6998de93ad56ff2693b59"
```
#### kibana-xpack.yml


```
- module: kibana
  xpack.enabled: true
  period: 10s
  hosts: ["http://srvelk8.local.ch:5601"]
  username: "${MB_USER}"
  password: "${MB_PW}"
  #api_key: "TxDLW38BJmYfkUzINFEF:0foLANnlS-qMNgx_jEXhGw"
  ssl:
    #certificate_authorities: /etc/elasticsearch/certs/http_ca.crt
    enabled: true
    ca_trusted_fingerprint: "a2929842b8920e5c0ebd91ea157c159f16b62df7e3b6998de93ad56ff2693b59"
```

#### beat-xpack.yml


```
- module: beat
  xpack.enabled: true
  period: 10s
  hosts: ["https://srvelk8.local.ch:5066"]
  metricsets:
    - stats
    - state
  username: "${MB_USER}"
  password: "${MB_PW}"
  #api_key: "TxDLW38BJmYfkUzINFEF:0foLANnlS-qMNgx_jEXhGw"
  ssl:
    #certificate_authorities: /etc/elasticsearch/certs/http_ca.crt
    enabled: true
    ca_trusted_fingerprint: "a2929842b8920e5c0ebd91ea157c159f16b62df7e3b6998de93ad56ff2693b59"
```

## setup metricbeat


```
metricbeat setup -e
```




