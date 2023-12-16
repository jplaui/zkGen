package templates

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"

	p "transpiler/zkpolicy"
)

type Circuit struct {
	FileName    string
	Template    *template.Template
	Constraints []string
	String      string
}

func NewCircuit(policy p.ZkPolicy) *Circuit {

	algorithm := policy.Relations[0].PrivateArgument.Protection.Algorithm
	circuitModel := getCircuitModel(algorithm)

	t := template.Must(template.New("zkcircuit").Parse(circuitModel))
	c := &Circuit{Template: t, FileName: policy.Name, Constraints: policy.Constraints}

	return c
}

func (t *Circuit) Transpile() error {

	// template data
	var data map[string]string
	if t.Constraint == "GT" {
		data = getComparatorGT(t.FileName)
	} else if t.Constraint == "LT" {
		data = getComparatorLT(t.FileName)
	} else if t.Constraint == "EQ" {
		data = getComparatorEQ(t.FileName)
	} else {
		err := errors.New("constraint not supported")
		log.Println("transpile error: constraint not found")
		return err
	}

	// transpile constraints into generator file
	builder := &strings.Builder{}
	if err := t.Template.Execute(builder, data); err != nil {
		log.Println("t.Template.Execute() error:", err)
		return err
	}
	t.String = builder.String()

	return nil
}

func (t *Circuit) StoreCircuit() error {

	// store generator file in jsnark-demo
	generatorFile, err := os.OpenFile("circuits/"+t.FileName+".java", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println("os.OpenFile() error:", err)
		return err
	}
	_, err = generatorFile.WriteString(t.FileName + ".go")
	defer generatorFile.Close()
	if err != nil {
		log.Println("generatorFile.WriteString error:", err)
		return err
	}

	return nil
}

func (t *Circuit) GenSolidity() error {

	// compile generator file
	cmd := exec.Command("bash", "-c", "javac -d bin -cp /usr/share/java/junit4.jar:bcprov-jdk15on-159.jar $(find ./src/* | grep \".java$\")")
	cmd.Dir = "dependencies/jsnark-demo/JsnarkCircuitBuilder/"
	_, err := cmd.Output()
	if err != nil {
		log.Println("cmd.Run() error:", err)
		return err
	}

	return nil
}

func dataAttributesLT(privateVariable, publicVariable string) map[string]string {

	data := map[string]string{
		"PrivateVariableDefinition": privateVariable + " frontend.Variable `gnark:\",public\"`",
		"Public":                    "`gnark:\",public\"`",
		"PublicVariableDefinition":  publicVariable + " frontend.Variable `gnark:\",public\"`",
		"Comparison":                "api.AssertIsLessOrEqual(circuit." + privateVariable + ", circuit." + publicVariable + ")",
	}

	return data
}

func dataAttributesGT(privateVariable, publicVariable string) map[string]string {

	data := map[string]string{
		"PrivateVariableDefinition": privateVariable + " frontend.Variable `gnark:\",public\"`",
		"Public":                    "`gnark:\",public\"`",
		"PublicVariableDefinition":  publicVariable + " frontend.Variable `gnark:\",public\"`",
		"Comparison":                "api.AssertIsLessOrEqual(circuit." + publicVariable + ", circuit." + privateVariable + ")",
	}

	return data
}

func getComparatorEQ() map[string]string {

	data := map[string]string{
		"GeneratorThreshold": "args[24]",
	}

	// sample_data := map[string]interface{}{
	// "Name":     "Bob",
	// "UserName": "bob92",
	// "Roles":    []string{"dbteam", "uiteam", "tester"},
	// }

	// `gnark:",public"`

	return data
}

func getCircuitModel(algorithm string) string {

	if algorithm == "commitment:mimc" {
		return `
			package gadgets
	
			import (
				"github.com/consensys/gnark/frontend"
				"github.com/consensys/gnark/std/hash/mimc"
			)
			
			// Circuit defines a pre-image knowledge proof
			// mimc(secret preImage) = public hash
			type MimcWrapper struct {
				{{.PrivateVariableDefinition}}
				Hash frontend.Variable {{.Public}}
				{{.PublicVariableDefinition}}
			}
			
			// Define declares the circuit's constraints
			func (circuit *MimcWrapper) Define(api frontend.API) error {
				// hash function
				mimc, _ := mimc.NewMiMC(api)
			
				// for i := 0; i < len(circuit.In); i++ {
				// 	mimc.Write(circuit.In[i])
				// }
				mimc.Write(circuit.In[:]...)
			
				result := mimc.Sum()
				api.AssertIsEqual(circuit.Hash, result)
	
				{{.Comparison}}
			
				return nil
			}	
		`
	} else {
		return ``
	}
}
