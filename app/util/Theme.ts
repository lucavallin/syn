import { extendTheme } from "native-base";

export default extendTheme({
  fontConfig: {
    Lato: {
      100: {
        normal: "Lato_100Thin",
        italic: "Lato_100Thin_Italic",
      },
      300: {
        normal: "Lato_300Light",
        italic: "Lato_300Light_Italic",
      },
      400: {
        normal: "Lato_400Regular",
        italic: "Lato_400Regular_Italic",
      },
      700: {
        normal: "Lato_700Bold",
        italic: "Lato_700Bold_Italic",
      },
      900: {
        normal: "Lato_900Black",
        italic: "Lato_900Black_Italic",
      },
    },
  },
  fonts: {
    heading: "Lato",
    body: "Lato",
    mono: "Lato",
  },
});
