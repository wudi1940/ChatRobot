# ChatRobot

本服务使用golang进行开发
网页测试地址：http://101.43.177.135:9999/

程序使用了微软认知服务的SpeechToText以及TextToSpeech对语音和文字进行相互转换，该ApiKey已经写入到服务中不需要填写
调用OpenApi服务仍需要输入ApiKey

针对没有ApiKey的用户提供一个试用key
`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`
当余额使用完之后该key失效，请各位仅限试用，让更多的人有机会体验


# 使用注意事项

由于没有域名也没有ssl，某些浏览器调用浏览器的录音组件需要给该地址加浏览器白名单，否则会无法运行
可以用以下方法加白名单
地址栏输入chrome://flags/, 搜索unsafely
enabled 并填入要授信的域名

# 支持

如果本代码对你有帮助，请多多star

![](docs/pic/goodjob.jpg)
