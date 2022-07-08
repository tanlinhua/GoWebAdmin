## 准备环境
- 安装目录：/opt/module
- 服务器：
   - 192.168.170.101 node1
   - 192.168.170.102 node2
   - 192.168.170.103 node3
   
## Java 安裝
&gt; 分别在node1，node2，node3三个节点上安装java环境

- Oracle-JDK官网下载地址：[https://www.oracle.com/java/technologies/javase/javase8-archive-downloads.html#license-lightbox](https://www.oracle.com/java/technologies/javase/javase8-archive-downloads.html#license-lightbox)
- 解压安装包
```bash
tar -zxvf jdk-8u161-linux-x64.tar.gz -C /opt/module
```

- 编辑环境变量
&gt;  vim /etc/profile.d/my.sh

```bash
#JAVA_HOME
export JAVA_HOME=/opt/module/jdk1.8.0_161
export PATH=$PATH:$JAVA_HOME/bin
```

- 查看Java版本
&gt; java -version

```bash
java version &#34;1.8.0_161&#34;
Java(TM) SE Runtime Environment (build 1.8.0_161-b12)
Java HotSpot(TM) 64-Bit Server VM (build 25.161-b12, mixed mode)
```
## Zookeeper 集群安裝

- 集群规划
   - 在node1,node2,node3 三台机器上安装Zookeeper
- Apache-Zookeeper 官网下载：
```bash
wget https://www.apache.org/dyn/closer.lua/zookeeper/zookeeper-3.7.0/apache-zookeeper-3.7.0-bin.tar.gz
```

- 在node1解压安装包
```bash
tar -zxvf apache-zookeeper-3.7.0-bin.tar.gz -C /opt/module
```

- 在/opt/module/apache-zookeeper-3.7.0-bin/这个目录下创建 data
```bash
mkdir /opt/module/apache-zookeeper-3.7.0-bin/data
```

- 配置node1服务器编号，分别在node2，node3上修改myid内容为2，3
&gt; 在/opt/module/apache-zookeeper-3.7.0-bin/data 目录下创建一个 myid 的文件

```bash
1
```

- 修改node1，node2，node3默认配置
&gt; 配置解读 server.A=B:C:D
&gt; A 是一个数字，表示这个是第几号服务器；
&gt; 集群模式下配置一个文件 myid，这个文件在 dataDir 目录下，这个文件里面有一个数据就是 A 的值，Zookeeper 启动时读取此文件，拿到里面的数据与 zoo.cfg 里面的配置信息比较从而判断到底是哪个 server。
&gt; B 是这个服务器的地址；
&gt; C 是这个服务器 Follower 与集群中的 Leader 服务器交换信息的端口；
&gt; D 是万一集群中的 Leader 服务器挂了，需要一个端口来重新进行选举，选出一个新的Leader，而这个端口就是用来执行选举时服务器相互通信的端口。
&gt; cp /opt/module/apache-zookeeper-3.7.0-bin/conf/zoo_sample.cfg /opt/module/apache-zookeeper-3.7.0-bin/conf/zoo.cfg
&gt; vim /opt/module/apache-zookeeper-3.7.0-bin/conf/zoo.cfg

```bash
# 数据存储目录，默认存储在/tmp目录下
dataDir=/data/modules/apache-zookeeper-3.7.0-bin/data

server.1=hadoop102:2888:3888
server.2=hadoop103:2888:3888
server.3=hadoop104:2888:3888
```

- Zoopkeeper 启动脚本

```bash
#!/bin/bash

case $1 in
&#34;start&#34;){
        for i in node1 node2 node3
        do
                echo  ------------- zookeeper $i 启动 ------------
                ssh $i &#34;/opt/module/apache-zookeeper-3.7.0-bin/bin/zkServer.sh start&#34;
        done
}
;;
&#34;stop&#34;){
        for i in node1 node2 node3
        do
                echo  ------------- zookeeper $i 停止 ------------
                ssh $i &#34;/opt/module/apache-zookeeper-3.7.0-bin/bin/zkServer.sh stop&#34;
        done
}
;;
&#34;status&#34;){
        for i in node1 node2 node3
        do
                echo  ------------- zookeeper $i 状态 ------------
                ssh $i &#34;/opt/module/apache-zookeeper-3.7.0-bin/bin/zkServer.sh status&#34;
        done
}
;;
esac

```
## Kafka 集群安裝

- Kafka 官方下载
```bash
wget https://www.apache.org/dyn/closer.cgi?path=/kafka/3.1.0/kafka_2.12-3.1.0.tgz
tar -zxvf kafka_2.12-3.1.0.tgz -C /opt/module
```

- 修改node1,node2,node3默认配置
&gt; vim /opt/module/kafka_2.12-3.1.0/config/server.properties
&gt; broker.id 分别为0，1，2

```bash
# 服务器，在集群模式需要保持唯一值
broker.id=0
# zookeeper 地址，自定义目录
zookeeper.connect=node1:2181,node2:2181,node3:2181/kafka
# 日志存储目录，默认在/tmp下面
log.dirs=/opt/module/kafka_2.12-3.1.0/kafka-logs
```

- Kafka 启动脚本
```bash
#! /bin/bash
case $1 in
&#34;start&#34;){
        for i in node1 node2 node3
        do
                echo &#34; --------启动 $i Kafka-------&#34;
                ssh  $i  &#34;/opt/module/kafka_2.12-3.1.1/bin/kafka-server-start.sh -daemon /opt/module/kafka_2.12-3.0.0/config/server.properties&#34;
        done
};;
&#34;stop&#34;){
        for i in node1 node2 node3
        do
                echo &#34; --------停止 $i Kafka-------&#34;
                ssh $i &#34;/opt/module/kafka_2.12-3.1.1/bin/kafka-server-stop.sh &#34;
        done
};;
esac
```
## Kafka 基础架构
![Kafka 基础架构.png](https://cdn.convee.cn/image/Kafka 基础架构.png)
## Kafka 生产者架构
![Kafka 生产者.png](https://cdn.convee.cn/image/Kafka 生产者.png)
## Topic操作
```bash
bin/kafka-topics.sh --bootstrap-server node1:9092,node2:9092,node3:9092 --topic first --create  --partitions 3 --replication-factor 2
```
```bash
bin/kafka-topics.sh --bootstrap-server node1:9092,node2:9092,node3:9092 --list
```
```bash
bin/kafka-topics.sh --bootstrap-server node1:9092,node2:9092,node3:9092 --topic second --describe
```
## Kafka 发送消息

```bash
bin/kafka-console-producer.sh --bootstrap-server node1:9092,node2:9092,node3:9092 --topic first
```
```go
package main

import (
	&#34;fmt&#34;

	&#34;github.com/Shopify/sarama&#34;
)

// 基于sarama第三方库开发的kafka client

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &amp;sarama.ProducerMessage{}
	msg.Topic = &#34;first&#34;
	msg.Value = sarama.StringEncoder(&#34;kafka test&#34;)
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{&#34;192.168.170.101:9092&#34;,&#34;192.168.170.102:9092&#34;,&#34;192.168.170.103:9092&#34;}, config)
	if err != nil {
		fmt.Println(&#34;producer closed, err:&#34;, err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println(&#34;send msg failed, err:&#34;, err)
		return
	}
	fmt.Printf(&#34;pid:%v offset:%v\n&#34;, pid, offset)
}

```
## Kafka 消费消息
```bash
bin/kafka-console-consumer.sh --bootstrap-server node1:9092,node2:9092,node3:9092 --topic first
```
```go
package main

import (
	&#34;fmt&#34;
	&#34;time&#34;

	&#34;github.com/Shopify/sarama&#34;
)

// kafka consumer

func main() {
	consumer, err := sarama.NewConsumer([]string{&#34;192.168.170.101:9092&#34;,&#34;192.168.170.102:9092&#34;,&#34;192.168.170.103:9092&#34;}, nil)
	if err != nil {
		fmt.Printf(&#34;fail to start consumer, err:%v\n&#34;, err)
		return
	}
	partitionList, err := consumer.Partitions(&#34;first&#34;) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf(&#34;fail to get list of partition:err%v\n&#34;, err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(&#34;first&#34;, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf(&#34;failed to start consumer for partition %d,err:%v\n&#34;, partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf(&#34;Partition:%d Offset:%d Key:%v Value:%v&#34;, msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
	time.Sleep(10*time.Second)
}
```