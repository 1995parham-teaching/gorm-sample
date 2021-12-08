package many2many

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	return &cobra.Command{Use: "many2many", Run: main}
}

func main(cmd *cobra.Command, _ []string) {}
