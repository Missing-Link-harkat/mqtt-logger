import { useState, useEffect } from "react";
import { useQuery } from "@tanstack/react-query";

import "./DataList.css";
import { fetchData } from "../../api/Data";

import { DataArray } from "../../schemas/DataSchema";
import { ChosenTopic } from "../../types/Types";

const DataList = ({chosenTopic}: ChosenTopic) => {

    const [isFetchEnabled, setIsFetchEnabled] = useState<boolean>(false);

    const {data, isLoading, error } = useQuery<DataArray>({
        queryKey: ["data"],
        queryFn: () => fetchData(chosenTopic),
        enabled: isFetchEnabled,
    })

    useEffect(() => {
        setIsFetchEnabled(false)
    }, [chosenTopic])

    if (isLoading) return <div>Loading topics...</div>
    if (error) return <div>Error loading topics</div>

    const handleButton = (): void => {
        setIsFetchEnabled(true);
    } 

    return (
        <div>
            <button onClick={handleButton}>Fetch</button>
            <div className="data-table">
                <table>
                    <thead>
                        <tr>
                            <th>Topic</th>
                            <th>Value</th>
                            <th>Timestamp</th>
                        </tr>
                    </thead>
                    <tbody>
                        {data?.map((item, index) => (
                            <tr key={index}>
                                <td>{item.topic}</td>
                                <td>{item.value}</td>
                                <td>{item.timestamp}</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </div>
    )
};

export default DataList;
