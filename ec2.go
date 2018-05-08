package awl

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Returns all of the EC2 instances in the given region.
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

// Returns an instance from the given region by private IP.
func (a *Account) InstanceByPrivateIP(region string, ip string) (*ec2.Instance, error) {
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name:   aws.String("private-ip-address"),
				Values: []*string{&ip},
			},
		},
	}

	resp, err := a.EC2(region).DescribeInstances(input)
	if err != nil {
		return nil, err
	}

	rv := []*ec2.Instance{}
	for _, reservation := range resp.Reservations {
		for _, data := range reservation.Instances {
			rv = append(rv, data)
		}
	}

	if len(rv) == 0 {
		return nil, errors.New("No such instance found")
	}

	// Just return the first one.
	return rv[0], nil
}

// Returns all VPCs in the given region.
func (a *Account) AllVpcs(region string) ([]*ec2.Vpc, error) {
	resp, err := a.EC2(region).DescribeVpcs(&ec2.DescribeVpcsInput{})
	if err != nil {
		return nil, err
	}

	return resp.Vpcs, nil
}

// Returns all subnets in the given region.
func (a *Account) AllSubnets(region string) ([]*ec2.Subnet, error) {
	resp, err := a.EC2(region).DescribeSubnets(&ec2.DescribeSubnetsInput{})
	if err != nil {
		return nil, err
	}

	return resp.Subnets, nil
}

// Returns all security groups in the given region.
func (a *Account) AllSecurityGroups(region string) ([]*ec2.SecurityGroup, error) {
	resp, err := a.EC2(region).DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{})
	if err != nil {
		return nil, err
	}

	return resp.SecurityGroups, nil
}
