package cmd

import (
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
		id, err := cmd.Flags().GetInt("MensaID")
		day, err := cmd.Flags().GetInt("DayOffset")
		name, err := cmd.Flags().GetString("MensaName")
		if err != nil {
			return
		}
		flags := config.NewFlagSet(name, id, day, configFileName)

		fetcher, err := usecase.NewFetcher(flags, c)
		fetcher.PrintMeals()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.MensaFetch.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().IntP("DayOffset", "d", 0, "Which day to fetch relative to today. 0 for Today, 1 for tomorrow and so on")
	rootCmd.PersistentFlags().IntP("MensaID", "m", 0, "The ID of the Canteen. Will try to use the string for the name if not given")
	rootCmd.PersistentFlags().StringP("MensaName", "n", "", "The Name of the Canteen. Will return a list of possibilites if more than one name matches")
	rootCmd.PersistentFlags().StringP("Configfile", "c", "config.yaml", "The Name of the config file in your mensafetch installation directory")
}
