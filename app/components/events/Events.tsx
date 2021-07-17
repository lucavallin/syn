import * as React from "react";

import { View, Box } from "native-base";
import { FunctionComponent } from "react";
import { EventList } from "./EventList";

export const Events: FunctionComponent = () => {
  return (
    <View>
      <Box>
        <EventList events={["ciao", "mondo"]} />
      </Box>
    </View>
  );
};
