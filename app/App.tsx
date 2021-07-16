import React from "react";
import "react-native-gesture-handler";
import { StatusBar } from "expo-status-bar";
import { NativeBaseProvider } from "native-base";
import Navigation from "./navigation";
import theme from "./components/Theme";

export default function App() {
  return (
    <NativeBaseProvider theme={theme}>
      <Navigation />
      <StatusBar />
    </NativeBaseProvider>
  );
}
