const url = import.meta.env.API_URL;

export const fetchTopics = async (): Promise<any> => {
  const response: Response = await fetch(url + "/topics");
  return await response.json();
};

export const fetchData = async (): Promise<any> => {
  const response: Response = await fetch(url + "/topics/data");
  return await response.json();
};
