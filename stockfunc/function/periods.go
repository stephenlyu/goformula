package function

import (
	. "github.com/stephenlyu/tds/period"
	"errors"
	"baiwenbao.com/arbitrage/util"
)

var (
	_, M1 = PeriodFromString("M1")
	_, M5 = PeriodFromString("M5")
	_, M15 = PeriodFromString("M15")
	_, M30 = PeriodFromString("M30")
	_, M60 = PeriodFromString("M60")
	_, D1 = PeriodFromString("D1")
	_, W1 = PeriodFromString("W1")
)

var systemPeriods = []Period {
	M1,
	M5,
	M15,
	M30,
	M60,
	D1,
	W1,
}

var periodIndices map[string]int

func init() {
	periodIndices = buildPeriodIndices(systemPeriods)
}

func buildPeriodIndices(periods []Period) map[string]int {
	m := make(map[string]int)
	for i, p := range periods {
		m[p.ShortName()] = i
	}
	return m
}

func SetCustomPeriods(periods []Period) error {
	var ps []Period
	ps = append(ps, systemPeriods...)
	ps = append(ps, periods...)
	m := buildPeriodIndices(ps)
	if len(m) < len(ps) {
		return errors.New("Duplicate period")
	}

	periodIndices = m
	return nil
}

func GetPeriodIndex(period Period) int {
	index, ok := periodIndices[period.ShortName()]
	util.Assert(ok, "Bad period")
	return index
}
