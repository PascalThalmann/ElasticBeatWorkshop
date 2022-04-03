# Elastic Beat Workshop #2 – create your own Metricbeat shipper

You can find here the code in full length for the workshop [Elasticsearch Beats Workshop #2 – create your own Metricbeat shipper](https://cdax.ch/2022/03/26/elasticsearch-beats-workshop-2-create-your-own-metricbeat-shipper/)

## Setting up Go and the dev environment

### Install Go

```
sudo apt-get install python3-venv
sudo mkdir -p /usr/local/go
wget https://go.dev/dl/go1.17.8.linux-amd64.tar.gz
sudo tar -C /usr/local/ -xzf go1.17.8.linux-amd64.tar.gz
export PATH=/usr/local/go/bin:$PATH
```

### Clone the official Beats Repo


```
mkdir go
export GOPATH=~/go
mkdir -p ${GOPATH}/src/github.com/elastic
git clone https://github.com/elastic/beats ${GOPATH}/src/github.com/elastic/beats
```
### Install Mage

```
cd ~
git clone https://github.com/magefile/mage
cd mage
go run bootstrap.go
sudo cp $GOPATH/bin/mage /usr/local/bin
cd ${GOPATH}/src/github.com/elastic/beats
make update
```
### Creating the module and metricset skeleton files

```
cd ${GOPATH}/src/github.com/elastic/beats/metricbeat
make create-metricset

```

### test the installation

```
mage update
mage build
chmod go-w ~/go/src/github.com/elastic/beats/metricbeat/metricbeat.yml
/metricbeat -e -d "*"
```



