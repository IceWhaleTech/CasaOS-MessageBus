package ysk_test

import (
	"context"
	"testing"

	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/IceWhaleTech/CasaOS-MessageBus/model"
	"github.com/IceWhaleTech/CasaOS-MessageBus/pkg/ysk"
	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
	"github.com/IceWhaleTech/CasaOS-MessageBus/service"
	"github.com/IceWhaleTech/CasaOS-MessageBus/utils"
	"gotest.tools/assert"
)

var ws *service.EventServiceWS

func setup(t *testing.T) (*service.EventServiceWS, *service.YSKService, func()) {
	repository, err := repository.NewDatabaseRepositoryInMemory()
	assert.NilError(t, err)
	s := service.NewServices(&repository)
	wsService := s.EventServiceWS
	yskService := s.YSKService

	ctx := context.Background()
	go s.Start(&ctx)
	return wsService, yskService, func() {
		repository.Close()
	}
}
func mockPublish(ctx context.Context, sourceID string, eventName string, body map[string]string) {
	if ws != nil {
		ws.Publish(model.Event{
			SourceID:   sourceID,
			Name:       eventName,
			Properties: body,
		})
	}
}

func TestUpdateProgress(t *testing.T) {
	logger.LogInitConsoleOnly()

	wsService, yskService, cleanup := setup(t)
	defer cleanup()
	ws = wsService

	err := ysk.NewYSKCard(context.Background(), utils.ApplicationInstallProgress.WithProgress(
		"Installing LinuxServer/Jellyfin", 50,
	), mockPublish)
	assert.NilError(t, err)

	cards, err := yskService.YskCardList(context.Background())
	assert.NilError(t, err)
	assert.Equal(t, len(cards), 1)

	assert.NilError(t, err)
	err = ysk.DeleteCard(context.Background(), utils.ApplicationInstallProgress.Id, mockPublish)
	assert.NilError(t, err)

	cards, err = yskService.YskCardList(context.Background())
	assert.NilError(t, err)
	assert.Equal(t, len(cards), 0)
}

func TestNoticeDiskInsert(t *testing.T) {
	err := ysk.NewYSKCard(context.Background(), utils.DiskInsertNotice, mockPublish)
	assert.NilError(t, err)
}
