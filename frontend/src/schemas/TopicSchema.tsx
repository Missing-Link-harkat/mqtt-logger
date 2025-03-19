import { z } from "zod";

export const SingleTopicSchema = z.object({
    topic: z.string().min(1),
})

export const TopicsSchema = z.array(SingleTopicSchema);


export type SingleTopic = z.infer<typeof SingleTopicSchema>
export type TopicsArray = z.infer<typeof TopicsSchema>