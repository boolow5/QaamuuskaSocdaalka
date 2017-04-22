package models

import (
	"fmt"
	"time"
)

type Country struct {
	Id                int    `json:"id" orm:"auto"`
	Name              string `json:"name" orm:"size(300)"`
	Citizen           string `json:"citizen" orm:"null"`
	OfficialLanguages string `json:"official_languages" orm:"null"`
	CapitalCity       *City  `json:"capital_city" orm:"rel(one)" orm:"null"`

	NorthernBorder string `json:"northern_border" orm:"size(500)" orm:"null"`
	EasternBorder  string `json:"eastern_border" orm:"size(500)" orm:"null"`
	SouthernBorder string `json:"southern_border" orm:"size(500)" orm:"null"`
	WesternBorder  string `json:"western_border" orm:"size(500)" orm:"null"`

	Location                   string  `json:"location"`
	Population                 int     `json:"population"`
	Area                       float32 `json:"area"`
	AverageCostOfLiving        float32 `json:"cost_of_living" orm:"null"`
	AverageVisaCost            float32 `json:"average_visa_cost" orm:"null"`
	NaturalizationPeriodLength float32 `json:"naturalization_period_length"  orm:"null"`

	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

type City struct {
	Id      int      `json:"id" orm:"auto"`
	Name    string   `json:"name" orm:"size(300)"`
	Country *Country `json:"country" orm:"reverse(one)"`

	Population          int     `json:"population"`
	Area                float32 `json:"area"`
	AverageCostOfLiving float32 `json:"cost_of_living"`

	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

func (City *City) TableName() string {
	return "Cities"
}
func (this *City) Valid() bool {
	return (len(this.Name) > 1 && this.Country.Id > 0)
}
func (this *City) String() string {
	return this.Name
}

// Country
func (img *Country) TableName() string {
	return "Countries"
}
func (this *Country) Valid() bool {
	return (len(this.Name) > 1 && this.CapitalCity.Id > 0)
}
func (this *Country) String() string {
	return this.Name
}
func (this *Country) About() string {
	output := fmt.Sprintf("%s is located in %s", this.Name, this.Location)
	if len(this.NorthernBorder) > 0 {
		output += fmt.Sprintf(", in its northern border is %s", this.NorthernBorder)
	}
	if len(this.EasternBorder) > 0 {
		output += fmt.Sprintf(", in its eastern border is %s", this.EasternBorder)
	}
	if len(this.SouthernBorder) > 0 {
		output += fmt.Sprintf(", to the south it is bordered with %s", this.SouthernBorder)
	}
	if len(this.WesternBorder) > 0 {
		output += fmt.Sprintf(" and to the west its bordered with %s", this.WesternBorder)
	}
	if this.Area > 0 {
		output += fmt.Sprintf("\nThe area of %s is %.0f, ", this.Name, this.Area)
	}
	if this.Population > 0 {
		output += fmt.Sprintf(", it has a population of %d", this.Population)
	}
	if this.AverageVisaCost > 0 {
		output += fmt.Sprintf("\nThe average cost of visa is %.2f ", this.AverageVisaCost)
	}
	if this.AverageCostOfLiving > 0 {
		output += fmt.Sprintf(", it costs %.2f/month to live there", this.AverageCostOfLiving)
	}
	if this.NaturalizationPeriodLength > 0 {
		output += fmt.Sprintf("\nIf you legally live in %s more than %.0f years you'll be elligeble to get %s citizenship", this.Name, this.NaturalizationPeriodLength, this.Citizen)
	}
	output += "."
	return output
}
