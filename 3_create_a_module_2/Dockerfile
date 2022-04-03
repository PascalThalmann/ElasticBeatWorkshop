#FROM docker.elastic.co/beats/metricbeat:8.1.2
FROM ubuntu:latest

#RUN rm -fr /usr/share/metricbeat/modules.d/*
RUN mkdir -p /metricbeat/modules.d

#COPY metricbeat /usr/share/metricbeat
#COPY metricbeat.yml /usr/share/metricbeat
#COPY metricbeat.reference.yml /usr/share/metricbeat
#COPY fields.yml /usr/share/metricbeat
#COPY modules.d/my_module.yml /usr/share/metricbeat/modules.d

COPY metricbeat /metricbeat
COPY metricbeat.yml /metricbeat
COPY metricbeat.reference.yml /metricbeat
COPY fields.yml /metricbeat
COPY modules.d/my_module.yml /metricbeat/modules.d/my_module.yml

CMD cd /metricbeat && ./metricbeat -e -d "*"
