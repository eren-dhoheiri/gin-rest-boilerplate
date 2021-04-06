package mobil

import "time"

type GetMobilDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateMobilInput struct {
	ID        int
	Brand     int
	Type      string
	Years     string
	Name      string
	Price	  string
	Condition string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetBrandDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateBrandInput struct {
	ID        int
	Brand     string
	CreatedAt time.Time
	UpdatedAt time.Time
}