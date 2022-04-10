FROM ubuntu:18.04

RUN useradd -rm -d /metricbeat -s /bin/bash -g root -G sudo -u 1000 metricbeat \
    && mkdir /metricbeat/modules.d \
    && chown metricbeat:root /metricbeat/modules.d \
    && apt update \
    && apt -y upgrade 

USER metricbeat
WORKDIR /metricbeat

COPY metricbeat /metricbeat

EXPOSE 9200

CMD cd /metricbeat && ./metricbeat -e -d "*"
