# XrChat的后端程序

## 1.HttpServer

## 2.TcpServer

tcp通信使用[tcpx](https://github.com/fwhezfwhez/tcpx)包来实现
tcpx消息基本结构如下

|     字段名      |   类型   |      大小      |
|:------------:|:------:|:------------:|
|    length    | uint32 |    固定4字节     |
|  messageID   | uint32 |    固定4字节     |
| headerLength | uint32 |    固定4字节     |
|  bodyLength  | uint32 |    固定4字节     |
|    header    |  map   | 固定以json格式编解码 |
|     body     | 自定义结构  |  使用自定义编解码器   |

本项目使用`protobuf`作为解码器，tcp协议与客户端交换的消息结构如下

```protobuf
message Message{
    uint64 src = 1; // 消息发起者
    uint64 tar = 2; // 消息目标
    string msg = 3; // 消息内容
    uint64 time = 4; // 消息时间
}
```

使用`messageID`作为路由，可以实现不同的功能

| messageID |  消息类型   |
|:---------:|:-------:|
|     1     |  用户上线   |
|     2     |  用户离线   |
|     3     |  好友消息   |
|     4     |  群聊消息   |
|     5     | 添加好友请求  |
|     6     | 加入群组请求  |
|     7     | 好友申请通过  |
|     8     | 群组申请通过  |
|     9     | 好友申请未通过 |
|    10     | 群组申请未通过 |
|    11     |  被删除好友  |
|    12     |  被踢出群聊  |
|    13     |  退出群聊   |

