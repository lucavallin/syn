import firebase from "firebase/app";
import "firebase/firestore";
import "firebase/auth";
import "firebase/storage";
import Constants from "expo-constants";

// Initialize Firebase
const firebaseConfig = {
  apiKey: Constants.manifest?.extra?.firebaseApiKey,
  authDomain: `${Constants.manifest?.extra?.firebaseProjectId}.firebaseapp.com`,
  databaseURL: `https://${Constants.manifest?.extra?.firebaseProjectId}.firebaseio.com`,
  projectId: Constants.manifest?.extra?.firebaseProjectId,
  storageBucket: `${Constants.manifest?.extra?.firebaseProjectId}.appspot.com`,
  messagingSenderId: Constants.manifest?.extra?.firebaseSenderId,
  appId: Constants.manifest?.extra?.firebaseProjectId,
  measurementId: Constants.manifest?.extra?.firebaseMeasurementId,
};

export default firebase.initializeApp(firebaseConfig);
export const firestore = firebase.firestore();
export const storage = firebase.storage();
