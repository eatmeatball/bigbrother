# init

# 依赖

项目依赖 `go 1.18`, 和 `cgo`。其中 go 使用了 1.18 `any` 关键字，暂时未使用泛型。如果需要尝试用低于 1.18 的版本进行编译，可以手动更改 go.mod 以及全局替换 `any`。（不推荐低于1.18版本）

# dev run

# build

```cmd
go build -ldflags -H=windowsgui .
```

```shell
go install fyne.io/fyne/v2/cmd/fyne@latest
# windows 打包
fyne package -os windows -icon logo.png
# mac 打包
fyne package -os darwin -icon logo.png
# linux 打包
fyne package -os linux -icon myapp.png 
```

```shell
cd ..
go work init .\bigbrother\
```
