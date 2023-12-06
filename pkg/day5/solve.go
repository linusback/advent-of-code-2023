package day5

import (
	"advent-of-code-2023/pkg/util"
	"embed"
	"fmt"
	"slices"
	"time"
)

//go:embed *.txt
var f embed.FS

type MapWrapper struct {
	Name string
	Rows [][4]uint64
}

const maxUint64 uint64 = 1 << 63

var (
	seedToSoil,
	soilToFertilizer,
	fertilizerToWater,
	waterToLight,
	lightToTemperature,
	temperatureToHumidity,
	humidityToLocation,
	totalMap MapWrapper
	seeds        []uint64
	currentSeeds [][2]uint64
)

func (m *MapWrapper) Get(src uint64) uint64 {
	for i := 0; i < len(m.Rows); i++ {
		if m.Rows[i][1] <= src && src <= m.Rows[i][2] {
			return m.Rows[i][0] + (src - m.Rows[i][1])
		}
	}
	return src
}

func (m *MapWrapper) Compact() {
	m.Rows = compact4(m.Rows)
	//fmt.Printf("rows for %s are %v\n", m.Name, m.Rows)
}

func Solve() (err error) {
	var (
		b            []byte
		currentParse *MapWrapper

		result1, result2 uint64
	)
	result1 = maxUint64
	result2 = maxUint64
	totalMap.Name = "total"
	b, err = f.ReadFile("input.txt")
	//b, err = f.ReadFile("example.txt")
	if err != nil {
		return
	}
	startParse := time.Now()
	err = util.DoEachRowAll(b, func(row []byte, rows [][]byte, nr, total int) error {
		if nr == 0 {
			seeds = parseLine(row[7:])
			currentSeeds = make([][2]uint64, len(seeds)/2)
			for i := 0; i < len(seeds); i = i + 2 {
				currentSeeds[i/2][0] = seeds[i]
				currentSeeds[i/2][1] = seeds[i] + seeds[i+1] - 1
			}
			//fmt.Println("seeds ", seeds)
			//fmt.Println("seeds ", currentSeeds)
			return nil
		}
		if len(row) == 0 {
			return nil
		}
		switch string(row) {
		case "seed-to-soil map:":
			currentParse = &seedToSoil
			currentParse.Name = "seed-to-soil map"
			return nil
		case "soil-to-fertilizer map:":
			currentParse.Compact()
			currentSeeds = calculateSeeds(currentSeeds, currentParse)

			totalMap.Rows = currentParse.Rows
			currentParse = &soilToFertilizer
			currentParse.Name = "soil-to-fertilizer map"
			return nil
		case "fertilizer-to-water map:":
			currentParse.Compact()
			currentSeeds = calculateSeeds(currentSeeds, currentParse)
			currentParse = &fertilizerToWater
			currentParse.Name = "fertilizer-to-water map"
			return nil
		case "water-to-light map:":
			currentParse.Compact()
			currentSeeds = calculateSeeds(currentSeeds, currentParse)

			currentParse = &waterToLight
			currentParse.Name = "water-to-light map"
			return nil
		case "light-to-temperature map:":
			currentParse.Compact()
			currentSeeds = calculateSeeds(currentSeeds, currentParse)

			currentParse = &lightToTemperature
			currentParse.Name = "light-to-temperature map"
			return nil
		case "temperature-to-humidity map:":
			currentParse.Compact()
			currentSeeds = calculateSeeds(currentSeeds, currentParse)

			currentParse = &temperatureToHumidity
			currentParse.Name = "temperature-to-humidity map"
			return nil
		case "humidity-to-location map:":
			currentParse.Compact()
			currentSeeds = calculateSeeds(currentSeeds, currentParse)

			currentParse = &humidityToLocation
			currentParse.Name = "humidity-to-location map"
			return nil
		default:
			l := parseLine4(row)
			parseMapLine(l, currentParse, currentSeeds)
			return nil
		}

	})
	currentParse.Compact()
	currentSeeds = calculateSeeds(currentSeeds, currentParse)
	slices.SortFunc(currentSeeds, func(a, b [2]uint64) int {
		return int(a[0]) - int(b[0])
	})
	result2 = currentSeeds[0][0]
	fmt.Println("time parsing", time.Since(startParse))
	//fmt.Println(currentSeeds)
	//fmt.Println("parsing done")
	//fmt.Println(seedToSoil)
	//fmt.Println(soilToFertilizer)
	//fmt.Println(fertilizerToWater)
	//fmt.Println(waterToLight)
	//fmt.Println(lightToTemperature)
	//fmt.Println(temperatureToHumidity)
	//fmt.Println(humidityToLocation)

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
		if v < result1 {
			result1 = v
		}
	}
	//var newTotal, changedValues uint64
	//for i := 0; i < len(seeds); i = i + 2 {
	//	var j uint64 = 0
	//	newTotal += seeds[i+1]
	//	fmt.Println("from ", seeds[i], " to ", seeds[i]+seeds[i+1], " total ", seeds[i+1])
	//	for ; j < seeds[i+1]; j++ {
	//		v := getValue(seeds[i]+j,
	//			&seedToSoil,
	//			&soilToFertilizer,
	//			&fertilizerToWater,
	//			&waterToLight,
	//			&lightToTemperature,
	//			&temperatureToHumidity,
	//			&humidityToLocation,
	//		)
	//		if v != seeds[i]+j {
	//			changedValues++
	//		}
	//		if v < result2 {
	//			result2 = v
	//		}
	//
	//	}
	//}
	//fmt.Printf("total number of peaple %d, people that changed value %d", newTotal, changedValues)
	fmt.Println()
	fmt.Println(result1)
	fmt.Println(result2)
	return
}

