package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"

	lp "transpiler/dsl"
	tp "transpiler/templates"
)

func PolicyGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policy-get",
		Short: "returns public policy of according to provided filename.",
		RunE: func(cmd *cobra.Command, args []string) error {

			// check for config name as input argument
			if len(args) < 1 {
				return errors.New("provide policy filename without extension")
			}
			policyName := args[0]

			// read config file
			jsonFile, err := os.Open("dsl/" + policyName + ".json")
			if err != nil {
				log.Error().Err(err).Msg("os.Open()")
				return err
			}
			defer jsonFile.Close()

			// parse json
			byteValue, _ := io.ReadAll(io.Reader(jsonFile))
			var policyJson lp.ZkPolicy
			json.Unmarshal(byteValue, &policyJson)

			// pretty print json string
			s, err := json.MarshalIndent(policyJson, "", "\t")
			if err != nil {
				log.Error().Err(err).Msg("json.MashalIndent()")
				return err
			}

			// print to console
			fmt.Print(string(s))
			fmt.Println()

			return nil
		},
	}

	return cmd
}

func ParseLibraryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lib-parse",
		Short: "parse gadget library via filename.",
		Run: func(cmd *cobra.Command, args []string) {

			// read folder
			files, err := os.ReadDir("./dsl")
			if err != nil {
				log.Error().Err(err).Msg("os.ReadDir()")
			}

			// print filename if not a directory
			for _, file := range files {
				if !file.IsDir() && strings.Contains(file.Name(), ".json") {
					fmt.Println(strings.Split(file.Name(), ".json")[0])
				}
			}
		},
	}

	return cmd
}

func PolicyTranspileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policy-transpile",
		Short: "transpiles policy and constraints into circuit.",
		RunE: func(cmd *cobra.Command, args []string) error {

			// check for config name as input argument
			if len(args) < 2 {
				fmt.Println()
				err := errors.New("please provide policy filename and generator filename without file extensions: hermes policy-transpile <policy-filename> <generator-filename>")
				return err
			}
			policyName := args[0]
			generatorFileName := args[1]

			// read config file
			jsonFile, err := os.Open("ledger_policy/" + policyName + ".json")
			if err != nil {
				log.Error().Err(err).Msg("os.Open()")
				return err
			}
			defer jsonFile.Close()

			// parse json
			byteValue, _ := io.ReadAll(io.Reader(jsonFile))
			var policyJson lp.ZkPolicy
			json.Unmarshal(byteValue, &policyJson)

			// check configs
			// fmt.Println("config values:", policyJson)

			// check if enough constraints in policy
			if len(policyJson.Constraints) < 1 {
				//log.Println("error: not enough constraints in selected policy.")
				err := errors.New("not enough constraints in selected policy")
				return err
			}

			// run transpiler
			t := tp.NewCircuit(generatorFileName, policyJson.Constraints[0].String)
			err = t.Transpile()
			if err != nil {
				log.Error().Err(err).Msg("t.Transpile()")
				return err
			}

			return nil
		},
	}

	return cmd
}

func PolicyListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policy-list",
		Short: "list currently available public policies.",
		Run: func(cmd *cobra.Command, args []string) {

			// read folder
			files, err := os.ReadDir("./dsl")
			if err != nil {
				log.Error().Err(err).Msg("os.ReadDir()")
			}

			// print filename if not a directory
			for _, file := range files {
				if !file.IsDir() && strings.Contains(file.Name(), ".json") {
					fmt.Println(strings.Split(file.Name(), ".json")[0])
				}
			}
		},
	}

	return cmd
}
