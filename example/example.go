package example

import "github.com/spf13/cobra"

func main() {
	cmd := &cobra.Command{
		Use:     "Example",
		Short:   "This is short",
		Long:    "This is long description of the command. As you can see it is so long",
		Example: "I have an example",
	}

	cmd.Execute()
}
