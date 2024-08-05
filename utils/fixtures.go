package utils

import (
	"github.com/IceWhaleTech/CasaOS-MessageBus/pkg/ysk"
)

var (
	ApplicationInstallProgress = ysk.YSKCard{
		Id:         "task:application:install",
		CardType:   ysk.CardTypeTask,
		RenderType: ysk.RenderTypeCardTask,
		Content: ysk.YSKCardContent{
			TitleIcon:        "jellyfin logo",
			TitleText:        "APP Installing",
			BodyProgress:     &ysk.YSKCardProgress{},
			BodyIconWithText: nil,
			BodyList:         nil,
			FooterActions:    nil,
		},
	}

	DiskInsertNotice = ysk.YSKCard{
		Id:         "long-notice:disk:insert",
		CardType:   ysk.CardTypeLongNote,
		RenderType: ysk.RenderTypeCardIconTextNotice,
		Content: ysk.YSKCardContent{
			TitleIcon:    "ZimaOS-Logo",
			TitleText:    "创建数据工作站",
			BodyProgress: nil,
			BodyIconWithText: &ysk.YSKCardIconWithText{
				Icon:        "disk",
				Description: "通过添加硬盘驱动器或固态硬盘来增强你的个人主机，并建立自己的个人数据中心。",
			},
			BodyList: nil,
			FooterActions: []ysk.YSKCardFooterAction{
				{
					Side:  "Right",
					Style: "primary",
					Text:  "创建数据工作站",
					MessageBus: ysk.YSKCardMessageBusAction{
						Key:     "open:disk:insert",
						Payload: "{'type':'disk'}",
					},
				},
			},
		},
	}

	ApplicationUpdateNotice = ysk.YSKCard{
		Id:         "short-notice:application:update",
		CardType:   ysk.CardTypeShortNote,
		RenderType: ysk.RenderTypeCardListNotice,
		Content: ysk.YSKCardContent{
			TitleIcon: "app store logo",
			TitleText: "有应用更新",
			BodyList: []ysk.YSKCardListItem{
				{
					Icon:        "jellyfin logo",
					Description: "Jellyfin 10.7.0",
					RightText:   "2 days ago",
				},
				{
					Icon:        "next-cloud logo",
					Description: "NextCloud 10.7.0",
					RightText:   "2 days ago",
				},
			},
			FooterActions: []ysk.YSKCardFooterAction{
				{
					Side:  "Right",
					Style: "primary",
					Text:  "更新所有",
					MessageBus: ysk.YSKCardMessageBusAction{
						Key:     "open:application:update",
						Payload: "{'type':'all'}",
					},
				},
			},
		},
	}
)
