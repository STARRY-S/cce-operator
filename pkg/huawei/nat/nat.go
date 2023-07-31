package nat

import (
	ccev1 "github.com/cnrancher/cce-operator/pkg/apis/cce.pandaria.io/v1"
	"github.com/cnrancher/cce-operator/pkg/huawei/common"
	"github.com/cnrancher/cce-operator/pkg/utils"
	nat "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/nat/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/nat/v2/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/nat/v2/region"
	"github.com/sirupsen/logrus"
)

func NewNatClient(auth *common.ClientAuth) *nat.NatClient {
	return nat.NewNatClient(
		nat.NatClientBuilder().
			WithRegion(region.ValueOf(auth.Region)).
			WithCredential(auth.Credential).
			Build())
}

func CreateNatGateway(
	client *nat.NatClient, name string, spec *ccev1.CCEClusterConfigSpec,
) (*model.CreateNatGatewayResponse, error) {
	req := &model.CreateNatGatewayRequest{
		Body: &model.CreateNatGatewayRequestBody{
			NatGateway: &model.CreateNatGatewayOption{
				Name:                name,
				RouterId:            spec.HostNetwork.VpcID,
				InternalNetworkId:   spec.HostNetwork.SubnetID,
				Spec:                model.GetCreateNatGatewayOptionSpecEnum().E_1,
				EnterpriseProjectId: nil,
			},
		},
	}
	res, err := client.CreateNatGateway(req)
	if err == nil {
		logrus.Debugf("CreateNatGateway failed: %v", utils.PrintObject(req))
	}
	return res, err
}

func ShowNetGateway(
	client *nat.NatClient, id string,
) (*model.ShowNatGatewayResponse, error) {
	req := &model.ShowNatGatewayRequest{
		NatGatewayId: id,
	}
	res, err := client.ShowNatGateway(req)
	if err == nil {
		logrus.Debugf("ShowNatGateway failed: %v", utils.PrintObject(req))
	}
	return res, err
}

func DeleteNetGateway(
	client *nat.NatClient, id string,
) (*model.DeleteNatGatewayResponse, error) {
	req := &model.DeleteNatGatewayRequest{
		NatGatewayId: id,
	}
	res, err := client.DeleteNatGateway(req)
	if err == nil {
		logrus.Debugf("DeleteNatGateway failed, ID [%v]", id)
	}
	return res, err
}

func CreateNatGatewaySnatRule(
	client *nat.NatClient, natID, networkID, eipID string, sourceType int32,
) (*model.CreateNatGatewaySnatRuleResponse, error) {
	req := &model.CreateNatGatewaySnatRuleRequest{
		Body: &model.CreateNatGatewaySnatRuleRequestOption{
			SnatRule: &model.CreateNatGatewaySnatRuleOption{
				NatGatewayId: natID,
				NetworkId:    &networkID,
				SourceType:   &sourceType,
				FloatingIpId: eipID,
			},
		},
	}
	res, err := client.CreateNatGatewaySnatRule(req)
	if err == nil {
		logrus.Debugf("CreateNatGatewaySnatRule failed: %v", utils.PrintObject(req))
	}
	return res, err
}
