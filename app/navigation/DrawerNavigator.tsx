/**
 * Learn more about createBottomTabNavigator:
 * https://reactnavigation.org/docs/bottom-tab-navigator
 */

import * as React from "react";

import {
  createDrawerNavigator,
  DrawerNavigationProp,
} from "@react-navigation/drawer";
import { Box, HStack, Pressable, HamburgerIcon } from "native-base";
import { FunctionComponent } from "react";
import { NavigationProp, useNavigation } from "@react-navigation/native";
import { DrawerNavigatorParamList } from "../types";
import EventsNavigator from "./EventsNavigator";
import { DrawerContent } from "../components/navigation/Drawer";

const DrawerNavigation = createDrawerNavigator<DrawerNavigatorParamList>();

export default function DrawerNavigator() {
  return (
    <Box safeArea flex={1}>
      <DrawerNavigation.Navigator
        initialRouteName="Events"
        drawerContent={() => <DrawerContent />}
      >
        <DrawerNavigation.Screen name="Events" component={EventsNavigator} />
      </DrawerNavigation.Navigator>
    </Box>
  );
}
