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
			TitleIcon:        ysk.AppStoreIcon,
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
			TitleIcon:    ysk.ZimaIcon,
			TitleText:    "Build data station",
			BodyProgress: nil,
			BodyIconWithText: &ysk.YSKCardIconWithText{
				Icon:        ysk.DiskIcon,
				Description: "For a data station with more storage capacity, it is recommended to add more hard drives.",
			},
			BodyList: nil,
			FooterActions: []ysk.YSKCardFooterAction{
				{
					Side:  ysk.ActionPositionRight,
					Style: "primary",
					Text:  "Learn more",
					MessageBus: ysk.YSKCardMessageBusAction{
						Key:     "casaos-ui/casaos-ui:app:mircoapp_communicate",
						Payload: `{"action":"open","peerType":"settings","name":"icewhale_settings","routerPath":"/storage"}`,
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
			TitleIcon:    ysk.ZimaIcon,
			TitleText:    "Remote Access",
			BodyProgress: nil,
			BodyIconWithText: &ysk.YSKCardIconWithText{
				Icon:        ysk.ZimaIcon,
				Description: "Configure Remote Access to access your home cloud remotely from anywhere.",
			},
			BodyList: nil,
			FooterActions: []ysk.YSKCardFooterAction{
				{
					Side:  ysk.ActionPositionRight,
					Style: "primary",
					Text:  "Learn more",
					MessageBus: ysk.YSKCardMessageBusAction{
						Key:     "casaos-ui/casaos-ui:app:mircoapp_communicate",
						Payload: `{"action":"open","peerType":"settings","name":"icewhale_settings","routerPath":"/network"}`,
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
			TitleIcon:    ysk.FileIcon,
			TitleText:    "File Management",
			BodyProgress: nil,
			BodyIconWithText: &ysk.YSKCardIconWithText{
				Icon:        ysk.FileIcon,
				Description: "Use Files to manage your data from different locations, such as your computer, phone, netdisk and server.",
			},
			BodyList: nil,
			FooterActions: []ysk.YSKCardFooterAction{
				{
					Side:  ysk.ActionPositionRight,
					Style: "primary",
					Text:  "Learn more",
					MessageBus: ysk.YSKCardMessageBusAction{
						Key:     "casaos-ui:open_files",
						Payload: `{"url": "/modules/icewhale_files/#"}`,
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
			TitleIcon: ysk.AppStoreIcon,
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
					Side:  ysk.ActionPositionRight,
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
			TitleIcon: ysk.ZimaIcon,
			TitleText: "Found a new device",
			BodyList: []ysk.YSKCardListItem{
				{
					Icon:        ysk.StorageIcon,
					Description: "ZimaOS-HD",
					RightText:   "2TB",
				}, {
					Icon:        ysk.StorageIcon,
					Description: "Safe-Storage",
					RightText:   "2TB",
				},
			},
			FooterActions: []ysk.YSKCardFooterAction{
				{
					Side:  ysk.ActionPositionRight,
					Style: "primary",
					Text:  "Manage",
					MessageBus: ysk.YSKCardMessageBusAction{
						Key:     "casaos-ui/casaos-ui:app:mircoapp_communicate",
						Payload: `{"action":"open","peerType":"settings","name":"icewhale_settings","routerPath":"/storage"}`,
					},
				},
			},
		},
	}
)
