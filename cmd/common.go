package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/b4b4r07/go-finder"
	"github.com/b4b4r07/go-finder/source"
	"github.com/olekukonko/tablewriter"
)

type requestData struct {
	url         string
	querystring string
	headerName  string
	headerValue string
}

type Dockbase struct {
	Posts []Post `json:"posts"`
	Meta  Meta   `json:"meta"`
}

type Meta struct {
	PreviousPage interface{} `json:"previous_page"`
	NextPage     string      `json:"next_page"`
	Total        int64       `json:"total"`
}

type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	// Body      string `json:"body"`
	Draft     bool   `json:"draft"`
	Archived  bool   `json:"archived"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	// Scope         *Scope `json:"scope"`
	Scope         string `json:"scope"`
	StarsCount    int64  `json:"stars_count"`
	GoodJobsCount int64  `json:"good_jobs_count"`
}

type Scope string

const (
	Group   Scope = "group"
	Private Scope = "private"
)

type Memo struct {
	Title string
	URL   string
	// Scope         *Scope
	Scope         string
	GoodJobsCount string
	StarsCount    string
}

func OpenUrl(url string) {
	url = strings.Trim(url, " ")
	command := exec.Command("open", url)
	err := command.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func fetchData(request requestData) []Memo {

	url := request.url + request.querystring
	var memos []Memo

	for {
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set(request.headerName, request.headerValue)

		client := new(http.Client)
		resp, _ := client.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)

		dockbase := Dockbase{}
		_ = json.Unmarshal(body, &dockbase)

		for _, post := range dockbase.Posts {
			memo := Memo{}

			memo.Title = post.Title
			memo.URL = post.URL
			memo.GoodJobsCount = strconv.FormatInt(post.GoodJobsCount, 10)
			memo.StarsCount = strconv.FormatInt(post.StarsCount, 10)
			memo.Scope = post.Scope

			memos = append(memos, memo)
		}

		if len(dockbase.Meta.NextPage) != 0 {
			url = dockbase.Meta.NextPage
		} else {
			break
		}
	}

	return memos

}

func createTable(memos []Memo) string {

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)

	table.SetHeader([]string{"no", "Title", "URL", "Scope", "Good Jobs Count", "Stars Count"})
	table.SetBorder(false)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	GoodJobsCountSum := 0
	StarsCountSum := 0
	index := 0
	scope := config.Scope

	for _, memo := range memos {
		if scope == memo.Scope {
			index++

			table.Append([]string{strconv.Itoa(index), memo.Title, memo.URL, memo.Scope, memo.GoodJobsCount, memo.StarsCount})

			GoodJobsCount, _ := strconv.Atoi(memo.GoodJobsCount)
			GoodJobsCountSum += GoodJobsCount

			StarsCount, _ := strconv.Atoi(memo.StarsCount)
			StarsCountSum += StarsCount
		}
	}

	table.SetFooter([]string{" ", " ", " ", "Total", strconv.Itoa(GoodJobsCountSum), strconv.Itoa(StarsCountSum)})
	table.Render()
	return tableString.String()
}

func showList(data string) {

	queryWord := ""

	for {
		query := "--query=" + queryWord
		peco, err := finder.New("peco", "--initial-index=1", "--prompt='>'", "--print-query", query)
		if err != nil {
			panic(err)
		}

		peco.Read(source.Text(data))
		selectedItems, _ := peco.Run()

		if len(selectedItems) > 0 {
			var urls []string
			for i, selectedItem := range selectedItems {
				if i == 0 {
					queryWord = selectedItem
					continue
				}
				urls = append(urls, strings.Split(selectedItem, "|")[3])
			}

			if len(urls) > 0 {
				for _, url := range urls {
					OpenUrl(url)
				}
			}

		} else {
			break
		}
	}
}
