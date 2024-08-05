package utils

import (
	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
)

var (
	ApplicationInstallProgress = codegen.YSKCard{
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

	DiskInsertNotice = codegen.YSKCard{
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
)
