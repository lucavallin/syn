/**
 * Learn more about createBottomTabNavigator:
 * https://reactnavigation.org/docs/bottom-tab-navigator
 */

import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import * as React from "react";

import { TabNavigatorParamList } from "../types";
import EventsNavigator from "./EventsNavigator";

const BottomTabNavigator = createBottomTabNavigator<TabNavigatorParamList>();

export default function TabNavigator() {
  return (
    <BottomTabNavigator.Navigator initialRouteName="Events">
      <BottomTabNavigator.Screen name="Events" component={EventsNavigator} />
    </BottomTabNavigator.Navigator>
  );
}
