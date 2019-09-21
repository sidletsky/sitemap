package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/sidletsky/sitemap"
	"github.com/sidletsky/sitemap/internal"
)

var RootCmd = &cobra.Command{
	Use:   "sitemap",
	Short: "sitemap can generate sitemap.xml for the provided url",
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		sitemap, err := sitemap.Parse(url)
		if err != nil {
			log.Fatal(err)
		}
		var urls []string
		for _, v := range sitemap {
			urls = append(urls, v.Loc)
		}
		internal.CreateFile("example/sitemap.xml", urls)
	},
}

func main() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}