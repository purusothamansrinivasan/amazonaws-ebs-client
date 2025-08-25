

# StartSnapshotResponse


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**description** | **String** | The description of the snapshot. |  [optional] |
|**snapshotId** | **String** | The ID of the snapshot. |  [optional] |
|**ownerId** | **String** | The Amazon Web Services account ID of the snapshot owner. |  [optional] |
|**status** | **Status** | The status of the snapshot. |  [optional] |
|**startTime** | **OffsetDateTime** | The timestamp when the snapshot was created. |  [optional] |
|**volumeSize** | **Integer** | The size of the volume, in GiB. |  [optional] |
|**blockSize** | **Integer** | The size of the blocks in the snapshot, in bytes. |  [optional] |
|**tags** | [**List&lt;Tag&gt;**](Tag.md) | The tags applied to the snapshot. You can specify up to 50 tags per snapshot. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html\&quot;&gt; Tagging your Amazon EC2 resources&lt;/a&gt; in the &lt;i&gt;Amazon Elastic Compute Cloud User Guide&lt;/i&gt;. |  [optional] |
|**parentSnapshotId** | **String** | The ID of the parent snapshot. |  [optional] |
|**kmsKeyArn** | **String** | The Amazon Resource Name (ARN) of the Key Management Service (KMS) key used to encrypt the snapshot. |  [optional] |
|**sseType** | **SSEType** | Reserved for future use. |  [optional] |



