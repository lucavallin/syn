import React, { FunctionComponent, useEffect } from "react";
import { useNavigation } from "@react-navigation/native";
import { StackHeaderLeftButtonProps } from "@react-navigation/stack";
import { Events } from "../components/events/Events";
import { useAppDispatch, useAppSelector } from "../store/hooks";
import { getEvents, selectAllEvents } from "../store/eventSlice";
import { Hamburger } from "../components/navigation/Drawer";

export const EventsScreen: FunctionComponent = () => {
  const events = useAppSelector(selectAllEvents);
  const dispatch = useAppDispatch();
  const navigation = useNavigation();

  useEffect(() => {
    dispatch(getEvents());
    navigation.setOptions({
      headerLeft: <Hamburger />,
    });
  });

  return <Events events={events} />;
};
