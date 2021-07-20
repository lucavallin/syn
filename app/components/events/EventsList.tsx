import * as React from "react";

import {
  Badge,
  Box,
  Image,
  ScrollView,
  Stack,
  Text,
  VStack,
  Center,
  Wrap,
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
                      uri: e.imageUrl,
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
                    <Wrap direction="row">
                      {e.labels.map((l, i) => (
                        <Badge
                          colorScheme="success"
                          mr={2}
                          mt={2}
                          key={`${e.id}.${i}`}
                        >
                          {l.description}
                        </Badge>
                      ))}
                    </Wrap>
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
