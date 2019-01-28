# 快速开始 Getting Started

滴滴云Golang开发者工具套件（didiyun-go-sdk）可让您在go语言环境下不用复杂编程即可访问滴滴云下计算产品线产品及账单类操作。本节介绍如何获取滴滴云go sdk并开始调用。

## 环境准备
* 滴滴云go sdk基于golang语言，因此，本文默认您已安装golang的基本语言环境，将不再进行赘述。
* 滴滴云go sdk使用OAuth 2.0协议Bearer Token(RFC 6750)形式进行API访问授权。为使用滴滴云Go SDK，您需要为账号生成一个滴滴云API Token。您可在滴滴云控制台中的API Token管理页面上创建您的Token。

## 安装滴滴云go sdk
执行以下命令，安装滴滴云go sdk。滴滴云go sdk依赖google grpc及protobuf3.x等package，已为您集成在vendor目录中，如有需要，您也可将其自行集成在您的工程项目文件中。

```
go get github.com/didiyun/didiyun-go-sdk
```

## 使用滴滴云go sdk
以下代码示例展示了调用滴滴云go sdk的四个主要步骤：

1. 使用oauth2 Token验证方式，调用grpc.Dial获取一个*grpc.ClientConn。
2. 使用此ClientConn初始化需要访问的产品线Client。
3. 组装请求体，并初始化context。
4. 发起请求并处理应答或错误。

```
package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"

	"github.com/didiyun/didiyun-go-sdk/base/v1"
	. "github.com/didiyun/didiyun-go-sdk/compute/v1"
)

const ServerAddr = "open.didiyunapi.com:8080"
const Token = "your token" //您的API Token

func main() {
	perRPC := oauth.NewOauthAccess(&oauth2.Token{
		AccessToken: Token,
		TokenType:   "bearer",
	})
	//step 1. 获取*grpc.ClientConn。
	clientConn, err := grpc.Dial(ServerAddr,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithPerRPCCredentials(perRPC))
	if err != nil {
		//异常处理
		panic(err)
	}
	//step 2. 获取相应Client。
	dc2Client := NewDc2Client(clientConn)
	//step 3. 组装请求体，并初始化context。
	req := &ListDc2Request{
		Header:   &base.Header{RegionId: "gz"},
		Start:    0,
		Limit:    10,
		Simplify: false,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	//step 4. 发起请求并处理应答或错误。
	out, err := dc2Client.ListDc2(ctx, req)
	if err != nil {
		//异常处理
		panic(err)
	} else if out.Error.Errno != 0 {
		panic(out.Error.Errmsg)
	}
	fmt.Println(out)
}
```

# 返回结构与错误处理
调用滴滴云go sdk中的所有Client的相应方法均会返回Response与一个go内置的error类型。其中，所有类型的Response均包含一个通用的滴滴云*Error类型和一个Data字段，如下所示。

```
type Error struct {
	Errno                int32 
	Errmsg               string
	RequestId            string
}

type ListDc2Response struct {
	Error                *v1.Error 
	Data                 []*Dc2Info
}
```

滴滴云go sdk在服务端或者sdk端出错时，会返回相应的的错误信息。在调用结束时，建议您遵循以下步骤对调用响应进行处理：
1. 对返回的内置error类型进行处理，确定sdk端的调用是否产生错误。
2. 对返回响应中的Error中的Errno进行判断，如果不为0，表示服务端产生了错误。
3. 若没有错误，处理返回响应中的Data部分。

```
out, err := dc2Client.ListDc2(ctx, req)
if err != nil {
    //异常处理，代表sdk端的调用出现error，包括但不限于网络断开、超时等错误
    panic(err)
} else if out.Error.Errno != 0 {
    //返回的Errno不为0，表示服务端出现错误，根据Errno与Errmsg有多种错误
    panic(out.Error.Errmsg)
}
```
   
# 异步调用
滴滴云go sdk中，所有对于资源的操作类请求都是异步实现的。在调用例如DC2开机等一系列异步操作类请求时，您可在返回值中获取到任务信息。

> 滴滴云API的任务（job）具有两个关键概念`done`和`success`，表示任务的执行状态。

| done | success | 说明 | 建议的操作 |
|---|---|---|---|
| false | false | 任务尚未结束运行 | 继续轮询此任务进度 |
| true | false | 任务结束但结果为失败 | 查看result字段中的失败原因，并重新发起操作请求 |
| true | true | 任务结束，结果为成功 | 进行下一步操作 | 

