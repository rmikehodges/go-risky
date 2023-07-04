import { UUID } from "crypto";

interface Business {
    id: UUID | null
    name: string
    revenue: number
    createdAt: Date | null
    }

export default Business;