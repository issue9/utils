utils
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fissue9%2Futils%2Fbadge%3Fref%3Dmaster&style=flat)](https://actions-badge.atrox.dev/issue9/utils/goto?ref=master)
[![Build Status](https://travis-ci.org/issue9/utils.svg?branch=master)](https://travis-ci.org/issue9/utils)
[![license](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://opensource.org/licenses/MIT)
[![codecov](https://codecov.io/gh/issue9/utils/branch/master/graph/badge.svg)](https://codecov.io/gh/issue9/utils)
======

utils 包提供了一些常用的 Go 函数

- MD5 生成 md5 编码；
- FileExists 判断文件或是目录是否存在；
- Merge 合并多个同类型的数据；
- GetSystemLanguageTag 获取当前系统的本地化信息；
- DumpGoSource 输出并格式化 Go 的源代码；
- CurrentFile 相当于部分语言的 `__FILE__`；
- CurrentDir 相当于部分语言的 `__DIR__`；
- CurrentLine 相当于部分语言的 `__LINE__`；
- CurrentFunction 相当于部分语言的 `__FUNCTION__`；

安装
----

```shell
go get github.com/issue9/utils
```

文档
----

[![Go Walker](http://gowalker.org/api/v1/badge)](http://gowalker.org/github.com/issue9/utils)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/issue9/utils)

版权
----

本项目采用 [MIT](http://opensource.org/licenses/MIT) 开源授权许可证，完整的授权说明可在 [LICENSE](LICENSE) 文件中找到。
