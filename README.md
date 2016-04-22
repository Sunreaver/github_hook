# github_hook
我的github更新hook


## github配置方式

在giahub项目的配置中，配置webhook到main.go中的Push配置

例如：

``` url
xxx.xxx/github.com/sunreaver/:url
```

## github_hook.conf

names 对应到github配置中的 :url
dolua 对应到./lua/目录中的.lua脚本文件

### tip

配置好后，github项目有push，会自动在url上push一些内容，详见github。

本程序会自动调用github_hook.conf中配置的lua脚本