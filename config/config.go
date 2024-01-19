package config

import (
	"time"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	PORT             = 8080
	PlusModels       = garray.NewStrArrayFrom([]string{"gpt-4", "gpt-4-browsing", "gpt-4-plugins", "gpt-4-mobile", "gpt-4-code-interpreter", "gpt-4-dalle", "gpt-4-gizmo", "gpt-4-magic-create"})
	NormalModels     = garray.NewStrArrayFrom([]string{"text-davinci-002-render-sha"})
	ForbiddenWords   = []string{}    // 禁止词
	LIMIT            = 40            // 限制次数
	PER              = time.Hour * 3 // 限制时间
	NORMALMODELLIMIT = 40            // Normal模型限制次数
	NORMALMODELPER   = time.Hour * 3 // Normal模型限制时间

)

func init() {
	ctx := gctx.GetInitCtx()
	port := g.Cfg().MustGetWithEnv(ctx, "PORT").Int()
	if port > 0 {
		PORT = port
	}
	g.Log().Info(ctx, "PORT:", PORT)

	limit := g.Cfg().MustGetWithEnv(ctx, "LIMIT").Int()
	if limit > 0 {
		LIMIT = limit
	}
	g.Log().Info(ctx, "LIMIT:", LIMIT)
	per := g.Cfg().MustGetWithEnv(ctx, "PER").Duration()
	if per > 0 {
		PER = per
	}

	// 普通模型
	normalModelLimit := g.Cfg().MustGetWithEnv(ctx, "NORMALMODELLIMIT").Int()
	if normalModelLimit > 0 {
		NORMALMODELLIMIT = normalModelLimit
	}
	g.Log().Info(ctx, "NORMALMODELLIMIT:", NORMALMODELLIMIT)
	normalModelPer := g.Cfg().MustGetWithEnv(ctx, "NORMALMODELPER").Duration()
	if normalModelPer > 0 {
		NORMALMODELPER = normalModelPer
	}
	g.Log().Info(ctx, "NORMALMODELPER:", NORMALMODELPER)
}
