# # syntax=docker/dockerfile:1.4
# FROM golang:1.22.1-bookworm AS build
# RUN --mount=type=cache,target=/go/pkg/mod \
#     CGO_ENABLED=0 \
#     go install -ldflags "-s -w" -trimpath google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0

# FROM scratch
# COPY --from=build --link /etc/passwd /etc/passwd
# COPY --from=build --link --chown=root:root /go/bin/protoc-gen-go .
# USER nobody
# ENTRYPOINT [ "/protoc-gen-go" ]

# syntax=docker/dockerfile:1.6
FROM golang as builder-local
ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0
WORKDIR /app
COPY buf.plugin.yaml ./
COPY ../../go.mod ./
COPY ../../go.sum ./
RUN go mod download
COPY ../../gen/redact ./gen/redact
COPY ../../cmd ./cmd
# COPY *.go ./

RUN go build -o bin/protoc-gen-redact ./cmd/protoc-gen-go-redact

FROM golang:1.21-bullseye AS plugin-builder
COPY --from=builder-local /app/cmd/protoc-gen-go-redact/internal/template.tmpl /go/bin/
COPY --from=builder-local /app/bin /go/bin/
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
#     go install -ldflags "-s -w" -trimpath github.com/jasonkolodziej/protoc-gen-go-redact@v1.1.0

# When building a Docker image on a host that does not match linux/amd64 (such as an M1),
# go install will put the binary in $GOPATH/bin/$GOOS_$GOARCH/. The mv command copies
# the binary to /go/bin so subsequent steps do not fail when copying from the builder.
# RUN mv /go/bin/linux_amd64/* /go/bin || true

FROM scratch
COPY --from=plugin-builder --link /etc/passwd /etc/passwd
COPY --from=plugin-builder /go/bin/ /
USER nobody
ENTRYPOINT [ "/protoc-gen-go-redact" ]
