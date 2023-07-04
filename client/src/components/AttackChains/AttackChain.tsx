import { UUID } from 'crypto';

export default interface AttackChain {
	id:          UUID | null
	name:     string 
	description: string  
	businessId:  UUID | null
	threatId: UUID | null
	createdAt:   Date | null
}