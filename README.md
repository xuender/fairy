# Fairy

文件目录管理精灵。

## 安装

```shell
go install github.com/xuender/fairy@latest
```

## 用法

在目录中运行命令，根据目录下的配置，将文件移动到对应目录中。

```shell
# config
fairy init
fairy
```

### init

初始化配置文件。

```shell
fairy init
```

### meta

显示文件或目录类型。

```shell
fairy meta [path...]
```

### move

移动文件或目录到分组目标目录。

```shell
fairy move [path...]
```
