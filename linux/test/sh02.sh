#!/bin/bash
# Program:
#   User inputs hist name and last name. Program shows his full name.
# History:
#   2017/04/04 Lou First release
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PATH

echo "please input your last name:"
read firstname
echo "please input your first name:"
read lastname
echo "your full name is $firstname $lastname"
