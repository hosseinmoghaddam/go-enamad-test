package main

import (
	"github.com/gocolly/colly"
	"go-enamad-test/config"
	"go-enamad-test/database"
	"go-enamad-test/database/migration"
	"go-enamad-test/models"
	"net/url"
	"strconv"
)

func GetCompanyPage(index int) {
	c := colly.NewCollector()

	c.OnHTML("#Div_Content>div.row", func(e *colly.HTMLElement) {
		company := models.Company{}
		company.Domain = e.ChildText("div:nth-child(2)")
		company.Name = e.ChildText("div:nth-child(3)")
		company.State = e.ChildText("div:nth-child(4)")
		company.City = e.ChildText("div:nth-child(5)")
		company.CreateDate = e.ChildText("div:nth-child(7)")
		company.ExpiryDate = e.ChildText("div:nth-child(8)")
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

func main() {

	config.Set()

	database.Connect()

	migration.Migrate()

	for i := 0; i < 3; i++ {
		GetCompanyPage(i)
	}

}
