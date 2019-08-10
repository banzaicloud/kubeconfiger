package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/banzaicloud/kubeconfiger"
)

const (
	sysexitsDataErr    = 65
	sysexitsNoInput    = 66
	sysexitsCantCreate = 73
)

func main() {

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read stdin: %v\n", err)
		os.Exit(sysexitsNoInput)
	}

	out, err := kubeconfiger.CleanKubeconfig(in)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(sysexitsDataErr)
	}

	_, err = os.Stdout.Write(out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to write output: %v\n", err)
		os.Exit(sysexitsCantCreate)
	}
}
