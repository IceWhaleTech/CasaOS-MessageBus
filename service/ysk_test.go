package service_test

import (
	"context"
	"testing"

	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
	"github.com/IceWhaleTech/CasaOS-MessageBus/service"
	"github.com/IceWhaleTech/CasaOS-MessageBus/utils"
	"gotest.tools/assert"
)

func TestInsertAndGetCardList(t *testing.T) {
	repository, err := repository.NewDatabaseRepositoryInMemory()
	assert.NilError(t, err)
	defer repository.Close()

	yskService := service.NewYSKService(&repository)

	cardList, err := yskService.YskCardList(context.Background())
	assert.NilError(t, err)
	assert.Equal(t, len(cardList), 0)

	err = yskService.UpsertYSKCard(context.Background(), utils.ApplicationInstallProgress)
	assert.NilError(t, err)

	err = yskService.UpsertYSKCard(context.Background(), utils.DiskInsertNotice)
	assert.NilError(t, err)

	cardList, err = yskService.YskCardList(context.Background())
	assert.NilError(t, err)
	assert.Equal(t, len(cardList), 2)
}
