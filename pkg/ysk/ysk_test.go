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
		Content: ysk.YSKCardContent{
			TitleIcon:        "jellyfin logo",
			TitleText:        "APP Installing",
			BodyProgress:     &ysk.YSKCardProgress{},
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
		Content: ysk.YSKCardContent{
			TitleIcon:        "jellyfin logo",
			TitleText:        "APP Installing",
			BodyProgress:     nil,
			BodyIconWithText: nil,
			BodyList:         nil,
			FooterActions:    nil,
		},
	}
	err := ysk.NewYSKCard(context.Background(), DiskInsertNotice, nil)
	assert.NilError(t, err)

}
