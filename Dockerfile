# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-entur"
LABEL REPO="https://github.com/larsw/entur"

ENV PROJPATH=/go/src/github.com/larsw/entur

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/larsw/entur
WORKDIR /go/src/github.com/larsw/entur

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/larsw/entur"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/entur/bin

WORKDIR /opt/entur/bin

COPY --from=build-stage /go/src/github.com/larsw/entur/bin/entur /opt/entur/bin/
RUN chmod +x /opt/entur/bin/entur

# Create appuser
RUN adduser -D -g '' entur
USER entur

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/entur/bin/entur"]
