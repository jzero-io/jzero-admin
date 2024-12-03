FROM registry.cn-hangzhou.aliyuncs.com/jaronnie/node:20.12.2-bullseye-slim as builder

WORKDIR /jzero-admin

COPY ./ ./

RUN --mount=type=cache,target=/node_modules npm config set registry https://registry.npmmirror.com \
  && npm install -g pnpm@8.15.6 \
  && rm -rf package-lock.json \
  && rm -rf pnpm-lock.yaml \
  && pnpm i \
  && pnpm build

FROM registry.cn-hangzhou.aliyuncs.com/jaronnie/nginx:1.26-alpine

WORKDIR /usr/share/nginx/

COPY --from=builder /jzero-admin/dist ./html
