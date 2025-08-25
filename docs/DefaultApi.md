# DefaultApi

All URIs are relative to *http://ebs.us-east-1.amazonaws.com*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**completeSnapshot**](DefaultApi.md#completeSnapshot) | **POST** /snapshots/completion/{snapshotId}#x-amz-ChangedBlocksCount |  |
| [**getSnapshotBlock**](DefaultApi.md#getSnapshotBlock) | **GET** /snapshots/{snapshotId}/blocks/{blockIndex}#blockToken |  |
| [**listChangedBlocks**](DefaultApi.md#listChangedBlocks) | **GET** /snapshots/{secondSnapshotId}/changedblocks |  |
| [**listSnapshotBlocks**](DefaultApi.md#listSnapshotBlocks) | **GET** /snapshots/{snapshotId}/blocks |  |
| [**putSnapshotBlock**](DefaultApi.md#putSnapshotBlock) | **PUT** /snapshots/{snapshotId}/blocks/{blockIndex}#x-amz-Data-Length&amp;x-amz-Checksum&amp;x-amz-Checksum-Algorithm |  |
| [**startSnapshot**](DefaultApi.md#startSnapshot) | **POST** /snapshots |  |


<a id="completeSnapshot"></a>
# **completeSnapshot**
> CompleteSnapshotResponse completeSnapshot(snapshotId, xAmzChangedBlocksCount, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, xAmzChecksum, xAmzChecksumAlgorithm, xAmzChecksumAggregationMethod)



&lt;p&gt;Seals and completes the snapshot after all of the required blocks of data have been written to it. Completing the snapshot changes the status to &lt;code&gt;completed&lt;/code&gt;. You cannot write new blocks to a snapshot after it has been completed.&lt;/p&gt; &lt;note&gt; &lt;p&gt;You should always retry requests that receive server (&lt;code&gt;5xx&lt;/code&gt;) error responses, and &lt;code&gt;ThrottlingException&lt;/code&gt; and &lt;code&gt;RequestThrottledException&lt;/code&gt; client error responses. For more information see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html\&quot;&gt;Error retries&lt;/a&gt; in the &lt;i&gt;Amazon Elastic Compute Cloud User Guide&lt;/i&gt;.&lt;/p&gt; &lt;/note&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://ebs.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String snapshotId = "snapshotId_example"; // String | The ID of the snapshot.
    Integer xAmzChangedBlocksCount = 56; // Integer | The number of blocks that were written to the snapshot.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    String xAmzChecksum = "xAmzChecksum_example"; // String | <p>An aggregated Base-64 SHA256 checksum based on the checksums of each written block.</p> <p>To generate the aggregated checksum using the linear aggregation method, arrange the checksums for each written block in ascending order of their block index, concatenate them to form a single string, and then generate the checksum on the entire string using the SHA256 algorithm.</p>
    String xAmzChecksumAlgorithm = "SHA256"; // String | The algorithm used to generate the checksum. Currently, the only supported algorithm is <code>SHA256</code>.
    String xAmzChecksumAggregationMethod = "LINEAR"; // String | The aggregation method used to generate the checksum. Currently, the only supported aggregation method is <code>LINEAR</code>.
    try {
      CompleteSnapshotResponse result = apiInstance.completeSnapshot(snapshotId, xAmzChangedBlocksCount, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, xAmzChecksum, xAmzChecksumAlgorithm, xAmzChecksumAggregationMethod);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#completeSnapshot");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **snapshotId** | **String**| The ID of the snapshot. | |
| **xAmzChangedBlocksCount** | **Integer**| The number of blocks that were written to the snapshot. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **xAmzChecksum** | **String**| &lt;p&gt;An aggregated Base-64 SHA256 checksum based on the checksums of each written block.&lt;/p&gt; &lt;p&gt;To generate the aggregated checksum using the linear aggregation method, arrange the checksums for each written block in ascending order of their block index, concatenate them to form a single string, and then generate the checksum on the entire string using the SHA256 algorithm.&lt;/p&gt; | [optional] |
| **xAmzChecksumAlgorithm** | **String**| The algorithm used to generate the checksum. Currently, the only supported algorithm is &lt;code&gt;SHA256&lt;/code&gt;. | [optional] [enum: SHA256] |
| **xAmzChecksumAggregationMethod** | **String**| The aggregation method used to generate the checksum. Currently, the only supported aggregation method is &lt;code&gt;LINEAR&lt;/code&gt;. | [optional] [enum: LINEAR] |

### Return type

[**CompleteSnapshotResponse**](CompleteSnapshotResponse.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **202** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ValidationException |  -  |
| **482** | ResourceNotFoundException |  -  |
| **483** | RequestThrottledException |  -  |
| **484** | ServiceQuotaExceededException |  -  |
| **485** | InternalServerException |  -  |

<a id="getSnapshotBlock"></a>
# **getSnapshotBlock**
> GetSnapshotBlockResponse getSnapshotBlock(snapshotId, blockIndex, blockToken, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



&lt;p&gt;Returns the data in a block in an Amazon Elastic Block Store snapshot.&lt;/p&gt; &lt;note&gt; &lt;p&gt;You should always retry requests that receive server (&lt;code&gt;5xx&lt;/code&gt;) error responses, and &lt;code&gt;ThrottlingException&lt;/code&gt; and &lt;code&gt;RequestThrottledException&lt;/code&gt; client error responses. For more information see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html\&quot;&gt;Error retries&lt;/a&gt; in the &lt;i&gt;Amazon Elastic Compute Cloud User Guide&lt;/i&gt;.&lt;/p&gt; &lt;/note&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://ebs.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String snapshotId = "snapshotId_example"; // String | <p>The ID of the snapshot containing the block from which to get data.</p> <important> <p>If the specified snapshot is encrypted, you must have permission to use the KMS key that was used to encrypt the snapshot. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html\"> Using encryption</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>.</p> </important>
    Integer blockIndex = 56; // Integer | The block index of the block in which to read the data. A block index is a logical index in units of <code>512</code> KiB blocks. To identify the block index, divide the logical offset of the data in the logical volume by the block size (logical offset of data/<code>524288</code>). The logical offset of the data must be <code>512</code> KiB aligned.
    String blockToken = "blockToken_example"; // String | The block token of the block from which to get data. You can obtain the <code>BlockToken</code> by running the <code>ListChangedBlocks</code> or <code>ListSnapshotBlocks</code> operations.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      GetSnapshotBlockResponse result = apiInstance.getSnapshotBlock(snapshotId, blockIndex, blockToken, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#getSnapshotBlock");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **snapshotId** | **String**| &lt;p&gt;The ID of the snapshot containing the block from which to get data.&lt;/p&gt; &lt;important&gt; &lt;p&gt;If the specified snapshot is encrypted, you must have permission to use the KMS key that was used to encrypt the snapshot. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html\&quot;&gt; Using encryption&lt;/a&gt; in the &lt;i&gt;Amazon Elastic Compute Cloud User Guide&lt;/i&gt;.&lt;/p&gt; &lt;/important&gt; | |
| **blockIndex** | **Integer**| The block index of the block in which to read the data. A block index is a logical index in units of &lt;code&gt;512&lt;/code&gt; KiB blocks. To identify the block index, divide the logical offset of the data in the logical volume by the block size (logical offset of data/&lt;code&gt;524288&lt;/code&gt;). The logical offset of the data must be &lt;code&gt;512&lt;/code&gt; KiB aligned. | |
| **blockToken** | **String**| The block token of the block from which to get data. You can obtain the &lt;code&gt;BlockToken&lt;/code&gt; by running the &lt;code&gt;ListChangedBlocks&lt;/code&gt; or &lt;code&gt;ListSnapshotBlocks&lt;/code&gt; operations. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**GetSnapshotBlockResponse**](GetSnapshotBlockResponse.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ValidationException |  -  |
| **482** | ResourceNotFoundException |  -  |
| **483** | RequestThrottledException |  -  |
| **484** | ServiceQuotaExceededException |  -  |
| **485** | InternalServerException |  -  |

<a id="listChangedBlocks"></a>
# **listChangedBlocks**
> ListChangedBlocksResponse listChangedBlocks(secondSnapshotId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, firstSnapshotId, pageToken, maxResults, startingBlockIndex, maxResults2, nextToken)



&lt;p&gt;Returns information about the blocks that are different between two Amazon Elastic Block Store snapshots of the same volume/snapshot lineage.&lt;/p&gt; &lt;note&gt; &lt;p&gt;You should always retry requests that receive server (&lt;code&gt;5xx&lt;/code&gt;) error responses, and &lt;code&gt;ThrottlingException&lt;/code&gt; and &lt;code&gt;RequestThrottledException&lt;/code&gt; client error responses. For more information see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html\&quot;&gt;Error retries&lt;/a&gt; in the &lt;i&gt;Amazon Elastic Compute Cloud User Guide&lt;/i&gt;.&lt;/p&gt; &lt;/note&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://ebs.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String secondSnapshotId = "secondSnapshotId_example"; // String | <p>The ID of the second snapshot to use for the comparison.</p> <important> <p>The <code>SecondSnapshotId</code> parameter must be specified with a <code>FirstSnapshotID</code> parameter; otherwise, an error occurs.</p> </important>
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    String firstSnapshotId = "firstSnapshotId_example"; // String | <p>The ID of the first snapshot to use for the comparison.</p> <important> <p>The <code>FirstSnapshotID</code> parameter must be specified with a <code>SecondSnapshotId</code> parameter; otherwise, an error occurs.</p> </important>
    String pageToken = "pageToken_example"; // String | <p>The token to request the next page of results.</p> <p>If you specify <b>NextToken</b>, then <b>StartingBlockIndex</b> is ignored.</p>
    Integer maxResults = 56; // Integer | <p>The maximum number of blocks to be returned by the request.</p> <p>Even if additional blocks can be retrieved from the snapshot, the request can return less blocks than <b>MaxResults</b> or an empty array of blocks.</p> <p>To retrieve the next set of blocks from the snapshot, make another request with the returned <b>NextToken</b> value. The value of <b>NextToken</b> is <code>null</code> when there are no more blocks to return.</p>
    Integer startingBlockIndex = 56; // Integer | <p>The block index from which the comparison should start.</p> <p>The list in the response will start from this block index or the next valid block index in the snapshots.</p> <p>If you specify <b>NextToken</b>, then <b>StartingBlockIndex</b> is ignored.</p>
    String maxResults2 = "maxResults_example"; // String | Pagination limit
    String nextToken = "nextToken_example"; // String | Pagination token
    try {
      ListChangedBlocksResponse result = apiInstance.listChangedBlocks(secondSnapshotId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, firstSnapshotId, pageToken, maxResults, startingBlockIndex, maxResults2, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listChangedBlocks");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **secondSnapshotId** | **String**| &lt;p&gt;The ID of the second snapshot to use for the comparison.&lt;/p&gt; &lt;important&gt; &lt;p&gt;The &lt;code&gt;SecondSnapshotId&lt;/code&gt; parameter must be specified with a &lt;code&gt;FirstSnapshotID&lt;/code&gt; parameter; otherwise, an error occurs.&lt;/p&gt; &lt;/important&gt; | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **firstSnapshotId** | **String**| &lt;p&gt;The ID of the first snapshot to use for the comparison.&lt;/p&gt; &lt;important&gt; &lt;p&gt;The &lt;code&gt;FirstSnapshotID&lt;/code&gt; parameter must be specified with a &lt;code&gt;SecondSnapshotId&lt;/code&gt; parameter; otherwise, an error occurs.&lt;/p&gt; &lt;/important&gt; | [optional] |
| **pageToken** | **String**| &lt;p&gt;The token to request the next page of results.&lt;/p&gt; &lt;p&gt;If you specify &lt;b&gt;NextToken&lt;/b&gt;, then &lt;b&gt;StartingBlockIndex&lt;/b&gt; is ignored.&lt;/p&gt; | [optional] |
| **maxResults** | **Integer**| &lt;p&gt;The maximum number of blocks to be returned by the request.&lt;/p&gt; &lt;p&gt;Even if additional blocks can be retrieved from the snapshot, the request can return less blocks than &lt;b&gt;MaxResults&lt;/b&gt; or an empty array of blocks.&lt;/p&gt; &lt;p&gt;To retrieve the next set of blocks from the snapshot, make another request with the returned &lt;b&gt;NextToken&lt;/b&gt; value. The value of &lt;b&gt;NextToken&lt;/b&gt; is &lt;code&gt;null&lt;/code&gt; when there are no more blocks to return.&lt;/p&gt; | [optional] |
| **startingBlockIndex** | **Integer**| &lt;p&gt;The block index from which the comparison should start.&lt;/p&gt; &lt;p&gt;The list in the response will start from this block index or the next valid block index in the snapshots.&lt;/p&gt; &lt;p&gt;If you specify &lt;b&gt;NextToken&lt;/b&gt;, then &lt;b&gt;StartingBlockIndex&lt;/b&gt; is ignored.&lt;/p&gt; | [optional] |
| **maxResults2** | **String**| Pagination limit | [optional] |
| **nextToken** | **String**| Pagination token | [optional] |

### Return type

[**ListChangedBlocksResponse**](ListChangedBlocksResponse.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ValidationException |  -  |
| **482** | ResourceNotFoundException |  -  |
| **483** | RequestThrottledException |  -  |
| **484** | ServiceQuotaExceededException |  -  |
| **485** | InternalServerException |  -  |

<a id="listSnapshotBlocks"></a>
# **listSnapshotBlocks**
> ListSnapshotBlocksResponse listSnapshotBlocks(snapshotId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, pageToken, maxResults, startingBlockIndex, maxResults2, nextToken)



&lt;p&gt;Returns information about the blocks in an Amazon Elastic Block Store snapshot.&lt;/p&gt; &lt;note&gt; &lt;p&gt;You should always retry requests that receive server (&lt;code&gt;5xx&lt;/code&gt;) error responses, and &lt;code&gt;ThrottlingException&lt;/code&gt; and &lt;code&gt;RequestThrottledException&lt;/code&gt; client error responses. For more information see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html\&quot;&gt;Error retries&lt;/a&gt; in the &lt;i&gt;Amazon Elastic Compute Cloud User Guide&lt;/i&gt;.&lt;/p&gt; &lt;/note&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://ebs.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String snapshotId = "snapshotId_example"; // String | The ID of the snapshot from which to get block indexes and block tokens.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    String pageToken = "pageToken_example"; // String | <p>The token to request the next page of results.</p> <p>If you specify <b>NextToken</b>, then <b>StartingBlockIndex</b> is ignored.</p>
    Integer maxResults = 56; // Integer | <p>The maximum number of blocks to be returned by the request.</p> <p>Even if additional blocks can be retrieved from the snapshot, the request can return less blocks than <b>MaxResults</b> or an empty array of blocks.</p> <p>To retrieve the next set of blocks from the snapshot, make another request with the returned <b>NextToken</b> value. The value of <b>NextToken</b> is <code>null</code> when there are no more blocks to return.</p>
    Integer startingBlockIndex = 56; // Integer | <p>The block index from which the list should start. The list in the response will start from this block index or the next valid block index in the snapshot.</p> <p>If you specify <b>NextToken</b>, then <b>StartingBlockIndex</b> is ignored.</p>
    String maxResults2 = "maxResults_example"; // String | Pagination limit
    String nextToken = "nextToken_example"; // String | Pagination token
    try {
      ListSnapshotBlocksResponse result = apiInstance.listSnapshotBlocks(snapshotId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, pageToken, maxResults, startingBlockIndex, maxResults2, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listSnapshotBlocks");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **snapshotId** | **String**| The ID of the snapshot from which to get block indexes and block tokens. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **pageToken** | **String**| &lt;p&gt;The token to request the next page of results.&lt;/p&gt; &lt;p&gt;If you specify &lt;b&gt;NextToken&lt;/b&gt;, then &lt;b&gt;StartingBlockIndex&lt;/b&gt; is ignored.&lt;/p&gt; | [optional] |
| **maxResults** | **Integer**| &lt;p&gt;The maximum number of blocks to be returned by the request.&lt;/p&gt; &lt;p&gt;Even if additional blocks can be retrieved from the snapshot, the request can return less blocks than &lt;b&gt;MaxResults&lt;/b&gt; or an empty array of blocks.&lt;/p&gt; &lt;p&gt;To retrieve the next set of blocks from the snapshot, make another request with the returned &lt;b&gt;NextToken&lt;/b&gt; value. The value of &lt;b&gt;NextToken&lt;/b&gt; is &lt;code&gt;null&lt;/code&gt; when there are no more blocks to return.&lt;/p&gt; | [optional] |
| **startingBlockIndex** | **Integer**| &lt;p&gt;The block index from which the list should start. The list in the response will start from this block index or the next valid block index in the snapshot.&lt;/p&gt; &lt;p&gt;If you specify &lt;b&gt;NextToken&lt;/b&gt;, then &lt;b&gt;StartingBlockIndex&lt;/b&gt; is ignored.&lt;/p&gt; | [optional] |
| **maxResults2** | **String**| Pagination limit | [optional] |
| **nextToken** | **String**| Pagination token | [optional] |

### Return type

[**ListSnapshotBlocksResponse**](ListSnapshotBlocksResponse.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ValidationException |  -  |
| **482** | ResourceNotFoundException |  -  |
| **483** | RequestThrottledException |  -  |
| **484** | ServiceQuotaExceededException |  -  |
| **485** | InternalServerException |  -  |

<a id="putSnapshotBlock"></a>
# **putSnapshotBlock**
> Object putSnapshotBlock(snapshotId, blockIndex, xAmzDataLength, xAmzChecksum, xAmzChecksumAlgorithm, putSnapshotBlockRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, xAmzProgress)



&lt;p&gt;Writes a block of data to a snapshot. If the specified block contains data, the existing data is overwritten. The target snapshot must be in the &lt;code&gt;pending&lt;/code&gt; state.&lt;/p&gt; &lt;p&gt;Data written to a snapshot must be aligned with 512-KiB sectors.&lt;/p&gt; &lt;note&gt; &lt;p&gt;You should always retry requests that receive server (&lt;code&gt;5xx&lt;/code&gt;) error responses, and &lt;code&gt;ThrottlingException&lt;/code&gt; and &lt;code&gt;RequestThrottledException&lt;/code&gt; client error responses. For more information see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html\&quot;&gt;Error retries&lt;/a&gt; in the &lt;i&gt;Amazon Elastic Compute Cloud User Guide&lt;/i&gt;.&lt;/p&gt; &lt;/note&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://ebs.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String snapshotId = "snapshotId_example"; // String | <p>The ID of the snapshot.</p> <important> <p>If the specified snapshot is encrypted, you must have permission to use the KMS key that was used to encrypt the snapshot. For more information, see <a href=\"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html\"> Using encryption</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>..</p> </important>
    Integer blockIndex = 56; // Integer | The block index of the block in which to write the data. A block index is a logical index in units of <code>512</code> KiB blocks. To identify the block index, divide the logical offset of the data in the logical volume by the block size (logical offset of data/<code>524288</code>). The logical offset of the data must be <code>512</code> KiB aligned.
    Integer xAmzDataLength = 56; // Integer | <p>The size of the data to write to the block, in bytes. Currently, the only supported size is <code>524288</code> bytes.</p> <p>Valid values: <code>524288</code> </p>
    String xAmzChecksum = "xAmzChecksum_example"; // String | A Base64-encoded SHA256 checksum of the data. Only SHA256 checksums are supported.
    String xAmzChecksumAlgorithm = "SHA256"; // String | The algorithm used to generate the checksum. Currently, the only supported algorithm is <code>SHA256</code>.
    PutSnapshotBlockRequest putSnapshotBlockRequest = new PutSnapshotBlockRequest(); // PutSnapshotBlockRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    Integer xAmzProgress = 56; // Integer | The progress of the write process, as a percentage.
    try {
      Object result = apiInstance.putSnapshotBlock(snapshotId, blockIndex, xAmzDataLength, xAmzChecksum, xAmzChecksumAlgorithm, putSnapshotBlockRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, xAmzProgress);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#putSnapshotBlock");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **snapshotId** | **String**| &lt;p&gt;The ID of the snapshot.&lt;/p&gt; &lt;important&gt; &lt;p&gt;If the specified snapshot is encrypted, you must have permission to use the KMS key that was used to encrypt the snapshot. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html\&quot;&gt; Using encryption&lt;/a&gt; in the &lt;i&gt;Amazon Elastic Compute Cloud User Guide&lt;/i&gt;..&lt;/p&gt; &lt;/important&gt; | |
| **blockIndex** | **Integer**| The block index of the block in which to write the data. A block index is a logical index in units of &lt;code&gt;512&lt;/code&gt; KiB blocks. To identify the block index, divide the logical offset of the data in the logical volume by the block size (logical offset of data/&lt;code&gt;524288&lt;/code&gt;). The logical offset of the data must be &lt;code&gt;512&lt;/code&gt; KiB aligned. | |
| **xAmzDataLength** | **Integer**| &lt;p&gt;The size of the data to write to the block, in bytes. Currently, the only supported size is &lt;code&gt;524288&lt;/code&gt; bytes.&lt;/p&gt; &lt;p&gt;Valid values: &lt;code&gt;524288&lt;/code&gt; &lt;/p&gt; | |
| **xAmzChecksum** | **String**| A Base64-encoded SHA256 checksum of the data. Only SHA256 checksums are supported. | |
| **xAmzChecksumAlgorithm** | **String**| The algorithm used to generate the checksum. Currently, the only supported algorithm is &lt;code&gt;SHA256&lt;/code&gt;. | [enum: SHA256] |
| **putSnapshotBlockRequest** | [**PutSnapshotBlockRequest**](PutSnapshotBlockRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **xAmzProgress** | **Integer**| The progress of the write process, as a percentage. | [optional] |

### Return type

**Object**

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ValidationException |  -  |
| **482** | ResourceNotFoundException |  -  |
| **483** | RequestThrottledException |  -  |
| **484** | ServiceQuotaExceededException |  -  |
| **485** | InternalServerException |  -  |

<a id="startSnapshot"></a>
# **startSnapshot**
> StartSnapshotResponse startSnapshot(startSnapshotRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



&lt;p&gt;Creates a new Amazon EBS snapshot. The new snapshot enters the &lt;code&gt;pending&lt;/code&gt; state after the request completes. &lt;/p&gt; &lt;p&gt;After creating the snapshot, use &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/ebs/latest/APIReference/API_PutSnapshotBlock.html\&quot;&gt; PutSnapshotBlock&lt;/a&gt; to write blocks of data to the snapshot.&lt;/p&gt; &lt;note&gt; &lt;p&gt;You should always retry requests that receive server (&lt;code&gt;5xx&lt;/code&gt;) error responses, and &lt;code&gt;ThrottlingException&lt;/code&gt; and &lt;code&gt;RequestThrottledException&lt;/code&gt; client error responses. For more information see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html\&quot;&gt;Error retries&lt;/a&gt; in the &lt;i&gt;Amazon Elastic Compute Cloud User Guide&lt;/i&gt;.&lt;/p&gt; &lt;/note&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://ebs.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    StartSnapshotRequest startSnapshotRequest = new StartSnapshotRequest(); // StartSnapshotRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      StartSnapshotResponse result = apiInstance.startSnapshot(startSnapshotRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#startSnapshot");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **startSnapshotRequest** | [**StartSnapshotRequest**](StartSnapshotRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**StartSnapshotResponse**](StartSnapshotResponse.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ValidationException |  -  |
| **482** | RequestThrottledException |  -  |
| **483** | ResourceNotFoundException |  -  |
| **484** | ServiceQuotaExceededException |  -  |
| **485** | InternalServerException |  -  |
| **486** | ConcurrentLimitExceededException |  -  |
| **487** | ConflictException |  -  |

