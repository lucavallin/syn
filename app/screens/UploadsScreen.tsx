import * as React from "react";

import { View, Box } from "native-base";
import UploadList from "../components/UploadList";
import firebase from "../services/firebase";

export default function UploadsScreen() {
  const uploads = firebase.firestore().collection("Uploads").get();
  console.log(uploads);
  return;
  return (
    <View>
      <Box>
        <UploadList uploads={uploads} />
      </Box>
    </View>
  );
}
