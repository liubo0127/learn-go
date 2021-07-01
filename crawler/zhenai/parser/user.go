package parser

import (
	"fmt"
	"regexp"
)

const userInfoRe = `<script>window.__INITIAL_STATE__=({"objectInfo".+}).*</script>`

func ParseUser(contents []byte) {
	re := regexp.MustCompile(userInfoRe)

	matches := re.FindAllSubmatch(contents, -1)

	fmt.Println(len(matches))

	//result := engine.ParserResult{}

	//for _, m := range matches {
	//	fmt.Println(string(m[2]), string(m[1]))
	//	result.Items = append(result.Items, "User "+string(m[2]))
	//	result.Requests = append(result.Requests, engine.Request{
	//		Url:        string(m[1]),
	//		ParserFunc: engine.NilParser,
	//	})
	//}

	//return result
}
