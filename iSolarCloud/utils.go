package iSolarCloud

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)


// writeTableTabular outputs tabular data to STDOUT
func writeTableTabular(data [][]string, cols ...string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(cols)
	table.SetBorder(false)
	table.AppendBulk(data)

	table.Render()
}

// writeTableJSON outputs JSON data to STDOUT
func writeTableJSON(data [][]string, cols ...string) {
	mappedData := make([]map[string]string, 0)
	for i := range data {
		rowData := make(map[string]string)
		for j := range data[i] {
			rowData[cols[j]] = data[i][j]
		}
		mappedData = append(mappedData, rowData)
	}
	jsonData, err := json.Marshal(mappedData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}

// writeTable outputs JSON or tabular data to STDOUT
func writeTable(outType OutputType, data [][]string, cols ...string) {
	writeTableTabular(data, cols...)
}

//func EntityFilename(entity string) string {
//	return strings.TrimSuffix(entity, ".json") + ".json"
//}
