import { z } from "zod";

import { TopicsArray, TopicsSchema  } from "../schemas/TopicSchema";
import { DataArray, DataSchema } from "../schemas/DataSchema";

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

export const fetchData = async (topic: string): Promise<any> => {

  const encodedTopic: string = encodeURIComponent(topic)
  const response: Response = await fetch(url + `/topics/data?topic=${encodedTopic}`);
  
  const data: unknown = await response.json();


  const result: z.SafeParseReturnType<unknown, DataArray> = DataSchema.safeParse(data);

  if (!result.success) {
    throw new Error("Data validation failed");
  }
  return result.data;

};
