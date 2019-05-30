const net = require("net")

let socket = net.Socket()

socket.connect(21700, "127.0.0.1", () => {
    console.log(`成功创建连接`)
    socket.write("我来啦~\n")
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
socket.on("end", () => {
    console.log("会话结束")
})