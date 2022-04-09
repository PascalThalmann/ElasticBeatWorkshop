# Elastic Beat Workshop #3 – a more sophisticated configuration for your Metricset

You can find here the code in full length for the workshop [Elastic Beat Workshop #3 – a more sophisticated configuration for your Metricset](https://cdax.ch/2022/04/09/elasticsearch-beats-workshop-3-a-more-sophisticated-configuration-for-your-metricset/)

## Prepare the workspace

```
cd ~/go/src/github.com/elastic/beats/metricbeat/
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
docker build -f Dockerfile . -t my_module:1.0
docker run --mount type=bind,source=/home/pascal/workspace/modules.d/,target=/metricbeat/modules.d -it [image id]
docker image save -o my_module.tar.gz [image id]
```
