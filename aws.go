package awl

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

var DefaultRegion string

// Session contains the AWS session, which is created on initialization.
var Session *session.Session

// Create the AWS session on init and set the default default region.
func init() {
	DefaultRegion = "us-east-1"
	Session = session.Must(session.NewSession())
}
