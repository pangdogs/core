package launcher

import (
	"encoding/json"
	"fmt"
	"github.com/golaxy-kit/golaxy/pt"
	"sort"
	"strings"
)

func (app *_App) printComp() {
	var compPts []pt.ComponentPt

	pt.RangeComponent(func(compPt pt.ComponentPt) bool {
		compPts = append(compPts, compPt)
		return true
	})

	sort.Slice(compPts, func(i, j int) bool {
		return strings.Compare(compPts[i].Name, compPts[j].Name) < 0
	})

	compPtsDs, err := json.MarshalIndent(compPts, "", "\t")
	if err != nil {
		panic(fmt.Errorf("marshal components info failed, %v", err))
	}

	fmt.Printf("%s", compPtsDs)
}

func (app *_App) printPt(ptPath string) {
	servicePt := app.loadPtConfig(ptPath)

	servicePtDs, err := json.MarshalIndent(servicePt, "", "\t")
	if err != nil {
		panic(fmt.Errorf("marshal service pt info failed, %v", err))
	}

	fmt.Printf("%s", servicePtDs)
}
