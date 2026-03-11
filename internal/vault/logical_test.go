package vault_test

import (
	"testing"

	"github.com/janosmiko/vlt/internal/vault"
	"github.com/janosmiko/vlt/internal/vault/vaultfakes"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	path := "testpath"

	fakeLogical := &vaultfakes.FakeLogical{}

	v := &vault.Vault{
		Logical: fakeLogical,
	}
	_, err := v.Logical.List(path)
	assert.NoError(t, err)

}
