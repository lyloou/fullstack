## sed
https://coolshell.cn/articles/9104.html

### 在某目录及其子目录下所有文件的最前面添加几行文字
```sh
grep -rl '' tmpdir\ | xargs sed -i "1 i hi 你好吗\n 你知道我是谁吗\n 是的，是我\n"
```
- [linux - Find and replace with sed in directory and sub directories - Stack Overflow](https://stackoverflow.com/questions/6758963/find-and-replace-with-sed-in-directory-and-sub-directories)

## nohub
用途：不挂断地运行命令。

语法：nohup Command [ Arg … ] [　& ]

描述：nohup 命令运行由 Command 参数和任何相关的 Arg 参数指定的命令，忽略所有挂断（SIGHUP）信号。在注销后使用 nohup 命令运行后台中的程序。要运行后台中的 nohup 命令，添加 & （ 表示”and”的符号）到命令的尾部。
http://www.cnblogs.com/allenblogs/archive/2011/05/19/2051136.html


## tail
查看实时日志： tail -fn 100 log_file_name.out
```sh
tail --help
-f: `--follow[HOW]` Output appended data as the file grows;
-n: `--lines=[+]NUM` Output the last NUM lines, instead of the last 10;
```

lsof -i 8088


## ps
查看运行的进程
`ps -aux | grep vim`
