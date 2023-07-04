import { UUID } from "crypto";

export default interface Action {
    id: UUID | null
    name: string
    description: string | null
    capabilityId: UUID | null
    vulnerabilityId: UUID | null
    businessId: string | null
    complexity: string
    assetId: UUID | null
    createdAt: Date | null
    }