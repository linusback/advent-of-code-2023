package day5

import (
	"advent-of-code-2023/pkg/util"
	"embed"
	"fmt"
)

//go:embed *.txt
var f embed.FS

type MapWrapper struct {
	Name string
	//Rows [][]uint64
	Map map[uint64]uint64
}

var (
	seedToSoil,
	soilToFertilizer,
	fertilizerToWater,
	waterToLight,
	lightToTemperature,
	temperatureToHumidity,
	humidityToLocation MapWrapper
	seeds []uint64
)

func (m *MapWrapper) Get(src uint64) uint64 {
	u, ok := m.Map[src]
	if ok {
		return u
	}
	return src
}

func Solve() (err error) {
	var (
		b []byte
	)
	b, err = f.ReadFile("input.txt")
	if err != nil {
		return
	}
	var currentParse *MapWrapper
	err = util.DoEachRowAll(b, func(row []byte, rows [][]byte, nr, total int) error {
		if nr == 0 {
			seeds = parseLine(row[7:])
			fmt.Println("parsed seeds")
			return nil
		}
		if len(row) == 0 {
			return nil
		}
		switch string(row) {
		case "seed-to-soil map:":
			currentParse = &seedToSoil
			currentParse.Map = make(map[uint64]uint64)
			fmt.Println("parsing seed to soil")
			return nil
		case "soil-to-fertilizer map:":
			currentParse = &soilToFertilizer
			currentParse.Map = make(map[uint64]uint64)
			fmt.Println("parsing seed to ferti")
			return nil
		case "fertilizer-to-water map:":
			currentParse = &fertilizerToWater
			currentParse.Map = make(map[uint64]uint64)
			fmt.Println("parsing seed to water")
			return nil
		case "water-to-light map:":
			currentParse = &waterToLight
			currentParse.Map = make(map[uint64]uint64)
			fmt.Println("parsing seed to light")
			return nil
		case "light-to-temperature map:":
			currentParse = &lightToTemperature
			currentParse.Map = make(map[uint64]uint64)
			fmt.Println("parsing seed to temp")
			return nil
		case "temperature-to-humidity map:":
			currentParse = &temperatureToHumidity
			currentParse.Map = make(map[uint64]uint64)
			fmt.Println("parsing seed to hum")
			return nil
		case "humidity-to-location map:":
			currentParse = &humidityToLocation
			currentParse.Map = make(map[uint64]uint64)
			fmt.Println("parsing seed to loc")
			return nil
		default:
			parseMapLine(row, currentParse)
			fmt.Println("parsing line")
			return nil
		}

	})
	fmt.Printf("parsing done")
	//fmt.Println(seedToSoil)
	//fmt.Println(soilToFertilizer)
	//fmt.Println(fertilizerToWater)
	//fmt.Println(waterToLight)
	//fmt.Println(lightToTemperature)
	//fmt.Println(temperatureToHumidity)
	//fmt.Println(humidityToLocation)
	var minVal uint64 = 1
	minVal = minVal << 63
	for _, seed := range seeds {
		v := getValue(seed,
			&seedToSoil,
			&soilToFertilizer,
			&fertilizerToWater,
			&waterToLight,
			&lightToTemperature,
			&temperatureToHumidity,
			&humidityToLocation,
		)
		if v < minVal {
			minVal = v
		}
	}
	fmt.Println(minVal)
	return
}

func getValue(src uint64, maps ...*MapWrapper) uint64 {
	for i := 0; i < len(maps); i++ {
		src = maps[i].Get(src)
	}
	return src
}

func parseMapLine(line []byte, m *MapWrapper) {
	l := parseLine(line)
	if len(l) != 3 {
		panic(fmt.Sprintf("map line not 3 got %d\n row: %s", len(l), string(line)))
	}
	var i uint64
	for ; i < l[2]; i++ {
		//fmt.Println("Setting", l[1]+i, l[0]+i)
		m.Map[l[1]+i] = l[0] + i
	}
}

func parseLine(line []byte) []uint64 {
	var r []uint64
	start := 0
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			u := util.ParseUint64NoError(line[start:i])
			r = append(r, u)
			start = i + 1
		}
	}
	u := util.ParseUint64NoError(line[start:])
	r = append(r, u)
	return r
}
