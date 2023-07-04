import { UUID } from 'crypto';

export default interface Asset {
	id:          UUID |	null
	name:     string 
	description: string  
	businessId:  UUID | null
	createdAt:   Date | null
}