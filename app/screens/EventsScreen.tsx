import React, { FunctionComponent, useEffect } from "react";
import { Events } from "../components/events/Events";
import { useAppDispatch, useAppSelector } from "../store/hooks";
import { getEvents, selectAllEvents } from "../store/eventSlice";

export const EventsScreen: FunctionComponent = () => {
  const events = useAppSelector(selectAllEvents);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(getEvents());
  }, [dispatch]);

  return <Events events={events} />;
};
