import { UUID } from 'crypto';

export interface LiabilityOutput {
	id:          UUID 
	name:     string 
	description: string 
    quantity:    number
    cost:        number
    type:        string 
    resourceType: string
	businessId:  UUID 
    detectionId: UUID
    mitigationId: UUID
    resourceId: UUID
    threatId: UUID
    impactId: UUID
	createdAt:   Date
}