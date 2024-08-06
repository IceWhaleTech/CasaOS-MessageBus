package ysk

import (
	"context"
	"encoding/json"

	"github.com/IceWhaleTech/CasaOS-MessageBus/common"
)

func DefineCard(ctx context.Context, cardID string) YSKCard {
	return YSKCard{}
}

func NewYSKCard(ctx context.Context, YSKCard YSKCard, publish func(context.Context, string, string, map[string]string)) error {
	yskCardBodyJSON, _ := json.Marshal(YSKCard)
	publish(ctx,
		common.SERVICENAME,
		common.EventTypeYSKCardUpsert.Name,
		map[string]string{
			common.PropertyTypeCardBody.Name: string(yskCardBodyJSON),
		},
	)
	return nil
}

func DeleteCard(ctx context.Context, cardID string, publish func(context.Context, string, string, map[string]string)) error {
	// do something
	publish(ctx,
		common.SERVICENAME,
		common.EventTypeYSKCardDelete.Name,
		map[string]string{
			common.PropertyTypeCardID.Name: cardID,
		})
	return nil
}
