/**
 * MCP Server function for No description available
 */

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.*;
import java.util.function.Function;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;

class Post_SnapshotsMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_SnapshotsHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("X-Amz-Content-Sha256")) {
            queryParams.add("X-Amz-Content-Sha256=" + args.get("X-Amz-Content-Sha256"));
        }
        if (args.containsKey("X-Amz-Date")) {
            queryParams.add("X-Amz-Date=" + args.get("X-Amz-Date"));
        }
        if (args.containsKey("X-Amz-Algorithm")) {
            queryParams.add("X-Amz-Algorithm=" + args.get("X-Amz-Algorithm"));
        }
        if (args.containsKey("X-Amz-Credential")) {
            queryParams.add("X-Amz-Credential=" + args.get("X-Amz-Credential"));
        }
        if (args.containsKey("X-Amz-Security-Token")) {
            queryParams.add("X-Amz-Security-Token=" + args.get("X-Amz-Security-Token"));
        }
        if (args.containsKey("X-Amz-Signature")) {
            queryParams.add("X-Amz-Signature=" + args.get("X-Amz-Signature"));
        }
        if (args.containsKey("X-Amz-SignedHeaders")) {
            queryParams.add("X-Amz-SignedHeaders=" + args.get("X-Amz-SignedHeaders"));
        }
        if (args.containsKey("ClientToken")) {
            queryParams.add("ClientToken=" + args.get("ClientToken"));
        }
        if (args.containsKey("Description")) {
            queryParams.add("Description=" + args.get("Description"));
        }
        if (args.containsKey("KmsKeyArn")) {
            queryParams.add("KmsKeyArn=" + args.get("KmsKeyArn"));
        }
        if (args.containsKey("ParentSnapshotId")) {
            queryParams.add("ParentSnapshotId=" + args.get("ParentSnapshotId"));
        }
        if (args.containsKey("VolumeSize")) {
            queryParams.add("VolumeSize=" + args.get("VolumeSize"));
        }
        if (args.containsKey("Timeout")) {
            queryParams.add("Timeout=" + args.get("Timeout"));
        }
        if (args.containsKey("Encrypted")) {
            queryParams.add("Encrypted=" + args.get("Encrypted"));
        }
        if (args.containsKey("Tags")) {
            queryParams.add("Tags=" + args.get("Tags"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_snapshots" + queryString;
                
                HttpClient client = HttpClient.newHttpClient();
                HttpRequest httpRequest = HttpRequest.newBuilder()
                    .uri(URI.create(url))
                    .header("Authorization", "Bearer " + config.getApiKey())
                    .header("Accept", "application/json")
                    .GET()
                    .build();
                
                HttpResponse<String> response = client.send(httpRequest, HttpResponse.BodyHandlers.ofString());
                
                if (response.statusCode() >= 400) {
                    return new MCPServer.MCPToolResult("API error: " + response.body(), true);
                }
                
                // Pretty print JSON
                ObjectMapper mapper = new ObjectMapper();
                JsonNode jsonNode = mapper.readTree(response.body());
                String prettyJson = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(jsonNode);
                
                return new MCPServer.MCPToolResult(prettyJson);
                
            } catch (IOException | InterruptedException e) {
                return new MCPServer.MCPToolResult("Request failed: " + e.getMessage(), true);
            } catch (Exception e) {
                return new MCPServer.MCPToolResult("Unexpected error: " + e.getMessage(), true);
            }
        };
    }
    
    public static MCPServer.Tool createPost_SnapshotsTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> X-Amz-Content-Sha256Property = new HashMap<>();
        X-Amz-Content-Sha256Property.put("type", "string");
        X-Amz-Content-Sha256Property.put("required", false);
        X-Amz-Content-Sha256Property.put("description", "");
        properties.put("X-Amz-Content-Sha256", X-Amz-Content-Sha256Property);
        Map<String, Object> X-Amz-DateProperty = new HashMap<>();
        X-Amz-DateProperty.put("type", "string");
        X-Amz-DateProperty.put("required", false);
        X-Amz-DateProperty.put("description", "");
        properties.put("X-Amz-Date", X-Amz-DateProperty);
        Map<String, Object> X-Amz-AlgorithmProperty = new HashMap<>();
        X-Amz-AlgorithmProperty.put("type", "string");
        X-Amz-AlgorithmProperty.put("required", false);
        X-Amz-AlgorithmProperty.put("description", "");
        properties.put("X-Amz-Algorithm", X-Amz-AlgorithmProperty);
        Map<String, Object> X-Amz-CredentialProperty = new HashMap<>();
        X-Amz-CredentialProperty.put("type", "string");
        X-Amz-CredentialProperty.put("required", false);
        X-Amz-CredentialProperty.put("description", "");
        properties.put("X-Amz-Credential", X-Amz-CredentialProperty);
        Map<String, Object> X-Amz-Security-TokenProperty = new HashMap<>();
        X-Amz-Security-TokenProperty.put("type", "string");
        X-Amz-Security-TokenProperty.put("required", false);
        X-Amz-Security-TokenProperty.put("description", "");
        properties.put("X-Amz-Security-Token", X-Amz-Security-TokenProperty);
        Map<String, Object> X-Amz-SignatureProperty = new HashMap<>();
        X-Amz-SignatureProperty.put("type", "string");
        X-Amz-SignatureProperty.put("required", false);
        X-Amz-SignatureProperty.put("description", "");
        properties.put("X-Amz-Signature", X-Amz-SignatureProperty);
        Map<String, Object> X-Amz-SignedHeadersProperty = new HashMap<>();
        X-Amz-SignedHeadersProperty.put("type", "string");
        X-Amz-SignedHeadersProperty.put("required", false);
        X-Amz-SignedHeadersProperty.put("description", "");
        properties.put("X-Amz-SignedHeaders", X-Amz-SignedHeadersProperty);
        Map<String, Object> ClientTokenProperty = new HashMap<>();
        ClientTokenProperty.put("type", "string");
        ClientTokenProperty.put("required", false);
        ClientTokenProperty.put("description", "Input parameter: <p>A unique, case-sensitive identifier that you provide to ensure the idempotency of the request. Idempotency ensures that an API request completes only once. With an idempotent request, if the original request completes successfully. The subsequent retries with the same client token return the result from the original successful request and they have no additional effect.</p> <p>If you do not specify a client token, one is automatically generated by the Amazon Web Services SDK.</p> <p>For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-direct-api-idempotency.html\"> Idempotency for StartSnapshot API</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p>");
        properties.put("ClientToken", ClientTokenProperty);
        Map<String, Object> DescriptionProperty = new HashMap<>();
        DescriptionProperty.put("type", "string");
        DescriptionProperty.put("required", false);
        DescriptionProperty.put("description", "Input parameter: A description for the snapshot.");
        properties.put("Description", DescriptionProperty);
        Map<String, Object> KmsKeyArnProperty = new HashMap<>();
        KmsKeyArnProperty.put("type", "string");
        KmsKeyArnProperty.put("required", false);
        KmsKeyArnProperty.put("description", "Input parameter: <p>The Amazon Resource Name (ARN) of the Key Management Service (KMS) key to be used to encrypt the snapshot.</p> <p>The encryption status of the snapshot depends on the values that you specify for <b>Encrypted</b>, <b>KmsKeyArn</b>, and <b>ParentSnapshotId</b>, and whether your Amazon Web Services account is enabled for <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-by-default\"> encryption by default</a>. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html\"> Using encryption</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> <important> <p>To create an encrypted snapshot, you must have permission to use the KMS key. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapi-permissions.html#ebsapi-kms-permissions\"> Permissions to use Key Management Service keys</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> </important>");
        properties.put("KmsKeyArn", KmsKeyArnProperty);
        Map<String, Object> ParentSnapshotIdProperty = new HashMap<>();
        ParentSnapshotIdProperty.put("type", "string");
        ParentSnapshotIdProperty.put("required", false);
        ParentSnapshotIdProperty.put("description", "Input parameter: <p>The ID of the parent snapshot. If there is no parent snapshot, or if you are creating the first snapshot for an on-premises volume, omit this parameter.</p> <p>You can't specify <b>ParentSnapshotId</b> and <b>Encrypted</b> in the same request. If you specify both parameters, the request fails with <code>ValidationException</code>.</p> <p>The encryption status of the snapshot depends on the values that you specify for <b>Encrypted</b>, <b>KmsKeyArn</b>, and <b>ParentSnapshotId</b>, and whether your Amazon Web Services account is enabled for <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-by-default\"> encryption by default</a>. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html\"> Using encryption</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> <important> <p>If you specify an encrypted parent snapshot, you must have permission to use the KMS key that was used to encrypt the parent snapshot. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapi-permissions.html#ebsapi-kms-permissions\"> Permissions to use Key Management Service keys</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> </important>");
        properties.put("ParentSnapshotId", ParentSnapshotIdProperty);
        Map<String, Object> VolumeSizeProperty = new HashMap<>();
        VolumeSizeProperty.put("type", "string");
        VolumeSizeProperty.put("required", true);
        VolumeSizeProperty.put("description", "Input parameter: The size of the volume, in GiB. The maximum size is <code>65536</code> GiB (64 TiB).");
        properties.put("VolumeSize", VolumeSizeProperty);
        Map<String, Object> TimeoutProperty = new HashMap<>();
        TimeoutProperty.put("type", "string");
        TimeoutProperty.put("required", false);
        TimeoutProperty.put("description", "Input parameter: <p>The amount of time (in minutes) after which the snapshot is automatically cancelled if:</p> <ul> <li> <p>No blocks are written to the snapshot.</p> </li> <li> <p>The snapshot is not completed after writing the last block of data.</p> </li> </ul> <p>If no value is specified, the timeout defaults to <code>60</code> minutes.</p>");
        properties.put("Timeout", TimeoutProperty);
        Map<String, Object> EncryptedProperty = new HashMap<>();
        EncryptedProperty.put("type", "string");
        EncryptedProperty.put("required", false);
        EncryptedProperty.put("description", "Input parameter: <p>Indicates whether to encrypt the snapshot.</p> <p>You can't specify <b>Encrypted</b> and <b> ParentSnapshotId</b> in the same request. If you specify both parameters, the request fails with <code>ValidationException</code>.</p> <p>The encryption status of the snapshot depends on the values that you specify for <b>Encrypted</b>, <b>KmsKeyArn</b>, and <b>ParentSnapshotId</b>, and whether your Amazon Web Services account is enabled for <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-by-default\"> encryption by default</a>. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html\"> Using encryption</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> <important> <p>To create an encrypted snapshot, you must have permission to use the KMS key. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapi-permissions.html#ebsapi-kms-permissions\"> Permissions to use Key Management Service keys</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> </important>");
        properties.put("Encrypted", EncryptedProperty);
        Map<String, Object> TagsProperty = new HashMap<>();
        TagsProperty.put("type", "string");
        TagsProperty.put("required", false);
        TagsProperty.put("description", "Input parameter: The tags to apply to the snapshot.");
        properties.put("Tags", TagsProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_snapshots",
            "No description available",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_SnapshotsHandler(config));
    }
    
}