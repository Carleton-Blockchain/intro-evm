package cmd

import (
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"

	"github.com/TheMarstonConnell/docute/gen"
	"github.com/spf13/cobra"
)

func GenerateCMD() *cobra.Command {
	c := cobra.Command{
		Use:   "generate",
		Short: "Generate a fully static site",
		Long:  `Generate the custom documentation site from your configuration.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			base, err := cmd.Flags().GetString("base")
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString("title")
			if err != nil {
				return err
			}
			err = gen.Gen(".out", base, title)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return &c
}

func GenColorFile() *cobra.Command {
	c := cobra.Command{
		Use:   "colors",
		Short: "Generate a color file",
		RunE: func(cmd *cobra.Command, args []string) error {
			col := gen.DefaultColors()

			data, err := yaml.Marshal(col)
			if err != nil {
				return err
			}

			err = os.WriteFile("colors.yaml", data, os.ModePerm)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return &c
}

func HostCMD() *cobra.Command {
	c := cobra.Command{
		Use:   "host",
		Short: "host the static site.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fs := http.FileServer(http.Dir(".out"))

			// Serve static files from the /static URL path
			http.Handle("/", fs)

			// Start the server on port 9797
			log.Println("Listening on :9797...")
			err := http.ListenAndServe(":9797", nil)
			if err != nil {
				log.Fatal(err)
			}

			return nil
		},
	}

	return &c
}

func WatchCMD() *cobra.Command {
	c := cobra.Command{
		Use:   "watch",
		Short: "Watch for changes and host them as a static test site.",
		RunE: func(cmd *cobra.Command, args []string) error {
			base, err := cmd.Flags().GetString("base")
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString("title")
			if err != nil {
				return err
			}

			err = gen.Gen(".out", base, title)
			if err != nil {
				return err
			}

			fs := http.FileServer(http.Dir("./out"))

			// Serve static files from the /static URL path
			http.Handle("/", fs)

			// Start the server on port 9797
			log.Println("Listening on :9797...")
			err = http.ListenAndServe(":9797", nil)
			if err != nil {
				log.Fatal(err)
			}

			return nil
		},
	}

	return &c
}
