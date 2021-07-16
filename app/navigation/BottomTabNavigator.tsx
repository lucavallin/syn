/**
 * Learn more about createBottomTabNavigator:
 * https://reactnavigation.org/docs/bottom-tab-navigator
 */

import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { createStackNavigator } from "@react-navigation/stack";
import * as React from "react";

import { WarningIcon } from "native-base";
import UploadsScreen from "../screens/UploadsScreen";
import { BottomTabParamList, UploadsStackParamList } from "../types";

const BottomTab = createBottomTabNavigator<BottomTabParamList>();

export default function BottomTabNavigator() {
  return (
    <BottomTab.Navigator initialRouteName="Uploads">
      <BottomTab.Screen
        name="Uploads"
        component={UploadsNavigator}
        options={{
          tabBarIcon: () => <WarningIcon />,
        }}
      />
    </BottomTab.Navigator>
  );
}

// Each tab has its own navigation stack, you can read more about this pattern here:
// https://reactnavigation.org/docs/tab-based-navigation#a-stack-navigator-for-each-tab
const UploadsStack = createStackNavigator<UploadsStackParamList>();

function UploadsNavigator() {
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
