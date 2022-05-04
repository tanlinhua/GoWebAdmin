# PHP与golang之间rpc调用实例

服务化了，有php代码，也有go代码，我们要相互调用，怎么办？走http，当然，这样子最简单，当并发高的时候，我们就可以走rpc了。
直接上代码，这样子可以好看一点，是吧
首先是golang的服务器

```go
// json_rpc_server project main.go
package main
import (
    "net/rpc"
    "net"
    "log"
    "net/rpc/jsonrpc"
)
//自己的数据类
type MyMath struct{
}
//加法--只能两个参数
func (mm *MyMath) Add(num map[string]float64,reply *float64) error {
    *reply = num["num1"] + num["num2"]
    return nil
}
//减法--只能两个参数
func (mm *MyMath) Sub(num map[string]string,reply *string) error {
    *reply = num["num1"] + num["num2"]
    return nil
}
func main() {
    //注册MyMath类，以代客户端调用
    rpc.Register(new(MyMath))
    listener, err := net.Listen("tcp", ":1215")
    if err != nil {
        log.Fatal("listen error:", err)
    }
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        //新协程来处理--json
        go jsonrpc.ServeConn(conn)
    }
}
```

然后是golang客户端调用服务器

```go
package main
import (
    "net"
    //"net/rpc"
    "log"
    "fmt"
    "net/rpc/jsonrpc"
)
type Args struct {
    A, B int
}
type Arg struct {
    A, B string
}
type Quotient struct {
    Quo, Rem string
}
type MyMath struct{
}
type RpcObj struct {
    Id   int    `json:"id"` // struct标签， 如果指定，jsonrpc包会在序列化json时，将该聚合字段命名为指定的字符串
    Name string `json:"name"`
}
func main()  {
    client, err := net.DialTimeout("tcp", "localhost:1215", 1000*1000*1000*30) // 30秒超时时间
    if err != nil {
        log.Fatal("client\t-", err.Error())
    }
    defer client.Close()
    m := map[string]int{"num1":2, "num2":3 }
    clientRpc := jsonrpc.NewClient(client)
    var reply interface{}
    // 请求数据，rpcObj对象会被填充
    err = clientRpc.Call("MyMath.Add", m, &reply)
    if err != nil {
        log.Fatal("MyMath error:", err)
    }
    fmt.Printf("MyMath: %d+%d=%f\n", m["num1"], m["num2"], reply)
    //当然，你也可以用协程来调用（Asynchronous Call），是不是很高大上
    m2 := map[string]int{"num1":4, "num2":3 }
    goCall := clientRpc.Go("MyMath.Add", m2, &reply, nil)
    replyCall := <-goCall.Done // will be equal to addCall
    fmt.Println(replyCall.Reply, reply)
}
```

接下来是大家关心的php掉golang的：

```php
<?php
class JsonRPC
{
    private $conn;
    function __construct($host, $port) {
        $this->conn = fsockopen($host, $port, $errno, $errstr, 3);
        if (!$this->conn) {
            return false;
        }
    }
    public function Call($method, $params) {
        if ( !$this->conn ) {
            return false;
        }
        $err = fwrite($this->conn, json_encode(array(
                'method' => $method,
                'params' => array($params),
                'id'     => 0,
            ))."\n");
        if ($err === false)
            return false;
        stream_set_timeout($this->conn, 0, 3000);
        $line = fgets($this->conn);
        if ($line === false) {
            return NULL;
        }
        return json_decode($line,true);
    }
}
$client = new JsonRPC("127.0.0.1", 1215);
$r = $client->Call("MyMath.Add",array('num1'=>1,'num2'=>2));
print_r($r);
$r = $client->Call("MyMath.Sub",array('num1'=>'1','num2'=>'2'));
print_r($r);
```

这个只是展示了它们之间可以如何调用，当然我们最好的约定和种协议，我们现在自己约定为json协议，你也可用thrift来做一个约定。这样子也是可以的，都走thrift。
