/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
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
		fmt.Printf(`----------------------------------------------------------------------------
|   \  |                                  \  |         |         _)        |
|  |\/ |   _ \  __ \    __|   _\ |       |\/ |   _\ |  __|   __|  |  \  /  |
|  |   |   __/  |   | \__ \  (   |       |   |  (   |  |    |     |    <   |
| _|  _| \___| _|  _| ____/ \__._|      _|  _| \__._| \__| _|    _| _/\_\  |
----------------------------------------------------------------------------

06.02.2024

-----------
category: Suppen
name: Porree-Weißkohleintopf mit Rindfleisch und Schweinefleisch
price: 2.18€
-----------
category: Angebot 1
name: Schweinekammbraten mit Möhrengemüse und Semmelknödel
price: 3.33€
-----------
category: Angebot 4
name: Linsen-Moussaka mit Tsatsiki, dazu Salat
price: 3.13€
-----------
category: Angebot 2
name: Penne mit Champignon-Tomaten-Rahmsoße und veganem Schmelz oder geriebener Goudakäse
price: 2.35€
-----------
category: Angebot 3
name: Penne mit Puten-Currysoße, dazu geriebener Käse
price: 2.35€
			`)
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

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