```
type JobInfo struct {
	JobUuid              string 
	ResourceUuid         string 
	Progress             float64
	Type                 string 
	Done                 bool   
	Success              bool   
	Result               string 
}
```

您需要初始化CommonClient，并调用JobResult方法，通过jobUuid来轮询获取此任务的进度。
其中，Done字段表示服务端是否还在处理此任务，Success字段表示处理结果是否成功。
建议您遵循以下步骤对异步任务进行处理：
1. 先判断调用响应是否有错误。
2. 对返回响应中的Done字段进行判断，若为false，则等待片刻重新轮询，若为true，表示任务完成，继续第3步。
3. 判断success字段，若为true，表示任务操作成功，若为false，表示任务失败，此时可读取result字段查看错误信息。

另外，我们在didiyun-go-sdk工程目录中简单实现了一个具有轮询任务功能的简易客户端`tests/common/client.go`，谨作参考。

# 常量与变量规则
本节对滴滴云go sdk中涉及到的一些常量与变量规则进行说明。

## 资源类型（resourceType）
在滴滴云go sdk中，如未进行特殊说明，名为`resourceType`的字段均遵循以下规则。

| resourceType | 说明 |
| ---- | ---- |
| dc2 | 弹性计算主机 |
| eip | 弹性公网IP |
| ebs | 弹性块存储 |
| snap | 磁盘快照，某一个时间点上某一个磁盘的数据备份 |
| vpc | 安全、隔离、IP地址可自定义配置的专有网络 |
| subnet | VPC下所划分的子网 |
| sg | 安全组产品，管控云环境的流量 |
| sgRule | 安全组规则，描述流量出入规则 |
| slb | 滴滴云负载均衡产品 | 

## 云盘类型（diskType）
在滴滴云go sdk中，在EBS以及相关产品中，云盘类型以及对应`diskType`字段共有以下三种。

| 云盘类型 | diskType |
| ---- | ---- |
| SSD云盘 | SSD |
| 高效云盘 | HE |
| 普通云盘 | HDD |

## 网络协议（protocol）
在滴滴云go sdk中，在安全组以及相关产品中，网络协议类型以及对应`protocol`字段共有以下三种。

| 协议类型 | protocol |
| ---- | ---- |
| TCP协议 | TCP |
| UDP协议 | UDP |
| ICMP协议 | ICMP |

## 安全组规则类型（type）
在滴滴云go sdk中，在安全组产品中，安全组规则类型机器对应`type`字段共有以下两种。
| 安全组规则类型 | type |
| 入方向 | Ingress |
| 出方向 | Egress |

## 地域Id与可用区Id（regionId&zoneId）
在滴滴云go sdk中，如未进行特殊说明，名为`regionId`与`zoneId`的字段均遵循以下规则。
目前

<table>
<tr>
	<th> 地域 </th>
	<th> 地域Id（regionId） </th>
	<th> 可用区 </th>
	<th> 可用区Id （zoneId）</th>
</tr>
<tr>
	<td rowspan="2"> 广州 </td>
	<td rowspan="2" > gz </td>
	<td> 广州一区</td>
	<td> gz01</td>
</tr>
<tr>
	<td> 广州二区</td>
	<td> gz02 </td>
</tr>
<tr>
	<td rowspan="1"> 北京</td>
	<td rowspan="1" > bj</td>
	<td> 北京一区 </td>
	<td> bj01 </td>
</tr>
</table>

## DC2密码规则
在滴滴云go sdk中，密码需要至少8位长度，并同时包含大写字母，小写字母和数字，且需要对密码进行16位编码后传输。

例如，若您想设置DC2的密码为`Aa123456`，则进行16位编码后，结果为`password: "4161313233343536"`。

您可使用`encoding/hex`包的`EncodeToString`方法编码您的密码。

## 时间戳（timestamp）规则
在滴滴云go sdk中，时间戳均遵循`Unix timestamp`标准，且单位为*毫秒*。

例如，`2019-01-01 00:00:00`对应时间戳为`1546272000000`。

## 网段（cidr）表示规则
在滴滴云go sdk中，在安全组以及VPC等相关产品中，对于网段的描述规则遵循`CIDR`表示法，使用子网掩码来划分。
例如，创建一个可用网段为`172.16.0.0/12`的VPC，则有其IP和掩码为：

