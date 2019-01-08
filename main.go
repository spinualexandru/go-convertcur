package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-convertcur"
	app.Usage = "VALUE BASE_CURRENCY to TARGET_CURRENCY"
	app.Action = func(c *cli.Context) error {
		amount := c.Args().Get(0)
		base := c.Args().Get(1)
		target := c.Args().Get(3)

		response, err := http.Get(fmt.Sprintf("https://api.exchangeratesapi.io/latest?symbols=%s&base=%s", target, base))
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			var result map[string]interface{}

			json.Unmarshal([]byte(string(contents)), &result)
			decodedResponse := result["rates"].(map[string]interface{})

			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}

			for _, value := range decodedResponse {
				amountNr, err := strconv.ParseFloat(amount, 10)
				if err != nil {
					fmt.Printf("%s", err)
					os.Exit(1)
				}
				fmt.Println(fmt.Sprintf("%s %s = %f %s", amount, base, float64(value.(float64)*float64(amountNr)), target))
			}
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
