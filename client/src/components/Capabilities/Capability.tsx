import { UUID } from "crypto";

interface Capability {
    id: UUID | null
    name: string
    description: string | null
    businessId: UUID | null
    createdAt: Date | null
    }

export default Capability;