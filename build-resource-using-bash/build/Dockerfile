# resource-template Dockerfile used to build docker image on concourse ci

FROM concourse/buildroot:base

# ADD TREE EXECUTABLE
# Not needed but illustrates how to add a command to build
COPY bin/tree /bin

# REQUIRED BY CONCROUSE RESOURCE
ADD /check-in-out/check /opt/resource/check
ADD /check-in-out/in /opt/resource/in
ADD /check-in-out/out /opt/resource/out

RUN chmod +x /opt/resource/*
