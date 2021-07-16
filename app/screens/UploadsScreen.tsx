import * as React from "react";

import { View, Box } from "native-base";
import UploadList from "../components/UploadList";

export default function UploadsScreen() {
  return (
    <View>
      <Box>
        <UploadList />
      </Box>
    </View>
  );
}
