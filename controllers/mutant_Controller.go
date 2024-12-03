package controllers

import (
	"net/http"
	"nicolascastro/go/isMutant/database"
	"nicolascastro/go/isMutant/models"

	"github.com/gin-gonic/gin"
)

func CheckMutant(c *gin.Context) {
	var dna models.DNA

	if err := c.ShouldBindJSON(&dna); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isMutant := models.IsMutant(dna.Sequences)

	// Convert the DNA sequences to a single string for storage
	sequence := ""
	for _, s := range dna.Sequences {
		sequence += s
	}

	dnaRecord := database.DNARecord{
		Sequence: sequence,
		IsMutant: isMutant,
	}

	if err := database.SaveDNARecord(dnaRecord); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save DNA sequence"})
		return
	}

	if isMutant {
		c.JSON(http.StatusOK, gin.H{"message": "Mutant detected"})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"message": "Not a mutant"})
	}
}
