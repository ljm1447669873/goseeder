package nexus

import (
	"github.com/mmcdole/gofeed"
	"seeder/src/config"
	"strconv"
)

type Client struct {
	baseURL string
	Rule config.NodeRule
}

type Torrent struct {
	GUID  string
	Title string
	URL   string
	Size  string
}

func NewClient(source string, limit int, passkey string,Rule config.NodeRule) Client {
	var baseURL = "https://" + source + "/torrentrss.php?rows=" + strconv.Itoa(limit) + "&tea6=1&tea1=1&tea28=1&tea9=1&tea31=1&tea18=1&tea22=1&tea34=1&tea35=1&tea36=1&iuplder=1&linktype=dl&passkey=" + passkey
	return Client{
		baseURL: baseURL,
		Rule:Rule,
	}
}

func (c *Client) Get() ([]Torrent, error) {
	var ts []Torrent
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(c.baseURL)
	if err == nil {
		for _, value := range feed.Items {
			ts = append(ts, Torrent{
				GUID:  value.GUID,
				Title: value.Title,
				URL:   value.Enclosures[0].URL,
				Size:  value.Enclosures[0].Length,
			})
		}
		return ts, nil
	}

	return nil, err
}
