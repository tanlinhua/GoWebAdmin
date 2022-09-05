# vscode

## settings.json
``` json
{
  // GO.START
  "go.useLanguageServer": true, //gopls解决vscode提示缓慢
  "[go]": {
    "editor.insertSpaces": false,
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "go.testFlags": ["-v"], // go test 附带参数,解决打印问题
  "go.toolsManagement.autoUpdate": true,
  // GO.END
  "search.followSymlinks": false, //字体大小
  "files.autoSave": "onFocusChange", //失去焦点自动保存
  "editor.formatOnSave": true, //保存代码自动格式化
  "editor.wordSeparators": "`~!@#%^&*()-=+[{]}\\|;:'\",.<>/?", // 执行文字相关的导航或操作时将用作文字分隔符的字符(`~!@#$%^&*()-=+[{]}\\|;:'\",.<>/?)
  "explorer.autoReveal": false, //取消打开文件目录的自动定位跟踪功能
  "explorer.compactFolders": false,
  "liveServer.settings.donotShowInfoMsg": true,
  "workbench.iconTheme": "vscode-icons", //资源文件夹图片主题
  "terminal.integrated.fontSize": 16, //控制终端的字号(以像素为单位)。
  // "terminal.external.osxExec": "iTerm.app",
  // "terminal.integrated.shell.osx": "zsh",
  // "editor.fontFamily": "Cascadia Code", // https://github.com/microsoft/cascadia-code
  "editor.fontFamily": "Fira Code", // https://github.com/tonsky/FiraCode
  "terminal.integrated.fontFamily": "Meslo LG S for Powerline", // https://github.com/powerline/fonts
  "editor.stickyScroll.enabled": true,
  // 一些编辑器颜色配置 START
  "workbench.colorCustomizations": {
    "editor.selectionHighlightBackground": "#0a76a8", //与所选内容具有相同内容的区域颜色
    "editorCursor.foreground": "#01cd78", //编辑器光标颜色
    // "editorBracketMatch.background": "#005c36", //匹配括号的背景色
    // "editor.findMatchBackground": "#e4393c" //搜索项颜色
    // "activityBar.foreground": "#03afff", //活动栏前景色，图标颜色
    // "activityBar.background":"e4393c",//活动栏背景色
    // "editor.background":"#e4393c",//编辑器背景颜色
    // "editor.findMatchHighlightBackground":"#01cd78",//其他搜索项颜色
    // "editor.lineHighlightBackground":"#e4393c",//光标所在行高亮文本的背景颜色
    // "editor.selectionBackground": "#01cd78",//编辑器所选内容的颜色
    // "editor.rangeHighlightBackground": "#e4393c",//突出显示范围的背景颜色，例如 "Quick Open" 和“查找”功能
    // "editorGutter.background": "#01cd78",//编辑器导航线的背景色，导航线包括边缘符号和行号
    // "editorLineNumber.foreground": "#01cd78",//编辑器行号颜色
    // "sideBar.background": "#01cd78",//侧边栏背景色
    // "sideBar.foreground": "#01cd78",//侧边栏前景色
    // "sideBarSectionHeader.background": "#01cd78",//侧边栏节标题的背景颜色
    // "statusBar.background": "#01cd78",//标准状态栏背景色
    // "statusBar.noFolderBackground": "#01cd78",//没有打开文件夹时状态栏的背景色
    // "statusBar.debuggingBackground": "#01cd78",//调试程序时状态栏的背景色
    // "tab.activeBackground": "#01cd78",//	活动选项卡的背景色
    // "tab.activeForeground": "#01cd78",//	活动组中活动选项卡的前景色
    // "tab.inactiveBackground": "#01cd78",// 非活动选项卡的背景色
    // "tab.inactiveForeground": "#01cd78",// 活动组中非活动选项卡的前景色
    // "terminal.foreground": "#01cd78", //终端的字体颜色
  },
  "editor.tokenColorCustomizations": {
    // "functions": "#e99f63", //函数颜色
    // "strings": "#7fb785", //字符串颜色
    "comments": {
      // "fontStyle": "italic", // 注释字体样式 bold italic underline
      // "foreground": "#aaaaaa" // 注释字体颜色
    }
    // "keywords": "#FF0000", //关键字颜色
    // "types": "#FF0000", //类型定义颜色
    // "variables": "#248e85", //变量颜色
    // "numbers": "#a683ff", //数字颜色
  },
  // 一些编辑器颜色配置 END
  "html.format.wrapLineLength": 0, //禁用html单行限制字符
  "workbench.startupEditor": "newUntitledFile",
  "git.autofetch": true,
  "todo-tree.general.tags": [
    "BUG",
    "HACK",
    "FIXME",
    "TODO",
    "XXX",
    "[ ]",
    "[x]"
  ],
  "todo-tree.regex.regex": "(//|#|<!--|;|/\\*|^|^\\s*(-|\\d+.))\\s*($TAGS)",
  "remote.SSH.remotePlatform": {
    "192.168.0.188": "linux"
  },
  "projectManager.sortList": "Recent",
  "editor.fontSize": 17,
  "[typescript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[vue]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[javascript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
}
```