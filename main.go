package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	files := os.Args
	file, err := excelize.OpenFile(files[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, _ := file.GetRows("Sheet1")
	header := rows[0]
	endpoint := os.Getenv("ENDPOINT_URL") + "/api/v2/issues"
	api_key := os.Getenv("API_KEY")

	for row_i, row := range rows {
		req, err := http.NewRequest("POST", endpoint, nil)
		if err != nil {
			log.Fatal(err)
		}
		ps := req.URL.Query()
		if row_i > 0 {
			for col_i, colCell := range row {
				ps.Add(header[col_i], colCell)
			}
			ps.Add("apiKey", api_key)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.URL.RawQuery = ps.Encode()
			fmt.Println(string(ps.Encode()))

			if err != nil {
				log.Fatal(err)
			}

			client := new(http.Client)
			res, err := client.Do(req)
			defer res.Body.Close()

			if err != nil {
				log.Fatal(err)
			}

			byteArray, _ := ioutil.ReadAll(res.Body)
			fmt.Println(string(byteArray))
		}
	}
}
