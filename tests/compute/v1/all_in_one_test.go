// all in one test 全流程测试
// 从创建至删除的一整套流程，涵盖所有接口

package v1

import (
	"testing"
)

var (
	imgUuid               string
	reinstallImgUuid      string
	vpcUuid               string
	subnetUuidToCreateDc2 string
	subnetUuids           []string
	sgUuid                string
	sgRuleUuids           []string
	dc2Uuid               string
	eipUuid               string
	eipUuids              []string
	ebsUuid               string
	ebsUuids              []string
	pubKeyUuid            string
	snapUuid              string
	slbUuid				  string
	slbListenerUuid       string
	poolUuid		      string
	poolMemberUuid        string
	slbAlgorithm          string
)

func TestAll(t *testing.T) {
	//step 0. ListImage
	TestListImage(t)

	//step 1. ListVpc
	TestListVpc(t)

	//step 2. CreateVpc
	TestCreateVpc(t)

	//step 3. getVPCByUuid
	TestGetVpcByUuid(t)

	//step 4. GetVpcTotalCnt
	TestGetVpcTotalCnt(t)

	//step 5. ChangeVpcName
	TestChangeVpcName(t)

	//step 6. ListAvailableCidr
	TestListAvailableCidr(t)

	//step 7. CheckCidr
	TestCheckSubnetCidrOverlap(t)

	//step 8. CreateSubnet
	TestCreateSubnet(t)

	//step 9. ListSubnet
	TestListSubnet(t)

	//step 10. GetSubnetTotalCnt
	TestGetSubnetTotalCnt(t)

	//step 11. ChangeSubnetName
	TestChangeSubnetName(t)

	//step 12. ListSg
	TestListSg(t)

	//step 13. GetSgTotalCnt
	TestGetSgTotalCnt(t)

	//step 14. CreateSg
	TestCreateSg(t)

	//step 15 ChangeSgName
	TestChangeSgName(t)

	//step 16. CreateSgRule
	TestCreateSgRule(t)

	//step 17. ListSgRule
	TestListSgRule(t)

	//step 18. GetSgRuleCnt
	TestGetSgRuleTotalCnt(t)

	//step 19. CreateDc2
	TestCreateDc2(t)

	//step 20 GetDc2ByUuid
	TestGetDc2ByUuid(t)

	//step 21 ListDc2
	TestListDc2(t)

	//step 22 GetDc2TotalCnt
	TestGetDc2TotalCnt(t)

	//step 23 TestStopDc2
	TestStopDc2(t)

	//step 24 TestStartDc2
	TestStartDc2(t)

	//step 25 TestRebootDc2
	TestRebootDc2(t)

	//step 26 TestChangeDc2Name
	TestChangeDc2Name(t)

	//step 26 TestChangeDc2Password
	TestChangeDc2Password(t)

	//step 27 ListSshKey
	TestListSSHKeys(t)

	//step 28 CreateSshKey
	TestCreateSSHKeys(t)

	//step 29 ReinstallDc2System
	TestReinstallDc2System(t)

	//step 30 DeleteSshKey
	TestDeleteSSHKeys(t)

	//step 31 ChangeDc2Spec
	TestChangeDc2Spec(t)

	//step 32 ListRegionAndZone
	TestListRegionAndZone(t)

	//step 33 ListEip
	TestListEip(t)

	//step 34 GetEipByUuid
	TestGetEipByUuid(t)

	//step 35 GetEipTotalCnt
	TestGetEipTotalCnt(t)

	//step 36 DetachEip
	TestDetachEipFromDc2(t)

	//step 37 AttachEip
	TestAttachEipToDc2(t)

	//step 36 DetachEip（把刚才绑定的EIP解绑）
	TestDetachEipFromDc2(t)

	//step 38 ChangeEipBandwidth
	TestChangeEipBandwidth(t)

	//step 39 CreateEip
	TestCreateEip(t)

	//step 40 DeleteEip
	TestDeleteEip(t)

	//step 41 ListEbs
	TestListEbs(t)

	//step 42 GetEbsByUuid
	TestGetEbsByUuid(t)

	//step 43 GetEbsTotalCnt
	TestGetEbsTotalCnt(t)

	//step 44 DetachEbs
	TestDetachEbsFromDc2(t)

	//step 45 ChangeEbsName
	TestChangeEbsName(t)

	//step 46 ChangeEbsSize
	TestChangeEbsSize(t)

	//step 47 AttachEbs
	TestAttachEbsToDc2(t)

	//step 48 CreateEbs
	TestCreateEbs(t)

	//step 49 DeleteEbs
	TestDeleteEbs(t)

	//step 50 ListSnapshot
	TestListSnap(t)

	//step 51 GetSnapshotCnt
	TestGetSnapTotalCnt(t)

	//step 52 CreateSnap
	TestCreateSnap(t)

	//step 53 RevertSnap
	TestRevertSnap(t)

	//step 54 ChangeSnapName
	TestChangeSnapshotName(t)

	//step 55 DeleteSnap
	TestDeleteSnap(t)

	//step 56 AttachDc2ToSg
	TestAttachDc2ToSg(t)

	//step 57 DetachDc2FromSg
	TestDetachDc2FromSg(t)

	//step 58 DeleteDc2
	TestDestroyDc2(t)

	//step 59. DeleteSgRule
	TestDeleteSgRule(t)

	//step 60. DeleteSg
	TestDeleteSg(t)

	//step 61. DeleteSubnet
	TestDeleteSubnet(t)

	//step 62. DeleteVpc
	TestDeleteVpc(t)
}

func TestSlb (t *testing.T) {
	TestGetSlbAlgorithm(t)
	TestCreateSLB(t)
	TestListSLB(t)
	TestChangeSLBName(t)
	TestCreateSLBListener(t)
	TestGetSLBByUuid(t)
	TestGetSLBTotalCnt(t)
	TestListSLBListener(t)
	TestUpdateSLBListener(t)
	TestAddSLBMemberToPool(t)
	TestUpdateSLBMember(t)
	TestListPoolMembers(t)
	TestDeleteSLBMember(t)
	TestDeleteSLBListener(t)
	TestDeleteSLB(t)
}
