package parser

import (
	"regexp"

	"github.com/liubo0127/learn-go/crawler/engine"
)

//type City struct {
//	Name string
//	Url  string
//}

const userListRe = `<a href="(http:[/.\w]+/u/\d+)"[^>]*>?([^<]+)</a>`
const nextPageRe = "<a href=\"(http:[/.\\w]+)(\\d+)\">\u4e0b\u4e00\u9875</a>"

var userRe = regexp.MustCompile(userListRe)
var nextRe = regexp.MustCompile(nextPageRe)

func ParseCity(contents []byte) engine.ParserResult {
	userMatches := userRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	for _, m := range userMatches {
		url := string(m[1])
		s := string(m[2])

		result.Items = append(result.Items, "User "+s)
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: engine.NilParser,
		})
	}

	nextMatches := nextRe.FindAllSubmatch(contents, -1)

	if len(nextMatches) > 0 {
		result.Items = append(result.Items, "Next Page "+string(nextMatches[0][2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(nextMatches[0][1]) + string(nextMatches[0][2]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
