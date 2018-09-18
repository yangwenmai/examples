# MeiqiaSDK-Push-Signature-Example
美洽 SDK 3.0 推送的数据结构签名算法，多语言示例。


该项目是 `美洽 SDK 3.0` 推送的数据结构签名的算法，具体推送的数据结构请移步 [美洽 SDK 推送说明](https://github.com/Meiqia/MeiqiaSDK-iOS#7-%E6%B6%88%E6%81%AF%E6%8E%A8%E9%80%81)。

开发者收到美洽的推送数据后，开发者可以选择进行 `签名验证`。

我们使用的签名算法是 [HMAC-SHA1](https://en.wikipedia.org/wiki/Hash-based_message_authentication_code)，key 使用的是美洽 App 的 `Secret-Key`，请到工作台获取你的 `Secret-Key`。**Warning**，请不要将 `Secret-Key` 暴露给任何人。

下图是获取 `Secret-Key` 截图：
![获取 Secret-Key 截图](https://s3.cn-north-1.amazonaws.com.cn/pics.meiqia.bucket/038ccdc55dafe225)

开发者收到美洽的推送数据请求后，请获取头信息的 `Authorization`，其中是美洽推送数据的签名。

我们提供 `Java、Python、Ruby、JavaScript、PHP` 五种语言的计算签名的代码，计算签名后，请和头信息的美洽数据签名进行比对，如果比对失败，则说明推送数据被修改过。
