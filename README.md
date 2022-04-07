# goshop

### 基于go开发的综合商城系统api

#### 软件架构
1. **config** 全局配置存放目录 **app.yml**
2. **common** 全局公用模块存放目录
3. **utils**  全局助手函数存放目录
4. **model**  数据库模型存放目录
5. **routes** 全局路由配置存放目录
6. **v1**     **api**版本控制存放目录
7. **v1/controller**  **api**控制器存放目录
8. **v1/route**       **api**版本路由地址配置存放目录



#### MAC安装教程

1.  安装 **homebrew**

```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

2.  安装 **golang**

```
brew install go
```

3.  克隆项目

```
git clone https://gitee.com/qiaoshuai951123/goshop.git
```

4.  进入目录初始化mod

```
go mod init goshop
go mod tidy
```
5.  运行项目

```
go run maim.go
```




#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