|  | 十进制表示 | 二进制表示 |
|---- | ---- | ---- |
| IP | 172.16.0.0 | 10101100 00010000 00000000 00000000 |
| 掩码 | 255.240.0.0 | 11111111 11110000 00000000 00000000 |

可用IP为172.16.0.1~172.31.255.254共1048574个。

# 调用与错误示例
对于滴滴云go sdk提供的所有接口，文件内均有调用示例，您可使用go test来运行每个接口的调用示例。（部分示例的正确运行需要您手动指定正确参数）。

以查看DC2列表为例，您可使用以下命令运行：

```
cd $GOPATH/src/github.com/didiyun/didiyun-go-sdk/tests/compute/v1
go test -v -run TestListDc2
```

在调用失败时，您可以通过错误码（Errno）与错误信息（Errmsg）得到调用错误的原因，若无法解决，可联系[滴滴云技术支持](#https://help.didiyun.com/hc/request/new/)。
常见错误码如下：

| 错误码  | 错误信息 |  描述  |
|-----|-----|-----|
| 0 | ok	| 调用成功 |
| 1 | 请求有点问题，请反馈客服 | 传入参数非法 |
| 2 | 服务开小差，请稍等| 系统错误，请联系技术支持 |
| 3 | 服务开小差，休息一会再试试 | 未知问题，请联系技术支持 |
| 4 | 认证错误/AccessDenied | token认证失败 |
| 13 | 调用方式有点小问题，请反馈客服 | 传入参数非法 |
| 40000 | 会话异常，请重新登录或稍后再试 | 会话异常，可稍后尝试，如仍有问题请联系技术支持 |
| 41017 | 查询SNAP信息失败 | - |
| 41025 | 查询套餐失败 | - |
| 41026 | 查询镜像失败 | - |
| 41039 | 查询EIP信息失败 | - |
| 41049 | 找不到SG | - |
| 41050 | 创建SG失败 | - |
| 41052 | 查询SG信息失败 | - |
| 41053 | 查询安全组规则失败 | - |
| 41054 | 删除SG失败 | - |
| 41056 | 绑定DC2至安全组失败 | - |
| 41057 | DC2从安全组解绑失败 | - |
| 41059 | 不允许删除默认SG | - |
| 41060 | 找不到指定SG规则 | - |
| 41063 | 找不到指定DC2 | - |
| 41088 | 创建EBS失败 | - |
| 41094 | 查询EBS信息失败 | - |
| 41099 | 创建DC2失败 | - |
| 41107 | 修改DC2规格失败 | - |
| 41108 | 查询DC2信息失败 | - |
| 41113 | 查询EBS总量失败 | - |
| 41115 | 查询SNAP总量失败 | - |
| 41117 | 查询DC2总量失败 | - |
| 41121 | 查询VPC信息失败 | - |
| 41126 | 查询SUBNET信息失败 | - |
| 41135 | 查询SUBNET总量失败 | - |
| 41139 | 校验子网网段失败 | - |
| 41153 | 请先解绑所有关联到该安全组的DC2 | - |
| 41164 | 不允许更改为不同类型的DC2规格 | - |
| 41165 | 不允许更改为更低配置的DC2规格 | - |
| 41172 | 通用型DC2根盘超过大小限制 | - |
| 41181 | 安全组规则存在重复项，请检查输入 | - |
| 1000011 | 账户已欠费，无法进行该操作 | 账户欠费被限制无法进行资源操作，请充值解除欠费状态后再试 |
| 1100017 | 包月到期策略设置失败 | 更改包月到期策略失败，请联系技术支持 |
| 1100018 | 包月续费设置失败 | 包月续费失败，请联系技术支持 |
| 1100020 | 余额不足 | 购买资源金额超过余额，请充值后再试 |
| 1100022 | 获取价格失败 | 获取资源价格失败，请联系技术支持 |
| 1100026 | 包月信息获取失败 | - |
| 1100027 | 包月资源不能删除 | - |
| 1100029 | 所选资源暂不允许包月购买 | - |
| 1100036 | 资源已转换为包月，无需重复操作 | - |
| 1300000 | 资源价格获取失败 | - |
| 2000001 | sshkey增加错误 | - |
| 2000002 | sshkey查询错误 | - |
| 2000003 | sshkey删除错误 | - |
| 2000004 | 不合法的sshkey | - |
| 10000001 | 没有权限，请联系管理员 | 无权限进行此操作，请咨询技术支持 |
| 16000002 | 查询地域与可用区信息失败 | - |
