// 简易实现的Client，实现了异步任务的简单轮询功能，仅供参考。

package common

import (
	"context"
	"fmt"
	"github.com/didiyun/didiyun-go-sdk/base/v1"
	"github.com/didiyun/didiyun-go-sdk/bill/v1"
	"github.com/didiyun/didiyun-go-sdk/compute/v1"
	"time"
)

func InitDicloudClient() *DicloudClient {
	conn, err := DialTCP(ServerAddr)
	if err != nil {
		panic(err)
	}
	return &DicloudClient{
		CommonClient: compute.NewCommonClient(conn),
		Dc2Client:    compute.NewDc2Client(conn),
		EipClient:    compute.NewEipClient(conn),
		EbsClient:    compute.NewEbsClient(conn),
		SgClient:     compute.NewSgClient(conn),
		SnapClient:   compute.NewSnapClient(conn),
		VpcClient:    compute.NewVpcClient(conn),
	}
}

type DicloudClient struct {
	compute.CommonClient
	bill.BillClient
	compute.Dc2Client
	compute.EipClient
	compute.EbsClient
	compute.SgClient
	compute.SnapClient
	compute.VpcClient
}

func (cli *DicloudClient) WaitForJobResult(ctx context.Context, jobUuids ...string) (allSuccess bool, resourceJobMap map[string]*base.JobInfo, err error) {
	var out *compute.JobResultResponse
	allSuccess = true
	resourceJobMap = make(map[string]*base.JobInfo)
	for errCnt, cnt := 0, 0; len(jobUuids) > 0 && errCnt < 3; cnt++ { //三次错误就放弃轮询
		time.Sleep(3 * time.Second)
		fmt.Println("Querying JobResult... times: ", cnt)
		out, err = cli.JobResult(ctx, &compute.JobResultRequest{
			Header:   &base.Header{RegionId: "gz"},
			JobUuids: jobUuids,
		})
		if err != nil {
			fmt.Println("Querying JobResult err:", err)
			errCnt++
			continue
		} else if out.Error.Errno != 0 {
			fmt.Println("Querying JobResult errno != 0:", out.Error.Errno, "Errmsg:", out.Error.Errmsg)
			errCnt++
			continue
		}
		for _, job := range out.Data {
			fmt.Println("jobUuid:", job.JobUuid, "resourceUuid:", job.ResourceUuid, "progress:", job.Progress, "%")
			if job.Done { //任务完成，轮询列表中去掉此任务。
				resourceJobMap[job.ResourceUuid] = job
				allSuccess = allSuccess && job.Success
				for idx, jobUuid := range jobUuids {
					if job.JobUuid == jobUuid {
						jobUuids = append(jobUuids[:idx], jobUuids[idx+1:]...)
					}
				}
			}
		}
	}
	return
}
