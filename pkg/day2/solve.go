package day2

import (
	"advent-of-code-2023/pkg/util"
	"bytes"
	"embed"
	"fmt"
)

//go:embed input.txt
//go:embed example.txt
var f embed.FS

type Game struct {
	Id     int64
	Rounds []Round
}

type Round struct {
	Red, Blue, Green int64
}

func Solve1() (err error) {
	var (
		b            []byte
		total, power int64
	)
	b, err = f.ReadFile("input.txt")
	if err != nil {
		return
	}
	var games []Game
	err = util.DoEachRow(b, func(row []byte, nr int) error {
		var game Game
		id, inc, err2 := ParseInt64(row[5:])
		if err2 != nil {
			return err2
		}
		game.Id = id
		//fmt.Println("game id", game.Id)
		err2 = ParseRounds(&game.Rounds, row[6+inc:])
		if err2 != nil {
			return err2
		}
		games = append(games, game)
		return nil
	})

	for _, game := range games {
		if ValidGame(&game, 12, 13, 14) {
			total += game.Id
		}
		power += PowerOfGame(&game)
	}
	fmt.Printf("power is %d\n", power)
	fmt.Printf("total is %d\n", total)
	return
}
func ValidGame(g *Game, red, green, blue int64) bool {
	for _, round := range g.Rounds {
		if round.Red > red || round.Green > green || round.Blue > blue {
			return false
		}
	}
	return true
}

func PowerOfGame(g *Game) int64 {
	var red, green, blue int64
	for _, round := range g.Rounds {
		if round.Red > red {
			red = round.Red
		}
		if round.Green > green {
			green = round.Green
		}
		if round.Blue > blue {
			blue = round.Blue
		}
	}
	return red * green * blue
}

func ParseInt64(sub []byte) (res int64, inc int, err error) {
	var total []byte
	i := 0
	for ; i < len(sub); i++ {
		if '0' > sub[i] || sub[i] > '9' {
			break
		}
		total = append(total, sub[i])
	}
	inc = i
	if len(total) == 0 {
		err = fmt.Errorf("could not parse int %s", string(sub))
		return
	}
	var mult int64 = 1

	for i = len(total) - 1; i >= 0; i-- {
		res += int64(total[i]-'0') * mult
		mult *= 10
	}

	return
}
func ParseRounds(g *[]Round, row []byte) (err error) {
	var color []byte
	round := bytes.Split(row, []byte{';'})
	*g = make([]Round, len(round))
	for i := 0; i < len(round); i++ {
		//fmt.Printf("round |%s|", string(round[i]))
	roundLoop:
		for k := 0; k < len(round[i]); k++ {
			if round[i][k] == ' ' {
				continue
			}
			color = color[:0]
			amount, inc, err2 := ParseInt64(round[i][k:])
			if err2 != nil {
				return err2
			}
			k += inc
			for ; k < len(round[i]); k++ {
				if round[i][k] == ',' || round[i][k] == ';' {
					break
				}
				if round[i][k] < 'a' || round[i][k] > 'z' {
					continue
				}
				color = append(color, round[i][k])
			}
			switch string(color) {
			case "red":
				(*g)[i].Red = amount
				continue roundLoop
			case "blue":
				(*g)[i].Blue = amount
				continue roundLoop
			case "green":
				(*g)[i].Green = amount
				continue roundLoop
			default:
				err = fmt.Errorf("failed to parse color %s", string(color))
			}
		}

	}
	return nil
}

func Solve2() (err error) {
	return
}
