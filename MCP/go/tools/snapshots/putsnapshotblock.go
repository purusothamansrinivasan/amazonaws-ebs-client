package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-elastic-block-store/mcp-server/config"
	"github.com/amazon-elastic-block-store/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func PutsnapshotblockHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		blockIndexVal, ok := args["blockIndex"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: blockIndex"), nil
		}
		blockIndex, ok := blockIndexVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: blockIndex"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.PutSnapshotBlockrequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/snapshots/%s/blocks/%s#x-amz-Data-Length&x-amz-Checksum&x-amz-Checksum-Algorithm", cfg.BaseURL, snapshotId, blockIndex)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		if val, ok := args["x-amz-Data-Length"]; ok {
			req.Header.Set("x-amz-Data-Length", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-amz-Progress"]; ok {
			req.Header.Set("x-amz-Progress", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-amz-Checksum"]; ok {
			req.Header.Set("x-amz-Checksum", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-amz-Checksum-Algorithm"]; ok {
			req.Header.Set("x-amz-Checksum-Algorithm", fmt.Sprintf("%v", val))
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
		var result models.PutSnapshotBlockResponse
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

func CreatePutsnapshotblockTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_snapshots_snapshotId_blocks_blockIndex#x-amz-Data-Length&x-amz-Checksum&x-amz-Checksum-Algorithm",
		mcp.WithDescription("<p>Writes a block of data to a snapshot. If the specified block contains data, the existing data is overwritten. The target snapshot must be in the <code>pending</code> state.</p> <p>Data written to a snapshot must be aligned with 512-KiB sectors.</p> <note> <p>You should always retry requests that receive server (<code>5xx</code>) error responses, and <code>ThrottlingException</code> and <code>RequestThrottledException</code> client error responses. For more information see <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html">Error retries</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> </note>"),
		mcp.WithString("X-Amz-Content-Sha256", mcp.Description("")),
		mcp.WithString("X-Amz-Date", mcp.Description("")),
		mcp.WithString("X-Amz-Algorithm", mcp.Description("")),
		mcp.WithString("X-Amz-Credential", mcp.Description("")),
		mcp.WithString("X-Amz-Security-Token", mcp.Description("")),
		mcp.WithString("X-Amz-Signature", mcp.Description("")),
		mcp.WithString("X-Amz-SignedHeaders", mcp.Description("")),
		mcp.WithString("snapshotId", mcp.Required(), mcp.Description("<p>The ID of the snapshot.</p> <important> <p>If the specified snapshot is encrypted, you must have permission to use the KMS key that was used to encrypt the snapshot. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html\"> Using encryption</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>..</p> </important>")),
		mcp.WithNumber("blockIndex", mcp.Required(), mcp.Description("The block index of the block in which to write the data. A block index is a logical index in units of <code>512</code> KiB blocks. To identify the block index, divide the logical offset of the data in the logical volume by the block size (logical offset of data/<code>524288</code>). The logical offset of the data must be <code>512</code> KiB aligned.")),
		mcp.WithNumber("x-amz-Data-Length", mcp.Required(), mcp.Description("<p>The size of the data to write to the block, in bytes. Currently, the only supported size is <code>524288</code> bytes.</p> <p>Valid values: <code>524288</code> </p>")),
		mcp.WithNumber("x-amz-Progress", mcp.Description("The progress of the write process, as a percentage.")),
		mcp.WithString("x-amz-Checksum", mcp.Required(), mcp.Description("A Base64-encoded SHA256 checksum of the data. Only SHA256 checksums are supported.")),
		mcp.WithString("x-amz-Checksum-Algorithm", mcp.Required(), mcp.Description("The algorithm used to generate the checksum. Currently, the only supported algorithm is <code>SHA256</code>.")),
		mcp.WithString("BlockData", mcp.Required(), mcp.Description("Input parameter: <p>The data to write to the block.</p> <p>The block data is not signed as part of the Signature Version 4 signing process. As a result, you must generate and provide a Base64-encoded SHA256 checksum for the block data using the <b>x-amz-Checksum</b> header. Also, you must specify the checksum algorithm using the <b>x-amz-Checksum-Algorithm</b> header. The checksum that you provide is part of the Signature Version 4 signing process. It is validated against a checksum generated by Amazon EBS to ensure the validity and authenticity of the data. If the checksums do not correspond, the request fails. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-accessing-snapshot.html#ebsapis-using-checksums\"> Using checksums with the EBS direct APIs</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PutsnapshotblockHandler(cfg),
	}
}
