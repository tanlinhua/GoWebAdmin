colorscheme hybrid "颜色主题
set background=dark "背景

let mapleader=','
inoremap jj <Esc>

set encoding=utf-8
set foldenable

"代码颜色区分
syntax enable
syntax on

"tab宽度和缩进设置
set tabstop=4
set softtabstop=4
set shiftwidth=4

"兼容老版本
set nocompatible
set backspace=indent,eol,start

"自动缩紧和对齐
set autoindent
set smartindent

"开启追踪列表选择
set cscopetag
set hlsearch

"默认显示行号
set number

"自动加载和保存折叠
au BufWinLeave * silent mkview
au BufWinEnter * silent loadview

"括号和引号自动补全
"inoremap ' ''<ESC>i
"inoremap " ""<ESC>i
"inoremap ( ()<ESC>i
"inoremap [ []<ESC>i
"inoremap { {}<ESC>i

"GoLang代码提示
imap <F2> <C-x><C-o>
"开启NerdTree
map <F3> :NERDTreeMirror<CR>
map <F3> :NERDTreeToggle<CR>
nmap <leader>v :NERDTreeFind<cr>
"nmap <leader>g :NERDTreeToggle<cr>
let NERDTreeShowHidden=1
let NERDTreeIgnore = [
            \ '\.git$', '\.hg$', '\.svn$', '\.stversions$', '\.pyc$', '\.pyo$', '\.svn$', '\.swp$',
            \ '\.DS_Store$', '\.sass-cache$', '__pycache__$', '\.egg-info$', '\.ropeproject$',
            \ ]

let g:NERDTreeWinPos="left"
"let g:NERDTreeWinSize=25
let g:NERDTreeShowLineNumbers=1
let g:neocomplcache_enable_at_startup = 1
"autocmd VimEnter * NERDTree
let NERDTreeWinSize=30

"文件结构显示
nmap <F4> :TagbarToggle<CR>
"let g:tagbar_autopreview = 1 "开启自动预览(随着光标在标签上的移动，顶部会出现一个实时的预览窗口
let g:tagbar_sort = 0 "关闭排序,即按标签本身在文件中的位置排序
let g:tagbar_left = 1 "默认在右边，设置在左边
let g:tagbar_width = 30 "宽度
let g:tagbar_type_go = {
    \ 'ctagstype' : 'go',
    \ 'kinds'     : [
        \ 'p:package',
        \ 'i:imports:1',
        \ 'c:constants',
        \ 'v:variables',
        \ 't:types',
        \ 'n:interfaces',
        \ 'w:fields',
        \ 'e:embedded',
        \ 'm:methods',
        \ 'r:constructor',
        \ 'f:functions'
    \ ],
    \ 'sro' : '.',
    \ 'kind2scope' : {
        \ 't' : 'ctype',
        \ 'n' : 'ntype'
    \ },
    \ 'scope2kind' : {
        \ 'ctype' : 't',
        \ 'ntype' : 'n'
    \ },
    \ 'ctagsbin'  : 'gotags',
    \ 'ctagsargs' : '-sort -silent'
\ }

"go函数追踪 ctrl+] 或 gd 追踪 ctrl+o返回
autocmd FileType go nnoremap <buffer> gd :call GodefUnderCursor()<cr>
autocmd FileType go nnoremap <buffer> <C-]> :call GodefUnderCursor()<cr>
let g:godef_split=3    "追踪打开新tab
let g:godef_same_file_in_same_window=1    "函数在同一个文件中时不需要打开新窗口

"保存自动goimports
let g:go_fmt_command = "goimports"
let g:goimports = 1
"golang
let g:go_version_warning = 0
let g:go_highlight_types = 1
let g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_function_calls = 1
let g:go_highlight_operators = 1
let g:go_highlight_extra_types = 1

"插件
call plug#begin('~/.vim/plugged')
Plug 'w0ng/vim-hybrid'
Plug 'scrooloose/nerdtree'
Plug 'Xuyuanp/nerdtree-git-plugin'
Plug 'fatih/vim-go'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'voldikss/vim-floaterm'
Plug 'mhinz/vim-startify'
Plug 'majutsushi/tagbar'
call plug#end()

"汇总一下快捷键命令：
"F2: 代码提示
"F3: 开启关闭NERDTree(也可以:NERDTree)
"F4: 右侧展示文件结构
"Ctrl + ]/o: 代码追踪/回退