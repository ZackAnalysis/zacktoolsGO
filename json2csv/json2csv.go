package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/rocketlaunchr/dataframe-go/exports"
	"github.com/rocketlaunchr/dataframe-go/imports"
)

var ctx = context.Background()

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input your json file path (records orient or jsonl file) : ")
	jsonfile, _ := reader.ReadString('\n')
	jsonfile = strings.TrimSpace(jsonfile)
	if !strings.HasSuffix(jsonfile, ".json") && !strings.HasSuffix(jsonfile, ".jsonl") {
		log.Fatal("only works with \".json\" or \".jsonl\" file")
	}
	fmt.Println("file path is \n", jsonfile)
	jsondata, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		log.Fatal(err)
	}
	outfile := strings.Replace(string(jsonfile), ".json", ".csv", 1)
	outfile = strings.Replace(string(outfile), ".jsonl", ".csv", 1)
	outfile = strings.TrimSpace(outfile)

	fmt.Println(outfile)

	data, err := imports.LoadFromJSON(ctx, strings.NewReader(string(jsondata)))
	if err != nil {
		log.Fatal(err)
	}
	out, err := os.Create(outfile)
	if err != nil {
		log.Fatal(err)
	}
	exports.ExportToCSV(ctx, out, data)
	fmt.Println("outfile is at ", outfile)

}
