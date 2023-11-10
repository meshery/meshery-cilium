FROM golang:1.21 as build-env
ARG VERSION
ARG GIT_COMMITSHA

WORKDIR /github.com/meshery/meshery-cilium
COPY go.mod go.sum ./
RUN go mod download
COPY main.go main.go
COPY internal/ internal/
COPY cilium/ cilium/
COPY build/ build/
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -ldflags="-w -s -X main.version=$VERSION -X main.gitsha=$GIT_COMMITSHA" -a -o meshery-cilium main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/nodejs:16
ENV DISTRO="debian"
ENV SERVICE_ADDR="meshery-cilium"
ENV MESHERY_SERVER="http://meshery:9081"
WORKDIR /
COPY templates/ ./templates
COPY --from=build-env /github.com/meshery/meshery-cilium/meshery-cilium .
ENTRYPOINT ["./meshery-cilium"]
