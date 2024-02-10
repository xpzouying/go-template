# README

## 1. 项目功能

### 1.1 管理页面功能

- [ ] React 页面
- [ ] 文件上传管理
- [ ] 用户管理
- [ ] 可视化的日志

### 1.2 后端功能

除了支持上面的功能外，还需要提供。

- [ ] 文件本地管理（用本地文件进行管理）
- [ ] 日志管理（日志单独保存到文件夹中）。
- [ ] API 接口：
  - 日志标准化
  - API 限流

# 2. 运行

## 生产环境

1. 编译前端

```bash
cd ./web
npm run build
```

2. 编译后端

由于后端会 serving 前端的页面，所以需要先编译前端的页面。

```bash
go build .
```

# 参考资料

## 前端框架

- shadcn/ui
- [daisyui](https://daisyui.com/)

## 后端组件

- [gin-static](https://github.com/soulteary/gin-static)
