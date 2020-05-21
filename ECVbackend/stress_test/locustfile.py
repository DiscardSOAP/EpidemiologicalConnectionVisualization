from locust import HttpLocust, TaskSet, task,between
import json
UserCount = 0
newsTemplate = {
    "title": "__TestTest",
    "text": "",
    "markdownText": "",
    "category": "湖北",
    "events": [
              {
                  "startDate": "2020-01-22",
                  "startTime": "11:45",
                  "endDate": "2020-01-22",
                  "endTime": "12:16",
                  "addr": "Tianjin No.1 High School",
                  "description": "Lorem ipsum dolor sit amet, no nam oblique veritus. Commune scaevola imperdiet nec ut, sed euismod convenire principes at. Est et nobis iisque percipit, an vim zril disputando voluptatibus, vix an salutandi sententiae."
              }
    ],
    "relationships": [
        {
            "relationship": 'From',
            "id": 0,
        }
    ]
}
cat_list=["河南","黑龙江","湖北","湖南","吉林","江苏","江西","辽宁","内蒙古","宁夏","青海","山东","山西","陕西","上海","四川","天津","西藏","新疆","云南","澳门","浙江"]
class RegisterTaskSet(TaskSet):
    def on_start(self):
        global UserCount
        self.id = UserCount
        UserCount = UserCount+1
        self.register()

    def register(self):
        response = self.client.post("/register", json.dumps({
            "username": "TestUser"+str(self.id),
            "password": "12345678",
            "confirmPassword": "12345678",
            "email": "TestUser"+str(self.id)+"@gmail.com",
            "invitationCode": "11111111111111111111",
        }))
        print(response)

    @task(1)
    def test(self):
        print("test")


import random

def randInt(mn=0,mx=1):
    return int(random.random()*(mx-mn))+mn


class MyTaskSet(TaskSet):
    def on_start(self):
        global UserCount
        self.id = UserCount
        self.username = "TestUser"+str(self.id)
        UserCount = UserCount+1
        self.login()

    def on_stop(self):
        """ on_stop is called when the TaskSet is stopping """
        self.logout()

    def login(self):
        response = self.client.post("/login", json.dumps({
            "username": self.username,
            "password": "12345678",
        }))
        self.token = json.loads(response.text)['token']
        self.header = {'Authorization': self.token}

    def logout(self):
        self.client.delete("/logout", headers=self.header)

    @task(1001)
    def test(self):
        res = self.client.get("/profile/"+self.username, headers=self.header)

    @task(1000)
    def getAll(self):
        res = self.client.post("/newslib/", headers=self.header, data=json.dumps({
            "categories": [],  # 限制新闻的category只能为这几类
            "username": "",  # username参数为空字符串时表示不限制这一维
            "startPublishDate": "2020-01-22",  # 新闻发布时间晚于该时间，为空字符串时表示不限制
            "endPublishDate": "2020-06-03",  # 新闻发布时间早于该时间，为空字符串时表示不限制
            "newsPerPage": 20,  # 每页的新闻数量
            "pageNum": 1,  # 从1开始
        }))

    @task(1000)
    def getTopo(self):
        global cat_list
        res = self.client.post("/topology", headers=self.header, data=json.dumps({
            "category": cat_list[randInt(0,len(cat_list)-1)],  # 限制新闻的category只能为这几类
        }))

    @task(1000)
    def getTrack(self):
        global cat_list
        res = self.client.post("/track", headers=self.header, data=json.dumps({
            "category": cat_list[randInt(0,len(cat_list)-1)],  # 限制新闻的category只能为这几类
        }))

    @task(3000)
    def getNews(self):
        x  = randInt(1,1000)
        res = self.client.get("/news/"+str(x), headers=self.header)
   
    @task(3)
    def modifyNews(self):
        global newsTemplate
        res = self.client.post("/news/1145", headers=self.header,data=json.dumps(newsTemplate))


class MyLocust(HttpLocust):
    task_set = MyTaskSet
    host = "http://47.93.180.83:8080/api"
    wait_time = between(1, 5)
