import { NavigationContainer } from "@react-navigation/native";
import * as React from "react";

import { createStackNavigator } from "@react-navigation/stack";
import LinkingConfiguration from "./LinkingConfiguration";
import { RootStackParamList } from "../types";
import DrawerNavigator from "./DrawerNavigator";

const RootStack = createStackNavigator<RootStackParamList>();

export default function Navigation() {
  return (
    <NavigationContainer linking={LinkingConfiguration}>
      <RootStack.Navigator screenOptions={{ headerShown: false }}>
        <RootStack.Screen name="Root" component={DrawerNavigator} />
      </RootStack.Navigator>
    </NavigationContainer>
  );
}
