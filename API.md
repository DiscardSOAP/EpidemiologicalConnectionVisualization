# API protocol

[TOC]

## 用户系统

### api/login/

POST

- 验证身份信息，通过则获取新发放的token

- 任何人

- ```json
    {username:"",password:""}
    ```

- 200 : 成功

    ```json
    {
        token: 		# access token
      'eyJ0eXA8iOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNTcxODk4MTk4LCJqdGkiOiIwODZiMGIxMDVkMjQ0NGVhODNlMTg0NTIyYzc1YzEwZiIsInVzZXJfaWQiOjJ9.RVxrGW5b24JpCnLSJhkyONixMllJZOIe4Hj86TpCBp',		
    }
    ```
    
- 400 : 登录失败

### api/logout/

DELETE

- 登出当前用户

- 已登录

- ```json
    
    ```

- 200 : 成功

    ```json
    {}
    ```



### api/register/

POST

- 注册

- 任何人

- ```json
    {username:""(长度>=8), password:""(长度>=8),invitationCode:""(长度=20),confirmPassword:""(长度>=8),email:""}
    ```

- 200 : 

    ```json
    {
        msg:"register success"
    }
    ```
    
-     400 : 注册失败

### api/profile/

GET

- 个人信息

- 所有人

- ```json
    {}
    ```

- 200 :

    ```json
    {
        user: {
            username: '', #用户名
            name: '', # 昵称
            birth: '', # 注册时间
            organization: '', #所属组织
            description: '', #自述
            email: '',
            email_md5:'',
        }
    }
    ```


POST

- 修改个人信息

- 已登录

- ```json
    {
        user:{
            name: '', # 昵称
            organization: '', #所属组织
            description: '', #自述
            password: '',
            newPassword: '',
            confirmNewPassword: '',
        }
}
    ```
    
- 200 : 

    ```json
    {
        msg:"edit profile success"
    }
    ```

- 400 : 修改个人信息失败

### api/profile/"username"

GET

- 查看用户名为username的用户信息

- 所有人

- ```json
    {}
    ```

- 200 :

    ```json
    {
        user: {
            username: '', #用户名
            name: '', # 昵称
            birth: '', # 注册时间
            organization: '', #所属组织
            description: '', #自述
            email: '',
            email_md5:'',
        }
    }
    ```




### api/root/manage/

GET

- 查询所有用户的信息

- **超级用户root**

- ```
    {}
    ```

- 200：

    

    ```json
    {
        users:[
            {
                username: '', #用户名
                name: '', # 昵称
                birth: '', # 注册时间
                organization: '', #所属组织
                description: '', #自述
                email: '',
                email_md5:'',
            }
        ]
    }
    ```

POST

- 删除某用户

- **超级用户root**

- ```json
    {
        action:'delete'
    	users:[
            {
                username: '', #用户名
                name: '', # 昵称
                birth: '', # 注册时间
                organization: '', #所属组织
                description: '', #自述
                email: '',
                email_md5:'',
    		}
    ]
    }
    ```
    
- 200：

    ```json
    {
        msg:'delete user success'
    }
    ```

## 新闻系统

### api/news/

POST

- 新建一个病例

- 已登录

- ```json
    {
              title: "",
              text: "",
        	markdownText:"",
              category: "",
              events: [
                {
                  startDate: "2020-01-22",
                  startTime: "11:45",
                  endDate: "2020-01-22",
                  endTime: "12:16",
                  addr: "Tianjin No.1 High School",
                  description: "Lorem ipsum dolor sit amet, no nam oblique veritus. Commune scaevola imperdiet nec ut, sed euismod convenire principes at. Est et nobis iisque percipit, an vim zril disputando voluptatibus, vix an salutandi sententiae."
                }
              ],
        	relationships:[
                {
                    relationship: 'From'
                    id: 0
                }
            ]
    }
    ```
    
- ```json
    # 200
    {
        msg:'delete user success',
        id:1 #该新闻编号
    }
    ```

### api/newslib/

POST：

- 请求某条件下所有的病例

- 不需要登录

