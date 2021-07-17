import { createStackNavigator } from "@react-navigation/stack";
import * as React from "react";
import { EventsScreen } from "../screens/EventsScreen";
import { EventsStackParamList } from "../types";

// Each tab has its own navigation stack, you can read more about this pattern here:
// https://reactnavigation.org/docs/tab-based-navigation#a-stack-navigator-for-each-tab
const Events = createStackNavigator<EventsStackParamList>();

export default function EventsNavigator() {
  return (
    <Events.Navigator>
      <Events.Screen
        name="EventsScreen"
        component={EventsScreen}
        options={{ headerTitle: "Events" }}
      />
    </Events.Navigator>
  );
}
