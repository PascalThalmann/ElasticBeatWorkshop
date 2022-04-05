# Elastic Beat Workshop #3 – a more sophisticated configuration for your Metricset

You can find here the code in full length for the workshop [Elastic Beat Workshop #3 – a more sophisticated configuration for your Metricset](https://cdax.ch/2022/03/05/elasticsearch-beat-workshop-1-secured-metricbeat/)

## Prepare the workspace

```
cd ~/go/src/github.com/elastic/beats/metricbeat/

mkdir ~/workspace
mkdir ~/workspace/modules.d
tar zcvf my_module.tar.gz module/my_module \
   fields.yml metricbeat.reference.yml \
   metricbeat.yml modules.d/my_module.yml metricbeat
mkdir ~/workspace
cp my_module.tar.gz ~/workspace
cd ~/workspace
tar zxvf my_module.tar.gz
~/workspace/metricbeat -e -d "*"

```

## Build a Docker image

```
cd ~/workspace/
docker pull docker.elastic.co/beats/metricbeat:8.1.2
docker build -f Dockerfile .
docker run --mount type=bind,source=/home/pascal/workspace/modules.d/,target=/metricbeat/modules.d -it 2ee827d94214
```