/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"

	"encoding/json"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.PersistentFlags().StringP("pat", "p", "", "peresonal access tken of the repository")
	//rootCmd.PersistentFlags().StringP("password", "p", "", "username of the repository")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "provides git repo details",
	Long:  `Get Repo information using the Cobra Command`,
	Run: func(cmd *cobra.Command, args []string) {
		pat, _ := cmd.Flags().GetString("pat")
		//password, _ := cmd.Flags().GetString("password")

		//cred := fmt.Sprint(username, ":", password)
		//base64Cred := b64.StdEncoding.EncodeToString([]byte(pat))
		//fmt.Println(base64Cred)
		getRepoDetails(pat)
	},
}

func getRepoDetails(pat string) {
	fmt.Println("Get Repo Details")
	repoURL := "https://api.github.com/user/repos"

	client := http.DefaultClient
	req, _ := http.NewRequest("GET", repoURL, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", pat))

	var result []interface{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error in getting the response :: %s", err.Error())
		return
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error in reading the response :: %s", err.Error())
		return
	}

	if res.StatusCode != 200 {
		fmt.Printf("Error in getting the response and Status code :: %d and data :: %s \n", res.StatusCode, string(data))
		return

	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Printf("Error in unmarshalling the response :: %s", err.Error())
		return
	}
	for i, repoDetails := range result {
		repo := repoDetails.(map[string]interface{})
		// fmt.Println(" name : ", repo["name"], " private :", repo["private"], "clone_url :", repo["clone_url"])

		fmt.Printf("Repo No :: %d \n and Repo Details :: %v \n", i+1, repo)
		fmt.Println("---------------------------------------------------")
	}

	fmt.Println("Total Repositories :: ", len(result))
}
