package service_test

import (
	"context"
	"testing"

	"github.com/IceWhaleTech/CasaOS-MessageBus/pkg/ysk"
	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
	"github.com/IceWhaleTech/CasaOS-MessageBus/service"
	"github.com/IceWhaleTech/CasaOS-MessageBus/utils"
	"gotest.tools/assert"
)

func setup(t *testing.T) *service.YSKService {
	repository, err := repository.NewDatabaseRepositoryInMemory()
	assert.NilError(t, err)
	defer repository.Close()

	yskService := service.NewYSKService(&repository)
	return yskService
}

func TestInsertAndGetCardList(t *testing.T) {
	yskService := setup(t)

	cardList, err := yskService.YskCardList(context.Background())
	assert.NilError(t, err)
	assert.Equal(t, len(cardList), 0)

	cardInsertQueue := []ysk.YSKCard{
		utils.ApplicationInstallProgress, utils.DiskInsertNotice,
	}

	for _, card := range cardInsertQueue {
		err = yskService.UpsertYSKCard(context.Background(), card)
		assert.NilError(t, err)
	}

	cardList, err = yskService.YskCardList(context.Background())
	assert.NilError(t, err)
	assert.Equal(t, len(cardList), 2)

	for _, card := range cardList {
		if card.Id == utils.ApplicationInstallProgress.Id {
			assert.DeepEqual(t, card, utils.ApplicationInstallProgress)
		} else if card.Id == utils.DiskInsertNotice.Id {
			assert.DeepEqual(t, card, utils.DiskInsertNotice)
		} else {
			t.Errorf("unexpected card: %v", card)
		}
	}
}

func TestInsertAllTypeCardList(t *testing.T) {
	yskService := setup(t)

	cardList, err := yskService.YskCardList(context.Background())
	assert.NilError(t, err)
	assert.Equal(t, len(cardList), 0)

	cardInsertQueue := []ysk.YSKCard{
		utils.ApplicationInstallProgress, utils.DiskInsertNotice, utils.ApplicationUpdateNotice,
		utils.ApplicationInstallProgress.WithProgress("Installing LinuxServer/Jellyfin", 50), utils.ApplicationInstallProgress.WithProgress("Installing LinuxServer/Jellyfin", 55),
		utils.ApplicationInstallProgress.WithProgress("Installing LinuxServer/Jellyfin", 80), utils.ApplicationInstallProgress.WithProgress("Installing LinuxServer/Jellyfin", 99),
		utils.ApplicationUpdateNotice,
	}

	for _, card := range cardInsertQueue {
		err = yskService.UpsertYSKCard(context.Background(), card)
		assert.NilError(t, err)
	}

	cardList, err = yskService.YskCardList(context.Background())
	assert.NilError(t, err)
	assert.Equal(t, len(cardList), 2)

}
