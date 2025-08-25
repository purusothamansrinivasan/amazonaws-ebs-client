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

func StartsnapshotHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.StartSnapshotrequest
		
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
		url := fmt.Sprintf("%s/snapshots", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
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
		var result models.StartSnapshotResponse
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

func CreateStartsnapshotTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_snapshots",
		mcp.WithDescription("<p>Creates a new Amazon EBS snapshot. The new snapshot enters the <code>pending</code> state after the request completes. </p> <p>After creating the snapshot, use <a href="https://docs.aws.amazon.com/ebs/latest/APIReference/API_PutSnapshotBlock.html"> PutSnapshotBlock</a> to write blocks of data to the snapshot.</p> <note> <p>You should always retry requests that receive server (<code>5xx</code>) error responses, and <code>ThrottlingException</code> and <code>RequestThrottledException</code> client error responses. For more information see <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html">Error retries</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> </note>"),
		mcp.WithString("X-Amz-Content-Sha256", mcp.Description("")),
		mcp.WithString("X-Amz-Date", mcp.Description("")),
		mcp.WithString("X-Amz-Algorithm", mcp.Description("")),
		mcp.WithString("X-Amz-Credential", mcp.Description("")),
		mcp.WithString("X-Amz-Security-Token", mcp.Description("")),
		mcp.WithString("X-Amz-Signature", mcp.Description("")),
		mcp.WithString("X-Amz-SignedHeaders", mcp.Description("")),
		mcp.WithNumber("VolumeSize", mcp.Required(), mcp.Description("Input parameter: The size of the volume, in GiB. The maximum size is <code>65536</code> GiB (64 TiB).")),
		mcp.WithString("ClientToken", mcp.Description("Input parameter: <p>A unique, case-sensitive identifier that you provide to ensure the idempotency of the request. Idempotency ensures that an API request completes only once. With an idempotent request, if the original request completes successfully. The subsequent retries with the same client token return the result from the original successful request and they have no additional effect.</p> <p>If you do not specify a client token, one is automatically generated by the Amazon Web Services SDK.</p> <p>For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-direct-api-idempotency.html\"> Idempotency for StartSnapshot API</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p>")),
		mcp.WithString("Description", mcp.Description("Input parameter: A description for the snapshot.")),
		mcp.WithBoolean("Encrypted", mcp.Description("Input parameter: <p>Indicates whether to encrypt the snapshot.</p> <p>You can't specify <b>Encrypted</b> and <b> ParentSnapshotId</b> in the same request. If you specify both parameters, the request fails with <code>ValidationException</code>.</p> <p>The encryption status of the snapshot depends on the values that you specify for <b>Encrypted</b>, <b>KmsKeyArn</b>, and <b>ParentSnapshotId</b>, and whether your Amazon Web Services account is enabled for <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-by-default\"> encryption by default</a>. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html\"> Using encryption</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> <important> <p>To create an encrypted snapshot, you must have permission to use the KMS key. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapi-permissions.html#ebsapi-kms-permissions\"> Permissions to use Key Management Service keys</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> </important>")),
		mcp.WithString("KmsKeyArn", mcp.Description("Input parameter: <p>The Amazon Resource Name (ARN) of the Key Management Service (KMS) key to be used to encrypt the snapshot.</p> <p>The encryption status of the snapshot depends on the values that you specify for <b>Encrypted</b>, <b>KmsKeyArn</b>, and <b>ParentSnapshotId</b>, and whether your Amazon Web Services account is enabled for <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-by-default\"> encryption by default</a>. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html\"> Using encryption</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> <important> <p>To create an encrypted snapshot, you must have permission to use the KMS key. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapi-permissions.html#ebsapi-kms-permissions\"> Permissions to use Key Management Service keys</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> </important>")),
		mcp.WithString("ParentSnapshotId", mcp.Description("Input parameter: <p>The ID of the parent snapshot. If there is no parent snapshot, or if you are creating the first snapshot for an on-premises volume, omit this parameter.</p> <p>You can't specify <b>ParentSnapshotId</b> and <b>Encrypted</b> in the same request. If you specify both parameters, the request fails with <code>ValidationException</code>.</p> <p>The encryption status of the snapshot depends on the values that you specify for <b>Encrypted</b>, <b>KmsKeyArn</b>, and <b>ParentSnapshotId</b>, and whether your Amazon Web Services account is enabled for <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-by-default\"> encryption by default</a>. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html\"> Using encryption</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> <important> <p>If you specify an encrypted parent snapshot, you must have permission to use the KMS key that was used to encrypt the parent snapshot. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapi-permissions.html#ebsapi-kms-permissions\"> Permissions to use Key Management Service keys</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> </important>")),
		mcp.WithArray("Tags", mcp.Description("Input parameter: The tags to apply to the snapshot.")),
		mcp.WithNumber("Timeout", mcp.Description("Input parameter: <p>The amount of time (in minutes) after which the snapshot is automatically cancelled if:</p> <ul> <li> <p>No blocks are written to the snapshot.</p> </li> <li> <p>The snapshot is not completed after writing the last block of data.</p> </li> </ul> <p>If no value is specified, the timeout defaults to <code>60</code> minutes.</p>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    StartsnapshotHandler(cfg),
	}
}
