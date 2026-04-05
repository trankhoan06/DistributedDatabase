package common

type Paging struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
	Total int    `json:"-"`
}
type Filler struct {
	Name        string `json:"name"`
	Day         int    `json:"day"`
	Coverage    string `json:"coverage"`
	PlanType    int    `json:"plan_type"`
	ProductType string `json:"product_type"`
}

func (p *Paging) SetPaging() {
	if p.Limit == 0 {
		p.Limit = 10
	}
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Sort == "" {
		p.Sort = "desc"
	}
}
