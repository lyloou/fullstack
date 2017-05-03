| git stash
> Git 还提供了一个 stash 功能，可以把当前工作状态“储藏”起来，等以后恢复现场后继续工作。
> [Bug分支](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/00137602359178794d966923e5c4134bc8bf98dfb03aea3000)

| git 设置别名：
```sh
git config --global alias.lg "log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"
```
> - [配置别名 - 廖雪峰的官方网站](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/001375234012342f90be1fc4d81446c967bbdc19e7c03d3000)

| git 移除文件夹的版本控制
```sh
git rm -r -n --cached "bin/" //-n：加上这个参数，执行命令时，是不会删除任何文件，而是展示此命令要删除的文件列表预览。
git rm -r --cached  "bin/"      //最终执行命令.
git commit -m" remove bin folder all file out of control"    //提交
git push origin master   //提交到远程服务器
```
> - [git如何移除某文件夹的版本控制](https://my.oschina.net/dlpinghailinfeng/blog/388606)
> - [Git 常用的不常用笔记](http://leoray.leanote.com/post/git)




| git 设置 socks5 代理
```sh
git config --global http.proxy 'socks5://127.0.0.1:1080'
git config --global https.proxy 'socks5://127.0.0.1:1080'
```
>  [git 设置和取消代理](https://gist.github.com/laispace/666dd7b27e9116faece6)
