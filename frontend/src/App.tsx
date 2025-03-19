import { useState } from "react";

import TopicDropDown from "./components/topicDropdown/TopicDropDown";

import "./App.css";

const App = () =>  {

  const [chosenTopic, setChosenTopic] = useState<string>("")
  return (

      <div>
        <h1>mqtt-logger</h1>
        <TopicDropDown setChosenTopic={setChosenTopic}/>
      </div>
  );
}

export default App;
