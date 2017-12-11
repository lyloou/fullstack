- [] The fox jumps the old lazy dog.
- [] The fox jumps the old lazy dog.

## 行尾结束符自动转换问题
如果依照文件的md5来作为标识，那么提交代码到远程仓库，文件的行尾结束符却自动发生了改变，那么之前的标识就发挥不到作用
（这样的bug很难找），这时候就需要取消git的行尾结束符自动转换功能；
```sh
$ git config --global core.autocrlf false
```
- [Git - 配置 Git](https://git-scm.com/book/zh/v1/%E8%87%AA%E5%AE%9A%E4%B9%89-Git-%E9%85%8D%E7%BD%AE-Git#格式化与空白)


## remove untracked files
```bash
$ git clean -d -f
$ git clean --help
```
> [branch - How to remove local (untracked) files from the current Git working tree? - Stack Overflow](https://stackoverflow.com/questions/61212/how-to-remove-local-untracked-files-from-the-current-git-working-tree)


## How to make Git “forget” about a file that was tracked but is now in .gitignore?
.gitignore will prevent untracked files from being added (without an add -f) to the set of files tracked by git, however git will continue to track any files that are already being tracked.
To stop tracking a file you need to remove it from the index. This can be achieved with this command.
    `git rm --cached <file>`
The removal of the file from the head revision will happen on the next commit.
> - https://stackoverflow.com/questions/1274057/how-to-make-git-forget-about-a-file-that-was-tracked-but-is-now-in-gitignore

## [How do I push a new local branch to a remote Git repository and track it too?](https://stackoverflow.com/questions/2765421/how-do-i-push-a-new-local-branch-to-a-remote-git-repository-and-track-it-too)
```sh
git push -u origin develop
```
## `.gitignore`文件中添加文件时 要注意`package.json`和`/package.json`这两种表示方式的区别：
  - 前者是所有目录中的 `package.json`文件都被忽略；
  - 后者是只有当前目录的 `package.json`文件被忽略；

## delete a git branch both locally and remotely.
- `git branch -d {the_local_branch}`
- `git push --delete {the_remote_branch}`

> [How do I delete a Git branch both locally and remotely? - Stack Overflow](https://stackoverflow.com/questions/2003505/how-do-i-delete-a-git-branch-both-locally-and-remotely)

## git reset 后丢弃远程的提交
```bash
git log # 获取log的某次提交commit id
git reset --hard 0301382 # 回退到0301382
git push --force # 强制推送到服务器端
```

- [github - git 怎样删除远程仓库的某次错误提交？ - SegmentFault](https://segmentfault.com/q/1010000002898735)

## 相对远程分支的内容变更情况

`git status`：查看本地未传送的提交次数；
```bash
$ git status
On branch master
Your branch is ahead of 'origin/master' by 1 commit.
  (use "git push" to publish your local commits)
```

`git cherry -v`： 查看未传送提交的描述/说明；
```bash
$ git cherry -v
+ d105e9bb16e606892863e6d8be3931174eedc601 update android.md
```

`git log master ^origin/master`：查看未传送提交的详细内容；
```bash
commit d105e9bb16e606892863e6d8be3931174eedc601 (HEAD -> master)
Author: Lou <lyloou6@gmail.com>
Date:   Fri Jun 2 09:47:19 2017 +0800

    update android.md
```

> 由第三点，其实可以想到，将远程和本地分支位置调换，即`git log origin/master ^origin`，
就可以查看远程库比本地库多的内容了。
不过，不过，不过，得先执行`git fetch origin master` 命令，将远程仓库的 commit 内容同步
到本地库。fetch 和 pull 的区别是： fetch 只同步远程库的 commit 信息（log信息），
但不会将文件同步到本地（没有执行merge）。pull 则会将文件也同步到本地。
使用 git log 的时候，提供参数 `-p` 可以查看更详细信息。

pull = fetch + merge

参考资料：
- [git查看远程版本库和本地库的差异 | 星辰](http://blog.kainaodong.com/?p=12)
- [git pull 和 git fetch 有什么区别？](https://ruby-china.org/topics/15729)


## [Which remote URL should I use?](https://help.github.com/articles/which-remote-url-should-i-use/)
- Cloning with HTTPS URLs **(recommended)**
- Cloning with SSH URLs
- Cloning with Subversion



## change origin's url
> Instead of removing and re-adding,  you can do this:
  `git remote set-url orign git:new.url.here`
  See this question: [Change the URI (URL) for a remote Git repository - Stack Overflow](http://stackoverflow.com/questions/16330404/how-to-remove-remote-origin-from-git-repo/16330439)
- [how to remove remote origin from git repo](http://stackoverflow.com/questions/16330404/how-to-remove-remote-origin-from-git-repo/16330439)

## generate SSH keys
```
ssh-keygen -t rsa -C "lyloou@qq.com"
```

## config ssh login
```
Step 1 -

Create SSH keys on your linux system using below command

ssh-keygen -t rsa -b 4096 -C "your_email"
It will ask for passphrase and file name (default will be ~/.ssh/id_rsa, ~/.ssh/id_rsa.pub)

Step 2 -

Once files created add public key id_rsa.pub to github account ssh section.

Step 3 -

On your machine add private key id_rsa to ssh-agent using below command

ssh-add ~/.ssh/id_rsa 
Step 4 -

Now add remote url git@github.com:user_name/repo_name.git to your local git repo using below command.

git remote remove origin
git remote add origin git@github.com:user_name/repo_name.git
Thats it.
```

## git stash
> Git 还提供了一个 stash 功能，可以把当前工作状态“储藏”起来，等以后恢复现场后继续工作。
> [Bug分支](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/00137602359178794d966923e5c4134bc8bf98dfb03aea3000)

## git 设置别名：
```sh
git config --global alias.lg "log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"
```
> - [配置别名 - 廖雪峰的官方网站](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/001375234012342f90be1fc4d81446c967bbdc19e7c03d3000)

## git 移除文件夹的版本控制
```sh
git rm -r -n --cached "bin/" //-n：加上这个参数，执行命令时，是不会删除任何文件，而是展示此命令要删除的文件列表预览。
git rm -r --cached  "bin/"      //最终执行命令.
git commit -m" remove bin folder all file out of control"    //提交
git push origin master   //提交到远程服务器
```
> - [git如何移除某文件夹的版本控制](https://my.oschina.net/dlpinghailinfeng/blog/388606)
> - [Git 常用的不常用笔记](http://leoray.leanote.com/post/git)




## git 设置 socks5 代理
```sh
git config --global http.proxy 'socks5://127.0.0.1:1080'
git config --global https.proxy 'socks5://127.0.0.1:1080'
```
>  [git 设置和取消代理](https://gist.github.com/laispace/666dd7b27e9116faece6)


## 当心Git的自作聪明
`warning: LF will be replaced by CRLF in xx.txt`

这一替换，会更改文件的hash值，所以如果在本地先生成文件，然后上传到服务器，另一端下载然后比对hash值，
因为这一转换，会导致文件的hash值发生变化。这是个很隐蔽的bug，当心。

解决办法是：
```sh
git config --global core.autocrlf  false
```
https://stackoverflow.com/questions/5834014/lf-will-be-replaced-by-crlf-in-git-what-is-that-and-is-it-important
