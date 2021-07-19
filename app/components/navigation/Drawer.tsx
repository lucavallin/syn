import { HamburgerIcon, HStack, Pressable } from "native-base";
import * as React from "react";
import {
  DrawerContentComponentProps,
  DrawerContentScrollView,
  DrawerItemList,
} from "@react-navigation/drawer";
import { FunctionComponent, useCallback } from "react";
import { DrawerActions, useNavigation } from "@react-navigation/native";

export const DrawerContent: FunctionComponent<DrawerContentComponentProps> = (
  props
) => {
  return (
    <DrawerContentScrollView>
      <DrawerItemList {...props} />
    </DrawerContentScrollView>
  );
};

export const Hamburger: FunctionComponent = () => {
  const navigation = useNavigation();

  const toggleDrawer = useCallback(() => {
    navigation.dispatch(DrawerActions.toggleDrawer());
  }, [navigation]);

  return (
    <HStack alignItems="center">
      <Pressable onPress={toggleDrawer}>
        <HamburgerIcon ml={6} size="sm" />
      </Pressable>
    </HStack>
  );
};
