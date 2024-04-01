package main

import (
	"github.com/gocolly/colly"
	ptime "github.com/yaa110/go-persian-calendar"
	"go-enamad-test/config"
	"go-enamad-test/database"
	"go-enamad-test/database/migration"
	"go-enamad-test/models"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func GetCompanyPage(index int) {
	c := colly.NewCollector()

	c.OnHTML("#Div_Content>div.row", func(e *colly.HTMLElement) {

		company := models.Company{}
		company.Domain = e.ChildText("div:nth-child(2)")
		company.Name = e.ChildText("div:nth-child(3)")
		company.State = e.ChildText("div:nth-child(4)")
		company.City = e.ChildText("div:nth-child(5)")
		//strPersianDate := strings.Split(e.ChildText("div:nth-child(7)"), "/")
		//year, _ := strconv.Atoi(strPersianDate[0])
		//month, _ := strconv.Atoi(strPersianDate[1])
		//day, _ := strconv.Atoi(strPersianDate[2])
		//var pt ptime.Time = ptime.Date(
		//	year,
		//	ptime.Month(month),
		//	day,
		//	0,
		//	0,
		//	0,
		//	0,
		//	ptime.Iran(),
		//)

		company.CreateDate, _ = StrToTime(e.ChildText("div:nth-child(7)"))
		company.ExpiryDate, _ = StrToTime(e.ChildText("div:nth-child(8)"))
		u, _ := url.Parse(e.ChildAttr("div:nth-child(2)>a:nth-child(1)", "href"))
		m, _ := url.ParseQuery(u.RawQuery)
		company.Code = m["code"][0]
		var newCompany models.Company

		database.Connection().Create(&company).Scan(&newCompany)
		//companies = append(companies, company)
	})

	c.Visit("https://enamad.ir/DomainListForMIMT/Index/" + strconv.Itoa(index))
	c.Wait()
}

func StrToTime(srtDate string) (time.Time, error) {
	strPersianDate := strings.Split(srtDate, "/")
	year, errYear := strconv.Atoi(strPersianDate[0])
	if errYear != nil {
		panic(errYear)
	}
	month, errMonth := strconv.Atoi(strPersianDate[1])
	if errMonth != nil {
		panic(errMonth)
	}
	day, errDay := strconv.Atoi(strPersianDate[2])
	if errDay != nil {
		panic(errDay)
	}
	var pt ptime.Time = ptime.Date(
		year,
		ptime.Month(month),
		day,
		0,
		0,
		0,
		0,
		ptime.Iran(),
	)

	return pt.Time(), nil
}

func main() {

	config.Set()

	database.Connect()

	migration.Migrate()

	for i := 0; i < 3; i++ {
		GetCompanyPage(i)
	}

}
