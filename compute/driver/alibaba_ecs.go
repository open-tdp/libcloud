package drivers

import (
	"github.com/open-tdp/go-libcloud/compute"
	"github.com/open-tdp/go-libcloud/provider"
	"github.com/open-tdp/go-libcloud/provider/alibaba"

	ecs "github.com/alibabacloud-go/ecs-20140526/v3/client"
)

type AlibabaEcsDriver struct {
	client *alibaba.Client
	ecs    *ecs.Client
	rq     *provider.ReqeustParam
}

func NewAlibabaEcsDriver(rq *provider.ReqeustParam) *AlibabaEcsDriver {

	client := alibaba.NewClient(rq)
	ecs, _ := client.Ecs()

	return &AlibabaEcsDriver{client, ecs, rq}

}

func (d *AlibabaEcsDriver) ListLocation() (*[]compute.Location, *provider.ResponseError) {

	locations := []compute.Location{}

	resp, err := d.ecs.DescribeRegions(&ecs.DescribeRegionsRequest{})

	if err != nil {
		return nil, d.client.Error(err)
	}

	for _, region := range resp.Body.Regions.Region {
		item := compute.Location{
			Id:      *region.RegionId,
			Name:    *region.LocalName,
			Country: "",
		}
		locations = append(locations, item)
	}

	return &locations, nil

}

func (d *AlibabaEcsDriver) ListNodes() (*[]compute.Node, *provider.ResponseError) {

	nodes := []compute.Node{}

	resp, err := d.ecs.DescribeInstances(&ecs.DescribeInstancesRequest{
		RegionId: &d.rq.RegionId,
	})

	if err != nil {
		return nil, d.client.Error(err)
	}

	for _, instance := range resp.Body.Instances.Instance {
		item := compute.Node{
			Id:   *instance.InstanceId,
			Name: *instance.InstanceName,
		}
		if len(instance.PublicIpAddress.IpAddress) > 0 {
			item.PublicIp = instance.PublicIpAddress.String()
		}
		if len(instance.VpcAttributes.PrivateIpAddress.IpAddress) > 0 {
			item.PublicIp = instance.VpcAttributes.PrivateIpAddress.String()
		}
		nodes = append(nodes, item)
	}

	return &nodes, nil
}

func (d *AlibabaEcsDriver) DetailNode(nodeId string) (*compute.Node, *provider.ResponseError) {

	node := compute.Node{}

	resp, err := d.ecs.DescribeInstanceAttribute(&ecs.DescribeInstanceAttributeRequest{
		InstanceId: &nodeId,
	})

	if err != nil {
		return nil, d.client.Error(err)
	}

	node.Id = *resp.Body.InstanceId
	node.Name = *resp.Body.InstanceName
	node.State = compute.NodeState(*resp.Body.Status)
	node.Size = &compute.NodeSize{}
	node.Image = &compute.NodeImage{
		Id: *resp.Body.ImageId,
	}
	node.PublicIp = resp.Body.PublicIpAddress.String()
	node.PrivateIp = resp.Body.VpcAttributes.PrivateIpAddress.String()
	node.Location = &compute.Location{
		Id: *resp.Body.RegionId,
	}

	return &node, nil

}
