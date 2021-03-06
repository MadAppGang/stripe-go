package loginlink

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/MadAppGang/stripe-go"
	_ "github.com/MadAppGang/stripe-go/testing"
)

func TestLoginLinkNew(t *testing.T) {
	link, err := New(&stripe.LoginLinkParams{
		Account: "acct_EXPRESS",
	})
	assert.Nil(t, err)
	assert.NotNil(t, link)
}
