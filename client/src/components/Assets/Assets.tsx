import { UUID } from 'crypto';

export interface AssetOutput {
	id:          UUID 
	name:     string 
	description: string  
	businessId:  UUID 
	createdAt:   Date
}