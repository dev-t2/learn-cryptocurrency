import React, { FC, memo, useCallback, useEffect, useState } from 'react';
import { Alert, StatusBar } from 'react-native';
import * as Location from 'expo-location';
import axios from 'axios';

import API_KEY from './apiKey';
import { Loading, Weather } from './src/components';

const App: FC = () => {
  const [isLoading, setIsLoading] = useState(true);
  const [condition, setCondition] = useState<condition>('Clear');
  const [temp, setTemp] = useState(0);

  const getWeather = useCallback(async (latitude: number, longitude: number) => {
    const {
      data: {
        main: { temp },
        weather,
      },
    } = await axios.get(
      `https://api.openweathermap.org/data/2.5/weather?lat=${latitude}&lon=${longitude}&units=metric&appid=${API_KEY}`
    );

    setCondition(weather[0].main);
    setTemp(Math.round(temp));
  }, []);

  const getLocation = useCallback(async () => {
    try {
      await Location.requestPermissionsAsync();

      const {
        coords: { latitude, longitude },
      } = await Location.getCurrentPositionAsync();

      await getWeather(latitude, longitude);

      setIsLoading(false);
    } catch (e) {
      Alert.alert('Location Error', 'Current location could not be found.');
    }
  }, []);

  useEffect(() => {
    getLocation();
  }, []);

  return (
    <>
      <StatusBar translucent backgroundColor="transparent" barStyle="light-content" />

      {isLoading ? <Loading /> : <Weather condition={condition} temp={temp} />}
    </>
  );
};

export default memo(App);
