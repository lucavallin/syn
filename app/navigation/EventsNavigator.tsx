import { createStackNavigator } from "@react-navigation/stack";
import * as React from "react";
import { EventsScreen } from "../screens/EventsScreen";
import { EventsStackParamList } from "../types";

const Events = createStackNavigator<EventsStackParamList>();

export default function EventsNavigator() {
  return (
    <Events.Navigator>
      <Events.Screen name="Events" component={EventsScreen} />
    </Events.Navigator>
  );
}
