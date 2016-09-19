# resource-template Dockerfile used to build docker image on concourse ci

FROM concourse/buildroot:base

# Put the binary into the container
COPY resource-template /

# REQUIRED BY CONCROUSE RESOURCE
COPY scripts/ /opt/resource/
RUN chmod +x /opt/resource/*