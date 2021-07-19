import * as React from "react";

import {
  Badge,
  Box,
  HStack,
  Image,
  ScrollView,
  Stack,
  Text,
  VStack,
  Center,
} from "native-base";
import { FunctionComponent } from "react";
import { EventData } from "../../store/eventSlice";

interface EventsProps {
  events: EventData[];
}

export const EventsList: FunctionComponent<EventsProps> = ({ events }) => {
  return (
    <Box>
      <ScrollView>
        <VStack>
          <Center>
            <Box pt={2}>
              {events.map((e) => (
                <Box
                  bg="white"
                  shadow={2}
                  rounded="lg"
                  key={e.id}
                  mb={5}
                  maxWidth="98%"
                >
                  <Image
                    source={{
                      uri: "https://images.unsplash.com/photo-1625772768856-982375e8cd17?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1950&q=80",
                    }}
                    alt="image base"
                    resizeMode="cover"
                    height={150}
                    roundedTop="md"
                  />
                  <Text
                    bold
                    position="absolute"
                    color="white"
                    top={0}
                    m={[4, 4, 8]}
                  >
                    ID: {e.id}
                  </Text>
                  <Stack space={4} p={[4, 4, 8]}>
                    <Text color="gray.400">{e.created.toString()}</Text>
                    <HStack>
                      {e.labels.map((l) => (
                        <Badge colorScheme="success" mr={2} key={l.score}>
                          {l.description}
                        </Badge>
                      ))}
                    </HStack>
                    <Text color="gray.700">Tap for more information.</Text>
                  </Stack>
                </Box>
              ))}
            </Box>
          </Center>
        </VStack>
      </ScrollView>
    </Box>
  );
};
