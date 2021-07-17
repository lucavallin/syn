import * as React from "react";

import { View, Box } from "native-base";
import EventList from "./EventList";
import firebase from "../../services/firebase";
import { useAppSelector, useAppDispatch } from "../../store/hooks";
import { selectCount } from "./eventSlice";

export function EventsScreen() {
  const count = useAppSelector(selectCount);
  const dispatch = useAppDispatch();

  firebase
    .firestore()
    .collection("Uploads")
    .get()
    .then((u) => console.log(u.docs));
  return (
    <View>
      <Box><EventList uploads={events}</Box>
    </View>
  );
}
