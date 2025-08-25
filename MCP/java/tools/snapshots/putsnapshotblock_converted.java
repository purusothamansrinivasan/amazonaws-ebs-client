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

class Put_Snapshots_Snapshot_Id_Blocks_Block_Index_X_Amz_Data_Length_X_Amz_Checksum_X_Amz_Checksum_AlgorithmMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPut_Snapshots_Snapshot_Id_Blocks_Block_Index_X_Amz_Data_Length_X_Amz_Checksum_X_Amz_Checksum_AlgorithmHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("snapshotId")) {
            queryParams.add("snapshotId=" + args.get("snapshotId"));
        }
        if (args.containsKey("x-amz-Checksum")) {
            queryParams.add("x-amz-Checksum=" + args.get("x-amz-Checksum"));
        }
        if (args.containsKey("x-amz-Checksum-Algorithm")) {
            queryParams.add("x-amz-Checksum-Algorithm=" + args.get("x-amz-Checksum-Algorithm"));
        }
        if (args.containsKey("BlockData")) {
            queryParams.add("BlockData=" + args.get("BlockData"));
        }
        if (args.containsKey("blockIndex")) {
            queryParams.add("blockIndex=" + args.get("blockIndex"));
        }
        if (args.containsKey("x-amz-Data-Length")) {
            queryParams.add("x-amz-Data-Length=" + args.get("x-amz-Data-Length"));
        }
        if (args.containsKey("x-amz-Progress")) {
            queryParams.add("x-amz-Progress=" + args.get("x-amz-Progress"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/put_snapshots_snapshot_id_blocks_block_index_x_amz_data_length_x_amz_checksum_x_amz_checksum_algorithm" + queryString;
                
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
    
    public static MCPServer.Tool createPut_Snapshots_Snapshot_Id_Blocks_Block_Index_X_Amz_Data_Length_X_Amz_Checksum_X_Amz_Checksum_AlgorithmTool(MCPServer.APIConfig config) {
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
        Map<String, Object> snapshotIdProperty = new HashMap<>();
        snapshotIdProperty.put("type", "string");
        snapshotIdProperty.put("required", true);
        snapshotIdProperty.put("description", "<p>The ID of the snapshot.</p> <important> <p>If the specified snapshot is encrypted, you must have permission to use the KMS key that was used to encrypt the snapshot. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html\"> Using encryption</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>..</p> </important>");
        properties.put("snapshotId", snapshotIdProperty);
        Map<String, Object> x-amz-ChecksumProperty = new HashMap<>();
        x-amz-ChecksumProperty.put("type", "string");
        x-amz-ChecksumProperty.put("required", true);
        x-amz-ChecksumProperty.put("description", "A Base64-encoded SHA256 checksum of the data. Only SHA256 checksums are supported.");
        properties.put("x-amz-Checksum", x-amz-ChecksumProperty);
        Map<String, Object> x-amz-Checksum-AlgorithmProperty = new HashMap<>();
        x-amz-Checksum-AlgorithmProperty.put("type", "string");
        x-amz-Checksum-AlgorithmProperty.put("required", true);
        x-amz-Checksum-AlgorithmProperty.put("description", "The algorithm used to generate the checksum. Currently, the only supported algorithm is <code>SHA256</code>.");
        properties.put("x-amz-Checksum-Algorithm", x-amz-Checksum-AlgorithmProperty);
        Map<String, Object> BlockDataProperty = new HashMap<>();
        BlockDataProperty.put("type", "string");
        BlockDataProperty.put("required", true);
        BlockDataProperty.put("description", "Input parameter: <p>The data to write to the block.</p> <p>The block data is not signed as part of the Signature Version 4 signing process. As a result, you must generate and provide a Base64-encoded SHA256 checksum for the block data using the <b>x-amz-Checksum</b> header. Also, you must specify the checksum algorithm using the <b>x-amz-Checksum-Algorithm</b> header. The checksum that you provide is part of the Signature Version 4 signing process. It is validated against a checksum generated by Amazon EBS to ensure the validity and authenticity of the data. If the checksums do not correspond, the request fails. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-accessing-snapshot.html#ebsapis-using-checksums\"> Using checksums with the EBS direct APIs</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p>");
        properties.put("BlockData", BlockDataProperty);
        Map<String, Object> blockIndexProperty = new HashMap<>();
        blockIndexProperty.put("type", "string");
        blockIndexProperty.put("required", true);
        blockIndexProperty.put("description", "The block index of the block in which to write the data. A block index is a logical index in units of <code>512</code> KiB blocks. To identify the block index, divide the logical offset of the data in the logical volume by the block size (logical offset of data/<code>524288</code>). The logical offset of the data must be <code>512</code> KiB aligned.");
        properties.put("blockIndex", blockIndexProperty);
        Map<String, Object> x-amz-Data-LengthProperty = new HashMap<>();
        x-amz-Data-LengthProperty.put("type", "string");
        x-amz-Data-LengthProperty.put("required", true);
        x-amz-Data-LengthProperty.put("description", "<p>The size of the data to write to the block, in bytes. Currently, the only supported size is <code>524288</code> bytes.</p> <p>Valid values: <code>524288</code> </p>");
        properties.put("x-amz-Data-Length", x-amz-Data-LengthProperty);
        Map<String, Object> x-amz-ProgressProperty = new HashMap<>();
        x-amz-ProgressProperty.put("type", "string");
        x-amz-ProgressProperty.put("required", false);
        x-amz-ProgressProperty.put("description", "The progress of the write process, as a percentage.");
        properties.put("x-amz-Progress", x-amz-ProgressProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "put_snapshots_snapshot_id_blocks_block_index_x_amz_data_length_x_amz_checksum_x_amz_checksum_algorithm",
            "No description available",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPut_Snapshots_Snapshot_Id_Blocks_Block_Index_X_Amz_Data_Length_X_Amz_Checksum_X_Amz_Checksum_AlgorithmHandler(config));
    }
    
}