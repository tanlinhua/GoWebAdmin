# NSQ

## 相关网址
- [官网](https://nsq.io/)
- [Github](https://github.com/nsqio/nsq)
- [下载地址](https://nsq.io/deployment/installing.html)
- [NSQ官方客户端](https://nsq.io/clients/client_libraries.html)
- [NSQ文档](https://nsq.io/overview/design.html)
- [NSQ协议](https://nsq.io/clients/tcp_protocol_spec.html)

## 启动
```
# 开三个终端,分别按顺序启动
./nsqlookupd 
./nsqd --lookupd-tcp-address=192.168.8.100:4160
./nsqadmin --lookupd-http-address=192.168.8.100:4161
# 访问
http://192.168.43.47:4171
# HTTP测试,启动nsqd后，可以用http来测试发送一条消息,可使用CURL来操作。
$ curl -d '这是一条测试消息' 'http://ip:port/pub?topic=test&channel=lc'
OK
```

## NSQ介绍
1. NSQ是Go语言编写的一个开源的实时分布式内存消息队列，其性能十分优异。
2. NSQ 是实时的分布式消息处理平台，其设计的目的是用来大规模地处理每天数以十亿计级别的消息。它具有分布式和去中心化拓扑结构，该结构具有无单点故障、故障容错、高可用性以及能够保证消息的可靠传递的特征.
3. 适合小型项目使用,用来学习消息队列实现原理、学习 golang channel知识以及如何用 go 来写分布式，为什么说适合小型小型项目使用因为，nsq 如果没有能力进行二次开发的情况存在的问题还是很多的。

## NSQ优势
1. NSQ提倡分布式和分散的拓扑,没有单点故障,支持容错和高可用性, 并提供可靠的消息交付保证.
2. NSQ支持横向扩展,没有任何集中式代理
3. NSQ易于配置和部署,并且内置了管理界面

## NSQ特性
1. 支持无 SPOF 的分布式拓扑
2. 水平扩展(没有中间件，无缝地添加更多的节点到集群)
3. 低延迟消息传递 (性能)
4. 结合负载均衡和多播消息路由风格
5. 擅长面向流媒体(高通量)和工作(低吞吐量)工作负载
6. 主要是内存中(除了高水位线消息透明地保存在磁盘上)
7. 运行时发现消费者找到生产者服务(nsqlookupd)
8. 传输层安全性 (TLS)
9. 数据格式不可知
10. 一些依赖项(容易部署)和健全的，有界，默认配置
11. 任何语言都有简单 TCP 协议支持客户端库
12. HTTP 接口统计、管理行为和生产者(不需要客户端库发布)
13. 为实时检测集成了 statsd
14. 健壮的集群管理界面 (nsqadmin)

### 注意点
1. 消息默认不持久化, 可以配置成持久化模式, nsq采用的方式是内存+硬盘的模式,当内存到一定程度就会持久化到硬盘.
    - 如果将 --mem-queue-size设置为0, 所有的消息将会存储到磁盘.
    - 服务器重启时也会将在内存中的消息持久化
2. 每条消息至少传递一次
3. 消息不保证有序.

## NSQ应用场景
>利用消息队列把业务流程中的非关键流程异步化, 从而显著降低业务请求的响应时间

### 应用解耦
>通过使用消息队列将不同业务逻辑解耦,降低系统间耦合,提高系统的健壮性, 后续有其他的业务要使用订单数据可直接订阅消息队列, 提高系统的灵活性.

### 流量削峰
>类似秒杀（大秒）等场景下，某一时间可能会产生大量的请求，使用消息队列能够为后端处理请求提供一定的缓冲区，保证后端服务的稳定性。

## NSQ模块
### nsqd
* nsqd是一个进程监听了http,tcp两种协议, 用来创建topic,channel, 分发消息给消费者,向nsqlooup 注册自己的元数据信息(topic、channel、consumer)，自己的服务信息，最核心模块。
```
nsqd 是一个守护进程，负责接收，排队，投递消息给客户端。也就是说这个服务是干活的。它可以独立运行，不过通常它是由 nsqlookupd 实例所在集群配置的。
```
* 特性:
    1. 对订阅了同一个topic，同一个channel的消费者使用负载均衡策略（不是轮询）
    2. 只要channel存在，即使没有该channel的消费者，也会将生产者的message缓存到队列中（注意消息的过期处理）
    3. 保证队列中的message至少会被消费一次，即使nsqd退出，也会将队列中的消息暂存磁盘上(结束进程等意外情况除外)
    4. 限定内存占用，能够配置nsqd中每个channel队列在内存中缓存的message数量，一旦超出，message将被缓存到磁盘中
    5. topic，channel一旦建立，将会一直存在，要及时在管理台或者用代码清除无效的topic和channel，避免资源的浪费

### nsqlookup
* 存储了nsqd的元数据和服务信息(endpoind),向消费者提供服务发现功能, 向nsqadmin提供数据查询功能.
```
nsqlookupd 是守护进程负责管理拓扑信息。客户端通过查询 nsqlookupd 来发现指定话题（topic）的生产者，并且 nsqd 节点广播话题（topic）和通道（channel）信息。也就是说nsqlookupd是管理者。
```
* 特性:
    1. 唯一性，在一个Nsq服务中只有一个nsqlookupd服务。当然也可以在集群中部署多个nsqlookupd，但它们之间是没有关联的.
    2. 去中心化，即使nsqlookupd崩溃，也会不影响正在运行的nsqd服务
    3. 充当nsqd和naqadmin信息交互的中间件
    4. 提供一个http查询服务，给客户端定时更新nsqd的地址目录.

### nsqadmin
* 简单的管理界面,展示了topic, channel以及channel上的消费者,也可以创建topic,channel
* 特性:
    1. 提供一个对topic和channel统一管理的操作界面以及各种实时监控数据的展示，界面设计的很简洁，操作也很简单
    2. 展示所有message的数量
    3. 能够在后台创建topic和channel
    4. nsqadmin的所有功能都必须依赖于nsqlookupd，nsqadmin只是向nsqlookupd传递用户操作并展示来自nsqlookupd的数据

## Topic和Channel
```
每个nsqd实例旨在一次处理多个数据流。
这些数据流称为“topics”，一个topic具有1个或多个“channels”。
每个channel都会收到topic所有消息的副本，实际上下游的服务是通过对应的channel来消费topic消息。
topic和channel不是预先配置的。topic在首次使用时创建，方法是将其发布到指定topic，或者订阅指定topic上的channel。
channel是通过订阅指定的channel在第一次使用时创建的。
topic和channel都相互独立地缓冲数据，防止缓慢的消费者导致其他chennel的积压（同样适用于topic级别）。
channel可以并且通常会连接多个客户端。假设所有连接的客户端都处于准备接收消息的状态，则每条消息将被传递到随机客户端。

生产者向某个topic中发送消息，如果topic有一个或者多个channle，那么该消息会被复制多分发送到每一个channel中。
类似 rabbitmq中的fanout类型，channle类似队列。
官方说 nsq 是分布式的消息队列服务，但是在我看来只有channel到消费者这部分提现出来分布式的感觉，nsqd 这个模块其实就是单点的，nsqd 将 topic、channel、以及消息都存储在了本地磁盘，官方还建议一个生产者使用一个 nsqd，这样不仅浪费资源还没有数据备份的保障。一旦 nsqd 所在的主机磁损坏，数据都将丢失。

总而言之,消息是从topic--> channel (每个channel接受该topic的所有消息的副本)多播的,但是从channel --> consumers均匀分布 (每个消费者接受该channel的一部分消息)
```

## 其他消息队列
```
RocketMQ
淘宝内部的交易系统使用了淘宝自主研发的Notify消息中间件，使用Mysql作为消息存储媒介，可完全水平扩容，为了进一步降低成本，我们认为存储部分可以进一步优化，2011年初，Linkin开源了Kafka这个优秀的消息中间件，淘宝中间件团队在对Kafka做过充分Review之后，Kafka无限消息堆积，高效的持久化速度吸引了我们，但是同时发现这个消息系统主要定位于日志传输，对于使用在淘宝交易、订单、充值等场景下还有诸多特性不满足，为此我们重新用Java语言编写了RocketMQ，定位于非日志的可靠消息传输（日志场景也OK），目前RocketMQ在阿里集团被广泛应用在订单，交易，充值，流计算，消息推送，日志流式处理，binglog分发等场景。

Kafka
Kafka是LinkedIn开源的分布式发布-订阅消息系统，目前归属于Apache定级项目。Kafka主要特点是基于Pull的模式来处理消息消费，追求高吞吐量，一开始的目的就是用于日志收集和传输。0.8版本开始支持复制，不支持事务，对消息的重复、丢失、错误没有严格要求，适合产生大量数据的互联网服务的数据收集业务。

RabbitMQ
RabbitMQ是使用Erlang语言开发的开源消息队列系统，基于AMQP协议来实现。AMQP的主要特征是面向消息、队列、路由（包括点对点和发布/订阅）、可靠性、安全。AMQP协议更多用在企业系统内，对数据一致性、稳定性和可靠性要求很高的场景，对性能和吞吐量的要求还在其次。
```

[优质文章1](https://www.cnblogs.com/you-men/p/13884645.html)