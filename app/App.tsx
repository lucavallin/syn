import React from "react";
import "react-native-gesture-handler";
import { Provider } from "react-redux";
import { NativeBaseProvider } from "native-base";
import { PersistGate } from "redux-persist/integration/react";
import AppLoading from "expo-app-loading";
import {
  useFonts,
  Lato_100Thin,
  Lato_100Thin_Italic,
  Lato_300Light,
  Lato_300Light_Italic,
  Lato_400Regular,
  Lato_400Regular_Italic,
  Lato_700Bold,
  Lato_700Bold_Italic,
  Lato_900Black,
  Lato_900Black_Italic,
} from "@expo-google-fonts/lato";
import Navigation from "./navigation";
import theme from "./util/Theme";
import { persistor, store } from "./store/store";

export default function App() {
  const [fontsLoaded] = useFonts({
    Lato_100Thin,
    Lato_100Thin_Italic,
    Lato_300Light,
    Lato_300Light_Italic,
    Lato_400Regular,
    Lato_400Regular_Italic,
    Lato_700Bold,
    Lato_700Bold_Italic,
    Lato_900Black,
    Lato_900Black_Italic,
  });

  if (!fontsLoaded) {
    return <AppLoading />;
  }
  return (
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        <NativeBaseProvider theme={theme}>
          <Navigation />
        </NativeBaseProvider>
      </PersistGate>
    </Provider>
  );
}
