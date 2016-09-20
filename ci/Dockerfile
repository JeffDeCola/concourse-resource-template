# resource-template Dockerfile used to build docker image on concourse ci

FROM concourse/buildroot:base

# Put the binary into the container
# This is not needed - but lets do it for fun.
COPY resource-template /

# Add tree 
COPY tree /bin

# REQUIRED BY CONCROUSE RESOURCE
ADD /assets-go/check /opt/resource/check
ADD /assets-go/in /opt/resource/in
ADD /assets-go/out /opt/resource/out

RUN chmod +x /opt/resource/*