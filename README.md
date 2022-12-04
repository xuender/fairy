# fairy

文件目录管理精灵。

## 安装

```shell
go install github.com/xuender/fairy@latest
```

## 用法

监听分组目录，自动移动文件到目标目录。

```shell
fairy
```

### init

初始化配置文件。

```shell
fairy init
```

### install

使用 crontab 设置自动启动。

```shell
fairy install
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

### scan

主动扫描分组目录，移动文件到目标目录。

```shell
fairy scan
```
