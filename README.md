#### 介绍

[AutoPy-Android](https://pypi.org/project/AutoPy-Android/)库的Go语言接口

AuoyPy Lite是安卓上基于无障碍服务实现模拟点击、滑动屏幕、根据ID/文本点击等功能的应用，酷安地址: https://www.coolapk.com/apk/chenlong.chenlong.autopylite

#### 安装

`go get -u github.com/cooolr/autopy`

#### import

`import "github.com/cooolr/autopy"`

#### 使用

autopy Go版本目前复刻接口如下

1. 点击  `autopy.Click(500, 500)`

2. 滑动  `autopy.Swipe(100, 0, 100, 600, 1000)`

3. 返回  `autopy.BACK()`

4. 主页  `auotpy.HOME()`

5. 菜单  `autopy.RECENTS()`

6. 截图  `autopy.Capturer()`

7. 等待  `autopy.Sleep(1)`

8. 根据ID点击 `autopy.ClickByID("arg")`

9. 根据文本点击 `autopy.ClickByText("今日头条")`

