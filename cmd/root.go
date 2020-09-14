package cmd

import "github.com/spf13/cobra"

/**
* @Author: super
* @Date: 2020-09-14 20:12
* @Description:
**/

var rootCmd = &cobra.Command{
	Use:"",
	Short:"",
	Long:"",
}

func Execute() error{
	return rootCmd.Execute()
}

func init(){
	rootCmd.AddCommand(wordCmd)
}