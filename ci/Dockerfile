# resource-template Dockerfile used to build docker image on concourse ci

FROM concourse/buildroot:base

# Put the binary into the container
# This is not needed - but lets do it for fun.
COPY resource-template /

# Add tree - Not needed but illustrates how to add a command to build
COPY tree /bin

# REQUIRED BY CONCROUSE RESOURCE
#ADD /assets-go/check /opt/resource/check
ADD /assets-bash/check /opt/resource/check
#ADD /assets-go/in /opt/resource/in
ADD /assets-bash/in /opt/resource/in
#ADD /assets-go/out /opt/resource/out
ADD /assets-bash/out /opt/resource/out

RUN chmod +x /opt/resource/*