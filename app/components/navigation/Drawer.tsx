import {
  Box,
  Divider,
  HamburgerIcon,
  HStack,
  Icon,
  Pressable,
  Text,
  VStack,
} from "native-base";
import * as React from "react";
import { DrawerContentScrollView } from "@react-navigation/drawer";
import { MaterialCommunityIcons } from "@expo/vector-icons";
import { FunctionComponent, useCallback } from "react";
import { DrawerActions, useNavigation } from "@react-navigation/native";

export const DrawerContent: FunctionComponent = () => {
  const navigation = useNavigation();

  return (
    <DrawerContentScrollView safeArea>
      <VStack space={6} my={2} mx={1}>
        <Box px={4}>
          <Text bold color="gray.700">
            Mail
          </Text>
          <Text fontSize={14} mt={1} color="gray.500" fontWeight={500}>
            john_doe@gmail.com
          </Text>
        </Box>
        <VStack divider={<Divider />} space={4}>
          <VStack space={3}>
            {navigation.routeNames.map((name, index) => (
              <Pressable
                px={5}
                py={3}
                rounded="md"
                bg={
                  index === props.state.index
                    ? "rgba(6, 182, 212, 0.1)"
                    : "transparent"
                }
                onPress={(event) => {
                  props.navigation.navigate(name);
                }}
              >
                <HStack space={7} alignItems="center">
                  <Icon
                    color={
                      index === props.state.index ? "primary.500" : "gray.500"
                    }
                    size={5}
                    as={<MaterialCommunityIcons name="email" />}
                  />
                  <Text
                    fontWeight={500}
                    color={
                      index === props.state.index ? "primary.500" : "gray.700"
                    }
                  >
                    {name}
                  </Text>
                </HStack>
              </Pressable>
            ))}
          </VStack>
        </VStack>
      </VStack>
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
        <HamburgerIcon ml={2} size="sm" />
      </Pressable>
    </HStack>
  );
};
