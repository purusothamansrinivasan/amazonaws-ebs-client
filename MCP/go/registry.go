package main

import (
	"github.com/amazon-elastic-block-store/mcp-server/config"
	"github.com/amazon-elastic-block-store/mcp-server/models"
	tools_snapshots "github.com/amazon-elastic-block-store/mcp-server/tools/snapshots"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_snapshots.CreateStartsnapshotTool(cfg),
		tools_snapshots.CreateCompletesnapshotTool(cfg),
		tools_snapshots.CreateListchangedblocksTool(cfg),
		tools_snapshots.CreateListsnapshotblocksTool(cfg),
		tools_snapshots.CreateGetsnapshotblockTool(cfg),
		tools_snapshots.CreatePutsnapshotblockTool(cfg),
	}
}
