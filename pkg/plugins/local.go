package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloud-native-application/rudrx/api/types"
)

func GetDefFromLocal(dir string, defType types.DefinitionType) ([]types.Template, error) {
	temps, err := LoadTempFromLocal(dir)
	if err != nil {
		return nil, err
	}
	var defs []types.Template
	for _, t := range temps {
		if t.Type != defType {
			continue
		}
		defs = append(defs, t)
	}
	return defs, nil
}

func SinkTemp2Local(templates []types.Template, dir string) int {
	success := 0
	for _, tmp := range templates {
		data, err := json.Marshal(tmp)
		if err != nil {
			fmt.Printf("sync %s err: %v\n", tmp.Name, err)
			continue
		}
		err = ioutil.WriteFile(filepath.Join(dir, tmp.Name), data, 0644)
		if err != nil {
			fmt.Printf("sync %s err: %v\n", tmp.Name, err)
			continue
		}
		success++
	}
	return success
}

func LoadTempFromLocal(dir string) ([]types.Template, error) {
	var tmps []types.Template
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("\"no definition files found, use 'vela refresh' to sync from cluster\"")
			return nil, nil
		}
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if strings.HasSuffix(f.Name(), ".cue") {
			continue
		}
		data, err := ioutil.ReadFile(filepath.Join(dir, f.Name()))
		if err != nil {
			fmt.Printf("read file %s err %v\n", f.Name(), err)
			continue
		}
		var tmp types.Template
		decoder := json.NewDecoder(bytes.NewBuffer(data))
		decoder.UseNumber()
		if err = decoder.Decode(&tmp); err != nil {
			fmt.Printf("ignore invalid format file: %s\n", f.Name())
			continue
		}
		tmps = append(tmps, tmp)
	}
	if len(tmps) == 0 {
		fmt.Println("\"no definition files found, use 'vela refresh' to sync from cluster\"")
	}
	return tmps, nil
}
