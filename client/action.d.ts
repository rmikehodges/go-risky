//convert ActionInput and ActionOutput from handlers/action/action.go to typescript interface
interface ActionInput {
    id: string
    name: string
    description: string
    capabilityId: string
    vulnerabilityId: string
    businessId: string
    complexity: string
    assetId: string
    createdAt: Date
  }

interface ActionOutput {
    id: string
    name: string
    description: string
    capabilityId: string
    vulnerabilityId: string
    businessId: string
    complexity: string
    assetId: string
    createdAt: Date
    }