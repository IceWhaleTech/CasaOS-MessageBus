package route

import (
	"net/http"

	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/pkg/ysk"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func (r *APIRoute) DeleteYskCard(ctx echo.Context, id string) error {
	err := r.services.YSKService.DeleteYSKCard(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
			Message: lo.ToPtr(err.Error()),
		})
	}
	return ctx.JSON(http.StatusOK, codegen.ResponseOK{
		Message: lo.ToPtr("success"),
	})
}

func (r *APIRoute) GetYskCard(ctx echo.Context) error {
	cardList, err := r.services.YSKService.YskCardList(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
			Message: lo.ToPtr(err.Error()),
		})
	}

	return ctx.JSON(http.StatusOK, codegen.ResponseGetYSKCardListOK{
		Data: lo.ToPtr(lo.Map(cardList, func(yskCard ysk.YSKCard, _ int) codegen.YSKCard {
			card, err := ysk.ToCodegenYSKCard(yskCard)
			if err != nil {
				return codegen.YSKCard{}
			}
			return card
		})),
	})
}
