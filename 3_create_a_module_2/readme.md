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
```

## Upload the Docker image to your registry

```
docker login -u xxx -p yyy [local-registry]:[port]
docker build -f Dockerfile . -t my_module:1.1
docker image tag my_module:1.1 srvnexus:8082/repository/dh/my_module:1.1
docker image push srvnexus:8082/repository/dh/my_module:1.1
```

## pull your docker image from your registry

```
docker pull srvnexus:8082/repository/dh/my_module:1.1
```

## save the docker image as a tarball

```
docker image save -o my_module.tar.gz [image id]
```

## Run the Docker container

```
METRICPATH=$HOME/workspace/metricbeat.yml
MODULEPATH=$HOME/workspace/modules.d

docker run \
--mount type=bind,source=$MODULEPATH,target=/metricbeat/modules.d \
--mount type=bind,source=$METRICPATH,target=/metricbeat/metricbeat.yml \
-it [image id]
```

## Pull the my_module Docker image from Dockerhub.io

```
docker pull cdax75/workshop:my_module
```