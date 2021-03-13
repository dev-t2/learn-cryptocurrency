import React, { memo, useEffect, useState } from 'react';
import { Alert } from 'react-native';
import * as Location from 'expo-location';

import { Loading } from './src/components';

const App = () => {
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    (async () => {
      try {
        await Location.requestPermissionsAsync();

        const {
          coords: { latitude, longitude },
        } = await Location.getCurrentPositionAsync();

        setIsLoading(false);
      } catch (e) {
        Alert.alert('Location Error', 'Current location could not be found.');
      }
    })();
  }, []);

  return isLoading ? <Loading /> : null;
};

export default memo(App);
