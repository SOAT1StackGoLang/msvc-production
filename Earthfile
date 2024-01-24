VERSION 0.7
all:
    BUILD --platform=linux/amd64 --platform=linux/arm64 +msvc
    BUILD --platform=linux/amd64 --platform=linux/arm64 +debug
amd64:
    BUILD --platform=linux/amd64 +msvc
    BUILD --platform=linux/amd64 +debug
arm64:
    BUILD --platform=linux/arm64 +msvc
    BUILD --platform=linux/arm64 +debug
file:
    LOCALLY
    SAVE ARTIFACT ./
deps:
    FROM golang:alpine
    ARG GITHUB_TOKEN
    ENV GOPRIVATE=github.com/SOAT1StackGoLang
    RUN echo "machine github.com login $GITHUB_TOKEN password x-oauth-basic" > ~/.netrc
    WORKDIR /build
    COPY +file/* ./
    RUN ls -althR
    RUN apk add --no-cache git
    #RUN git config --global url."ssh://git@github.com/".insteadOf https://github.com/
    RUN git config --global credential.helper 'store --file=/build/.git-credentials'
    RUN echo "https://$GITHUB_TOKEN:x-oauth-basic@github.com" > /build/.git-credentials
    RUN go mod tidy
    RUN go mod download



compile:
    FROM +deps
    ARG GITHUB_TOKEN
    ENV GOPRIVATE=github.com/SOAT1StackGoLang
    RUN echo "machine github.com login $GITHUB_TOKEN password x-oauth-basic" > ~/.netrcs
    ARG GOOS=linux
    ARG GOARCH=amd64
    ARG VARIANT
    RUN ls -alth && pwd
    RUN GOARM=${VARIANT#v} CGO_ENABLED=0 go build \
        -installsuffix 'static' \
        -o compile/app cmd/server/*.go
    SAVE ARTIFACT compile/app /app AS LOCAL compile/app
#--ldflags "-X 'msvc.Version=v0.0.3' -X 'msvc.BuildTime=$(date "+%H:%M:%S--%d/%m/%Y")' -X 'msvc.GitCommit=$(git rev-parse --short HEAD)'" \

msvc:
    ARG EARTHLY_TARGET_TAG_DOCKER
    ARG EARTHLY_GIT_SHORT_HASH
    ARG TARGETPLATFORM
    ARG TARGETARCH
    ARG TARGETVARIANT
    FROM --platform=$TARGETPLATFORM gcr.io/distroless/static
    ## enable multiple debug version with shell
    #FROM --platform=$TARGETPLATFORM gcr.io/distroless/static:debug
    #FROM --platform=$TARGETPLATFORM alpine:latest
    LABEL org.opencontainers.image.source=https://github.com/soat1stackgolang/msvc-production
    LABEL org.opencontainers.image.description="Main App Image only have the app binary and nothing else"
    WORKDIR /
    COPY \
        --platform=linux/amd64 \
        (+compile/app --GOARCH=$TARGETARCH --VARIANT=$TARGETVARIANT) /app
    ENV GIN_MODE=release
    ENTRYPOINT ["/app"]
    EXPOSE 8000
    SAVE IMAGE --push ghcr.io/soat1stackgolang/msvc-production:msvc-$EARTHLY_TARGET_TAG_DOCKER
    SAVE IMAGE --push ghcr.io/soat1stackgolang/msvc-production:msvc-$EARTHLY_GIT_SHORT_HASH

debug:
    ARG EARTHLY_TARGET_TAG_DOCKER
    ARG EARTHLY_GIT_SHORT_HASH
    ARG TARGETPLATFORM
    ARG TARGETARCH
    ARG TARGETVARIANT
    #FROM --platform=$TARGETPLATFORM gcr.io/distroless/static
    ## enable multiple debug version with shell
    #FROM --platform=$TARGETPLATFORM gcr.io/distroless/static:debug
    FROM --platform=$TARGETPLATFORM alpine:latest
    LABEL org.opencontainers.image.source=https://github.com/soat1stackgolang/msvc-production
    LABEL org.opencontainers.image.description="Debug Image will have all binaries"
    WORKDIR /
    COPY \
        --platform=linux/amd64 \
        (+compile/app --GOARCH=$TARGETARCH --VARIANT=$TARGETVARIANT) /app
    ENV GIN_MODE=release
    CMD /app
    EXPOSE 8000
    SAVE IMAGE --push ghcr.io/soat1stackgolang/msvc-production:debug-$EARTHLY_TARGET_TAG_DOCKER
    SAVE IMAGE --push ghcr.io/soat1stackgolang/msvc-production:debug-$EARTHLY_GIT_SHORT_HASH
