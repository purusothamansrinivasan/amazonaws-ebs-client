

# ChangedBlock

A block of data in an Amazon Elastic Block Store snapshot that is different from another snapshot of the same volume/snapshot lineage.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**blockIndex** | **Integer** | The block index. |  [optional] |
|**firstBlockToken** | **String** | The block token for the block index of the &lt;code&gt;FirstSnapshotId&lt;/code&gt; specified in the &lt;code&gt;ListChangedBlocks&lt;/code&gt; operation. This value is absent if the first snapshot does not have the changed block that is on the second snapshot. |  [optional] |
|**secondBlockToken** | **String** | The block token for the block index of the &lt;code&gt;SecondSnapshotId&lt;/code&gt; specified in the &lt;code&gt;ListChangedBlocks&lt;/code&gt; operation. |  [optional] |



