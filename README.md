# jzero-admin

[![Build Status](https://img.shields.io/github/actions/workflow/status/jzero-io/jzero-admin/web.yaml?branch=main&label=web&logo=github&style=flat-square)](https://github.com/jzero-io/jzero-admin/actions?query=workflow%3Aweb)
[![Build Status](https://img.shields.io/github/actions/workflow/status/jzero-io/jzero-admin/web.yaml?branch=main&label=server&logo=github&style=flat-square)](https://github.com/jzero-io/jzero-admin/actions?query=workflow%3Aserver)
[![GitHub release](https://img.shields.io/github/release/jzero-io/jzero-admin.svg?style=flat-square)](https://github.com/jzero-io/jzero-admin/releases/latest)
[![GitHub package version](https://img.shields.io/github/v/release/jzero-io/jzero-admin?include_prereleases&sort=semver&label=Docker%20Image%20version)](https://github.com/jzero-io/jzero-admin/pkgs/container/jzero)

<p align="center">
<img align="center" width="150px" src="https://oss.jaronnie.com/jzero-admin.jpg">
</p>

基于 [jzero](https://github.com/jzero-io/jzero) 脚手架搭建后台服务, 基于 [soybean](https://github.com/soybeanjs/soybean-admin) 搭建前端服务的下一代后台管理系统.

## 演示地址

* [web 服务部署仓库地址](https://github.com/jaronnie/jzero-admin-deploy-web)
* [server 服务部署仓库地址](https://github.com/jaronnie/jzero-admin-deploy-server)

### 部署在 [vercel](https://vercel.com) 平台

[demo 演示地址](https://admin.jzero.io)

### 部署在阿里云函数计算

[demo 演示地址](https://jzero-admin.jaronnie.com)

## 本地一键部署

```shell
git clone https://github.com/jzero-io/jzero-admin.git
cd jzero-admin/deploy/docker-compose
docker-compose up
# 修改源码后重新编译
# docker-compose up --build
```

## 感谢

* 使用在线 postgres 实例 [neon](https://neon.tech/)
