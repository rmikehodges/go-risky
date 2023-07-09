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
	controller.DetectionRoutes(router)
	controller.ImpactRoutes(router)
	controller.LiabilityRoutes(router)
	controller.MitigationRoutes(router)
	controller.ThreatRoutes(router)
	controller.VulnerabilityRoutes(router)
	controller.ActionRoutes(router)
}

func (controller *PublicController) ActionRoutes(router *gin.Engine) {
	router.GET("/actions", controller.GetActions)
	router.GET("/action", controller.GetAction)
	router.DELETE("/action", controller.DeleteAction)
	router.PATCH("/action", controller.UpdateAction)
	router.POST("/action", controller.CreateAction)
}

func (controller *PublicController) AssetRoutes(router *gin.Engine) {
	router.GET("/assets", controller.GetAssets)
	router.GET("/asset", controller.GetAsset)
	router.DELETE("/asset", controller.DeleteAsset)
	router.PUT("/asset", controller.UpdateAsset)
	router.POST("/assets", controller.CreateAsset)
}

func (controller *PublicController) AttackChainRoutes(router *gin.Engine) {
	router.GET("/attackChains", controller.GetAttackChains)
	router.GET("/attackChain", controller.GetAttackChain)
	router.DELETE("/attackChain", controller.DeleteAttackChain)
	router.PATCH("/attackChain", controller.UpdateAttackChain)
	router.POST("/attackChain", controller.CreateAttackChain)
}

func (controller *PublicController) AttackChainStepRoutes(router *gin.Engine) {
	router.GET("/attackChainSteps", controller.GetAttackChainSteps)
	router.GET("/attackChainStep", controller.GetAttackChainStep)
	router.DELETE("/attackChainStep", controller.DeleteAttackChainStep)
	router.PATCH("/attackChainStep/:id", controller.UpdateAttackChainStep)
	router.POST("/attackChainStep", controller.CreateAttackChainStep)
}

func (controller *PublicController) BusinessRoutes(router *gin.Engine) {
	router.GET("/businesses", controller.GetBusinesses)
	router.GET("/business", controller.GetBusiness)
	router.DELETE("/business", controller.DeleteBusiness)
	router.PATCH("/business", controller.UpdateBusiness)
	router.POST("/business", controller.CreateBusiness)
}

func (controller *PublicController) CapabilityRoutes(router *gin.Engine) {
	router.GET("/capabilities", controller.GetCapabilities)
	router.GET("/capability", controller.GetCapability)
	router.DELETE("/capability", controller.DeleteCapability)
	router.PATCH("/capability", controller.UpdateCapability)
	router.POST("/capabilities", controller.CreateCapability)
}

func (controller *PublicController) DetectionRoutes(router *gin.Engine) {
	router.GET("/detections", controller.GetDetections)
	router.GET("/detection", controller.GetDetection)
	router.DELETE("/detection", controller.DeleteDetection)
	router.PATCH("/detection", controller.UpdateDetection)
	router.POST("/detections", controller.CreateDetection)
}

func (controller *PublicController) ImpactRoutes(router *gin.Engine) {
	router.GET("/impacts", controller.GetImpacts)
	router.GET("/impact/:id", controller.GetImpact)
	router.DELETE("/impact/:id", controller.DeleteImpact)
	router.PATCH("/impact/:id", controller.UpdateImpact)
	router.POST("/impacts", controller.CreateImpact)
}

func (controller *PublicController) LiabilityRoutes(router *gin.Engine) {
	router.GET("/liabilities", controller.GetLiabilities)
	router.GET("/liability", controller.GetLiability)
	// router.GET("/liabilityByImpactId", controller.GetLiabilityByImpactId)
	router.DELETE("/liability", controller.DeleteLiability)
	router.PATCH("/liability", controller.UpdateLiability)
	router.POST("/liabilities", controller.CreateLiability)
}

func (controller *PublicController) MitigationRoutes(router *gin.Engine) {
	router.GET("/mitigations", controller.GetMitigations)
	router.GET("/mitigation/:id", controller.GetMitigation)
	router.DELETE("/mitigation/:id", controller.DeleteMitigation)
	router.PATCH("/mitigation/:id", controller.UpdateMitigation)
	router.POST("/mitigations", controller.CreateMitigation)
}

func (controller *PublicController) ResourceRoutes(router *gin.Engine) {
	router.GET("/resources", controller.GetResources)
	router.GET("/resource/:id", controller.GetResource)
	router.DELETE("/resource/:id", controller.DeleteResource)
	router.PATCH("/resource/:id", controller.UpdateResource)
	router.POST("/resources", controller.CreateResource)
}

func (controller *PublicController) ThreatRoutes(router *gin.Engine) {
	router.GET("/threats", controller.GetThreats)
	router.GET("/threat/:id", controller.GetThreat)
	router.DELETE("/threat/:id", controller.DeleteThreat)
	router.PATCH("/threat/:id", controller.UpdateThreat)
	router.POST("/threats", controller.CreateThreat)
}

func (controller *PublicController) VulnerabilityRoutes(router *gin.Engine) {
	router.GET("/vulnerabilities", controller.GetVulnerabilities)
	router.GET("/vulnerability/:id", controller.GetVulnerability)
	router.DELETE("/vulnerability/:id", controller.DeleteVulnerability)
	router.PATCH("/vulnerability/:id", controller.UpdateVulnerability)
	router.POST("/vulnerabilities", controller.CreateVulnerability)
}
