import { z } from "zod";

import { TopicsArray, TopicsSchema  } from "../schemas/TopicSchema";

const url = import.meta.env.VITE_API_URL;

export const fetchTopics = async (): Promise<TopicsArray> => {
  const response: Response = await fetch(url + "/topics");
  const data: unknown = await response.json();

  const result: z.SafeParseReturnType<unknown, TopicsArray> = TopicsSchema.safeParse(data);

  if (!result.success) {
    throw new Error("Data validation failed");
  }
  return result.data;
};

export const fetchData = async (): Promise<any> => {
  const response: Response = await fetch(url + "/topics/data");
  return await response.json();
};
