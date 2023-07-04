import { UUID } from 'crypto';

export default interface Liability {
	id:          UUID  | null
	name:     string 
	description: string 
    quantity:    number
    cost:        number
    type:        string 
    resourceType: string
	businessId:  UUID  | null
    detectionId: UUID | null
    mitigationId: UUID | null
    resourceId: UUID  | null
    threatId: UUID | null
    impactId: UUID  | null
	createdAt:   Date | null
}