import { UUID } from 'crypto';

export default interface Threat {
	id:          UUID | null
	name:     string 
	description: string  
	businessId:  UUID | null
	createdAt:   Date | null
}