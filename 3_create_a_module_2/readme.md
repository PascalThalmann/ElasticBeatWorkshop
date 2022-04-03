# Elastic Beat Workshop #3 – a more sophisticated configuration for your Metricset

You can find here the code in full length for the workshop [Elastic Beat Workshop #3 – a more sophisticated configuration for your Metricset](https://cdax.ch/2022/03/05/elasticsearch-beat-workshop-1-secured-metricbeat/)

## Prepare the workspace

```
mkdir ~/workspace
mkdir ~/workspace/modules.d
cp ~/go/src/github.com/elastic/beats/metricbeat/metricbeat ~/workspace
cp ~/go/src/github.com/elastic/beats/metricbeat/metricbeat.yml ~/workspace/
cp ~/go/src/github.com/elastic/beats/metricbeat/modules.d/my_module.yml ~/workspace/modules.d/
~/workspace/metricbeat -e -d "*"

```

## Create the Dockerfile

```

```