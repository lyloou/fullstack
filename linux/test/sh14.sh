#!/bin/bash
# Program:
#   test 'while'
# History:
#   2017/04/07 Lou  First release
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PATH

until [ "$yn" = "yes" -o "$yn" = "YES" ]
do
  read -p "Please input yes/YES to stop this program: " yn
done
echo "OK!"
