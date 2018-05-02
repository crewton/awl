package awl

import (
	"github.com/aws/aws-sdk-go/service/iam"
)

// Returns (and caches) the account alias. If none is set, then returns the
// account ID.
func (a *Account) Alias() (string, error) {
	if a.alias != "" {
		return a.alias, nil
	}

	resp, err := a.IAM().ListAccountAliases(&iam.ListAccountAliasesInput{})
	if err != nil {
		a.alias = a.Id
		return a.alias, err
	}

	if len(resp.AccountAliases) > 0 {
		a.alias = *resp.AccountAliases[0]
	} else {
		// if no alias is set, just use the Account ID
		a.alias = a.Id
	}

	return a.alias, nil
}
