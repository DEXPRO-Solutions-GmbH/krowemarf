package cli

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// InitConfig uses the global viper instance and binds all flags to viper keys.
//
// If you use cobra, call this function in the PersistentPreRun hook on the root command.
//
// When used, simply use viper.GetString("you-config") to fetch the value of a config flag.
func InitConfig(flags *pflag.FlagSet) {
	if err := godotenv.Load(".env"); err != nil {
		// ignore if the .env file does not exist
		if !errors.Is(err, os.ErrNotExist) {
			panic(fmt.Errorf("loading .env file failed: %w", err))
		}

		slog.Info("no .env file found")
	}

	v := viper.GetViper()

	viper.SetEnvPrefix("FRZ")
	viper.AutomaticEnv()

	// Replace viper keys to ensure that they conform to what users are used to
	viper.SetEnvKeyReplacer(strings.NewReplacer(
		"-", "_", // some-value => some_value
	))

	flags.VisitAll(func(f *pflag.Flag) {
		key := f.Name

		if err := v.BindPFlag(key, f); err != nil {
			panic(err)
		}
	})
}
