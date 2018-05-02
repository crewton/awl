package awl

import (
	"testing"
)

type TestNewAccountData struct {
	id   string
	role string
	arn  string
}

func TestNewAccount(t *testing.T) {
	tests := []TestNewAccountData{
		{id: "123456789012", role: "my-role-name", arn: "arn:aws:iam::123456789012:role/my-role-name"},
		{id: "000011112222", role: "xacct/pathed-role", arn: "arn:aws:iam::000011112222:role/xacct/pathed-role"},
	}

	for _, test := range tests {
		acct := NewAccount(test.id, test.role)
		if acct.AssumeRoleArn != test.arn {
			t.Errorf("Role ARN was incorrectly constructed. Got '%s', expected '%s'.", acct.AssumeRoleArn, test.arn)
		}
	}
}
