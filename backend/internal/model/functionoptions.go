package model

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var (
	functionOptionColumns = []string{"functionoptions.functionoptionid", "functionid", "score", "optionname", "description"}
)

type FunctionOption struct {
	FunctionOptionID int32  `json:"functionoptionid"`
	FunctionID       int32  `json:"functionid"`
	Score            int32  `json:"score"`
	OptionName       string `json:"optionname"`
	Description      string `json:"description"`
}

type FindFunctionOptionsInput struct {
	FunctionID *int32
}

func FindFunctionOptions(ctx context.Context, input FindFunctionOptionsInput) ([]*FunctionOption, error) {
	sqlb := stmntBuilder.
		Select(functionOptionColumns...).
		From("functionoptions")

	if input.FunctionID != nil {
		sqlb = sqlb.Where("functionid=?", *input.FunctionID)
	}

	return query(ctx, sqlb, pgx.RowToAddrOfStructByName[FunctionOption])
}
