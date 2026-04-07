// @ts-ignore
import { sidebar } from "vuepress-theme-hope";

export const zhSidebarConfig = sidebar({
    "/": [
        "",
        {
            text: "Sidebar_1",
            icon: "vscode-icons:file-type-codekit",
            prefix: "sidebar_1/",
            children: "structure",
            collapsible: true,
            expanded: true,
        },
    ]
});
