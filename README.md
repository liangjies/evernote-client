# evernote-client

#### 介绍
笔记应用,前后端分离,具有搜索笔记、富媒体编辑器、历史记录、笔记标签等功能。



#### 软件架构
Gin、Gorm、MySQL、Redis、JWT、Zap

原项目有一些错误会影响运行

而且原博主有许多重要的部分没有写在README，本人在学习golang期间查看此项目发现了部分问题所在并做了修改，现已可以轻松部署



后续如有进展我会继续更新

### 本地部署
#### 修改配置文件(假如您使用默认的config.yaml)
修改config.yaml中mysql的username，password；redis的addr和password(如果您采用的不是默认)七牛云和邮箱验证码按需求配置即可
#### 运行后端
在终端输入go run main.go然后打开您的mysql可视化工具查看确保存在evernote数据库
在env_users表中添加id为1，deleted_at为NULL，uuid为11111111-1111-1111-1111-111111111111，username为demo，password为fe01ce2a7fbac8fafaed7c982a04e229的用户(这里的密码是加密后的"demo")
#### 运行前端
进入前端目录并在命令行执行yarn serve(若没安装请自行安装)

执行成功即可访问用户端

#### 项目演示

##### 登录注册

![](https://icewx-1251138640.cos.ap-guangzhou.myqcloud.com/github%2Fevernote-client%2F%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20220524145721.png)

##### 笔记列表
![](https://icewx-1251138640.cos.ap-guangzhou.myqcloud.com/github%2Fevernote-client%2F%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20220524114547.png)

##### 笔记历史版本
![](https://icewx-1251138640.cos.ap-guangzhou.myqcloud.com/github%2Fevernote-client%2F%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20220524114517.png)

##### 搜索笔记
![](https://icewx-1251138640.cos.ap-guangzhou.myqcloud.com/github%2Fevernote-client%2F%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20220524114626.png)

##### 笔记本
![](https://icewx-1251138640.cos.ap-guangzhou.myqcloud.com/github%2Fevernote-client%2F%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20220524114653.png)

##### 设置界面
![](https://icewx-1251138640.cos.ap-guangzhou.myqcloud.com/github%2Fevernote-client%2F%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20220524114709.png)


