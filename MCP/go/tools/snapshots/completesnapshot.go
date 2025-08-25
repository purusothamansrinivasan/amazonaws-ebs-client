package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amazon-elastic-block-store/mcp-server/config"
	"github.com/amazon-elastic-block-store/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CompletesnapshotHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		snapshotIdVal, ok := args["snapshotId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: snapshotId"), nil
		}
		snapshotId, ok := snapshotIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: snapshotId"), nil
		}
		url := fmt.Sprintf("%s/snapshots/completion/%s#x-amz-ChangedBlocksCount", cfg.BaseURL, snapshotId)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Amz-Content-Sha256"]; ok {
			req.Header.Set("X-Amz-Content-Sha256", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Date"]; ok {
			req.Header.Set("X-Amz-Date", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Algorithm"]; ok {
			req.Header.Set("X-Amz-Algorithm", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Credential"]; ok {
			req.Header.Set("X-Amz-Credential", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Security-Token"]; ok {
			req.Header.Set("X-Amz-Security-Token", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Signature"]; ok {
			req.Header.Set("X-Amz-Signature", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-SignedHeaders"]; ok {
			req.Header.Set("X-Amz-SignedHeaders", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-amz-ChangedBlocksCount"]; ok {
			req.Header.Set("x-amz-ChangedBlocksCount", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-amz-Checksum"]; ok {
			req.Header.Set("x-amz-Checksum", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-amz-Checksum-Algorithm"]; ok {
			req.Header.Set("x-amz-Checksum-Algorithm", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-amz-Checksum-Aggregation-Method"]; ok {
			req.Header.Set("x-amz-Checksum-Aggregation-Method", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCompletesnapshotTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_snapshots_completion_snapshotId#x-amz-ChangedBlocksCount",
		mcp.WithDescription("<p>Seals and completes the snapshot after all of the required blocks of data have been written to it. Completing the snapshot changes the status to <code>completed</code>. You cannot write new blocks to a snapshot after it has been completed.</p> <note> <p>You should always retry requests that receive server (<code>5xx</code>) error responses, and <code>ThrottlingException</code> and <code>RequestThrottledException</code> client error responses. For more information see <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html">Error retries</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> </note>"),
		mcp.WithString("X-Amz-Content-Sha256", mcp.Description("")),
		mcp.WithString("X-Amz-Date", mcp.Description("")),
		mcp.WithString("X-Amz-Algorithm", mcp.Description("")),
		mcp.WithString("X-Amz-Credential", mcp.Description("")),
		mcp.WithString("X-Amz-Security-Token", mcp.Description("")),
		mcp.WithString("X-Amz-Signature", mcp.Description("")),
		mcp.WithString("X-Amz-SignedHeaders", mcp.Description("")),
		mcp.WithString("snapshotId", mcp.Required(), mcp.Description("The ID of the snapshot.")),
		mcp.WithNumber("x-amz-ChangedBlocksCount", mcp.Required(), mcp.Description("The number of blocks that were written to the snapshot.")),
		mcp.WithString("x-amz-Checksum", mcp.Description("<p>An aggregated Base-64 SHA256 checksum based on the checksums of each written block.</p> <p>To generate the aggregated checksum using the linear aggregation method, arrange the checksums for each written block in ascending order of their block index, concatenate them to form a single string, and then generate the checksum on the entire string using the SHA256 algorithm.</p>")),
		mcp.WithString("x-amz-Checksum-Algorithm", mcp.Description("The algorithm used to generate the checksum. Currently, the only supported algorithm is <code>SHA256</code>.")),
		mcp.WithString("x-amz-Checksum-Aggregation-Method", mcp.Description("The aggregation method used to generate the checksum. Currently, the only supported aggregation method is <code>LINEAR</code>.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CompletesnapshotHandler(cfg),
	}
}
