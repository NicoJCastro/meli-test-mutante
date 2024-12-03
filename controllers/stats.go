package controllers

import (
	"net/http"

	"nicolascastro/go/isMutant/database"

	"github.com/gin-gonic/gin"
)

func GetStats(c *gin.Context) {
	var countMutantDNA int64
	var countHumanDNA int64

	database.DB.Model(&database.DNARecord{}).Where("is_mutant = ?", true).Count(&countMutantDNA)
	database.DB.Model(&database.DNARecord{}).Where("is_mutant = ?", false).Count(&countHumanDNA)

	var ratio float64
	if countHumanDNA == 0 {
		ratio = 0
	} else {
		ratio = float64(countMutantDNA) / float64(countHumanDNA)
	}

	c.JSON(http.StatusOK, gin.H{
		"count_mutant_dna": countMutantDNA,
		"count_human_dna":  countHumanDNA,
		"ratio":            ratio,
	})
}
