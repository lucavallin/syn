import React, { FunctionComponent, useEffect } from "react";
import { Center, Box } from "native-base";
import { EventsList } from "../components/events/EventsList";
import { useAppDispatch, useAppSelector } from "../store/hooks";
import { getEvents, selectAllEvents } from "../store/eventsSlice";

export const EventsScreen: FunctionComponent = () => {
  const events = useAppSelector(selectAllEvents);
  const user = useAppSelector((state) => state.user.user);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(getEvents());
  }, [dispatch]);

  return (
    <Box>
      <Center>
        <EventsList events={events} />
      </Center>
    </Box>
  );
};
