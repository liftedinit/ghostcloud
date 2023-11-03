package cli

import (
	"fmt"

	"ghostcloud/x/ghostcloud/types"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

func CmdUpdateDeployment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update name",
		Short: "Update a deployment",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			argName := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argDescription := cmd.Flag(FlagDescription).Value.String()
			argDomain := cmd.Flag(FlagDomain).Value.String()
			argWebsitePayload := cmd.Flag(FlagWebsitePayload).Value.String()

			payload, err := createPayload(argWebsitePayload)
			if err != nil {
				return fmt.Errorf("unable to create payload: %v", err)
			}

			meta := types.Meta{
				Creator:     clientCtx.GetFromAddress().String(),
				Name:        argName,
				Description: argDescription,
				Domain:      argDomain,
			}

			msg := types.NewMsgUpdateDeploymentRequest(
				&meta,
				payload,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	addUpdateFlags(cmd)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}