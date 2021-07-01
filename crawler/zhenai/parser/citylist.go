package parser

import (
	"regexp"

	"github.com/liubo0127/learn-go/crawler/engine"
)

//type CityList struct {
//	City  []City
//	Count int
//}
//
//func (c *CityList) Get(contents []byte) engine.ParserResult {
//	re := regexp.MustCompile(`<a href="(http:[/.\w]+zhenghun/\w+)"[^>]*>([^<]+)</a>`)
//
//	matches := re.FindAllSubmatch(contents, -1)
//
//	for _, m := range matches {
//		c.City = append(c.City, City{
//			Name: string(m[2]),
//			Url:  string(m[1]),
//		})
//	}
//}

const cityListRe = `<a href="(http:[/.\w]+zhenghun/\w+)"[^>]*>([^<]+)</a>`

var re = regexp.MustCompile(cityListRe)

func ParseCityList(contents []byte) engine.ParserResult {
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	for _, m := range matches {
		url := string(m[1])
		s := string(m[2])

		result.Items = append(result.Items, "City "+s)
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: ParseCity,
		})
	}

	return result
}
