FROM ubuntu:18.04

RUN mkdir -p /metricbeat/modules.d
RUN apt update
RUN apt -y upgrade

COPY metricbeat /metricbeat
COPY metricbeat.yml /metricbeat
COPY metricbeat.reference.yml /metricbeat
COPY fields.yml /metricbeat
COPY modules.d/my_module.yml /metricbeat/modules.d/my_module.yml
EXPOSE 9200

CMD cd /metricbeat && ./metricbeat -e -d "*"
