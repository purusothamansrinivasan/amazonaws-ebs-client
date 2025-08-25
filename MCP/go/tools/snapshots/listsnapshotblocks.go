package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-elastic-block-store/mcp-server/config"
	"github.com/amazon-elastic-block-store/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListsnapshotblocksHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["pageToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageToken=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["startingBlockIndex"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startingBlockIndex=%v", val))
		}
		if val, ok := args["MaxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("MaxResults=%v", val))
		}
		if val, ok := args["NextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("NextToken=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/snapshots/%s/blocks%s", cfg.BaseURL, snapshotId, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.ListSnapshotBlocksResponse
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

func CreateListsnapshotblocksTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_snapshots_snapshotId_blocks",
		mcp.WithDescription("<p>Returns information about the blocks in an Amazon Elastic Block Store snapshot.</p> <note> <p>You should always retry requests that receive server (<code>5xx</code>) error responses, and <code>ThrottlingException</code> and <code>RequestThrottledException</code> client error responses. For more information see <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html">Error retries</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> </note>"),
		mcp.WithString("X-Amz-Content-Sha256", mcp.Description("")),
		mcp.WithString("X-Amz-Date", mcp.Description("")),
		mcp.WithString("X-Amz-Algorithm", mcp.Description("")),
		mcp.WithString("X-Amz-Credential", mcp.Description("")),
		mcp.WithString("X-Amz-Security-Token", mcp.Description("")),
		mcp.WithString("X-Amz-Signature", mcp.Description("")),
		mcp.WithString("X-Amz-SignedHeaders", mcp.Description("")),
		mcp.WithString("snapshotId", mcp.Required(), mcp.Description("The ID of the snapshot from which to get block indexes and block tokens.")),
		mcp.WithString("pageToken", mcp.Description("<p>The token to request the next page of results.</p> <p>If you specify <b>NextToken</b>, then <b>StartingBlockIndex</b> is ignored.</p>")),
		mcp.WithNumber("maxResults", mcp.Description("<p>The maximum number of blocks to be returned by the request.</p> <p>Even if additional blocks can be retrieved from the snapshot, the request can return less blocks than <b>MaxResults</b> or an empty array of blocks.</p> <p>To retrieve the next set of blocks from the snapshot, make another request with the returned <b>NextToken</b> value. The value of <b>NextToken</b> is <code>null</code> when there are no more blocks to return.</p>")),
		mcp.WithNumber("startingBlockIndex", mcp.Description("<p>The block index from which the list should start. The list in the response will start from this block index or the next valid block index in the snapshot.</p> <p>If you specify <b>NextToken</b>, then <b>StartingBlockIndex</b> is ignored.</p>")),
		mcp.WithString("MaxResults", mcp.Description("Pagination limit")),
		mcp.WithString("NextToken", mcp.Description("Pagination token")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListsnapshotblocksHandler(cfg),
	}
}
