package main

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/boolow5/QaamuuskaSocdaalka/controllers"
	"github.com/boolow5/QaamuuskaSocdaalka/g"
	_ "github.com/boolow5/QaamuuskaSocdaalka/routers"
	"github.com/boolow5/bolow/boldate"
	"github.com/russross/blackfriday"
)

func main() {
	g.InitEnv()
	controllers.InitLocales()
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.AddFuncMap("neq", notEqual)
	beego.AddFuncMap("timeSince", timeSince)
	beego.AddFuncMap("title", strings.Title)
	beego.AddFuncMap("lessthan", lessthan)
	beego.AddFuncMap("shorten_words", shorten_words)
	beego.AddFuncMap("markdown", markdown)
	beego.Run()
}

func markdown(s string) template.HTML {
	output := blackfriday.MarkdownCommon([]byte(s))
	return beego.Str2html(string(output))
}

func shorten_words(w string, limit int) string {
	words := strings.Split(w, " ")
	newWords := []string{}
	for i := 0; i < len(words); i++ {
		if i < limit {
			newWords = append(newWords, words[i])
		}
	}
	w = strings.Join(newWords, " ")
	if len(words) > len(newWords) {
		w += "..."
	}
	return strings.Title(w)
}

func lessthan(number1, number2 int) bool {
	return number1 < number2
}

func notEqual(val1, val2 interface{}) bool {
	return val1 != val2
}

func timeSince(date time.Time, currentLang string) string {
	now := time.Now()
	year, month, day, hour, minute, second := boldate.Difference(date, now)
	result := ""
	if year > 0 {
		result += fmt.Sprintf("%d %s", year, i18n.Tr(currentLang, "years"))
	}
	if month > 0 {
		result += fmt.Sprintf("%d %s", month, i18n.Tr(currentLang, "months"))
	}
	if day > 0 {
		result += fmt.Sprintf("%d %s", day, i18n.Tr(currentLang, "days"))
	}
	if hour > 0 {
		result += fmt.Sprintf("%d %s", hour, i18n.Tr(currentLang, "hours"))
	}
	if minute > 0 {
		result += fmt.Sprintf("%d %s", minute, i18n.Tr(currentLang, "minutes"))
	}
	if second > 0 {
		// result += i18n.Tr(currentLang, "%d seconds ", second)
	}
	return result
}
