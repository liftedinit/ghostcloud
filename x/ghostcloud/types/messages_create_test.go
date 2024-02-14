package types_test

import (
	"testing"

	"github.com/liftedinit/ghostcloud/testutil/sample"
	"github.com/liftedinit/ghostcloud/x/ghostcloud/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/stretchr/testify/require"
)

func TestMsgCreateDeployment_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgCreateDeploymentRequest
		err  error
	}{
		{
			name: "invalid address",
			msg:  types.MsgCreateDeploymentRequest{Meta: sample.CreateMetaInvalidAddress()},
			err:  sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg:  types.MsgCreateDeploymentRequest{Meta: sample.CreateMeta(0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
