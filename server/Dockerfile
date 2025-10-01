FROM --platform=$BUILDPLATFORM ghcr.io/jzero-io/jzero:latest as builder

ARG TARGETARCH
ARG LDFLAGS

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /usr/local/go/src/app

COPY ./ ./

RUN jzero serverless build

RUN --mount=type=cache,target=/go/pkg CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -a -ldflags="$LDFLAGS" -o /dist/app main.go \
    && cp -r etc /dist \
    && mkdir -p /dist/desc && cp -r desc/swagger /dist/desc && cp -r desc/sql_migration /dist/desc \
    && for plugin in $(ls -d plugins/*/); do \
         mkdir -p /dist/$plugin && cp -r $plugin/etc /dist/$plugin; \
         [ -d $plugin/desc/sql_migration ] && mkdir -p /dist/$plugin/desc && cp -r $plugin/desc/sql_migration /dist/$plugin/desc || true; \
       done

FROM --platform=$TARGETPLATFORM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /dist

COPY --from=builder /dist .

EXPOSE 8001

CMD ["./app", "server"]