package commands

import (
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"

	t "transpiler/templates"
	p "transpiler/zkpolicy"
)

func PolicyTranspileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "zkpolicy-transpile",
		Short: "transpiles policy and constraints into circuit.",
		RunE: func(cmd *cobra.Command, args []string) error {

			// check for input arguments
			if len(args) < 1 {
				return errors.New("provide policy filename without extension")
			}
			policyName := args[0]

			// fetch policy
			zkPolicy, err := p.ParseZkPolicy(policyName)
			if err != nil {
				log.Error().Err(err).Msg("p.ParseZkPolicy()")
				return err
			}

			// check configs
			// fmt.Println("config values:", zkPolicy)

			// check if enough constraints in policy
			if len(zkPolicy.Constraints) < 1 {
				//log.Println("error: not enough constraints in selected policy.")
				err := errors.New("not enough constraints in selected policy")
				return err
			}

			// run transpiler
			template := t.NewCircuit(zkPolicy)
			err = template.Transpile()
			if err != nil {
				log.Error().Err(err).Msg("template.Transpile()")
				return err
			}

			// store output
			err = template.StoreCircuit()
			if err != nil {
				log.Error().Err(err).Msg("template.StoreCircuit()")
				return err
			}

			// generate solidity
			err = template.GenSolidity()
			if err != nil {
				log.Error().Err(err).Msg("template.GenSolidity()")
				return err
			}

			return nil
		},
	}

	return cmd
}
