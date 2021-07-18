import firebase from "firebase/app";
import "firebase/firestore";
import "firebase/auth";
import Constants from "expo-constants";

// Initialize Firebase
const firebaseConfig = {
  apiKey: Constants.manifest?.extra?.firebaseApiKey,
  authDomain: `${Constants.manifest?.extra?.firebaseProjectId}.firebaseapp.com`,
  databaseURL: `https://${Constants.manifest?.extra?.firebaseProjectId}.firebaseio.com`,
  projectId: Constants.manifest?.extra?.firebaseProjectId,
  storageBucket: `${Constants.manifest?.extra?.firebaseProjectId}.appspot.com`,
  messagingSenderId: "sender-id",
  appId: Constants.manifest?.extra?.firebaseProjectId,
  measurementId: "G-measurement-id",
};

export default firebase.initializeApp(firebaseConfig);
export const firestore = firebase.firestore();
