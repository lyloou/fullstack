## How to set socks5 proxy in the terminal (ubuntu)
```sh
# install proxychains
sudo apt install proxychains
sudo proxychains apt-get update

# now you can config your proxy in /etc/proxychains.conf
socks5 127.0.0.1 1080
```

## 卸载软件
1. 查找安装名称：`dpkg -l | grep sogoupinyin`
2. 卸载`sudo apt-get remove package_name`, 具体输入`apt-get`命令查看

## 安装网易云音乐
1. 从官网下载deb安装文件
2. 执行下列命令
```sh
$ sudo dpkg -i netease*.dbg
$ sudo apt -f install
```

## install wine
https://wiki.winehq.org/Ubuntu


## 安装搜狗输入法
1. 安装搜狗输入法之前，先移除fcitx (有冲突)
```sh
$ sudo apt remove fcitx*
$ sudo apt autoremove
```
2. 从官网下载deb文件: http://pinyin.sogou.com/linux/?r=pinyin
3. 执行安装命令
```sh
$ sudo dpkg -i sogoupinyin*.deb
$ sudo apt -f install
```
4. 完成上面的步骤之后，重启电脑
5. 查看状态栏，如果还没有显示搜狗输入法,选择配置，添加输入法，选择搜狗输入法，确定。



## history之后执行指定行的命令
> `$ history // 查看命令历史`  
> `$ !334  //表示执行第334行的命令`


## 在命令行中用默认程序打开文件
`xdg-open { file | URL }`
> [Ubuntu下用命令行快速打开各类型文件(转)-bough22-ChinaUnix博客](http://blog.chinaunix.net/uid-27025492-id-3376626.html)


## 安装WenQuanYi Zen Hei 字体
1. download file: wqy-zenhei-0.9.45.deb
2. install: dpkg -i wqy-zenhei-0.9.45.deb


## Terminator
- [Terminator – Multiple GNOME terminals in one window | Ubuntu Geek](http://www.ubuntugeek.com/terminator-multiple-gnome-terminals-in-one-window.html)
- [安装 Terminator：一个支持多终端的终端-软件 ◆ 分享|Linux.中国-开源社区](https://linux.cn/article-2978-1.html)
> #sudo apt-get install terminator

## 去除 输入密码以解锁密码环
- 终端输入`seahorse`
- 右键删除 `登录`
- 退出 chrome
- 提示输入密码时，不输入任何内容，直接按下一步。
> [解决打开Chrome出现 输入密码以解锁您的登录密钥环 - Android/Linux的专栏 - 博客频道 - CSDN.NET](http://blog.csdn.net/kangear/article/details/20789451)

## Ubuntu 系统强制关闭进程。
> $ps -aux | grep [应用名]  # 抓取指定应用的进程信息，几下 应用的pid
> $kill -9 [应用的pid]

## [Top 6 Things To Do After Installing Ubuntu 16.04](https://www.youtube.com/watch?v=ZcpWofRAs-A)
1. Be sure you are up to date:
  - dash -> input `update` -> 去更新软件吧。
2. Install VLC player
  - ubuntu软件管理工具 -> input `vlc`  
3. Install Chrome
  - download `.deb` file  
  - run `sudo dpkg -i file_name.deb`
  - run `sudo apt-get -f install`
4. Choose default application
  - dash ->  input `application`
5. Install unity tweak tool
  - ubuntu软件管理工具 -> input `unity tweak`  
6. Install your theme
  - open your browser -> youtube.com -> search by key `macbuntu`  