func calculateSeeds(source [][2]uint64, parse *MapWrapper) [][2]uint64 {
	var (
		diff uint64
		s    [2]uint64
		r    [4]uint64
	)
	res := make([][2]uint64, 0, len(source)*2)
outer:
	for i := 0; i < len(source); i++ {
		s = source[i]
		//fmt.Println("checking", s)
		for j := 0; j < len(parse.Rows); j++ {

			r = parse.Rows[j]
			//fmt.Println("against", r)
			//should be sorted end of seed is less than start of map so break
			if s[1] < r[2] {
				//fmt.Println("less")
				res = append(res, s)
				continue outer
			}
			if s[0] <= r[3] && r[2] <= s[1] {
				//fmt.Println("match")
				if s[0] < r[2] {
					res = append(res, [2]uint64{s[0], r[2] - 1})
					s[0] = r[2]
				}

				if s[1] <= r[3] {
					//fmt.Println("perfect")
					diff = s[1] - s[0]
					res = append(res, [2]uint64{r[0], r[0] + diff})
					continue outer
				}
				res = append(res, [2]uint64{r[0], r[1]})
				s[0] = r[2]
			}
		}
		res = append(res, s)
	}
	//fmt.Println("new seeds ", res)
	return res
}

func compact(res [][2]uint64) [][2]uint64 {
	lenR := len(res)
outer:
	for i := 0; i < lenR; i++ {
		for j := i + 1; j < lenR; j++ {
			//has any overlap
			if res[i][0] <= res[j][1] && res[j][0] <= res[i][1] {
				fmt.Println("compacting")
				if res[j][0] < res[i][0] {
					res[i][0] = res[j][0]
				}
				if res[j][1] > res[i][1] {
					res[i][1] = res[j][1]
				}
				copy(res[j:], res[j+1:])
				i--
				lenR--
				res = res[:lenR]
				continue outer
			}
		}
	}
	return res
}

func compact4(res [][4]uint64) [][4]uint64 {
	lenR := len(res)
outer:
	for i := 0; i < lenR; i++ {
		for j := i + 1; j < lenR; j++ {
			//has any overlap
			if res[i][2] <= res[j][3] && res[j][2] <= res[i][3] {
				//fmt.Println("compacting")
				if res[j][2] < res[i][2] {
					res[i][2] = res[j][2]
					res[i][0] = res[j][0]
				}
				if res[j][3] > res[i][3] {
					res[i][3] = res[j][3]
					res[i][1] = res[j][1]
				}
				copy(res[j:], res[j+1:])
				i--
				lenR--
				res = res[:lenR]
				continue outer
			}
		}
	}
	slices.SortFunc(res, func(a, b [4]uint64) int {
		return int(a[2]) - int(b[2])
	})
	return res
}

func contains(src []uint64, val uint64) bool {
	for i := 0; i < len(src); i = i + 2 {
		if src[i] <= val && val < src[i+1] {
			return true
		}
	}
	return false
}

func populateTotalMap(total *MapWrapper, m *MapWrapper) {
	existing := make([][3]uint64, len(m.Rows))
	for i := 0; i < len(m.Rows); i++ {
		copy(existing[i][:], m.Rows[i][:])
	}

	//copy(existing)
	//existing := m.Rows
	//fmt.Println("populating", m.Rows)

	//fmt.Println(m.Rows)
	//var i uint64
	//for ; i < line[2]; i++ {
	//	fmt.Println("is run")
	//}
}

func getValue(src uint64, maps ...*MapWrapper) uint64 {
	for i := 0; i < len(maps); i++ {
		src = maps[i].Get(src)
	}
	return src
}

func parseMapLine(l [4]uint64, m *MapWrapper, currentSeed [][2]uint64) {
	var (
		l2   [4]uint64
		diff uint64
	)
	for i := 0; i < len(currentSeed); i++ {
		//has overlap
		if currentSeed[i][0] <= l[3] && l[2] <= currentSeed[i][1] {
			copy(l2[:], l[:])
			if l2[2] < currentSeed[i][0] {
				diff = currentSeed[i][0] - l2[2]
				l2[2] += diff
				l2[0] += diff
			}
			if l2[3] > currentSeed[i][1] {
				diff = l2[3] - currentSeed[i][1]
				l2[3] -= diff
				l2[1] -= diff
			}
			m.Rows = append(m.Rows, l2)
		}
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

func parseLine4(line []byte) [4]uint64 {
	var (
		r [4]uint64
		j int
	)

	start := 0
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			u := util.ParseUint64NoError(line[start:i])
			r[j] = u
			j++
			start = i + 1
		}
	}
	u := util.ParseUint64NoError(line[start:])
	r[j] = u
	r[3] = r[2]
	r[2] = r[1]
	r[1] = r[0] + r[3] - 1
	r[3] = r[2] + r[3] - 1
	return r
}
