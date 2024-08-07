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

	ZimaOSDataStationNotice = ysk.YSKCard{
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
						Key:     "casaos-ui/casaos-ui:app:mircoapp_communicate",
						Payload: `{"access_id":"qWUS_pKWefbN-Bcxu3_nG","casaos_lang":"zh_cn","action":"open","peerType":"settings","name":"icewhale_settings","routerPath":"/storage"}`,
					},
				},
			},
		},
	}

	ZimaOSRemoteAccessNotice = ysk.YSKCard{
		Id:         "long-notice:remote:access",
		CardType:   ysk.CardTypeLongNote,
		RenderType: ysk.RenderTypeCardIconTextNotice,
		Content: ysk.YSKCardContent{
			TitleIcon:    "ZimaOS-Logo",
			TitleText:    "远程访问",
			BodyProgress: nil,
			BodyIconWithText: &ysk.YSKCardIconWithText{
				Icon:        "remote access",
				Description: "通过远程访问，您可以随时随地访问您的个人主机。",
			},
			BodyList: nil,
			FooterActions: []ysk.YSKCardFooterAction{
				{
					Side:  "Right",
					Style: "primary",
					Text:  "learn more",
					MessageBus: ysk.YSKCardMessageBusAction{
						Key:     "casaos-ui/casaos-ui:app:mircoapp_communicate",
						Payload: `{"access_id":"1733L6fM4PHol8kRssFvK","casaos_lang":"zh_cn","action":"open","peerType":"settings","name":"icewhale_settings","routerPath":"/network"}`,
					},
				},
			},
		},
	}

	ZimaOSFileManagementNotice = ysk.YSKCard{
		Id:         "long-notice:file:management",
		CardType:   ysk.CardTypeLongNote,
		RenderType: ysk.RenderTypeCardIconTextNotice,
		Content: ysk.YSKCardContent{
			TitleIcon:    "ZimaOS-Logo",
			TitleText:    "文件管理",
			BodyProgress: nil,
			BodyIconWithText: &ysk.YSKCardIconWithText{
				Icon:        "file management",
				Description: "使用Files来管理分散的数据，包括备份、 云存储、 NAS 或本地网络中的其他个人数据。",
			},
			BodyList: nil,
			FooterActions: []ysk.YSKCardFooterAction{
				{
					Side:  "Right",
					Style: "primary",
					Text:  "learn more",
					MessageBus: ysk.YSKCardMessageBusAction{
						Key:     "casaos-ui/casaos-ui:app:mircoapp_communicate",
						Payload: "{'type':'file'}",
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

	DiskInsertNotice = ysk.YSKCard{
		Id:         "long-notice:disk:insert",
		CardType:   ysk.CardTypeLongNote,
		RenderType: ysk.RenderTypeCardListNotice,
		Content: ysk.YSKCardContent{
			TitleIcon: "disk logo",
			TitleText: "硬盘插入",
			BodyList: []ysk.YSKCardListItem{
				{
					Icon:        "disk",
					Description: "ZimaOS-HD",
					RightText:   "2TB",
				}, {
					Icon:        "disk",
					Description: "Safe-Storage",
					RightText:   "2TB",
				},
			},
			FooterActions: nil,
		},
	}
)
