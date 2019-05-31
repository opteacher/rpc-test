const net = require("net")

for (let i = 0; i < 10; i++) {
    let socket = net.Socket()

    socket.connect(21700, "127.0.0.1", () => {
        console.log(`成功创建连接`)
        socket.write(JSON.stringify({
            method: "HelloSvc.SayHello",
            params: ["opower"]
        }) + "\n")
    })
    socket.on("error", err => {
        console.log(`发生错误：${err}`)
    })
    socket.on("data", buff => {
        console.log(`收到数据：${buff}`)
        socket.destroy()
    })
    socket.on("close", hadErr => {
        console.log(`连接终端：${hadErr ? "异常关闭" : "正常关闭"}`)
    })
}