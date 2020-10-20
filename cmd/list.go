package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	token  string
	author string
	domain string
	scope  string
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List DocBase memos",
	Run: func(cmd *cobra.Command, args []string) {

		author := config.Author
		domain := config.Domain
		token := config.Token
		page := "30"

		request := requestData{
			url:         "https://api.docbase.io/teams/",
			querystring: domain + "/posts?age=2&per_page=" + page + "&q=author:" + author,
			headerName:  "X-DocBaseToken",
			headerValue: token,
		}

		memos := fetchData(request)
		data := createTable(memos)
		showList(data)
	},
}

func init() {
	localCmd := listCmd
	rootCmd.AddCommand(listCmd)

	localCmd.Flags().StringVarP(&author, "author", "a", "", "Memo author to list")
	localCmd.Flags().StringVarP(&scope, "scope", "s", "", "DocBase scope set group(sharing) / private(not sharing) / \"\"(draft)")

	viper.BindPFlag("author", localCmd.Flags().Lookup("author"))
	viper.BindPFlag("scope", localCmd.Flags().Lookup("scope"))
}
