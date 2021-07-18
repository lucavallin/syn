import React, { FunctionComponent } from "react";
import { ScrollView, VStack, Center } from "native-base";
import { EventData } from "../../store/eventSlice";

interface EventListProps {
  events: EventData[];
}

export const EventList: FunctionComponent<EventListProps> = ({ events }) => (
  <ScrollView>
    <VStack>
      {events.map((e) => (
        <Center key={e.id} rounded="lg" p={7} bg="primary.400" my={5} mb={3}>
          <div>{e.id}</div>
          <div>{e.created.toString()}</div>
          <div>
            {e.labels.map((l) => (
              <div>{l.description}</div>
            ))}
          </div>
        </Center>
      ))}
    </VStack>
  </ScrollView>
);
