package awl

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"sync"
)

// An Account represents an AWS account you will use via an assumed role. You
// should use the NewAccount constructor to set it up.
//
// If you need to use more than one role with one AWS account, simply construct
// additional Account objects, specifying different roles.
type Account struct {
	Id            string
	Alias         string
	AssumeRoleArn string

	creds  *credentials.Credentials
	ec2svc map[string]*ec2.EC2
	iamsvc *iam.IAM
	lock   sync.Mutex
}

// Constructs a new Account object from an AWS account ID and the name (not the
// full ARN) of a role, including the path, if any.
func NewAccount(id string, role string) *Account {
	return &Account{Id: id, Alias: id, AssumeRoleArn: fmt.Sprintf("arn:aws:iam::%s:role/%s", id, role)}
}

// Fetches a lazily provisioned STS credentials manager which will renew the
// credentials as necessary. Usually this function will only be called
// internally, but it's here if you need it.
func (a *Account) Credentials() *credentials.Credentials {
	if a.creds == nil {
		a.lock.Lock()
		a.creds = stscreds.NewCredentials(Session, a.AssumeRoleArn)
		a.lock.Unlock()
	}
	return a.creds
}

// Returns a lazily provisioned IAM client for the default region, since IAM
// is a global service.
func (a *Account) IAM() *iam.IAM {
	if a.iamsvc == nil {
		a.lock.Lock()
		a.iamsvc = iam.New(Session, &aws.Config{Credentials: a.Credentials(), Region: aws.String(DefaultRegion)})
		a.lock.Unlock()
	}
	return a.iamsvc
}

// Returns a lazily provisioned EC2 client for the given region.
func (a *Account) EC2(region string) *ec2.EC2 {
	if a.ec2svc == nil {
		a.lock.Lock()
		a.ec2svc = map[string]*ec2.EC2{}
		a.lock.Unlock()
	}

	if rv, ok := a.ec2svc[region]; ok {
		return rv
	} else {
		a.lock.Lock()
		a.ec2svc[region] = ec2.New(Session, &aws.Config{Credentials: a.Credentials(), Region: aws.String(region)})
		a.lock.Unlock()
		return a.ec2svc[region]
	}
}
