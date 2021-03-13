import React, { FC, memo, useCallback, useEffect, useState } from 'react';
import { Alert } from 'react-native';
import * as Location from 'expo-location';
import axios from 'axios';

import { API_KEY } from './apiKey';
import { Loading } from './src/components';

const App: FC = () => {
  const [isLoading, setIsLoading] = useState(true);

  const getWeather = useCallback(async (latitude, longitude) => {
    const { data } = await axios.get(
      `https://api.openweathermap.org/data/2.5/weather?lat=${latitude}&lon=${longitude}&appid=${API_KEY}`
    );
    console.log(data);
  }, []);

  const getLocation = useCallback(async () => {
    try {
      await Location.requestPermissionsAsync();

      const {
        coords: { latitude, longitude },
      } = await Location.getCurrentPositionAsync();

      getWeather(latitude, longitude);

      setIsLoading(false);
    } catch (e) {
      Alert.alert('Location Error', 'Current location could not be found.');
    }
  }, []);

  useEffect(() => {
    getLocation();
  }, []);

  return isLoading ? <Loading /> : null;
};

export default memo(App);
