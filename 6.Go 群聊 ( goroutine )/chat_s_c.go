// Go 群聊 ( goroutine )
// 聊天室的组成
// 聊天室分为两个部分，分别是：
// 服务端
// 客户端
  
// 然后，一般情况下我们互相聊天使用的都只是客户端而已，服务端只是起到调度的作用
// 信息发送与接收的流程
//     假设我们有 服务端（S） 客户端（C1） 客户端（C2） 客户端（C3）并且 S 已经 与 C1 C2 C3 建立了连接
// 理论上的流程是这样的：
// C1 向 S 发出信息
// S 接收到信息
// S 将接收到的信息广播给 C2 C3
// C2 C3 接收信息
 // 服务器代码
package main
 
import (
    "time"
    "fmt"
    "net"
)
 
// 客户端 map 
var client_map =  make(map[string]*net.TCPConn)
 
// 监听请求
func listen_client(ip_port string) {
    tcpAddr, _ := net.ResolveTCPAddr("tcp", ip_port)
    tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
    for {// 不停地接收
        client_con, _ := tcpListener.AcceptTCP()// 监听请求连接
        client_map[client_con.RemoteAddr().String()] = client_con// 将连接添加到 map
        go add_receiver(client_con)
        fmt.Println("用户 : ", client_con.RemoteAddr().String(), " 已连接.")
    }
}
 
// 向连接添加接收器
func add_receiver(current_connect *net.TCPConn) {
    for {
        byte_msg := make([]byte, 2048)
        len, err := current_connect.Read(byte_msg)
        if err != nil { current_connect.Close() }
        fmt.Println(string(byte_msg[:len]))
        msg_broadcast(byte_msg[:len], current_connect.RemoteAddr().String())
    }
}
 
// 广播给所有 client
func msg_broadcast(byte_msg []byte, key string) {
    for k, con := range client_map {
        if k != key { con.Write(byte_msg) }
    }
}
 
// 主函数
func main() {
    fmt.Println("服务已启动...")
    time.Sleep(1 * time.Second)
    fmt.Println("等待客户端请求连接...")
    go listen_client("127.0.0.1:1801")
    select{}
}    

// 客户端代码

package main
 
import (
    "fmt"
    "net"
    "os"
    "bufio"
)
 
// 用户名
var login_name string
 
// 本机连接
var self_connect *net.TCPConn
 
// 读取行文本
var reader = bufio.NewReader(os.Stdin)
 
// 建立连接
func connect(addr string) {
    tcp_addr, _ := net.ResolveTCPAddr("tcp", addr) // 使用tcp
    con, err := net.DialTCP("tcp", nil, tcp_addr) // 拨号
    self_connect = con
    if err != nil {
        fmt.Println("服务器连接失败")
        os.Exit(1)
    }
    go msg_sender()
    go msg_receiver()
}
 
// 消息接收器
func msg_receiver() {
    buff := make([]byte, 2048)
    for {
        len, _ := self_connect.Read(buff) // 读取消息
        fmt.Println(string(buff[:len]))
    }
}
 
// 消息发送器
func msg_sender() {
    for {
        read_line_msg, _, _ := reader.ReadLine()
        read_line_msg = []byte(login_name + " : " + string(read_line_msg))
        self_connect.Write(read_line_msg)
    }
}
 
// 主函数
func main() {
    fmt.Println("请问您怎么称呼？")
    name, _, _ := reader.ReadLine()
    login_name = string(name)
    connect("127.0.0.1:1801")
    select{}
}
