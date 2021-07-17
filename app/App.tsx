import React from "react";
import "react-native-gesture-handler";
import { Provider } from "react-redux";
import { NavigationContainer } from "@react-navigation/native";
import { NativeBaseProvider } from "native-base";
import Navigation from "./navigation";
import theme from "./util/Theme";
import { store } from "./store/store";

export default function App() {
  return (
    <Provider store={store}>
      <NavigationContainer>
        <NativeBaseProvider theme={theme}>
          <Navigation />
        </NativeBaseProvider>
      </NavigationContainer>
    </Provider>
  );
}
