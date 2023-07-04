import { UUID } from 'crypto';

export default interface Resource {
  id: UUID | null;
  name: string;
  description: string;
  cost: number;
  unit: string;
  total: number;
  resourceType: string;
  businessId: UUID | null;
  createdAt: Date| null;
}
