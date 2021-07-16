import React from "react";
import "react-native-gesture-handler";
import { NativeBaseProvider } from "native-base";
import Navigation from "./navigation";
import theme from "./util/Theme";

export default function App() {
  return (
    <NativeBaseProvider theme={theme}>
      <Navigation />
    </NativeBaseProvider>
  );
}
