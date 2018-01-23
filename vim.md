```vim
" Vim-Plug {{{
call plug#begin('~/.vim/plugged')
Plug 'junegunn/vim-easy-align'
Plug 'https://github.com/junegunn/vim-github-dashboard.git'
Plug 'SirVer/ultisnips' | Plug 'honza/vim-snippets'
Plug 'scrooloose/nerdtree', {'on': 'NERDTreeToggle'}
Plug 'tpope/vim-fireplace', {'for': 'clojure'}
Plug 'rdnetto/YCM-Generator', {'branch': 'stable'}
Plug 'fatih/vim-go', {'tag': '*'}
Plug 'nsf/gocode', {'tag': 'v.20150303', 'rtp':'vim'}
call plug#end()
" }}}

" General {{{
set nocompatible   
set nobackup
set noswapfile
set history=1024
set autochdir
set whichwrap=b,s,<,>,[,]
set nobomb
set backspace=indent,eol,start whichwrap+=<,>,[,]
set clipboard+=unnamed " Vim的默认寄存器和系统剪贴板共享
set winaltkeys=no
" }}}

" GUI {{{
"colorscheme tomorrow-night
set cursorline
set hlsearch
set incsearch
set number
set splitbelow
set splitright
set guifont=Inconsolata:h20
" }}}

" Format {{{
set autoindent
set smartindent
set tabstop=4
set expandtab
set softtabstop=4
set foldmethod=indent
set shiftwidth=4
set ls=2
syntax on
" }}}

" Keymap {{{
```
