package ysk_test

import (
	"context"
	"testing"

	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/pkg/ysk"
	"gotest.tools/assert"
)

func TestUpdateProgress(t *testing.T) {
	ApplicationInstallProgress := codegen.YSKCard{
		Id:         "task:application:install",
		CardType:   codegen.YSKCardCardTypeTask,
		RenderType: codegen.YSKCardRenderTypeTask,
		Content: codegen.YSKCardContent{
			TitleIcon:        "jellyfin logo",
			TitleText:        "APP Installing",
			BodyProgress:     &codegen.YSKCardProgress{},
			BodyIconWithText: nil,
			BodyList:         nil,
			FooterActions:    nil,
		},
	}

	err := ysk.NewYSKCard(context.Background(), ysk.TaskWithProgress(
		ApplicationInstallProgress,
		"Installing LinuxServer/Jellyfin",
		50,
	), nil)
	assert.NilError(t, err)
	err = ysk.DeleteCard(context.Background(), ApplicationInstallProgress.Id, nil)
	assert.NilError(t, err)

}

func TestNoticeDiskInsert(t *testing.T) {
	DiskInsertNotice := codegen.YSKCard{
		Id:         "long-notice:disk:insert",
		CardType:   codegen.YSKCardCardTypeLongNotice,
		RenderType: codegen.YSKCardRenderTypeListNotice,
		Content: codegen.YSKCardContent{
			TitleIcon:    "ZimaOS-Logo",
			TitleText:    "创建数据工作站",
			BodyProgress: nil,
			BodyIconWithText: &codegen.YSKCardIconWithText{
				Icon:        "disk",
				Description: "通过添加硬盘驱动器或固态硬盘来增强你的个人主机，并建立自己的个人数据中心。",
			},
			BodyList: nil,
			FooterActions: &[]codegen.YSKCardFooterAction{
				{
					Side:  "Right",
					Style: "primary",
					Text:  "创建数据工作站",
					MessageBus: codegen.YSKCardMessageBusAction{
						Key:     "open:disk:insert",
						Payload: "{'type':'disk'}",
					},
				},
			},
		},
	}
	err := ysk.NewYSKCard(context.Background(), DiskInsertNotice, nil)
	assert.NilError(t, err)

}
