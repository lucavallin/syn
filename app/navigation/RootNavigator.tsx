import { createStackNavigator } from "@react-navigation/stack";
import * as React from "react";

import { RootStackParamList } from "../types";
import DrawerNavigator from "./DrawerNavigator";

// A root stack navigator is often used for displaying modals on top of all other content
// Read more here: https://reactnavigation.org/docs/modal
const Stack = createStackNavigator<RootStackParamList>();

export default function RootNavigator() {
  return (
    <Stack.Navigator screenOptions={{ headerShown: false }}>
      <Stack.Screen name="Root" component={DrawerNavigator} />
    </Stack.Navigator>
  );
}
