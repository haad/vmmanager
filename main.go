package main

import (
	"github.com/haad/vmmanager/cmd"
)

func main() {
	// Set up Cobra and Viper
	// cobra.OnInitialize(initConfig)

	// viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	// viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))

	cmd.Execute()
}

// func initConfig() {
// 	// Set up Viper to read configuration values from environment variables and a configuration file
// 	viper.SetEnvPrefix("vmmanager")
// 	viper.AutomaticEnv()
// 	viper.SetConfigName("config")
// 	viper.AddConfigPath("/etc/vmmanager/")
// 	viper.AddConfigPath("$HOME/.vmmanager")
// 	viper.AddConfigPath(".")
// 	if err := viper.ReadInConfig(); err != nil {
// 		fmt.Println("Error reading configuration:", err)
// 		os.Exit(1)
// 	}
// }
