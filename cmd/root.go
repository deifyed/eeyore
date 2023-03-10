package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/deifyed/eeyore/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "eeyore",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.eeyore.yaml)")

	rootCmd.PersistentFlags().StringP(config.OpenAIToken, "t", "", "API token to use")
	err := viper.BindPFlag(config.OpenAIToken, rootCmd.PersistentFlags().Lookup(config.OpenAIToken))
	cobra.CheckErr(err)
	err = viper.BindEnv(config.OpenAIToken, "EEYORE_OPENAI_TOKEN")
	cobra.CheckErr(err)

	rootCmd.PersistentFlags().IntP(config.MaxTokens, "m", 1024, "Maximum number of tokens to generate")
	err = viper.BindPFlag(config.MaxTokens, rootCmd.PersistentFlags().Lookup(config.MaxTokens))
	cobra.CheckErr(err)
	err = viper.BindEnv(config.MaxTokens, "EEYORE_MAX_TOKENS")
	cobra.CheckErr(err)

	rootCmd.PersistentFlags().Float32P(config.Temperature, "T", 0.5, "Temperature to use")
	err = viper.BindPFlag(config.Temperature, rootCmd.PersistentFlags().Lookup(config.Temperature))
	cobra.CheckErr(err)
	err = viper.BindEnv(config.Temperature, "EEYORE_TEMPERATURE")
	cobra.CheckErr(err)

	rootCmd.PersistentFlags().StringP(config.Model, "M", "text-davinci-003", "Model to use")
	err = viper.BindPFlag(config.Model, rootCmd.PersistentFlags().Lookup(config.Model))
	cobra.CheckErr(err)
	err = viper.BindEnv(config.Model, "EEYORE_MODEL")
	cobra.CheckErr(err)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		configPath := path.Join(home, ".config", "eeyore")

		viper.AddConfigPath(configPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.SetEnvPrefix("eeyore")
	viper.EnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
