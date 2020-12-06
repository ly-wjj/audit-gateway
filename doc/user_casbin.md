## 用户权限表设计
我将分为用户，角色，权限的设计模式，可以参考RABC模型。


## Casbin权限模块

目前暂时先不使用这种可以适用多种场景的模型。我们先自己从简单的实现方式开始。

参考文档：
- casbin文档： https://github.com/casbin/casbin
- gorm适配器： https://github.com/casbin/gorm-adapter