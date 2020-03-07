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
        'refresh': 		# refresh token
        'eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImV4cCI6MTU3MTk4MDk5OCwianRpIjoiM2UwNzU0YWU2NmQ2NDZkYWJmZDc4OWE4YjQ1YjNjOGYiLCJ1c2VyX2lkIjoyfQ.YIEn2aJ9zkzjgpw0aN03VIHCZAmgpvohMkMAmMeQi9U',	
        'access': 		# access token
        'eyJ0eXA8iOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNTcxODk4MTk4LCJqdGkiOiIwODZiMGIxMDVkMjQ0NGVhODNlMTg0NTIyYzc1YzEwZiIsInVzZXJfaWQiOjJ9.RVxrGW5b24JpCnLSJhkyONixMllJZOIe4Hj86TpCBp',		
    }
    ```

- 401 : 登录失败

### api/refresh/

POST

- 更新token

- 已登录

- ```json
    {'refresh':""}
    ```

- 200 : 

    ```json
    {
        'access': 	# new access token
        'eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNTcxODk4NDA3LCJqdGkiOiJmMjE1Y2ZiMzhjMTA0MDNlYTQ2ZTkwMjNhNGRiZDA2MSIsInVzZXJfaWQiOjJ9.QS4-AKNhnEEzlzCKJ2clFsjHzb5hSAp4Uk1waLyraRg'
    }
    ```

    

     