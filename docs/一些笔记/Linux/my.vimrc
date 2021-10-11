inoremap jj <Esc>
set number
set encoding=utf-8
syntax on

call plug#begin('~/.vim/plugged')

Plug 'w0ng/vim-hybrid'
Plug 'scrooloose/nerdtree'
Plug 'fatih/vim-go'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'voldikss/vim-floaterm'
Plug 'mhinz/vim-startify'

call plug#end()

let mapleader=','

let g:go_version_warning = 0

nmap <leader>v :NERDTreeFind<cr>
nmap <leader>g :NERDTreeToggle<cr>
let NERDTreeShowHidden=1
let NERDTreeIgnore = [
            \ '\.git$', '\.hg$', '\.svn$', '\.stversions$', '\.pyc$', '\.pyo$', '\.svn$', '\.swp$',
            \ '\.DS_Store$', '\.sass-cache$', '__pycache__$', '\.egg-info$', '\.ropeproject$',
            \ ]


set background=dark
colorscheme hybrid