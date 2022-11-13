package launcher

import (
	"encoding/json"
	"fmt"
	"github.com/galaxy-kit/galaxy/pt"
	"sort"
	"strings"
)

func (app *App) printComp() {
	var compPts []pt.ComponentPt

	pt.RangeComponentPts(func(compPt pt.ComponentPt) bool {
		compPts = append(compPts, compPt)
		return true
	})

	sort.Slice(compPts, func(i, j int) bool {
		return strings.Compare(compPts[i].Name+compPts[i].Path, compPts[j].Name+compPts[j].Path) < 0
	})

	compPtsData, err := json.MarshalIndent(compPts, "", "\t")
	if err != nil {
		panic(fmt.Errorf("marshal components prototype info failed, %v", err))
	}

	fmt.Printf("%s", compPtsData)
}
