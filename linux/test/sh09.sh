#!/bin/bash
# Program:
#  Check $1 is equal to 'hello' 
# History:
#   2017/04/06 Lou First release
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PATH

if [ "$1" = "hello" ];then
  echo "Hello, how are you ?"
elif [ "$1" = "" ]; then
  echo  "You MUST input parameters, ex>{$0 someword}"
else 
  echo "The only parameter is 'helo' ex>{$0 hello}"
fi
