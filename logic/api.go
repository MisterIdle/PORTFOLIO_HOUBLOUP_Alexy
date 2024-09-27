package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCategoryData is a function that retrieves data from a specific category
func GetCategoryData(c *gin.Context) {
	category := c.Param("category")
	var entries [][]string

	tables := GetAllTablesNames()

	var tableExists bool
	for _, table := range tables {
		if table == category {
			tableExists = true
			break
		}
	}

	if !tableExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	entries, err := GetValuesFromTable(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving data from table"})
		return
	}

	columns, err := GetColumnNames(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving column names"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": category,
		"columns":  columns,
		"data":     entries,
	})
}
