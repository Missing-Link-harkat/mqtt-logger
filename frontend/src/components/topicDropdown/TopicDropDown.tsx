import { useQuery } from "@tanstack/react-query";

import { fetchTopics } from "../../api/Data";

import { TopicsArray } from "../../schemas/TopicSchema";
import { TopicDropDownProps } from "../../types/Types";

import "./TopicDropDown.css";

const TopicDropDown = ({ setChosenTopic}: TopicDropDownProps) => {
    const {data: topics, isLoading, error} = useQuery<TopicsArray>({
        queryKey: ["topics"],
        queryFn: fetchTopics,
    })

    if (isLoading) return <div>Loading topics...</div>
    if (error) return <div>Error loading topics</div>

    return (
        <select onChange={(e) => setChosenTopic(e.target.value)} defaultValue="">
            <option value="" disabled>
                Select a topic
            </option>
            {topics?.map((topic) => (
                <option key={topic.topic} value={topic.topic}>
                    {topic.topic}
                </option>
            ))}
        </select>
    )
}

export default TopicDropDown;