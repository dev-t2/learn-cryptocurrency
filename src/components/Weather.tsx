import React, { FC, memo } from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { LinearGradient } from 'expo-linear-gradient';

import options from '../options';
import WeatherIcon from './WeatherIcon';

const styles = StyleSheet.create({
  linearGradient: {
    flex: 1,
  },

  top: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  temp: {
    color: '#fff',
    fontSize: 40,
  },

  bottom: {
    flex: 1,
    justifyContent: 'center',
    paddingHorizontal: 32,
  },
  title: {
    color: '#fff',
    fontSize: 32,
    marginBottom: 8,
  },
  subTitle: {
    color: '#fff',
    fontSize: 16,
    fontWeight: 'bold',
  },
});

const Weather: FC<weather> = ({ condition, temp }) => {
  return (
    <LinearGradient colors={options[condition].gradient} style={styles.linearGradient}>
      <View style={styles.top}>
        <WeatherIcon name={condition} size={96} color="#fff" />
        <Text style={styles.temp}>{temp}°</Text>
      </View>

      <View style={styles.bottom}>
        <Text style={styles.title}>{options[condition].title}</Text>
        <Text style={styles.subTitle}>{options[condition].subTitle}</Text>
      </View>
    </LinearGradient>
  );
};

export default memo(Weather);
