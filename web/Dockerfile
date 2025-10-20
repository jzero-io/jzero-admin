FROM node:20.12.2-bullseye-slim AS builder

WORKDIR /jzero-admin

COPY ./ ./

RUN --mount=type=cache,target=/node_modules npm install -g pnpm@8.15.6 \
  && pnpm i \
  && pnpm build

FROM nginx:alpine

WORKDIR /usr/share/nginx/

COPY --from=builder /jzero-admin/dist ./html
