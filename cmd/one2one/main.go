package one2one

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	return &cobra.Command{Use: "one2one", Run: main}
}

func main(cmd *cobra.Command, _ []string) {}
