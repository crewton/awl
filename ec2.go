package awl

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Returns _all_ of the EC2 instances in the given region.
func (a *Account) AllInstances(region string) ([]*ec2.Instance, error) {
	resp, err := a.EC2(region).DescribeInstances(&ec2.DescribeInstancesInput{})
	if err != nil {
		return nil, err
	}

	rv := []*ec2.Instance{}
	for _, reservation := range resp.Reservations {
		for _, data := range reservation.Instances {
			rv = append(rv, data)
		}
	}

	return rv, nil
}
