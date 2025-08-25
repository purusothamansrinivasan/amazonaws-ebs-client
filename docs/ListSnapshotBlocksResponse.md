

# ListSnapshotBlocksResponse


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**blocks** | [**List&lt;Block&gt;**](Block.md) | An array of objects containing information about the blocks. |  [optional] |
|**expiryTime** | **OffsetDateTime** | The time when the &lt;code&gt;BlockToken&lt;/code&gt; expires. |  [optional] |
|**volumeSize** | **Integer** | The size of the volume in GB. |  [optional] |
|**blockSize** | **Integer** | The size of the blocks in the snapshot, in bytes. |  [optional] |
|**nextToken** | **String** | The token to use to retrieve the next page of results. This value is null when there are no more results to return. |  [optional] |



