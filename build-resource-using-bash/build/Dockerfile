# resource-template Dockerfile used to build docker image on concourse ci

# FROM concourse/buildroot:base
FROM alpine

LABEL maintainer="Jeff DeCola https://github.com/JeffDeCola/concourse-resource-template"

# UPDATE AND UPGRADE
RUN apk update

# INSTALL PACKAGES
RUN apk add --no-cache curl
RUN apk add --no-cache git
RUN apk add --no-cache bash
RUN apk add --no-cache jq

# ADD TREE EXECUTABLE
# Not needed but illustrates how to add a command to build
# COPY bin/tree /bin

# REQUIRED BY CONCOURSE RESOURCE
ADD /check-in-out/check /opt/resource/check
ADD /check-in-out/in /opt/resource/in
ADD /check-in-out/out /opt/resource/out

# MAKE EXECUTABLE
