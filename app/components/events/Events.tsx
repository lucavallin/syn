import * as React from "react";

import { View, Box } from "native-base";
import { FunctionComponent } from "react";
import { EventList } from "./EventList";
import { EventData } from "../../store/eventSlice";

interface EventsProps {
  events: EventData[];
}

export const Events: FunctionComponent<EventsProps> = ({ events }) => {
  return (
    <View>
      <Box>
        <EventList events={events} />
      </Box>
    </View>
  );
};
