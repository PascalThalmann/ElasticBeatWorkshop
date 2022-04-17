# Elastic Beat Workshop #4 – Metricbeat versus Filebeat

You can find here the code in full length for the workshop [Elastic Beat Workshop #4 – Metricbeat versus Filebeat](https://cdax.ch/2022/04/09/elasticsearch-beats-workshop-3-a-more-sophisticated-configuration-for-your-metricset/)

## Installing Filebeat

This works on Ubuntu 18.04

```
sudo apt-get install apt-transport-https
echo "deb https://artifacts.elastic.co/packages/8.x/apt stable main" \
  | sudo tee -a /etc/apt/sources.list.d/elastic-8.x.list
wget -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch \
  | sudo apt-key add -
sudo apt-get update && sudo apt install filebeat=8.0.0
apt-mark hold filebeat
```

## Add a user and a password to the Keystore

create the keystore for the user filebeat_user 

```
filebeat keystore add FB_USER
The keystore does not exist. Do you want to create it? [y/N]: y
Created keystore
Enter value for MB_USER: 
Successfully updated the keystone

filebeat keystore add FB_PW
Enter value for FB_PW: 
Successfully updated the keystone
```

## Create the filebeat roles and a user filebeat_user

create a role filebeat_setup_role:

  - cluster privileges: monitor, manage, all
  - indices: testdata, privileges: manage

create a role filebeat_role:

  - indices: testdata, privileges: create, create_doc, delete, write, all

create a user filebeat_user:

  - privileges: ingest_admin, filebeat_role, filebeat_setup_role

## Configure Filebeat

This is for a minimal configuration with secured communication. Save it to ~/filebeat.yml

```
filebeat.inputs:
- type: log
  enabled: true
  paths:
    - /data/filebeat.test.log
  exclude_lines: ["^\"\""]
  processors:
  - decode_csv_fields:
      fields:
        message: decoded.csv
      separator: ","
  - extract_array:
        field: decoded.csv
        mappings:
          timestamp: 0
          file_name: 1
          mod_time: 2
  - drop_fields:
        fields: ["decoded"]

setup.template.enabled: false
setup.ilm.enabled: false

output.elasticsearch:
  hosts: ["https://your.elastic.server.com:9200"]
  protocol: "https"
  username: "${FB_USER}"
  password: "${FB_PW}"
  index: "testdata"
  allow_older_versions: true
  ssl:
    enabled: true
    ca_trusted_fingerprint: "a2929842b8920e5c0ebd91ea157c159f16b62df7e3b6998de93ad56ff2693b59"
```

## script fbeat_test.ksh

This script simulates log entries

```
#!/bin/bash

while true; do   
	DATE=`date --rfc-3339=seconds`  
	FILE_NAME=filebeat.log;   
	MOD_TIME=`ls -l --time-style=full-iso ${FILE_NAME} | awk '{print $6 " " $7}'`;   
	echo "${DATE},${FILE_NAME},${MOD_TIME}" >> ${FILE_NAME};    
	sleep 10; 
done
```

## run filebeat

```
filebeat -v -e -c ~/filebeat.yml
```