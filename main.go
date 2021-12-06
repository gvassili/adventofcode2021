package main

import (
	"fmt"
	"github.com/gvassili/adventofcode2021/calendar"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"time"
)

type result struct {
	label    string
	result   string
	error    error
	duration time.Duration
}

func runChallenge(challenge calendar.Challenge) []result {
	var r *os.File
	if challenge.Day() != 0 {
		file, err := os.Open(fmt.Sprintf("./calendar/day%02d/input", challenge.Day()))
		if err != nil {
			return []result{{"prepare", "", err, 0}}
		}
		defer file.Close()
		r = file
	}
	startTs := time.Now()
	if err := challenge.Prepare(r); err != nil {
		return []result{{"prepare", "", err, time.Now().Sub(startTs)}}
	}
	prepareTs := time.Now()
	result1, err1 := challenge.Part1()
	part1Ts := time.Now()
	result2, err2 := challenge.Part2()
	part2Ts := time.Now()
	return []result{
		{"prepare", "", nil, prepareTs.Sub(startTs)},
		{"part1", result1, err1, part1Ts.Sub(prepareTs)},
		{"part2", result2, err2, part2Ts.Sub(part1Ts)},
	}
}

func main() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetAutoMergeCells(true)
	table.SetHeader([]string{"Day", "Step", "Result", "Error", "Excl time", "Incl time"})

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor, tablewriter.BgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor, tablewriter.BgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor, tablewriter.BgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor, tablewriter.BgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor, tablewriter.BgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor, tablewriter.BgCyanColor},
	)
	var challenges []calendar.Challenge
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			day, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}
			challenge, err := calendar.Load(day)
			if err != nil {
				panic(err)
			}
			challenges = append(challenges, challenge)
		}
	} else {
		challenges = calendar.LoadAllChallenges()
	}
	for idx, challenge := range challenges {
		blockFgColor := tablewriter.FgHiBlackColor
		blockMod := tablewriter.Normal
		if idx&1 == 1 {
			blockFgColor = tablewriter.FgWhiteColor
			blockMod = 2
		}
		var inclTime time.Duration
		for _, row := range runChallenge(challenge) {
			inclTime += row.duration
			fgColor := blockFgColor
			errMsg := ""
			if row.error != nil {
				fgColor = tablewriter.FgRedColor
				errMsg = row.error.Error()
			}
			table.Rich([]string{strconv.Itoa(challenge.Day()), row.label, row.result, errMsg, row.duration.String(), inclTime.String()}, []tablewriter.Colors{
				{blockMod, blockFgColor},
				{blockMod, fgColor},
				{blockMod, fgColor},
				{blockMod, fgColor},
				{blockMod, fgColor},
				{blockMod, fgColor},
			})
		}
	}
	table.Render()
}
