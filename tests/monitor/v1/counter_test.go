package v1

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/didiyun/didiyun-go-sdk/base/v1"

	. "github.com/didiyun/didiyun-go-sdk/monitor/v1"
	c "github.com/didiyun/didiyun-go-sdk/tests/common"
)

func panicError(err2 *base.Error, err error) {
	if err != nil {
		panic("Query Err: " + err.Error())
	}
	if err2 != nil && err2.Errno != 0 {
		panic("Errno != 0: " + strconv.Itoa(int(err2.Errno)) + "ErrMsg: " + err2.Errmsg + " requestId: " + err2.RequestId)
	}
}

func GetMonitorClient() MonitorClient {
	conn, err := c.DialTCP(c.ServerAddr)
	if err != nil {
		panic(err)
	}
	return NewMonitorClient(conn)
}

func TestListCounter(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListCounterRequest{
		Header: &base.Header{RegionId: "gz"},
		Resource: []*CounterResource{
			{
				ResourceUuids:[]string{"3a4044a5d9b2555d8617a853a3bec8ef"},
				ResourceType: "dc2",
				Metric : []string{"cpu.util", "disk.read", "disk.write"},
			},
			{
				ResourceUuids:[]string{"fcde51c1c4115e2487ee04fcb81e7531"},
				ResourceType: "eip",
				Metric : []string{"rxbytes", "txbytes"},
			},
		},
	}
	out, err := GetMonitorClient().ListCounter(ctx, in)
	panicError(out.Error, err)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestListCounterData(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListCounterDataRequest{
		Header: &base.Header{RegionId: "gz"},
		Counter: []*CounterDataInput{
			{
				ResourceType: "dc2",
				ResourceUuid: "3a4044a5d9b2555d8617a853a3bec8ef",
				MonitorTags: "device=vda",
				Metric:  "disk.write",
				StartTime : 1560944187,
				EndTime: 1560947787,
			},
			{
				ResourceType: "dc2",
				ResourceUuid: "3a4044a5d9b2555d8617a853a3bec8ef",
				MonitorTags: "device=vdb",
				Metric:  "disk.read",
				StartTime : 1560944187,
				EndTime: 1560947787,
			},
			{
				ResourceType: "eip",
				ResourceUuid: "fcde51c1c4115e2487ee04fcb81e7531",
				MonitorTags: "",
				Metric:  "rxbytes",
				StartTime : 1560944187,
				EndTime: 1560947787,
			},
		},
	}
	out, err := GetMonitorClient().ListCounterData(ctx, in)
	panicError(out.Error, err)
	fmt.Println(c.ToPrettyJsonString(out))
}