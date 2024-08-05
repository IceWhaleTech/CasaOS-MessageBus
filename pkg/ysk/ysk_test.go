package ysk_test

import (
	"context"
	"testing"

	"github.com/IceWhaleTech/CasaOS-MessageBus/pkg/ysk"
	"github.com/IceWhaleTech/CasaOS-MessageBus/utils"

	"gotest.tools/assert"
)

func TestUpdateProgress(t *testing.T) {

	err := ysk.NewYSKCard(context.Background(), ysk.TaskWithProgress(
		utils.ApplicationInstallProgress,
		"Installing LinuxServer/Jellyfin",
		50,
	), nil)
	assert.NilError(t, err)
	err = ysk.DeleteCard(context.Background(), utils.ApplicationInstallProgress.Id, nil)
	assert.NilError(t, err)

}

func TestNoticeDiskInsert(t *testing.T) {
	err := ysk.NewYSKCard(context.Background(), utils.DiskInsertNotice, nil)
	assert.NilError(t, err)
}
