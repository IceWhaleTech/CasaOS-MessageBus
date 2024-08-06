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
