package main

import (
	"fmt"
	"github.com/amassanet/goct"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func DisplayDot(dot []byte) error {
	var file *os.File
	file, err := ioutil.TempFile("", "rx")
	if err != nil {
		return err
	}
	imagefilename := file.Name() + ".png"
	ioutil.WriteFile(file.Name(), dot, 0744)
	log.Println(imagefilename)
	cmd := exec.Command("/usr/local/bin/dot", "-x", "-Tpng", "-o"+imagefilename, file.Name())
	if err := cmd.Run(); err != nil {
		return err
	}
	exec.Command("open", imagefilename).Run()
	return nil
}
func main() {
	rx := goct.NewRadixGraph(9182, 7)
	if err := DisplayDot(rx.CreateDot()); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("value is %v\n", rx.CalcCT())
	rx.CreateCode()
}
