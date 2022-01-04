package main

import (
	"Liellarpc/src/testproto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接服务端失败: %s", err)
		return
	}
	defer conn.Close()

	// 新建一个客户端
	c := testproto.NewGoTestClient(conn)

	// 调用服务端函数
	r, err := c.GetAddressResponse(context.Background(), &testproto.SendAddress{Address: "1"})
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return
	}

	fmt.Printf("调用成功: %d", r.HttpCode)
}
