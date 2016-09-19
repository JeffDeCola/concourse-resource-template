FROM concourse/buildroot:base

COPY resource-template /

COPY scripts/ /opt/resource/

RUN chmod +x /opt/resource/*