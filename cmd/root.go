package cmd

import (
	"fmt"
	"mensafetch/config"
	"mensafetch/usecase"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "MensaFetch",
	Short: "Look at what your canteen has to offer in the terminal",
	Long: `MensaFetch was created to avoid having to scroll through abysmal GUIs
Developers can now check what their favourite canteen has to offer without leaving the comfort of the Terminal
Usage without any subcommand will print out todays meals for a specific canteen with the configured information`,
	Run: func(cmd *cobra.Command, args []string) {
		configFileName, err := cmd.Flags().GetString("Configfile")
		c, err := config.ReadConfig(configFileName)
		id, err := cmd.Flags().GetInt("mensaID")
		day, err := cmd.Flags().GetInt("dayOffset")
		name, err := cmd.Flags().GetString("mensaName")
		if err != nil {
			fmt.Printf("Error while parsing flags: %s", err)
		}
		flags := config.NewFlagSet(name, id, day, configFileName)
		fmt.Println(c)
		fetcher, err := usecase.NewFetcher(flags, c)
		//fetcher.PrintMeals()
		fetcher.GetUIView()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.MensaFetch.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().IntP("dayOffset", "d", 0, "Which day to fetch relative to today. 0 for Today, 1 for tomorrow and so on")
	rootCmd.PersistentFlags().IntP("mensaID", "m", 6, "The ID of the Canteen. Will try to use the string for the name if not given")
	rootCmd.PersistentFlags().StringP("mensaName", "n", "", "The Name of the Canteen. Will return a list of possibilities if more than one name matches")
	rootCmd.PersistentFlags().StringP("configFile", "c", "config.yaml", "The Name of the config file in your mensafetch installation directory")
}
