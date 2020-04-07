package v1

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/didiyun/didiyun-go-sdk/base/v1"

	. "github.com/didiyun/didiyun-go-sdk/compute/v1"
	c "github.com/didiyun/didiyun-go-sdk/tests/common"
)

var regionId = "gz"

// 获取负载均衡算法
func TestGetSlbAlgorithm(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetSlbAlgorithmRequest{
		Header: &base.Header{RegionId: regionId},
	}
	out, err := iaasCli.GetSlbAlgorithm(ctx, in)
	panicErr(err)
	if len(out.Data) < 1 {
		panicErr(errors.New("no records"))
	}
	slbAlgorithm = out.Data[0].Code
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

// 创建SLB
func TestCreateSLB(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &CreateSLBRequest{
		Header:       &base.Header{RegionId: regionId},
		Count:        1,
		AutoContinue: false,
		PayPeriod:    0,
		Name:         "test_slb",
		VpcUuid:      vpcUuid,
		AddressType:  "internet",
		Eip: &CreateSLBRequest_Eip{
			Bandwidth: 1,
		},
	}
	out, err := iaasCli.CreateSLB(ctx, in)
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, regionId, getJobUuids(out.Data)...) //轮询查询进度
	for _, job := range result {
		slbUuid = job.ResourceUuid
	}
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

// 获取slb列表
func TestListSLB(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListSlbRequest{
		Header: &base.Header{RegionId: regionId},
		Start:  0,
		Limit:  10,
	}
	out, err := iaasCli.ListSLB(ctx, in)
	panicErr(err)
	if len(out.Data) < 1 {
		err = errors.New("no records")
		panicErr(err)
	}
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

// 更改slb名称
func TestChangeSLBName(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ChangeSLBNameRequest{
		Header: &base.Header{RegionId: regionId},
		Slb: []*ChangeSLBNameRequest_Slb{{
			SlbUuid: slbUuid,
			Name:    "changeName_test",
		}},
	}
	out, err := iaasCli.ChangeSLBName(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

// 创建slb listener
func TestCreateSLBListener(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &CreateSLBListenerRequest{
		Header:  &base.Header{RegionId: regionId},
		SlbUuid: slbUuid,
		SlbListener: []*ListenerInputInfo{
			{
				Name:         "sdk_test",
				Protocol:     "TCP",
				ListenerPort: 8001,
				BackProtocol: "TCP",
				Algorithm:    slbAlgorithm,
				Monitor: &MonitorInputInfo{
					Protocol:           "TCP",
					Interval:           10,
					Timeout:            5,
					HealthyThreshold:   3,
					UnhealthyThreshold: 3,
				},
			},
		},
	}
	out, err := iaasCli.CreateSLBListener(ctx, in)
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, regionId, out.Data.JobUuid) //轮询查询进度
	for _, job := range result {
		slbListenerUuid = job.ResourceUuid
	}
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}


// 获取slb监听器
func TestListSLBListener(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListSLBListenerRequest{
		Header: &base.Header{RegionId: regionId},
		Start:  0,
		Limit:  10,
		Condition: &ListSLBListenerRequest_Condition{
			SlbUuid: slbUuid,
		},
	}
	out, err := iaasCli.ListSLBListener(ctx, in)
	panicErr(err)
	if len(out.Data) < 1 {
		panicErr(errors.New("no records"))
	}
	for _, l := range out.Data {
		slbListenerUuid = l.SlbListenerUuid
		poolUuid = l.PoolUuid
		if l.Job == nil || l.Job.JobUuid == "" || l.Job.Done == true {
			continue
		}
		success, result, err := iaasCli.WaitForJobResult(ctx, regionId, l.Job.JobUuid) //轮询查询进度
		for _, job := range result {
			slbListenerUuid = job.ResourceUuid
		}
		panicErr(err)
		fmt.Println("Success: ", success, "result: ", result)
		if !success {
			panic("not success")
		}
	}
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

// 根据uuid获取slb
func TestGetSLBByUuid(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetSLBByUuidRequest{
		Header:  &base.Header{RegionId: regionId},
		SlbUuid: slbUuid,
	}
	out, err := iaasCli.GetSLBByUuid(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

// 获取slb数量
func TestGetSLBTotalCnt(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetSLBTotalCntRequest{
		//Header: &base.Header{RegionId: regionId},
		VpcUuid: vpcUuid,
	}
	out, err := iaasCli.GetSLBTotalCnt(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

// 更新SLB监听器
func TestUpdateSLBListener(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &UpdateSLBListenerRequest{
		Header: &base.Header{RegionId: regionId},
		SlbListener: []*ListenerInput{
			{
				Algorithm:    slbAlgorithm,
				BackProtocol: "TCP",
				ListenerPort: 81,
				Monitor: &MonitorInputInfo{
					HealthyThreshold:   3,
					Interval:           23,
					Protocol:           "TCP",
					Timeout:            5,
					UnhealthyThreshold: 3,
				},
				Name:            "test_sdk",
				Protocol:        "TCP",
				SlbListenerUuid: slbListenerUuid,
			},
		},
	}
	out, err := iaasCli.UpdateSLBListener(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

// 增加slb pool member
func TestAddSLBMemberToPool(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &AddSLBMemberToPoolRequest{
		Header:   &base.Header{RegionId: regionId},
		PoolUuid: poolUuid,
		Members: []*MemberInputInfo{
			{
				Dc2Uuid: dc2Uuid,
				Weight:  80,
				Port:    80,
			},
		},
	}
	out, err := iaasCli.AddSLBMemberToPool(ctx, in)
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, regionId, getJobUuids(out.Data)...) //轮询查询进度
	for _, job := range result {
		poolMemberUuid = job.ResourceUuid
	}
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

// 更新slb pool member
func TestUpdateSLBMember(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &UpdateSLBMemberRequest{
		Header: &base.Header{RegionId: regionId},
		Members: []*MemberInputInfo{
			{
				SlbMemberUuid: poolMemberUuid,
				Dc2Uuid:       dc2Uuid,
				Weight:        200,
				Port:          8080,
			},
		},
	}
	out, err := iaasCli.UpdateSLBMember(ctx, in)
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, regionId, getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

// 获取pool member
func TestListPoolMembers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListPoolMembersRequest{
		Header: &base.Header{RegionId: regionId},
		Start:  0,
		Limit:  10,
		Condition: &ListPoolMembersRequest_Condition{
			PoolUuid: poolUuid,
		},
	}
	out, err := iaasCli.ListPoolMembers(ctx, in)
	panicErr(err)
	if len(out.Data) < 1 {
		panicErr(errors.New("no records"))
	}
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

// 删除SLB Pool Member
func TestDeleteSLBMember(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &DeleteSLBMemberRequest{
		Header: &base.Header{RegionId: regionId},
		Members: []*DeleteSLBMemberRequest_Member{
			{
				SlbMemberUuid: poolMemberUuid,
			},
		},
	}
	out, err := iaasCli.DeleteSLBMember(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

// 删除slb监听器
func TestDeleteSLBListener(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &DeleteSLBListenerRequest{
		Header: &base.Header{RegionId: regionId},
		SlbListener: []*DeleteSLBListenerRequest_SlbListener{
			{SlbListenerUuid: slbListenerUuid},
		},
	}
	out, err := iaasCli.DeleteSLBListener(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

// 删除slb
func TestDeleteSLB(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &DeleteSLBRequest{
		Header: &base.Header{RegionId: regionId},
		Slb: []*DeleteSLBRequest_Slb{
			{SlbUuid: slbUuid},
		},
	}
	out, err := iaasCli.DeleteSLB(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}
