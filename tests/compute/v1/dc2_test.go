package v1

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/didiyun/didiyun-go-sdk/base/v1"
	. "github.com/didiyun/didiyun-go-sdk/compute/v1"
	c "github.com/didiyun/didiyun-go-sdk/tests/common"
	"strconv"
	"testing"
	"time"
)

func panicErr(err error) {
	if err != nil {
		panic("Query Err: " + err.Error())
	}
}

func panicError(err *base.Error) {
	if err != nil && err.Errno != 0 {
		panic("Errno != 0: " + strconv.Itoa(int(err.Errno)) + "ErrMsg: " + err.Errmsg + " requestId: " + err.RequestId)
	}
}

func getJobUuids(jobs []*base.JobInfo) (jobUuids []string) {
	for _, job := range jobs {
		if len(job.JobUuid) > 0 {
			jobUuids = append(jobUuids, job.JobUuid)
		}
	}
	return
}

func TestListRegionAndZone(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListRegionAndZoneRequest{
		Header: &base.Header{RegionId: "gz", ZoneId: "gz01"},
		Condition: &ListRegionAndZoneRequest_Condition{
			Product: "dc2",
		},
	}
	out, err := iaasCli.ListRegionAndZone(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestListDc2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListDc2Request{
		Header:   &base.Header{RegionId: "gz"},
		Start:    0,
		Limit:    10,
		Simplify: false,
	}
	out, err := iaasCli.ListDc2(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestGetDc2ByUuid(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetDc2ByUuidRequest{
		Header:  &base.Header{RegionId: "gz"},
		Dc2Uuid: dc2Uuid,
	}
	out, err := iaasCli.GetDc2ByUuid(ctx, in)
	panicErr(err)
	panicError(out.Error)
	eipUuid = out.Data[0].Eip.EipUuid
	eipUuids = append(eipUuids, eipUuid)
	ebsUuid = out.Data[0].Ebs[0].EbsUuid
	ebsUuids = append(ebsUuids, ebsUuid)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestGetDc2TotalCnt(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetDc2TotalCntRequest{
		Header: &base.Header{RegionId: "gz"},
		//Ip:     "10.255.0.182",
	}
	out, err := iaasCli.GetDc2TotalCnt(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestCreateDc2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	passwd := "AbddS134bff"
	passwdByte := make([]byte, len(passwd)*2)
	hex.Encode(passwdByte, []byte(passwd))
	in := &CreateDc2Request{
		Header:     &base.Header{RegionId: "gz"},
		ImgUuid:    imgUuid,
		SubnetUuid: subnetUuidToCreateDc2,
		Dc2Model:   "dc2.s1.small1.d20",
		Name:       "test-SDK",
		Password:   string(passwdByte),
		Eip: &CreateDc2Request_Eip{
			Bandwidth: 1,
		},
		Ebs: []*CreateDc2Request_Ebs{{
			Count:    1,
			Size:     25,
			DiskType: "SSD",
			Name:     "test-ebs-dc2Create",
		}},
	}
	out, err := iaasCli.CreateDc2(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, getJobUuids(out.Data)...) //轮询查询进度
	for _, job := range result {
		dc2Uuid = job.ResourceUuid
	}
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestChangeDc2Spec(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ChangeDc2SpecRequest{
		Header: &base.Header{RegionId: "gz"},
		Dc2: []*ChangeDc2SpecRequest_Input{
			{
				Dc2Uuid:  dc2Uuid,
				Dc2Model: "dc2.s1.small2.d40",
			},
		},
	}
	out, err := iaasCli.ChangeDc2Spec(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestStartDc2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &StartDc2Request{
		Header: &base.Header{RegionId: "gz"},
		Dc2: []*StartDc2Request_Input{
			{
				Dc2Uuid: dc2Uuid,
			},
		},
	}
	out, err := iaasCli.StartDc2(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestStopDc2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &StopDc2Request{
		Header: &base.Header{RegionId: "gz"},
		Dc2: []*StopDc2Request_Input{
			{
				Dc2Uuid: dc2Uuid,
			},
		},
	}
	out, err := iaasCli.StopDc2(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestRebootDc2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &RebootDc2Request{
		Header: &base.Header{RegionId: "gz"},
		Dc2: []*RebootDc2Request_Input{
			{
				Dc2Uuid: dc2Uuid,
			},
		},
	}
	out, err := iaasCli.RebootDc2(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestReinstallDc2System(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	passwd := "Fsrv1324fUIs"
	passwdByte := make([]byte, len(passwd)*2)
	hex.Encode(passwdByte, []byte(passwd))
	in := &ReinstallDc2SystemRequest{
		Header: &base.Header{RegionId: "gz"},
		Dc2: []*ReinstallDc2SystemRequest_Input{
			{
				Dc2Uuid: dc2Uuid,
				ImgUuid: reinstallImgUuid, //ubuntu16.04
				//Password: string(passwdByte),
				PubKeyUuids: []string{pubKeyUuid},
			},
		},
	}
	out, err := iaasCli.ReinstallDc2System(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestChangeDc2Password(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	passwd := "SrrSttt123456"
	passwdByte := make([]byte, len(passwd)*2)
	hex.Encode(passwdByte, []byte(passwd))
	in := &ChangeDc2PasswordRequest{
		Header: &base.Header{RegionId: "gz"},
		Dc2: []*ChangeDc2PasswordRequest_Input{
			{
				Dc2Uuid:  dc2Uuid,
				Password: string(passwdByte),
			},
		},
	}
	out, err := iaasCli.ChangeDc2Password(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestChangeDc2Name(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ChangeDc2NameRequest{
		Header: &base.Header{RegionId: "gz"},
		Dc2: []*ChangeDc2NameRequest_Input{
			{
				Dc2Uuid: dc2Uuid,
				Name:    "testChangeDc2Name",
			},
		},
	}
	out, err := iaasCli.ChangeDc2Name(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
}

//测试环境的DC2不能随便删啊，删了可就创建不了了
func TestDestroyDc2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &DestroyDc2Request{
		Header: &base.Header{RegionId: "gz"},
		Dc2: []*DestroyDc2Request_Input{
			{
				Dc2Uuid: dc2Uuid,
			},
		},
	}
	out, err := iaasCli.DestroyDc2(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestListImage(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListImageRequest{
		Header: &base.Header{RegionId: "gz"},
	}
	out, err := iaasCli.ListImage(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
	for _, img := range out.Data {
		if len(imgUuid) > 0 && len(reinstallImgUuid) > 0 {
			break
		} else if img.Name == "CentOS-7.4" {
			imgUuid = img.ImgUuid
		} else if img.Name == "Ubuntu-16.04" {
			reinstallImgUuid = img.ImgUuid
		}
	}
}

func TestListSSHKeys(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListSshKeyRequest{
		Header: &base.Header{RegionId: "gz"},
	}
	out, err := iaasCli.ListSshKey(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestCreateSSHKeys(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &CreateSshKeyRequest{
		Header: &base.Header{RegionId: "gz"},
		Name:   "testCreateSshKey",
		Key:    "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDLPkD03IMvLLWMCO1R3M8xEIcePWj9MPKpFh/dOuraLWnP9tBNIgtEjFXHzomO1i72z8dwEpBy+Xk15RWoMV+C8F4eR9fUpl75On433ji4mLVfIGxDb4CYhWeT0O4KG7fkr4GU6266DBxHVX0HiykNjxHCjO5+fCJ6eeeHVPqfEDO+ZLXE92mxMbdb647wjrTIg94E4sJ6LhRmqHml/W8gS+L0TCcbhNbhyp71hsYrDM/2NTLeU7ehZrhUYNoTxgcHtLI24QT5W+vYLvWTasv0dTsK/CHMlewwjFEJJuhQ9LTSjffPB19xEMgc265a7TolBWEja8L+1VgqhHH3lh35 renlixiang@didichuxing.com",
	}
	out, err := iaasCli.CreateSshKey(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
	pubKeyUuid = out.Data[0].PubKeyUuid
}

func TestDeleteSSHKeys(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &DeleteSshKeyRequest{
		Header:     &base.Header{RegionId: "gz"},
		PubKeyUuid: pubKeyUuid,
	}
	out, err := iaasCli.DeleteSshKey(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}
