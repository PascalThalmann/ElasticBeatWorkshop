# Elastic Beat Workshop #1 – secured Metricbeat

You can find here the code in full length for the workshop [Elastic Beat Workshop #1 – secured metricbeat](https://cdax.ch/2022/02/20/elasticsearch-python-workshop-1-the-basics/)

## Install metricbeat

```
sudo apt install metricbeat=8.0.0
```

## Prereqs

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


