package externalapis

import (
	"strings"

	"TPs-DDS/TP1/development/entities"
)

var contributorMap = map[string]entities.AFIPContributor{
	"20-37282752-2": {"20-37282752-2", "Mauricio", "Truchet"},
	"21-12345678-2": {"21-12345678-2", "Pepe", "Diaz"},
}

func GetContributor(cuit string) *entities.AFIPContributor {
	for contributor, _ := range contributorMap {
		if strings.Compare(cuit, contributor) == 0 {
			return &entities.AFIPContributor{
				CUIT:     contributorMap[cuit].CUIT,
				Name:     contributorMap[cuit].Name,
				LastName: contributorMap[cuit].LastName,
			}
		}
	}
	return nil
}
