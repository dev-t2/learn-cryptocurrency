import React, { FC, memo } from 'react';
import { MaterialCommunityIcons } from '@expo/vector-icons';

const WeatherIcon: FC<weatherIcon> = ({ name, size, color }) => {
  if (name === 'Thunderstorm') {
    return <MaterialCommunityIcons name={'weather-lightning-rainy'} size={size} color={color} />;
  }

  if (name === 'Drizzle') {
    return <MaterialCommunityIcons name={'weather-rainy'} size={size} color={color} />;
  }

  if (name == 'Rain') {
    return <MaterialCommunityIcons name={'weather-pouring'} size={size} color={color} />;
  }

  if (name === 'Snow') {
    return <MaterialCommunityIcons name={'weather-snowy'} size={size} color={color} />;
  }

  if (
    name === 'Mist' ||
    name === 'Smoke' ||
    name === 'Haze' ||
    name === 'Dust' ||
    name === 'Fog' ||
    name === 'Sand' ||
    name === 'Ash'
  ) {
    return <MaterialCommunityIcons name={'weather-fog'} size={size} color={color} />;
  }

  if (name === 'Squall') {
    return <MaterialCommunityIcons name={'weather-tornado'} size={size} color={color} />;
  }

  if (name === 'Tornado') {
    return <MaterialCommunityIcons name={'weather-tornado'} size={size} color={color} />;
  }

  if (name === 'Clouds') {
    return <MaterialCommunityIcons name={'weather-cloudy'} size={size} color={color} />;
  }

  return <MaterialCommunityIcons name={'weather-sunny'} size={size} color={color} />;
};

export default memo(WeatherIcon);
