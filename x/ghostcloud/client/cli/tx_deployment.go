package cli

import (
	"archive/zip"
	"bytes"
	"fmt"
	"ghostcloud/x/ghostcloud/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

const (
	FlagDescription = "description"
	FlagDomain      = "domain"
)

func addDeploymentFlags(cmd *cobra.Command) {
	f := cmd.Flags()
	f.String(FlagDescription, "", "Description of the deployment")
	f.String(FlagDomain, "", "Custom domain of the deployment")
}

func ReadWebsiteRoot(path string) ([]*types.File, error) {
	// Walk through the directory and process each file
	var files []*types.File
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}
		// Skip directories
		if info.IsDir() {
			return nil
		}
		fileBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		files = append(files, &types.File{
			Meta:    &types.FileMeta{Name: info.Name()},
			Content: &types.FileContent{Content: fileBytes},
		})

		if err != nil {
			fmt.Printf("Error encoding file %q: %v\n", path, err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

func CmdCreateDeployment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-deployment name website-root",
		Short: "Create a new deployment",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexName := args[0]
			argWebsiteRoot := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			websiteFiles, err := ReadWebsiteRoot(argWebsiteRoot)
			if err != nil {
				return err
			}

			argDescription := cmd.Flag(FlagDescription).Value.String()
			argDomain := cmd.Flag(FlagDomain).Value.String()

			meta := types.DeploymentMeta{
				Creator:     clientCtx.GetFromAddress().String(), // TODO: Use []byte?
				Name:        indexName,
				Description: argDescription,
				Domain:      argDomain,
			}

			msg := types.NewMsgCreateDeployment(
				&meta,
				websiteFiles,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	addDeploymentFlags(cmd)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// checkAndLoadWebsiteArchive checks if the website archive is valid and loads it into memory
func checkAndLoadWebsiteArchive(argWebsiteArchive string) (bool, []byte, error) {
	fileInfo, err := os.Stat(argWebsiteArchive)
	if err != nil {
		return false, nil, err
	}
	if fileInfo.Size() > types.DefaultMaxArchiveSize {
		return false, nil, fmt.Errorf("website archive is too big")
	}

	// Read website archive
	websiteArchiveBytes, err := os.ReadFile(argWebsiteArchive)
	if err != nil {
		return false, nil, err
	}

	r := bytes.NewReader(websiteArchiveBytes)
	zipReader, err := zip.NewReader(r, int64(len(websiteArchiveBytes)))
	if err != nil {
		return false, nil, err
	}

	found := false
	for _, f := range zipReader.File {
		if f.Name == "index.html" {
			found = true
		}
	}

	return found, websiteArchiveBytes, nil

}

func CmdCreateDeploymentArchive() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-deployment-archive [name] [website-archive]",
		Short: "Create a new deployment from a website zip archive",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argWebsiteArchive := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			found, websiteArchiveBytes, err := checkAndLoadWebsiteArchive(argWebsiteArchive)
			if err != nil {
				return err
			}

			if !found {
				return fmt.Errorf("website archive does not contain index.html")
			}

			argDescription := cmd.Flag(FlagDescription).Value.String()
			argDomain := cmd.Flag(FlagDomain).Value.String()

			meta := types.DeploymentMeta{
				Creator:     clientCtx.GetFromAddress().String(), // TODO: Use []byte here?
				Name:        argName,
				Description: argDescription,
				Domain:      argDomain,
			}

			msg := types.NewMsgCreateDeploymentArchive(
				&meta,
				websiteArchiveBytes,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	addDeploymentFlags(cmd)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateDeployment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-deployment name website-root",
		Short: "Update a deployment",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexName := args[0]
			argWebsiteRoot := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			websiteFiles, err := ReadWebsiteRoot(argWebsiteRoot)
			if err != nil {
				return err
			}

			argDescription := cmd.Flag(FlagDescription).Value.String()
			argDomain := cmd.Flag(FlagDomain).Value.String()

			meta := types.DeploymentMeta{
				Creator:     clientCtx.GetFromAddress().String(), // TODO: Use []byte here?
				Name:        indexName,
				Description: argDescription,
				Domain:      argDomain,
			}

			msg := types.NewMsgUpdateDeployment(
				&meta,
				websiteFiles,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	addDeploymentFlags(cmd)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateDeploymentMeta() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-deployment-meta name description domain",
		Short: "Update a deployment's meta",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexName := args[0]
			argDescription := args[1]
			argDomain := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			meta := types.DeploymentMeta{
				Name:        indexName,
				Description: argDescription,
				Domain:      argDomain,
			}

			msg := types.NewMsgUpdateDeploymentMeta(
				clientCtx.GetFromAddress().String(),
				&meta,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteDeployment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-deployment name",
		Short: "Delete a deployment",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexName := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteDeployment(
				clientCtx.GetFromAddress().String(),
				indexName,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
