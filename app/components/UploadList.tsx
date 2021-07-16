import React from "react";
import { ScrollView, VStack, Center } from "native-base";

export default () => {
  return (
    <ScrollView>
      <VStack>
        {[
          "one",
          "two",
          "three",
          "four",
          "five",
          "six",
          "seven",
          "eight",
          "nine",
          "ten",
          "eleven",
        ].map((val) => (
          <Center rounded="lg" p={7} bg="primary.400" my={5} mb={3}>
            {val}
          </Center>
        ))}
      </VStack>
    </ScrollView>
  );
};
