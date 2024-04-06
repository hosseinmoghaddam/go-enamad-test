package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"go-enamad-test/config"
	"go-enamad-test/database"
	"go-enamad-test/models"
	"strings"
)

func saveDetails() {
	var companies []models.Company

	database.Connection().
		Where("code != ?", "").
		Where("phone = ?", "").Find(&companies)

	for _, company := range companies {
		getCompanyData(company.EnamadID, company.Code, company.ID)
	}

}

func getCompanyData(id string, code string, compamyID uint) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("referer", "https://enamad.ir/")
	})

	c.OnHTML("div.row.person_details > div.mainul", func(e *colly.HTMLElement) {

		var company models.Company

		database.Connection().First(&company, compamyID)
		company.Address = strings.TrimSpace(e.ChildText("div.row:nth-child(1) > div.col-sm-12.col-md-8.contentinformation.licontent.mobiledes"))
		company.Phone = strings.TrimSpace(e.ChildText("div.row:nth-child(2) > div.col-sm-12.col-md-8.contentinformation.licontent.mobiledes"))
		company.Email = strings.TrimSpace(e.ChildText("div.row:nth-child(3) > div.col-sm-12.col-md-8.contentinformation.licontent.mobiledes"))
		company.AnswerTime = strings.TrimSpace(e.ChildText("div.row:nth-child(4) > div.col-sm-12.col-md-8.contentinformation.licontent.mobiledes"))
		database.Connection().Save(&company)
	})

	err := c.Visit(fmt.Sprintf("https://trustseal.enamad.ir/?id=%s&code=%s", id, code))
	if err != nil {
		return
	}
	c.Wait()
}

func main() {

	config.Set()

	database.Connect()

	saveDetails()

}
