package ysk

import (
	"context"
	"encoding/json"

	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
)

const (
	SERVICENAME = "ysk"
)

func DefineCard(ctx context.Context, cardID string) YSKCard {
	return YSKCard{}
}

func NewYSKCard(ctx context.Context, YSKCard codegen.YSKCard, publish func(context.Context, string, string, map[string]string)) error {
	// do something
	yskCardBodyJSON, _ := json.Marshal(YSKCard)
	publish(ctx, SERVICENAME, "ysk:card:create", map[string]string{"body": string(yskCardBodyJSON)})
	return nil
}

func DeleteCard(ctx context.Context, cardID string, publish func(context.Context, string, string, map[string]string)) error {
	// do something
	publish(ctx, SERVICENAME, "ysk:card:delete", map[string]string{"body": string(``)})
	return nil
}

func TaskWithProgress(card codegen.YSKCard, label string, progress int) codegen.YSKCard {
	if card.Content.BodyProgress != nil {
		card.Content.BodyProgress = &codegen.YSKCardProgress{
			Label: label,
			Value: progress,
		}
		return card
	}
	return card
}
