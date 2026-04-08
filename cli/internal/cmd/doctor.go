package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

type healthResponse struct {
	Healthy bool   `json:"healthy"`
	Message string `json:"message"`
}

func newDoctorCmd(opts *rootOptions) *cobra.Command {
	return &cobra.Command{
		Use:   "doctor",
		Short: "Check BatchJob server reachability and token wiring",
		RunE: func(cmd *cobra.Command, args []string) error {
			httpClient, err := newHTTPClient(opts)
			if err != nil {
				return err
			}
			ctx, cancel := context.WithTimeout(cmd.Context(), opts.timeout)
			defer cancel()

			var resp healthResponse
			if err := httpClient.GetJSON(ctx, "/v1/health", &resp); err != nil {
				return err
			}

			if opts.output == "json" {
				payload := map[string]any{
					"server":     opts.server,
					"token_set":  opts.token != "",
					"healthy":    resp.Healthy,
					"message":    resp.Message,
					"base_usage": "set BATCHJOB_SERVER and BATCHJOB_TOKEN before running template commands",
				}
				enc := json.NewEncoder(cmd.OutOrStdout())
				enc.SetIndent("", "  ")
				return enc.Encode(payload)
			}

			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "server: %s\n", opts.server)
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "token: %t\n", opts.token != "")
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "healthy: %t\n", resp.Healthy)
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "message: %s\n", resp.Message)
			return nil
		},
	}
}
