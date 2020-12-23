FROM registry.redhat.io/ubi8/ubi-minimal

LABEL version="1.0" \
      description="Minimal GO server" \
      creationDate="2020-12-23" \
      updatedDate="2020-12-23"

RUN mkdir /opt/go-server
COPY ./app /opt/go-server/

RUN chgrp -R 0 /opt/go-server && chmod -R g=u /opt/go-server

EXPOSE 8080
USER 1001

CMD /opt/go-server/app