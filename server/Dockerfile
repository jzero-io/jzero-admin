FROM --platform=$BUILDPLATFORM registry.cn-hangzhou.aliyuncs.com/jaronnie/jzero:latest as builder

ARG TARGETARCH
ARG LDFLAGS

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /usr/local/go/src/app

COPY ./ ./

RUN --mount=type=cache,target=/go/pkg CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -a -ldflags="$LDFLAGS" -o /dist/app main.go \
    && jzero gen swagger \
    && cp -r etc /dist/etc \
    && mkdir -p /dist/desc && cp -r desc/swagger /dist/desc


FROM --platform=$TARGETPLATFORM registry.cn-hangzhou.aliyuncs.com/jaronnie/alpine:latest

WORKDIR /dist

COPY --from=builder /dist .

EXPOSE 8001

CMD ["./app", "server"]