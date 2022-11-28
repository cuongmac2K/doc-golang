package luu

import "fmt"

func FirstOrCreate(botID int, position int, rd RdScriptGroup) (RdScriptGroup, error) {
	fmt.Println("rdf la ", rd)
	err := DB.Where(map[string]interface{}{
		"block_id": botID,
		"position": position,
	}).FirstOrCreate(&rd).Debug()
	return rd, err.Error
}
