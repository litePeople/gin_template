### gin模板 

### 路由跟控制器修改成beego类似的写法


##### golang居于gin框架二次封装，改写路由的定于方式，路由定义分组等参考beego风格
##### 同时改写路由的请求绑定控制器方式，同样是类似beego风格
##### 由于个人比较喜欢beego的mvc分割，但是gin又没有类似封装
##### 所以本人自己封装了一套，本质上是gin框架的二次封装，改成beego封装，但是又没有beego那么重
##### 同时按照swagger的注释方式，可以自动生成API文档 
##### orm使用，看个人偏好，本模板集成了gorm，如果喜欢beego的orm，可以直接拿过来用即可

#### 规范
1. 每个文件夹都需要配备：README.md，用于说明这个文件夹下的所有文件/文件夹都用于做什么的
2. 常量命名使用驼峰写法

#### 项目的文件夹结构说明
1. conf 配置文件
2. consts 常量文件
3. controllers 协议层，处理前端的接口协议
4. svr 服务层，处理服务的文件
5. das 数据库层，操作数据库的文件
6. models 数据库文件
7. docs 文档
8. dto 与前端交互的数据结构文件
9. middle gin的中间件
10. modules 系统自己封装的模块文件
11. router 路由文件
12. service 后台运行的服务

swagger文档生成或更新
~~~
    1.安装swag文档生成工具
        go install github.com/swaggo/swag/cmd/swag@latest
    2.在项目的根目录下执行命令
        swag init --output docs/swagger
~~~