import { UUID } from 'crypto';

export default interface AttackChainStep {
	id : UUID | null
	businessId:  UUID | null
	actionId: UUID | null
	assetId: UUID | null
	attackChainId: UUID | null
	nextStep: UUID | null
	previousStep: UUID | null
	createdAt:   Date | null
}