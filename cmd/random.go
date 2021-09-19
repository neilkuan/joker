/*
Copyright © 2021 Neil Kuan <neilkuanpro@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
//nolint:unparam
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random Joke",
	Long: `A longer description that Get a random Joke. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := cmd.Flags().GetInt64("count")
		if args == nil {
			fmt.Println("This is Joke Command: ...")
		}
		fmt.Println("This is Joke Command: ...")
		if err != nil {
			fmt.Println(err)
		}
		JokeList, err := getRandomJoke(c)
		if err != nil {
			fmt.Println(err)
		}
		for _, j := range JokeList {
			fmt.Println(string(j.Joke))
		}
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Need to define count flag
	randomCmd.PersistentFlags().Int64("count", 1, "Input the number of joke")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke(count int64) ([]Joke, error) {
	var joke Joke
	var jokeList []Joke

	for i := int64(0); i < count; i++ {
		err := json.Unmarshal(getJokeData("https://icanhazdadjoke.com/"), &joke)
		if err != nil {
			log.Printf("Cloud not unmarshal joke - %v", err)
		}
		jokeList = append(jokeList, joke)
	}

	return jokeList, nil
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Cloud not request a dadjoke - %v", err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "joker CLI (github.com)")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Cloud not request a dadjoke - %v", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Cloud not Read Response Body - %v", err)
	}
	return respBytes
}
