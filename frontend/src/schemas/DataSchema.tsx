import { z } from "zod";

export const SingleDataSchema = z.object({
    topic: z.string(),
    value: z.string(),
    timestamp: z.string(),
})

export const DataSchema = z.array(SingleDataSchema);

export type SingleData = z.infer<typeof SingleDataSchema>
export type DataArray = z.infer<typeof DataSchema>