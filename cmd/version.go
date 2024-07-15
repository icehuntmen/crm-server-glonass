package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the current version number from VERSION file.`,
	Run: func(cmd *cobra.Command, args []string) {
		version, err := os.ReadFile("./VERSION")
		if err != nil {
			fmt.Println("Error reading VERSION file:", err)
			os.Exit(1)
		}
		fmt.Println(strings.TrimSpace(string(version)))
	},
}

var bumpCmd = &cobra.Command{
	Use:   "patch",
	Short: "Update the patch version",
	Long:  `Update the patch version in the VERSION file.`,
	Run: func(cmd *cobra.Command, args []string) {
		version, err := os.ReadFile("./VERSION")
		if err != nil {
			fmt.Println("Error reading VERSION file:", err)
			os.Exit(1)
		}
		currentVersion := strings.TrimSpace(string(version))
		parts := strings.Split(currentVersion, ".")
		if len(parts) != 3 {
			fmt.Println("Invalid version format in VERSION file")
			os.Exit(1)
		}
		patch := parts[2]
		newPatchInt, err := strconv.Atoi(patch)
		if err != nil {
			fmt.Println("Error parsing patch version:", err)
			os.Exit(1)
		}
		newPatch := fmt.Sprintf("%d", newPatchInt+1)
		newVersion := fmt.Sprintf("%s.%s.%s", parts[0], parts[1], newPatch)
		err = os.WriteFile("./VERSION", []byte(newVersion), 0644)
		if err != nil {
			fmt.Println("Error writing VERSION file:", err)
			os.Exit(1)
		}
		fmt.Printf("Updated version to: %s\n", newVersion)
	},
}
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the version file",
	Long:  `Check the version file without modifying it.`,
	Run: func(cmd *cobra.Command, args []string) {
		version, err := os.ReadFile("./VERSION")
		if err != nil {
			fmt.Println("Error reading VERSION file:", err)
			os.Exit(1)
		}
		fmt.Println("Current version:", strings.TrimSpace(string(version)))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(VersionCmd)
	rootCmd.AddCommand(bumpCmd)
	rootCmd.AddCommand(checkCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
