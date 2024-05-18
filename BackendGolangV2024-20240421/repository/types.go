// This file contains types that are used in the repository layer.
package repository

type (
	FindStatsResponse struct {
		Count int `json:"count"`
		Max   int `json:"max"`
		Min   int `json:"min"`
	}
)
