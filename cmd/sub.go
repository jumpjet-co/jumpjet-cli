package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"../api"
	"strings"
)

var outputFile string
var env string

var subCmd = &cobra.Command{
	Use:   "sub --env env [--output-file file] input-file",
	DisableFlagsInUseLine:true,
	Short: "Replace placeholders with jumpjet variables in a file",
	Example: "  jumpjet sub --env production --output .env .env.template",
	Run: func(cmd *cobra.Command, args []string) {

		apiKey := os.Getenv("JUMPJET_APIKEY")
		appSlug := os.Getenv("JUMPJET_APPSLUG")

		if apiKey == "" {
			println("JUMPJET_APIKEY environment variable is not set")
			os.Exit(1)
		}

		if appSlug == "" {
			println("JUMPJET_APPSLUG environment variable is not set")
			os.Exit(1)
		}

		file,err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		b,_ := ioutil.ReadAll(file)
		fileContent := string(b)
		stat,_ := file.Stat()

		vars := api.GetEnvVariables(apiKey, appSlug, env)

		for name, value := range vars {
			placeholder := fmt.Sprintf("{{%s}}", name)

			fileContent = strings.Replace(fileContent, placeholder, value, -1)
		}

		if outputFile != "" {
			ioutil.WriteFile(outputFile, []byte(fileContent), stat.Mode())
		} else {
			println(fileContent)
		}
	},
}

func initSubCmd() {
	subCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file, if not specified print to stdout")
	subCmd.Flags().StringVar(&env, "env", "", "Environment to use, mandatory")
	subCmd.Args = cobra.ExactArgs(1)

	subCmd.MarkFlagRequired("env")

	rootCmd.AddCommand(subCmd)
}