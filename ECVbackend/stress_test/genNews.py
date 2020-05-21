import json
import random
provinceList = ["安徽","北京","重庆","福建","甘肃","广东","广西","贵州","海南","河北","河南","黑龙江","湖北","湖南","吉林","江苏","江西","辽宁","内蒙古","宁夏","青海","山东","山西","陕西","上海","四川","天津","西藏","新疆","云南","澳门","浙江","香港","台湾"]
locationList = ["饭店","酒店","小区","公园","医院","超市","车站","地铁站","广场","大学","学校"]

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
def randInt(mn=0,mx=1):
    return int(random.random()*(mx-mn))+mn

def get_links(arr):
    n = len(arr)+2
    degrees = [1 for i in range(n)]
    for i in arr:
        degrees[i] += 1
    res = []
    for i in range(n):
        res.append([])
    for k in range(n-2):
        for j in range(n):
            if degrees[j]==1:
                degrees[j] -= 1
                degrees[arr[k]] -= 1
                if j<arr[k]:
                    res[j].append(arr[k])
                else:
                    res[arr[k]].append(j)
                break
    tmp=[]
    for j in range(n):
        if degrees[j]==1:
            tmp.append(j)
    res[tmp[0]].append(tmp[1])
    return res
import math
import requests
start_id = 1144

response = requests.post("http://47.93.180.83:8080/api/login", json.dumps({
            "username": "rootroot",
            "password": "12345678",
        }))
token = json.loads(response.text)['token']
header = {'Authorization': token}
import time
for item in provinceList:
    num = 60+randInt(0,30)
    ids = range(start_id+1,start_id+num+1)
    start_id += num
    prufer = []
    for i in range(num-2):
        prufer.append(randInt(0,num-1))
    link = get_links(prufer)
    for i in range(num):
        k = random.random()*800
        news = newsTemplate
        news['title'] = "病例"+str(ids[i])
        news['category'] = item
        news['text'] = "fake news"
        news['markdownText'] = "fake news"
        events_num = int(random.random()*5)+1
        events=[]
        for j in range(events_num):
            month = randInt(1,4)
            day = randInt(1,28)
            start_date = "2020-%02d-%02d" % (month, day)
            end_date = "2020-%02d-%02d" % (month, day+1)
            t = locationList[randInt(0,len(locationList)-1)]
            try:
                res = requests.get("https://restapi.amap.com/v3/assistant/inputtips?parameters",params={
                    "key":"8eac8768fec1d1525ddd8d3a5702e6bf",
                    "keywords":item+t
                })
                resjson = json.loads(res.text)["tips"]
                if len(resjson)<=1:
                    continue
                name = resjson[randInt(1,len(resjson)-1)]["name"]
                district = resjson[randInt(1,len(resjson)-1)]["district"]
                if item not in district:
                    continue
                event = {
                    "startDate": start_date,
                    "startTime": "%02d:%02d" % (randInt(8,22), randInt(1,28)),
                    "endDate": end_date,
                    "endTime": "%02d:%02d" % (randInt(8,22), randInt(1,28)),
                    "addr": name,
                    "description": "none",
                }
                events.append(event)
            except:
                print("err")
        news['events'] = events
        rels = []
        for k in link[i]:
            rel = {
                "relationship":"To",
                "id":ids[k]
            }
            rels.append(rel)
        news["relationships"] = rels
        res = requests.post("http://47.93.180.83:8080/api/news/", headers=header, data=json.dumps(news))
        print(news)
        time.sleep(0.1)


