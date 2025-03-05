import { useQuery } from "@tanstack/react-query";

import { fetchTopics } from "../../api/Data";

import { Topic } from "../../types/types";

import "./TopicDropDown.css";

const TopicDropDown = ({ onTopicSelect }) => {
    const {data: topics, isLoading, error} = useQuery<Topic[]>('topics', fetchTopics)



    if (isLoading) return <div>Loading topics...</div>
    if (error) return <div>Error loading topics</div>

    return (
        <select>
            <option value="" disabled>
                Select a topic
            </option>
            {topics?.map(topic => (
                <option key={topic.topic} value={topic.topic}}>
                {topic.topic}
            ))}
        </select>
    )
}

export default TopicDropDown;