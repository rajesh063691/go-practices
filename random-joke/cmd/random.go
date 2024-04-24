/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long: `It gives you a random joke or a random quote:

	It can be used to lighten up your mood or to get some motivation. It can be used in your scripts to get a random joke or quote.`,
	Run: func(cmd *cobra.Command, args []string) {
		joke := getRandomJoke()

		fmt.Println(joke)
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getRandomJoke() string {
	endURL := "https://icanhazdadjoke.com/"

	data, err := callEndPointURI(endURL)
	if err != nil {
		return "Error in calling endpoint - getting joke"
	}

	joke := &Joke{}
	err = json.Unmarshal(data, &joke)
	if err != nil {
		return "Error in unmarshalling -  getting joke"
	}

	return joke.Joke
}

func callEndPointURI(endpoint string) ([]byte, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("Accept", "application/json")

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

type Joke struct {
	ID     string `json:"id,omitempty"`
	Joke   string `json:"joke,omitempty"`
	Status int    `json:"status,omitempty"`
}
