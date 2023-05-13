package riskyrouter

import (
	"go-risky/handlers/action"
	"go-risky/handlers/asset"
	"go-risky/handlers/attackChain"
	"go-risky/handlers/attackChainStep"
	"go-risky/handlers/business"
	"go-risky/handlers/capability"
	"go-risky/handlers/impact"
	"go-risky/handlers/liability"
	"go-risky/handlers/mitigation"
	"go-risky/handlers/resource"
	"go-risky/handlers/threat"
	"go-risky/handlers/vulnerability"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(router *gin.Engine) {
	action.ActionRoutes(router)
	attackChainStep.AttackChainStepRoutes(router)
	asset.AssetRoutes(router)
	attackChain.AttackChainRoutes(router)
	business.BusinessRoutes(router)
	capability.CapabilityRoutes(router)
	impact.ImpactRoutes(router)
	liability.LiabilityRoutes(router)
	mitigation.MitigationRoutes(router)
	resource.ResourceRoutes(router)
	threat.ThreatRoutes(router)
	vulnerability.VulnerabilityRoutes(router)
}
