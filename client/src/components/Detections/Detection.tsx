import { UUID } from 'crypto';

export default interface Detection {
	id:          UUID | null
	name:     string 
	description: string  
	businessId:  UUID | null
	implemented: boolean
	createdAt:   Date | null
}