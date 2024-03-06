package lib

import (
	"github.com/abinercl/try-gin/src/model"
)

// album slice to seed record album  data.
var Albums = []model.Album{
	{ID: "1", Title: "Blue train", Artist: "jhon coltrane", Price: 56.99},
	{ID: "2", Title: "jeru", Artist: "Gerry Muligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vauhan and Cliford Brown", Artist: "Sarah Vauhan", Price: 39.99},
}
