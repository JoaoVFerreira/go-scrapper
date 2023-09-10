package output

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/JoaoVFerreira/GOScrapper/db"
)

type Result struct {
	Code         string         `json:"code"`
	ResultDetail resultDetail   `json:"result_detail"`
}

type resultDetail struct {
	DividendYield float64 `json:"dividen_yield"`
	Liquidity     float64 `json:"liquidity"`
	Pvp           float64 `json:"pvp"`
	Type          string  `json:"type"`
	Description   string  `json:"description"`
}


func Decision() {
	var f db.RealStateFund
	funds, err := f.FindAllWithData(); if err != nil {
		panic(err)
	} 
	resultRaw := []Result{}

	for _, fund := range funds {
		r := Result{}

		dy, err := formatPercentage(fund.DividendYield); if err != nil {
			fmt.Println(err)
		}

		pvp, err := formatPvp(fund.Pvp); if err != nil {
			fmt.Println(err)
		} 

		liq, err := parseStringToFloat(fund.Liquidity); if err != nil {
			fmt.Println(err)
		}

		r.Code = fund.Code
		r.ResultDetail.DividendYield = dy
		r.ResultDetail.Liquidity = liq
		r.ResultDetail.Pvp = pvp
		r.ResultDetail.Type = fund.Type
		r.ResultDetail.Description = fund.Description

		resultRaw = append(resultRaw, r)
	}

	resultJSON, _ := json.MarshalIndent(resultRaw, "", "  ")
	os.WriteFile("result.json", resultJSON, 0666) 
}

func formatPercentage(input string) (float64, error) {
	input = strings.ReplaceAll(input, " ", "")
	input = strings.Replace(input, ",", ".", 1)
	input = strings.TrimRight(input, "%")
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0.0, err
	}
	return value, nil
}

func parseStringToFloat(input string) (float64, error) {
	input = strings.ReplaceAll(input, " ", "")
	input = strings.Replace(input, ",", ".", 1)
	input = strings.TrimRight(input, "KkMm")
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0.0, err
	}
	if strings.HasSuffix(input, "K") || strings.HasSuffix(input, "k") {
		value *= 1000.0
	} else if strings.HasSuffix(input, "M") || strings.HasSuffix(input, "m") {
		value *= 1000000.0
	}
	return value, nil
}

func formatPvp(input string) (float64, error) {
	input = strings.TrimSpace(input)
	input = strings.TrimRight(input, " \t\n\r")
	input = strings.Replace(input, ",", ".", 1)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0.0, err
	}
	return value, nil
}