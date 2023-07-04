import { UUID } from 'crypto';

export default interface Mitigation {
	id:          UUID | null
	name:     string 
	description: string  
	businessId:  UUID | null
	actionId: UUID | null
	implemented: boolean
	createdAt:   Date | null
}