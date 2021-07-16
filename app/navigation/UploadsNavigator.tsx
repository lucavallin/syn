import { createStackNavigator } from "@react-navigation/stack";
import * as React from "react";

import UploadsScreen from "../screens/UploadsScreen";
import { UploadsStackParamList } from "../types";

// Each tab has its own navigation stack, you can read more about this pattern here:
// https://reactnavigation.org/docs/tab-based-navigation#a-stack-navigator-for-each-tab
const UploadsStack = createStackNavigator<UploadsStackParamList>();

export default function UploadsNavigator() {
  return (
    <UploadsStack.Navigator>
      <UploadsStack.Screen
        name="UploadsScreen"
        component={UploadsScreen}
        options={{ headerTitle: "Uploads" }}
      />
    </UploadsStack.Navigator>
  );
}
