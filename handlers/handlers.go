package handlers

import (
	"go-risky/database"

	"github.com/gin-gonic/gin"
)

type PublicController struct {
	DBManager *database.DBManager
}

func (controller *PublicController) RegisterRoutes(router *gin.Engine) {
	controller.AssetRoutes(router)
	controller.AttackChainRoutes(router)
	controller.AttackChainStepRoutes(router)
	controller.BusinessRoutes(router)
	controller.CapabilityRoutes(router)
	// controller.DetectionRoutes(router)
	controller.ImpactRoutes(router)
	controller.LiabilityRoutes(router)
	// controller.MitigationRoutes(router)
	controller.ThreatRoutes(router)
	controller.VulnerabilityRoutes(router)
	controller.ActionRoutes(router)
}

func (controller *PublicController) ActionRoutes(router *gin.Engine) {
	router.GET("/v0/actions", controller.GetActions)
	router.GET("/v0/action", controller.GetAction)
	router.DELETE("/v0/action", controller.DeleteAction)
	router.PATCH("/v0/action", controller.UpdateAction)
	router.POST("/v0/action", controller.CreateAction)
}

func (controller *PublicController) AssetRoutes(router *gin.Engine) {
	router.GET("/v0/assets", controller.GetAssets)
	router.GET("/v0/asset", controller.GetAsset)
	router.DELETE("/v0/asset", controller.DeleteAsset)
	router.PUT("/v0/asset", controller.UpdateAsset)
	router.POST("/v0/assets", controller.CreateAsset)
}

func (controller *PublicController) AttackChainRoutes(router *gin.Engine) {
	router.GET("/v0/attackChains", controller.GetAttackChains)
	router.GET("/v0/attackChain", controller.GetAttackChain)
	router.DELETE("/v0/attackChain", controller.DeleteAttackChain)
	router.PATCH("/v0/attackChain", controller.UpdateAttackChain)
	router.POST("/v0/attackChain", controller.CreateAttackChain)
}

func (controller *PublicController) AttackChainStepRoutes(router *gin.Engine) {
	router.GET("/v0/attackChainSteps", controller.GetAttackChainSteps)
	router.GET("/v0/attackChainStep", controller.GetAttackChainStep)
	router.DELETE("/v0/attackChainStep", controller.DeleteAttackChainStep)
	router.PATCH("/v0/attackChainStep", controller.UpdateAttackChainStep)
	router.POST("/v0/attackChainStep", controller.CreateAttackChainStep)
}

func (controller *PublicController) BusinessRoutes(router *gin.Engine) {
	router.GET("/v0/businesses", controller.GetBusinesses)
	router.GET("/v0/business", controller.GetBusiness)
	router.DELETE("/v0/business", controller.DeleteBusiness)
	router.PATCH("/v0/business", controller.UpdateBusiness)
	router.POST("/v0/business", controller.CreateBusiness)
}

func (controller *PublicController) CapabilityRoutes(router *gin.Engine) {
	router.GET("/v0/capabilities", controller.GetCapabilities)
	router.GET("/v0/capability", controller.GetCapability)
	router.DELETE("/v0/capability", controller.DeleteCapability)
	router.PATCH("/v0/capability", controller.UpdateCapability)
	router.POST("/v0/capabilities", controller.CreateCapability)
}

// func (controller *PublicController) DetectionRoutes(router *gin.Engine) {
// 	router.GET("/v0/detections", controller.GetDetections)
// 	router.GET("/v0/detection", controller.GetDetection)
// 	router.DELETE("/v0/detection", controller.DeleteDetection)
// 	router.PATCH("/v0/detection", controller.UpdateDetection)
// 	router.POST("/v0/detections", controller.CreateDetection)
// }

func (controller *PublicController) ImpactRoutes(router *gin.Engine) {
	router.GET("/v0/impacts", controller.GetImpacts)
	router.GET("/v0/impact/:id", controller.GetImpact)
	router.DELETE("/v0/impact/:id", controller.DeleteImpact)
	router.PATCH("/v0/impact/:id", controller.UpdateImpact)
	router.POST("/v0/impacts", controller.CreateImpact)
}

func (controller *PublicController) LiabilityRoutes(router *gin.Engine) {
	router.GET("/v0/liabilities", controller.GetLiabilities)
	router.GET("/v0/liability", controller.GetLiability)
	// router.GET("/v0/liabilityByImpactId", controller.GetLiabilityByImpactId)
	router.DELETE("/v0/liability", controller.DeleteLiability)
	router.PATCH("/v0/liability", controller.UpdateLiability)
	router.POST("/v0/liabilities", controller.CreateLiability)
}

// func (controller *PublicController) MitigationRoutes(router *gin.Engine) {
// 	router.GET("/v0/mitigations", controller.GetMitigations)
// 	router.GET("/v0/mitigation/:id", controller.GetMitigation)
// 	router.DELETE("/v0/mitigation/:id", controller.DeleteMitigation)
// 	router.PATCH("/v0/mitigation/:id", controller.UpdateMitigation)
// 	router.POST("/v0/mitigations", controller.CreateMitigation)
// }

func (controller *PublicController) ResourceRoutes(router *gin.Engine) {
	router.GET("/v0/resources", controller.GetResources)
	router.GET("/v0/resource/:id", controller.GetResource)
	router.DELETE("/v0/resource/:id", controller.DeleteResource)
	router.PATCH("/v0/resource/:id", controller.UpdateResource)
	router.POST("/v0/resources", controller.CreateResource)
}

func (controller *PublicController) ThreatRoutes(router *gin.Engine) {
	router.GET("/v0/threats", controller.GetThreats)
	router.GET("/v0/threat/:id", controller.GetThreat)
	router.DELETE("/v0/threat/:id", controller.DeleteThreat)
	router.PATCH("/v0/threat/:id", controller.UpdateThreat)
	router.POST("/v0/threats", controller.CreateThreat)
}

func (controller *PublicController) VulnerabilityRoutes(router *gin.Engine) {
	router.GET("/v0/vulnerabilities", controller.GetVulnerabilities)
	router.GET("/v0/vulnerability/:id", controller.GetVulnerability)
	router.DELETE("/v0/vulnerability/:id", controller.DeleteVulnerability)
	router.PATCH("/v0/vulnerability/:id", controller.UpdateVulnerability)
	router.POST("/v0/vulnerabilities", controller.CreateVulnerability)
}
