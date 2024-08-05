package ysk_test

import (
	"context"
	"testing"

	"github.com/IceWhaleTech/CasaOS-MessageBus/pkg/ysk"
	"github.com/IceWhaleTech/CasaOS-MessageBus/utils"
	"gotest.tools/assert"
)

func mockPublish(context.Context, string, string, map[string]string) {
}

func TestUpdateProgress(t *testing.T) {

	err := ysk.NewYSKCard(context.Background(), utils.ApplicationInstallProgress.WithProgress(
		"Installing LinuxServer/Jellyfin", 50,
	), mockPublish)
	assert.NilError(t, err)
	err = ysk.DeleteCard(context.Background(), utils.ApplicationInstallProgress.Id, mockPublish)
	assert.NilError(t, err)

}

func TestNoticeDiskInsert(t *testing.T) {
	err := ysk.NewYSKCard(context.Background(), utils.DiskInsertNotice, mockPublish)
	assert.NilError(t, err)
}
