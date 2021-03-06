
## alias
```sh
alias cdt='cd ~/t'
alias cdw='cd ~/w'
alias e.='caja .'
alias egrep='egrep --color=auto'
alias fgrep='fgrep --color=auto'
alias gl='git lg'
alias grep='grep --color=auto'
alias l='ls -CF'
alias la='ls -A'
alias ll='ls -alF'
alias ls='ls --color=auto'
```

## java
```sh
#set oracle jdk environment
export JAVA_HOME=/usr/lib/jvm/jdk1.8.0_121  ## 这里要注意目录要换成自己解
压的jdk 目录 
export JRE_HOME=${JAVA_HOME}/jre   
export CLASSPATH=.:${JAVA_HOME}/lib:${JRE_HOME}/lib   
export ${JAVA_HOME}/bin:$PATH
```

## go
```sh
export GOROOT=$HOME/c/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=$HOME/w/go
export PATH=${GOPATH}/bin:$PATH
```

## git
### to set
```sh
git config --global user.name lilou
git config --global user.email lilou@lyloou.com
git config --global core.autocrlf  false
git config --global alias.lg "log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"
git config --global http.proxy 'socks5://127.0.0.1:1080'
git config --global https.proxy 'socks5://127.0.0.1:1080'
```

### result
```sh
[user]
        email = lilou@lyloou.com
        name = lilou
[alias]
        lg = log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit
[core]        
        autocrlf = false
```