- ```json
    # 请求参数
    {
        categories:[], #限制新闻的category只能为这几类
        username:"", #username参数为空字符串时表示不限制这一维
        startPublishDate:"2020-01-22", #新闻发布时间晚于该时间，为空字符串时表示不限制
        endPublishDate:"2020-03-03", #新闻发布时间早于该时间，为空字符串时表示不限制
        newsPerPage:20,#每页的新闻数量
        pageNum:1,#从1开始
}
    ```
    

```json
{
    totalPage:10, #总新闻数
    news:[ #给出的新闻顺序按时间逆序排列
        {
            id: 3,
            title:"",
            text:"",
            markdownText:"",
            category:"",
            name:"",
            username:"",
            publishTime:"",
        }
    ]
}
```

### api/news/\<id\>

GET

- 获取id编号的新闻

- 所有人

- ```
    {}
    ```

- 200：

- ```
    {
              title: "",
              name:"",用户昵称
              username:"",用户名
              publishTime:"",上传时间
              text: "",
              markdownText:"",
              category: "",
              prevArticle: {
              title: '上一个标题',
              id: 103,
            },
            nextArticle: {
              title: '下一个标题',
              id: 105,
            },
              events: [
                {
                  startDate: "2020-01-22",
                  startTime: "11:45",
                  endDate: "2020-01-22",
                  endTime: "12:16",
                  addr: "Tianjin No.1 High School",
                  description: "Lorem ipsum dolor sit amet, no nam oblique veritus. Commune scaevola imperdiet nec ut, sed euismod convenire principes at. Est et nobis iisque percipit, an vim zril disputando voluptatibus, vix an salutandi sententiae."
                }
              ],
    }
    ```

POST

- 修改该新闻

- 新闻作者，超级用户

- ```
    {
              title: "",
              text: "",
              markdownText: "",
              category: "",
              events: [
                {
                  startDate: "2020-01-22",
                  startTime: "11:45",
                  endDate: "2020-01-22",
                  endTime: "12:16",
                  addr: "Tianjin No.1 High School",
                  description: "Lorem ipsum dolor sit amet, no nam oblique veritus. Commune scaevola imperdiet nec ut, sed euismod convenire principes at. Est et nobis iisque percipit, an vim zril disputando voluptatibus, vix an salutandi sententiae."
                }
              ],
              relationships:[
                {
                    relationship: 'From'
                    id: 0
                }
            ]
    }
    ```
```
    
- 200

- ```json
    {
        msg:"change success"
    }
```

delete

- 删除该新闻

- 新闻作者，超级用户

- ```
    {
    }
    ```

- 200

- ```json
    {
        msg:"delete success"
    }
    ```

## 拓扑图

### api/topology/

POST：获取一个category中所有的病例信息和关系图

```json
# 请求参数
{
    category : "tianjin"
}
```

```json
# 结果
{
    [
    	{
            id: 1, #对应新闻编号
            relationships: [
            	{
     	     		relationship: "from",
            		id: 0
	    	    }
            ]
        }
    ]
}
```

GET : 获取全国的疫情情况

```json
{
    data:{
        tianjin: 2,
        beijing: 4,
        ...
    }
}
```

## 群体路径追踪图

### api/track/

POST：获取一个category中所有的病例轨迹信息，不要往里放别的信息

```json
# 请求参数
{
    category : "tianjin"
}
```

```json
# 结果
{
    news:[
        {
            id: 2,
            events:[
                ...
            ]
        }
        ...
    ]
}
```

## 病例人数变化

### api/trend/

POST：获取某省份在很长一段时间的病例人数随时间的变化

```json
# 请求参数
{
    category: "tianjin" //空字符串表示全国
}

# 结果
{
	date:[
        '1/23',
        '1/24',
        '1/25',
    ],
    active:[ //表示在此日期以前有多少病例还没被关起来
        2,3,4,    
    ],
    locked:[ //表示在此日期以前有多少病例已经被关起来了
        4,6,9
    ],
}
```

注：根据病例的活跃日期，假设病例轨迹第一个事件在1/23开始，最后一个事件结束在2/25，那么在1/23之前该病例相当于不存在，不统计，在1/23-2/25该病例为活跃病例，2/25以后就被关了