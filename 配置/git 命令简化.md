# git 配置别名简化操作命令

修改 `~/.gitconfig`文件，在末尾按自定义喜好新增别名配置，推荐如下：

```shell
[alias]
	st = status
	a = add
	aa = add .
	ba = branch -a
	br = branch
	ck = checkout
	cka = checkout .
	cb = checkout -b
	ca = commit --amend --no-edit
	cm = commit -m
	clog = log --graph --pretty=oneline --abbrev-commit
	last = log -1 HEAD
	d = diff
	rb = rebase
	rs = reset
	rh = reset --hard
	v = version
```

