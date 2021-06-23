package main

import "github.com/spf13/cobra"

func main() {
	cmd := &cobra.Command{
		Use:     "Example",
		Short:   "The Short Description",
		Long:    "wsw",
		Example: "I have an example",
	}

	cmd.Execute()
}
