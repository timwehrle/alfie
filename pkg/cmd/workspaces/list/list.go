package list

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"

	"github.com/spf13/cobra"
	"github.com/timwehrle/asana/api"
	"github.com/timwehrle/asana/internal/auth"
	"github.com/timwehrle/asana/utils"
)

func NewCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List available workspaces",
		Long: heredoc.Doc(`
				Retrieve and display a list of all workspaces associated 
				with your Asana account.
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun()
		},
	}

	return cmd
}

func listRun() error {
	token, err := auth.Get()
	if err != nil {
		return err
	}

	client := api.New(token)

	workspaces, err := client.GetWorkspaces()
	if err != nil {
		return err
	}

	if len(workspaces) == 0 {
		fmt.Println("No workspaces found.")
		return nil
	}

	fmt.Println(utils.BoldUnderline().Sprint("Your Workspaces:"))
	for i, ws := range workspaces {
		fmt.Printf("%d. %s\n", i+1, ws.Name)
	}

	return nil
}