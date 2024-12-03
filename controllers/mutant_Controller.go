package controllers

import (
	"net/http"
	"nicolascastro/go/isMutant/database"
	"nicolascastro/go/isMutant/models"

	"github.com/gin-gonic/gin"
)

func CheckMutant(c *gin.Context) {
	var dna models.DNA

	// Validar el JSON recibido
	if err := c.ShouldBindJSON(&dna); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar si es un mutante
	isMutant := models.IsMutant(dna.Sequences)

	// Convertir las secuencias de ADN en una única cadena para almacenamiento
	sequence := ""
	for _, s := range dna.Sequences {
		sequence += s
	}

	// Crear el registro de ADN
	dnaRecord := database.DNARecord{
		Sequence: sequence,
		IsMutant: isMutant,
	}

	// Guardar en la base de datos
	if err := database.SaveDNARecord(dnaRecord); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to save DNA sequence",
			"details": err.Error(), // Opcional: Detalle del error para depuración
		})
		return
	}

	// Responder según si es mutante o no
	if isMutant {
		c.JSON(http.StatusOK, gin.H{"message": "Mutant detected"})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"message": "Not a mutant, is a human"})
	}
}
