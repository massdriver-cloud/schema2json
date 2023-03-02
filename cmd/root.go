/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/massdriver-cloud/schema2json"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "schema2json",
	Short:        "A CLI tool to generate a JSON document based on a JSON Schema document",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	RunE:         run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}

func run(cmd *cobra.Command, args []string) error {
	schema := new(schema2json.Schema)
	bytes, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, schema)
	if err != nil {
		return err
	}

	blob, err := schema2json.GenerateJSON(schema)
	if err != nil {
		return err
	}

	out, err := json.MarshalIndent(blob, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(out))

	return nil
}
