# 百度TTS
```
https://tts.baidu.com/text2audio?tex=你好世界&cuid=baike&lan=ZH&ie=utf-8&ctp=1&pdt=301&vol=9&rate=32&per=0
```

tex参数：要转为语音的文本内容

spd参数：朗读速度，取值范围 [1,9]

lan参数：语言码，取值及其含义如下表：

|语言码|含义|
|-----|----|
|ZH|普通话|
|EN|美式英语|
|UK|英式英语|
|CTE|粤语|

per参数：音色码，取值及其实际音色如下表（只有普通话有音色码，其他语音只有一种标准音音色，标准音音色的音色码可以设为0）：

|音色码|实际音色|
|-----|-------|
|0	  |标准女音|
|1	  |标准男音|
|3	  |斯文男音|
|4	  |小萌萌|
|5	  |知性女音|
|6	  |老教授|
|8	  |葛平音|
|9	  |播音员|
|10	  |京腔|
|11	  |温柔大叔|


```python
#!/usr/bin/env python
# -*- coding: utf-8 -*-

import requests
#import os


def TTS(text, speed, lan, per):
    """普通话-音色：
          标准女音
          标准男音
          斯文男音
          小萌萌
          知性女音
          老教授
          葛平音
          播音员
          京腔
          温柔大叔
       英式英语-音色：
          标准音
       美式英语-音色：
          标准音
       粤语-音色：
          标准音
    """
    convertTable = {                       #建立易读文本和音色码/语言码的关系表
        '中文': ('ZH', {
            '标准女音': 0,
            '标准男音': 1,
            '斯文男音': 3,
            '小萌萌': 4,
            '知性女音': 5,
            '老教授': 6,
            '葛平音': 8,
            '播音员': 9,
            '京腔': 10,
            '温柔大叔': 11
        }),
        '英式英语': ('UK', {
            '标准音': 0
        }),
        '美式英语': ('EN', {
            '标准音': 0
        }),
        '粤语': ('CTE', {
            '标准音': 0
        })
    }
    data = {
        'tex': text,
        'spd': speed,
        'lan': convertTable[lan][0],
        'per': convertTable[lan][1][per],
        'ctp': 1,
        'cuid': 'baike',
        'ie': 'UTF-8',
        'pdt': 301,
        'vol': 9,
        'rate': 32
    }
    result = requests.get('https://tts.baidu.com/text2audio', params=data)
    try:
        result.json()
    except:
        return result.content
    else:
        raise ValueError


if __name__ == '__main__':
    try:
        bindata = TTS('''测试一下，你好世界！''', 5, '中文', '小萌萌')
    except:
        print('Error')
    else:
        with open('result.mp3', 'wb+') as f:
            f.write(bindata)            #在同级目录写入为mp3文件
        #os.startfile('result.mp3')     #自动运行生成的mp3
```