import React, { FC, memo } from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { MaterialCommunityIcons } from '@expo/vector-icons';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  temp: {
    fontSize: 40,
  },
});

interface IWeather {
  condition: condition;
  temp: number;
}

const Weather: FC<IWeather> = ({ condition, temp }) => {
  return (
    <View style={styles.container}>
      <View style={styles.container}>
        <MaterialCommunityIcons name="weather-lightning-rainy" size={96} color="black" />
        <Text style={styles.temp}>{temp}</Text>
      </View>

      <View style={styles.container}></View>
    </View>
  );
};

export default memo(Weather);
