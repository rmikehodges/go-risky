package main

import (
	"go-risky/handlers/action"
	"go-risky/handlers/attackChain"
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

func initializeRouter(router *gin.Engine) {
	action.ActionRoutes(router)
	attackChainStep.attackChainStepRoutes(router)
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

func main() {
	router := gin.Default()

	initializeRouter(router)

	router.Run("localhost:8081")
}
