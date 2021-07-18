import * as React from "react";

import { createDrawerNavigator } from "@react-navigation/drawer";
import { DrawerNavigatorParamList } from "../types";
import EventsNavigator from "./EventsNavigator";
import { DrawerContent } from "../components/navigation/Drawer";

const DrawerNavigation = createDrawerNavigator<DrawerNavigatorParamList>();

export default function DrawerNavigator() {
  return (
    <DrawerNavigation.Navigator
      initialRouteName="Events"
      drawerContent={(props) => <DrawerContent {...props} />}
    >
      <DrawerNavigation.Screen name="Events" component={EventsNavigator} />
    </DrawerNavigation.Navigator>
  );
}
