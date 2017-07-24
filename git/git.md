- [] The fox jumps the old lazy dog.
- [] The fox jumps the old lazy dog.

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
