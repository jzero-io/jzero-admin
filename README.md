# jzero-admin

[![Build Status](https://img.shields.io/github/actions/workflow/status/jzero-io/jzero-admin/web.yaml?branch=main&label=web&logo=github&style=flat-square)](https://github.com/jzero-io/jzero-admin/actions?query=workflow%3Aweb)
[![Build Status](https://img.shields.io/github/actions/workflow/status/jzero-io/jzero-admin/web.yaml?branch=main&label=server&logo=github&style=flat-square)](https://github.com/jzero-io/jzero-admin/actions?query=workflow%3Aserver)
[![GitHub release](https://img.shields.io/github/release/jzero-io/jzero-admin.svg?style=flat-square)](https://github.com/jzero-io/jzero-admin/releases/latest)
[![GitHub package version](https://img.shields.io/github/v/release/jzero-io/jzero-admin?include_prereleases&sort=semver&label=Docker%20Image%20version)](https://github.com/jzero-io/jzero-admin/pkgs/container/jzero)

<p align="center">
<img align="center" width="150px" src="https://oss.jaronnie.com/jzero-admin.jpg">
</p>

基于 [jzero](https://github.com/jzero-io/jzero) 脚手架搭建后台服务, 基于 [soybean](https://github.com/soybeanjs/soybean-admin) 搭建前端服务的下一代后台管理系统.

## demo 演示

前后端均部署在 [vercel](https://vercel.com)

* [web 服务部署仓库地址](https://github.com/jaronnie/jzero-admin-deploy-web)
* [server 服务部署仓库地址](https://github.com/jaronnie/jzero-admin-deploy-server)

[demo 演示地址](https://demo.jzero-admin.jaronnie.com)

## 感谢

* 使用在线 mysql 实例 [sqlpub](https://sqlpub.com)

## 路线图

- [ ] 替换 apifox 接口为 jzero 框架接口
  - [ ] /auth/login
  - [ ] /auth/getUserInfo
  - [ ] /auth/refreshToken
  - [ ] /auth/error
  - [ ] /route/getConstantRoutes
  - [ ] /route/getUserRoutes
  - [ ] /route/isRouteExist
  - [ ] /systemManage/getUserList
  - [ ] /systemManage/getRoleList
  - [ ] /systemManage/getAllRoles
  - [ ] /systemManage/getAllPages
  - [ ] /systemManage/getMenuList/v2
  - [ ] /systemManage/getMenuTree