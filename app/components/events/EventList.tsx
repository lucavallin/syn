import React, { FunctionComponent } from "react";
import { ScrollView, VStack, Center } from "native-base";

interface EventListProps {
  events: Array<string>;
}

export const EventList: FunctionComponent<EventListProps> = ({ events }) => (
  <ScrollView>
    <VStack>
      {events.map((e) => (
        <Center key={e} rounded="lg" p={7} bg="primary.400" my={5} mb={3}>
          {e}
        </Center>
      ))}
    </VStack>
  </ScrollView>
);
