FROM ubuntu:18.04

RUN useradd -rm -d /metricbeat -s /bin/bash -g root -G sudo -u 1000 metricbeat \
    && mkdir /metricbeat/modules.d \
    && chown metricbeat:root /metricbeat/modules.d \
    && apt update \
    && apt -y upgrade 

USER metricbeat
WORKDIR /metricbeat

COPY metricbeat /metricbeat
COPY metricbeat.yml /metricbeat
COPY metricbeat.reference.yml /metricbeat
COPY fields.yml /metricbeat
#COPY modules.d/my_module.yml /metricbeat/modules.d/my_module.yml
EXPOSE 9200


CMD cd /metricbeat && ./metricbeat -e -d "*"
