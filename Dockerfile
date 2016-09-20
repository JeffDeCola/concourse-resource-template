# resource-template Dockerfile used to build docker image on concourse ci

FROM concourse/buildroot:base

# Put the binary into the container
# This is not needed - but lets do it for fun.
COPY resource-template /

# Add tree 
COPY bin/tree /bin

# REQUIRED BY CONCROUSE RESOURCE
ADD /assets-bash/check /opt/resource/check
ADD /assets-bash/in /opt/resource/in
ADD /assets-bash/out /opt/resource/out

RUN chmod +x /opt/resource/*
