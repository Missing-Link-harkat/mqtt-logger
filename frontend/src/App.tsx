import { useState } from "react";

import TopicDropDown from "./components/topicDropdown/TopicDropDown";
import DataList from "./components/dataList/DataList";

import "./App.css";

const App = () =>  {

  const [chosenTopic, setChosenTopic] = useState<string>("")
  return (

      <div>
        <h1>mqtt-logger</h1>
        <TopicDropDown setChosenTopic={setChosenTopic}/>
        <DataList chosenTopic={chosenTopic} />
      </div>
  );
}

export default App;
