import { UUID } from 'crypto';

export interface ThreatOutput {
	id:          UUID 
	name:     string 
	description: string  
	businessId:  UUID 
	createdAt:   Date
}