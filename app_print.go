package galaxy

import (
	"encoding/json"
	"fmt"
	"github.com/pangdogs/galaxy/pt"
	"sort"
	"strings"
)

func (app *App) printComp() {
	var compPts []pt.ComponentPt

	pt.RangeCompPts(func(compPt pt.ComponentPt) bool {
		compPts = append(compPts, compPt)
		return true
	})

	sort.Slice(compPts, func(i, j int) bool {
		return strings.Compare(compPts[i].Api+compPts[i].Tag, compPts[j].Api+compPts[j].Tag) < 0
	})

	compPtsData, err := json.MarshalIndent(compPts, "", "\t")
	if err != nil {
		panic(fmt.Errorf("marshal components prototype info failed, %v", err))
	}

	fmt.Printf("%s", compPtsData)
}
