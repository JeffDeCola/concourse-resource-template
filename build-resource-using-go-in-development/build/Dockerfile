# resource-template Dockerfile used to build docker image on concourse ci

FROM concourse/buildroot:base

# Put the binary into the container
# This is need with go only
COPY concourse-resource-template /

# Add tree - Not needed but illustrates how to add a command to build
COPY bin/tree /bin

# REQUIRED BY CONCROUSE RESOURCE
ADD /assets-go/check /opt/resource/check
# ADD /check-in-out/check /opt/resource/check
ADD /assets-go/in /opt/resource/in
# ADD /check-in-out/in /opt/resource/in
ADD /assets-go/out /opt/resource/out
# ADD /check-in-out/out /opt/resource/out

RUN chmod +x /opt/resource/*
