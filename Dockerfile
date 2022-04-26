FROM golang:1.18-bullseye AS  build-env
RUN apt install -y git

ENV CGO_ENABLED=0, GO111MODULE=on
WORKDIR /go/src/github.com/zv0n/csi-webdav

ADD . /go/src/github.com/zv0n/csi-webdav

RUN go mod download
RUN export BUILD_TIME=`date -R` && \
    export VERSION=`cat /go/src/github.com/zv0n/csi-webdav/version.txt 2&> /dev/null` && \
    go build -o /csi-webdav -ldflags "-X 'github.com/zv0n/csi-webdav/pkg/webdav.BuildTime=${BUILD_TIME}' -X 'github.com/zv0n/csi-webdav/pkg/webdav.Version=${VERSION}'" github.com/zv0n/csi-webdav/cmd/csi-webdav

FROM debian:bullseye
COPY --from=build-env /csi-webdav /bin/csi-webdav
ENTRYPOINT ["/bin/csi-webdav"]
CMD [""]
