package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type listTemplatesResponse struct {
	Templates []templateSummary `json:"templates"`
}

type templateSummary struct {
	TemplateID   string `json:"templateId"`
	Name         string `json:"name"`
	Scenario     string `json:"scenario"`
	InputSummary string `json:"inputSummary"`
	OutputType   string `json:"outputType"`
	Version      string `json:"version"`
}

type templateSchemaResponse struct {
	TemplateID   string `json:"templateId"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Scenario     string `json:"scenario"`
	InputSummary string `json:"inputSummary"`
	OutputType   string `json:"outputType"`
}

func newTemplateCmd(opts *rootOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "template",
		Short: "Discover hosted BatchJob templates",
	}

	cmd.AddCommand(
		newTemplateListCmd(opts),
		newTemplateSchemaCmd(opts),
	)
	return cmd
}

func newTemplateListCmd(opts *rootOptions) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List published templates",
		RunE: func(cmd *cobra.Command, args []string) error {
			httpClient, err := newHTTPClient(opts)
			if err != nil {
				return err
			}
			ctx, cancel := context.WithTimeout(cmd.Context(), opts.timeout)
			defer cancel()

			var resp listTemplatesResponse
			if err := httpClient.GetJSON(ctx, "/v1/templates", &resp); err != nil {
				return err
			}

			if opts.output == "json" {
				enc := json.NewEncoder(cmd.OutOrStdout())
				enc.SetIndent("", "  ")
				return enc.Encode(resp)
			}

			if len(resp.Templates) == 0 {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "no templates")
				return nil
			}
			for _, tmpl := range resp.Templates {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\n", tmpl.TemplateID, tmpl.Name, tmpl.OutputType)
			}
			return nil
		},
	}
}

func newTemplateSchemaCmd(opts *rootOptions) *cobra.Command {
	return &cobra.Command{
		Use:   "schema <template-id>",
		Short: "Inspect one template schema",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			httpClient, err := newHTTPClient(opts)
			if err != nil {
				return err
			}
			ctx, cancel := context.WithTimeout(cmd.Context(), opts.timeout)
			defer cancel()

			var resp templateSchemaResponse
			path := "/v1/templates/" + strings.TrimSpace(args[0]) + "/schema"
			if err := httpClient.GetJSON(ctx, path, &resp); err != nil {
				return err
			}

			if opts.output == "json" {
				enc := json.NewEncoder(cmd.OutOrStdout())
				enc.SetIndent("", "  ")
				return enc.Encode(resp)
			}

			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "template: %s\n", resp.TemplateID)
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "name: %s\n", resp.Name)
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "output: %s\n", resp.OutputType)
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "scenario: %s\n", resp.Scenario)
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "input: %s\n", resp.InputSummary)
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "description: %s\n", resp.Description)
			return nil
		},
	}
}
