#!/bin/bash
# Program:
#   test 'while'
# History:
#   2017/04/07 Lou  First release
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PATH

while [ "$yn" != "yes" -a "$yn" != "YES" ]
do
  read -p "Please input yes to stop this program: " yn
done
echo "OK!"
