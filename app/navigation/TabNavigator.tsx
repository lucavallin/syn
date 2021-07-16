/**
 * Learn more about createBottomTabNavigator:
 * https://reactnavigation.org/docs/bottom-tab-navigator
 */

import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import * as React from "react";

import { TabNavigatorParamList } from "../types";
import UploadsNavigator from "./UploadsNavigator";

const BottomTabNavigator = createBottomTabNavigator<TabNavigatorParamList>();

export default function TabNavigator() {
  return (
    <BottomTabNavigator.Navigator initialRouteName="Uploads">
      <BottomTabNavigator.Screen name="Uploads" component={UploadsNavigator} />
    </BottomTabNavigator.Navigator>
  );
}
