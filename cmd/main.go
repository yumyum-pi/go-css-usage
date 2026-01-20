package cmd

import (
	"fmt"
	"os"
	"yumyum-pi/go-css-usage/pkg/html"
	"yumyum-pi/go-css-usage/pkg/ignorefile"

	"github.com/spf13/cobra"
)

var (
	cssFilePath *string
	gitignore   *string
)

var rootCmd = &cobra.Command{
	Use:   "go-css-usage",
	Short: "find all the css selectors used in a html file",
	Run: func(cmd *cobra.Command, args []string) {
		if *cssFilePath == "" {
			fmt.Fprintln(os.Stderr, "css flag not set")
			os.Exit(1)

		}
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "ignore and html files not given in args")
			os.Exit(1)
		}

		ignoreList, err := ignorefile.ReadFile(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while searching for files:%s", err.Error())
			os.Exit(1)
		}

		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "html files not given in args")
			os.Exit(1)
		}

		// check if css file exist
		// check if args is a dir or file
		paths, err := html.GetFiles(args[1], ignoreList)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while searching for files:%s", err.Error())
			os.Exit(1)
		}

		for _, p := range paths {
			fmt.Println(p)
		}

		fmt.Println("total no. of files found", len(paths))
	},
}

func init() {
	cssFilePath = rootCmd.PersistentFlags().StringP("css", "c", "", "Css file path")
	gitignore = rootCmd.PersistentFlags().StringP("gitignore", "i", "", "git ignore file path")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
