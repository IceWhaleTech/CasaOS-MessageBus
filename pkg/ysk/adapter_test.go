package ysk_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/pkg/ysk"
	"gotest.tools/assert"
)

func compareJSON(json1, json2 string) (bool, error) {
	var v1, v2 interface{}

	if err := json.Unmarshal([]byte(json1), &v1); err != nil {
		return false, err
	}

	if err := json.Unmarshal([]byte(json2), &v2); err != nil {
		return false, err
	}

	return reflect.DeepEqual(v1, v2), nil
}

func TestJsonIsSame(t *testing.T) {
	domainObject := ysk.YSKCard{
		Id:         "test-card",
		CardType:   ysk.CardTypeTask,
		RenderType: ysk.RenderTypeCardIconTextNotice,
		Content: ysk.YSKCardContent{
			TitleIcon: "test-icon",
			TitleText: "test-title",
		},
	}
	codegenObject := codegen.YSKCard{
		Id:         "test-card",
		CardType:   codegen.YSKCardCardTypeTask,
		RenderType: codegen.YSKCardRenderTypeIconTextNotice,
		Content: codegen.YSKCardContent{
			TitleIcon: "test-icon",
			TitleText: "test-title",
		},
	}

	json1, err := json.Marshal(domainObject)
	assert.NilError(t, err)
	json2, err := json.Marshal(codegenObject)
	assert.NilError(t, err)

	equal, err := compareJSON(string(json1), string(json2))
	assert.NilError(t, err)
	assert.Equal(t, equal, true)
}

func TestWithFunc(t *testing.T) {
	t.Run("WithId", func(t *testing.T) {
		card := ysk.YSKCard{}
		updatedCard := card.WithId("test-id")
		assert.Equal(t, "test-id", updatedCard.Id)
	})

	t.Run("WithTaskContent", func(t *testing.T) {
		card := ysk.YSKCard{}
		updatedCard := card.WithTaskContent(ysk.FileIcon, "Test Title")
		assert.Equal(t, ysk.FileIcon, updatedCard.Content.TitleIcon)
		assert.Equal(t, "Test Title", updatedCard.Content.TitleText)
	})

	t.Run("WithProgress", func(t *testing.T) {
		card := ysk.YSKCard{
			Content: ysk.YSKCardContent{
				BodyProgress: &ysk.YSKCardProgress{
					Label:    "hello",
					Progress: 20,
				},
			},
		}
		updatedCard := card.WithProgress("Progress", 50)
		assert.Assert(t, updatedCard.Content.BodyProgress != nil)
		assert.Equal(t, "Progress", updatedCard.Content.BodyProgress.Label)
		assert.Equal(t, 50, updatedCard.Content.BodyProgress.Progress)
	})

	t.Run("WithList", func(t *testing.T) {
		card := ysk.YSKCard{}
		listItems := []ysk.YSKCardListItem{
			{Icon: ysk.FileIcon, Description: "Item 1", RightText: "Right Text 1"},
			{Icon: ysk.DiskIcon, Description: "Item 2", RightText: "Right Text 2"},
		}
		updatedCard := card.WithList(listItems)
		assert.DeepEqual(t, listItems, updatedCard.Content.BodyList)
	})

	t.Run("WithIconText", func(t *testing.T) {
		card := ysk.YSKCard{}
		updatedCard := card.WithIconText(ysk.StorageIcon, "Storage Description")
		assert.Assert(t, updatedCard.Content.BodyIconWithText != nil)
		assert.Equal(t, ysk.StorageIcon, updatedCard.Content.BodyIconWithText.Icon)
		assert.Equal(t, "Storage Description", updatedCard.Content.BodyIconWithText.Description)
	})

	t.Run("WithFooterActions", func(t *testing.T) {
		card := ysk.YSKCard{}
		actions := []ysk.YSKCardFooterAction{
			{Side: ysk.ActionPositionLeft, Style: "primary", Text: "Confirm", MessageBus: ysk.YSKCardMessageBusAction{Key: "action1", Payload: "payload1"}},
			{Side: ysk.ActionPositionRight, Style: "secondary", Text: "Cancel", MessageBus: ysk.YSKCardMessageBusAction{Key: "action2", Payload: "payload2"}},
		}
		updatedCard := card.WithFooterActions(actions)
		assert.DeepEqual(t, actions, updatedCard.Content.FooterActions)
	})

	t.Run("UpsertFooterAction", func(t *testing.T) {
		card := ysk.YSKCard{
			Content: ysk.YSKCardContent{
				FooterActions: []ysk.YSKCardFooterAction{
					{Side: ysk.ActionPositionLeft, Style: "primary", Text: "Old Button", MessageBus: ysk.YSKCardMessageBusAction{Key: "old", Payload: "old"}},
				},
			},
		}
		newAction := ysk.YSKCardFooterAction{Side: ysk.ActionPositionLeft, Style: "primary", Text: "New Button", MessageBus: ysk.YSKCardMessageBusAction{Key: "new", Payload: "new"}}
		updatedCard := card.UpsertFooterAction(newAction)
		assert.Equal(t, 1, len(updatedCard.Content.FooterActions))
		assert.DeepEqual(t, newAction, updatedCard.Content.FooterActions[0])
	})
}
