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
	Rows [][]uint64
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
	for i := 0; i < len(m.Rows); i++ {
		if m.Rows[i][1] <= src && src < m.Rows[i][3] {
			return m.Rows[i][0] + (src - m.Rows[i][1])
		}
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
			//fmt.Println("parsed seeds")
			return nil
		}
		if len(row) == 0 {
			return nil
		}
		switch string(row) {
		case "seed-to-soil map:":
			currentParse = &seedToSoil
			currentParse.Name = "seed-to-soil map"
			//fmt.Println("parsing seed to soil")
			return nil
		case "soil-to-fertilizer map:":
			currentParse = &soilToFertilizer
			currentParse.Name = "soil-to-fertilizer map"
			//fmt.Println("parsing seed to ferti")
			return nil
		case "fertilizer-to-water map:":
			currentParse = &fertilizerToWater
			currentParse.Name = "fertilizer-to-water map"
			//fmt.Println("parsing seed to water")
			return nil
		case "water-to-light map:":
			currentParse = &waterToLight
			currentParse.Name = "water-to-light map"
			//fmt.Println("parsing seed to light")
			return nil
		case "light-to-temperature map:":
			currentParse = &lightToTemperature
			currentParse.Name = "light-to-temperature map"
			//fmt.Println("parsing seed to temp")
			return nil
		case "temperature-to-humidity map:":
			currentParse = &temperatureToHumidity
			currentParse.Name = "temperature-to-humidity map"
			//fmt.Println("parsing seed to hum")
			return nil
		case "humidity-to-location map:":
			currentParse = &humidityToLocation
			currentParse.Name = "humidity-to-location map"
			//fmt.Println("parsing seed to loc")
			return nil
		default:
			parseMapLine(row, currentParse)
			//fmt.Println("parsing line")
			return nil
		}

	})
	//fmt.Println("parsing done")
	//fmt.Println(seedToSoil)
	//fmt.Println(soilToFertilizer)
	//fmt.Println(fertilizerToWater)
	//fmt.Println(waterToLight)
	//fmt.Println(lightToTemperature)
	//fmt.Println(temperatureToHumidity)
	//fmt.Println(humidityToLocation)
	var minVal, minVal2 uint64
	minVal = 1 << 63
	minVal2 = minVal
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
	for i := 0; i < len(seeds); i = i + 2 {
		var j uint64 = 0
		fmt.Println("from ", seeds[i], " to ", seeds[i]+seeds[i+1], " total ", seeds[i+1])
		for ; j < seeds[i+1]; j++ {
			v := getValue(seeds[i]+j,
				&seedToSoil,
				&soilToFertilizer,
				&fertilizerToWater,
				&waterToLight,
				&lightToTemperature,
				&temperatureToHumidity,
				&humidityToLocation,
			)
			if v < minVal2 {
				minVal2 = v
			}

		}
	}

	fmt.Println(minVal)
	fmt.Println(minVal2)
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
	l = append(l, l[1]+l[2])
	m.Rows = append(m.Rows, l)
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
