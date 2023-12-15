package templates

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type Circuit struct {
	FileName   string
	Template   *template.Template
	Constraint string
}

func NewCircuit(filename, constraint string) *Circuit {
	generatorTempl := getTemplate1()
	t := template.Must(template.New("generator").Parse(generatorTempl))
	c := &Circuit{Template: t, FileName: filename, Constraint: constraint}
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
	generatorStr := builder.String()

	// store generator file in jsnark-demo
	generatorFile, err := os.OpenFile("dependencies/jsnark-demo/JsnarkCircuitBuilder/src/examples/generators/transpiled/"+t.FileName+".java", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println("os.OpenFile() error:", err)
		return err
	}
	_, err = generatorFile.WriteString(generatorStr)
	if err != nil {
		log.Println("generatorFile.WriteString error:", err)
		generatorFile.Close()
		return err
	}
	generatorFile.Close()

	// debug
	log.Println("generator file has been written successfully.")

	// compile generator file
	cmd := exec.Command("bash", "-c", "javac -d bin -cp /usr/share/java/junit4.jar:bcprov-jdk15on-159.jar $(find ./src/* | grep \".java$\")")
	cmd.Dir = "dependencies/jsnark-demo/JsnarkCircuitBuilder/"
	_, err = cmd.Output()
	if err != nil {
		log.Println("cmd.Run() error:", err)
		return err
	}

	// debug command output of compile
	// log.Println("compile java generator, output:", string(out))
	// log.Println("java generator code has been compiled successfully.")

	// // overwrite bin folder zk-build/jsnark
	// cmd2 := exec.Command("cp", "-r", "dependencies/jsnark-demo/JsnarkCircuitBuilder/bin", "prover/zksnark_build/jsnark")
	// cmd.Dir = "../"
	// if err := cmd2.Run(); err != nil {
	// 	log.Println("cmd2.Run() error:", err)
	// 	return err
	// }
	// log.Println("compiled generator code has been copied successfully to prover folder..")

	return nil
}

func getComparatorGT(filename string) map[string]string {

	data := map[string]string{
		"GenName":                filename,
		"Import":                 "GTFloatThresholdComparatorGadget",
		"PrivateThresholdType":   "BigInteger",
		"ThresholdWireArraySize": "1",
		"ComparatorCall":         "GTFloatThresholdComparatorGadget",
		"SetThreshold":           "circuitEvaluator.setWireValue(thresholdValue[0], threshold)",
		"TestThreshold":          "new BigInteger(\"1406000\")",
		"GeneratorThreshold":     "new BigInteger(args[24])",
	}

	return data
}

func getComparatorLT(filename string) map[string]string {

	data := map[string]string{
		"GenName":                filename,
		"Import":                 "LTFloatThresholdComparatorGadget",
		"PrivateThresholdType":   "BigInteger",
		"ThresholdWireArraySize": "1",
		"ComparatorCall":         "LTFloatThresholdComparatorGadget",
		"SetThreshold":           "circuitEvaluator.setWireValue(thresholdValue[0], threshold)",
		"TestThreshold":          "new BigInteger(\"1406000\")",
		"GeneratorThreshold":     "new BigInteger(args[24])",
	}

	return data
}

func getComparatorEQ(filename string) map[string]string {

	data := map[string]string{
		"GenName":                filename,
		"Import":                 "EQFloatComparatorGadget",
		"PrivateThresholdType":   "String",
		"ThresholdWireArraySize": "floatStringLen",
		"ComparatorCall":         "EQFloatComparatorGadget",
		"SetThreshold":           "setWires2(thresholdValue, threshold, circuitEvaluator)",
		"TestThreshold":          "\"140.6000\"",
		"GeneratorThreshold":     "args[24]",
	}

	// sample_data := map[string]interface{}{
	// "Name":     "Bob",
	// "UserName": "bob92",
	// "Roles":    []string{"dbteam", "uiteam", "tester"},
	// }

	return data
}

func getTemplate1() string {
	return `

`
}
