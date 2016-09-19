# resource-template Dockerfile used to build docker image on concourse ci

FROM concourse/buildroot:base

# Put the binary into the container
# This is not needed - but lets do it for fun.
COPY resource-template /

# REQUIRED BY CONCROUSE RESOURCE
ADD assets/check /opt/resource/check
ADD assets/in /opt/resource/in
ADD assets/out /opt/resource/out

RUN chmod +x /opt/resource/*