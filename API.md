# API protocol

[TOC]

## 用户系统

### api/token/

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

### api/register/

POST

- 注册

- 任何人

- ```json
    {username:"", password:"",invitationCode:"",confirmPassword:""}
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

- 已登录

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

    

- 



​    
