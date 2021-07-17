import React, { FunctionComponent } from "react";
import { Events } from "../components/events/Events";
import { useAppSelector } from "../store/hooks";
import { selectAllEvents } from "../store/eventSlice";

export const EventsScreen: FunctionComponent = () => {
  const events = useAppSelector(selectAllEvents);

  return <Events events={events} />;
};
