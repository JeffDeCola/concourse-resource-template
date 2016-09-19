# resource-template Dockerfile used to build docker image on concourse ci

FROM concourse/buildroot:base

# Put the binary into the container
COPY resource-template /

# REQUIRED BY CONCROUSE RESOURCE
COPY scripts/check /opt/resource/check
COPY scripts/in /opt/resource/in
COPY scripts/out /opt/resource/out

RUN chmod +x /opt/resource/*