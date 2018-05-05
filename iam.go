package awl

import (
	"github.com/aws/aws-sdk-go/service/iam"
)

// Returns (and caches) the account alias. If none is set, then returns the
// account ID.
func (a *Account) CacheAlias() (string, error) {
	a.aliasLock.Lock()
	defer a.aliasLock.Unlock()

	resp, err := a.IAM().ListAccountAliases(&iam.ListAccountAliasesInput{})
	if err == nil && len(resp.AccountAliases) > 0 {
		a.Alias = *resp.AccountAliases[0]
	} else {
		a.Alias = a.Id
	}

	return a.Alias, err
}
