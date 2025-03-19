import { Dispatch, SetStateAction } from 'react';

export type TopicDropDownProps = {
  setChosenTopic: Dispatch<SetStateAction<string>>;
};
