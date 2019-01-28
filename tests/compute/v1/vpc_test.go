package v1

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/didiyun/didiyun-go-sdk/base/v1"

	. "github.com/didiyun/didiyun-go-sdk/compute/v1"
	c "github.com/didiyun/didiyun-go-sdk/tests/common"
)

func TestListVpc(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListVpcRequest{
		Header: &base.Header{RegionId: "gz", ZoneId: "gz01"},
		Start:  0,
		Limit:  10,
	}
	out, err := iaasCli.ListVpc(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestGetVpcByUuid(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetVpcByUuidRequest{
		Header:  &base.Header{RegionId: "gz"},
		VpcUuid: vpcUuid,
	}
	out, err := iaasCli.GetVpcByUuid(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestGetVpcTotalCnt(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetVpcTotalCntRequest{
		Header: &base.Header{RegionId: "gz"},
	}
	out, err := iaasCli.GetVpcTotalCnt(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestCreateVpc(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &CreateVpcRequest{
		Header: &base.Header{RegionId: "gz"},
		Name:   "test-sdk-create-vpc",
		Cidr:   "172.16.0.0/12",
		Subnet: []*CreateSubnetInput{{
			Name:   "test-create-subnet(vpc)",
			Cidr:   "172.16.0.0/16",
			ZoneId: "gz02",
		}},
	}
	out, err := iaasCli.CreateVpc(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, "gz", getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	for _, job := range result {
		vpcUuid = job.ResourceUuid
	}
	if !success {
		panic("not success")
	}
}

func TestDeleteVpc(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &DeleteVpcRequest{
		Header: &base.Header{RegionId: "gz"},
		Vpc: []*DeleteVpcRequest_Input{{
			VpcUuid: vpcUuid,
		}},
	}
	out, err := iaasCli.DeleteVpc(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, "gz", getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestChangeVpcName(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ChangeVpcNameRequest{
		Header: &base.Header{RegionId: "gz"},
		Vpc: []*ChangeVpcNameRequest_Input{{
			VpcUuid: vpcUuid,
			Name:    "test-change-vpc-name-2",
		}},
	}
	out, err := iaasCli.ChangeVpcName(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, "gz", getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestListAvailableCidr(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListVpcAvailableCidrRequest{
		Header: &base.Header{RegionId: "gz"},
	}
	out, err := iaasCli.ListVpcAvailableCidr(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestListSubnet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListSubnetRequest{
		Header:   &base.Header{RegionId: "gz"},
		Start:    0,
		Limit:    10,
		Simplify: false,
		Condition: &ListSubnetCondition{
			VpcUuid: vpcUuid,
		},
	}
	out, err := iaasCli.ListSubnet(ctx, in)
	for _, subnet := range out.Data {
		subnetUuids = append(subnetUuids, subnet.SubnetUuid)
	}
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestGetSubnetByUuid(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetSubnetByUuidRequest{
		Header:     &base.Header{RegionId: "gz"},
		VpcUuid:    vpcUuid,
		SubnetUuid: subnetUuidToCreateDc2,
	}
	out, err := iaasCli.GetSubnetByUuid(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestGetSubnetTotalCnt(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetSubnetTotalCntRequest{
		Header:  &base.Header{RegionId: "gz"},
		VpcUuid: vpcUuid,
	}
	out, err := iaasCli.GetSubnetTotalCnt(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestCreateSubnet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &CreateSubnetRequest{
		Header:  &base.Header{RegionId: "gz"},
		VpcUuid: vpcUuid,
		Subnet: []*CreateSubnetInput{{
			Name:   "testCreateSubnet",
			Cidr:   "172.18.0.1/16",
			ZoneId: "gz01",
		}},
	}
	out, err := iaasCli.CreateSubnet(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, "gz", getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	for _, job := range result {
		subnetUuidToCreateDc2 = job.ResourceUuid
	}
	if !success {
		panic("not success")
	}
}

func TestDeleteSubnet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &DeleteSubnetRequest{
		Header:  &base.Header{RegionId: "gz"},
		VpcUuid: vpcUuid,
	}
	for _, uuid := range subnetUuids {
		in.Subnet = append(in.Subnet, &DeleteSubnetRequest_Input{
			SubnetUuid: uuid,
		})
	}
	fmt.Println(in)
	out, err := iaasCli.DeleteSubnet(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, "gz", getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestChangeSubnetName(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ChangeSubnetNameRequest{
		Header:  &base.Header{RegionId: "gz"},
		VpcUuid: vpcUuid,
		Subnet: []*ChangeSubnetNameRequest_Input{{
			SubnetUuid: subnetUuidToCreateDc2,
			Name:       "test-change-subnet-name",
		}},
	}
	out, err := iaasCli.ChangeSubnetName(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, "gz", getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestCheckSubnetCidrOverlap(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &CheckSubnetCidrOverlapRequest{
		Header:  &base.Header{RegionId: "gz"},
		VpcUuid: vpcUuid,
		Cidr:    "172.18.0.1/16",
	}
	out, err := iaasCli.CheckSubnetCidrOverlap(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}
