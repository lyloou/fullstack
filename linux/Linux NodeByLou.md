

| 挂载点是什么？一定是目录。 --p235

| 开机挂载
`/etc/fstab (file system table)`就是将我们利用 mount 命令进行挂载时，
将所有的参数写入到这个文件中就可以了。

| 使用`hdparm` 在sata上做测试：
```sh
hdparm -Tt /dev/sda
hdparm -Tt /dev/sdb
```

| 可以通过Label的方式进行挂载： --p231
```sh
mount -L "vbird_logical" /mnt/ccc
```

| e.g., 挂载一个iso文件到`/mnt/ccc`
```sh
mount xxxx.iso /mnt/ccc
```
一般这种iso文件都是只读形式的，因此可能会有下面这种提示：
`mount: /dev/loop0 is write-protected, mounting read-only`

| 进入到挂载目录中，进行卸载操作是不成功的。
`umount: /media/cdrome: device is busy`
解决办法是退出挂载目录，重新`umount`
--p230

| 磁盘挂载： mount

| 磁盘检验：fsck, badblocks

| 磁盘分区完成后，进行格式化操作：`mkfs`

| fdisk 没有办法处理大于2TB以上的磁盘分区。（可以使用parted这个命令） --p223

| fdisk
需要以管理员身份运行
```sh
fdisk /dev/sda
fdisk /dev/sdb # 注意不需要带具体数值
```

| 要制作连接文件必须要使用`ln`这个命令；--p215
```sh
ln [-sf] 源文件 目标文件
```

| `symbolic link`符号连接；（类似于windows上的快捷方式）

| hard link连接有限制：
- 不能跨文件系统；
- 不能连接到目录；

| 注意硬链接和软连接的区别；--p213

| 整个filesystem其实都是VFS（Virtual Filesystem Switch）进行管理；

| 根目录的上一级就是他自己。 `ls -ild / /. /..`，通过这个命令发现`inode`都是2；

| Ext2文件系统 --p197

## 第8章

| 通过find来按时间属性查找：
- mtime
- ctime
- atime

| 更新locate数据库：
`updatedb`

| 按文件名查找：
- `whereis`
- `locate`

| 查找`commond`的路径:
`which -a ifconfig`

| root用户身份的umask权限默认是：022
  一般用户身份的umask权限默认是：002 --p183

| 设置umask: `umask 002`
- 具体运算规则，是默认的权限减去umask数值。

| 新建文件的默认权限：-rw-rw-rw-
  新建文件夹的默认权限：drwxrwxrwx
注意：具体的生成结果还得依赖于umask的值(拿掉umask的那一部分)。

| `ll; ls`:
`;`分号表示连续命令的执行；

| `ll`命令是`ls -l`的缩写。--p179

| `touch`可以只用来修改时间，使用参数`-a`。 --p179
可以指定时间，使用`-d`参数。
还可以用来创建空文件。

| tail命令：显示文件后面几行
`-f`参数表示，持续检测文件的内容(实时性)。按ctrl-c停止检测。

| man 命令就是通过调用 less 来显示说明文件的内容的。

| `nl`显示文件的时候，顺便显示行号。

| `rename`更改大量文件的文件名

| `mv`命令
移动文件 & 重命名文件

| `rm`命令
删除文件/目录
- `-r` 遍历删除参数。（慎用）

| 你是要站在巨人的肩膀上瞭望的。

| `mv`命令
用途：
- 移动文件/文件夹；
- 用来重命名；

| `cp`命令
用途：
- 复制文件；
加上参数`-a` 讲文件的所有特性一起复制过来。（包括属性/权限）
这个功能常用来备份文件。
注意：复制给别人的文件要注意到文件的权限（否则别人无法修订）。
要复制连接文件的属性，就得使用 -d 的参数了。

- 创建链接（快捷方式）


# 《鸟哥的Linux私房菜》读书笔记

=============

## 错误
| p181
> 修改属性：chown -R dmtsai:users/tmp/chapter7_1
  修改属性：chown -R dmtsai:users /tmp/chapter7_1 // 少了一个空格

> 答： mkfs-tvfat/dev/hdc6
  答： mkfs -t vfat /dev/hdc6  
