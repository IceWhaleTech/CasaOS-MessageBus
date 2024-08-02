package ysk_test

import (
	"context"
	"testing"

	"github.com/IceWhaleTech/ZimaOS/pkg/ysk"
)

func TestUpdateProgress(t *testing.T) {
	ApplicationInstallProgress := ysk.YSKCard{
		ID:         "task:application:install",
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

	ysk.NewYSKCard(context.Background(), ApplicationInstallProgress.WithProgress(
		"Installing LinuxServer/Jellyfin",
		50,
	), nil)

	ysk.DeleteCard(context.Background(), ApplicationInstallProgress.ID, nil)
}

func TestNoticeInUser
