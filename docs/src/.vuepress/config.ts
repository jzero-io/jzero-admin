// @ts-ignore
import { defineUserConfig } from "vuepress";
import theme from "./theme.js";

export default defineUserConfig({
  base: "/",

  locales: {
    "/": {
      lang: "zh-CN",
      title: "jzero-admin",
      description: "jzero-admin 文档",
    },
    "/en/": {
      lang: "en-US",
      title: "jzero-admin",
      description: "docs for jzero-admin",
    },
  },

  theme,

  // 和 PWA 一起启用
  // shouldPrefetch: false,
});
