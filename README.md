# GoUtil

## 简介

GoUtil 是一个 Go 语言工具集合库，提供了丰富的工具函数，帮助开发者更高效地进行 Go 语言开发。

## 功能特点

该库包含多个工具包，每个工具包专注于不同的功能领域：

- **byteutil**: 字节处理工具
- **cmdutil**: 命令行工具
- **colorutil**: 颜色处理工具
- **compressutil**: 压缩/解压工具
- **convutil**: 数据类型转换工具
- **cryptutil**: 加密/解密工具
- **ctxutil**: 上下文处理工具
- **dtoutil**: 数据传输对象处理工具
- **fileutil**: 文件操作工具
- **headerutil**: HTTP 头处理工具
- **httputil**: HTTP 请求处理工具
- **iputil**: IP 地址处理工具
- **itfutil**: 接口处理工具
- **jsonutil**: JSON 处理工具
- **maputil**: Map 操作工具
- **msgpackutil**: MessagePack 编解码工具
- **osutil**: 操作系统工具
- **randomutil**: 随机数生成工具
- **regexutil**: 正则表达式工具
- **sliutil**: 切片操作工具
- **strutil**: 字符串处理工具
- **stuutil**: 结构体处理工具
- **timeutil**: 时间处理工具
- **urlutil**: URL 处理工具
- **uuidutil**: UUID 生成工具
- **validutil**: 数据验证工具

## 安装

```bash
go get -u github.com/fzf-labs/goutil
```

## 使用示例

```go
package main

import (
    "fmt"
    "github.com/fzf-labs/goutil/strutil"
    "github.com/fzf-labs/goutil/timeutil"
)

func main() {
    // 字符串工具示例
    str := "hello world"
    fmt.Println(strutil.Capitalize(str)) // 输出: Hello world
    
    // 时间工具示例
    fmt.Println(timeutil.NowDatetimeStr()) // 输出当前时间的格式化字符串
}
```

## 贡献指南

1. Fork 本仓库
2. 创建你的功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交你的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 开启一个 Pull Request

## 代码风格

本项目使用 `golangci-lint` 进行代码检查，确保代码质量和一致性。提交前请运行：

```bash
make lint
```

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 致谢

- 感谢所有贡献者的付出
- 感谢 Go 社区的支持
