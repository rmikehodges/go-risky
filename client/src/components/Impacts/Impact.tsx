import { UUID } from 'crypto';

export default interface Mitigation {
	id:          UUID | null
	name:     string 
	description: string  
	businessId:  UUID | null
	threatId: UUID | null
	exploitationCost: number
	mitigationCost: number
	createdAt:   Date | null
}