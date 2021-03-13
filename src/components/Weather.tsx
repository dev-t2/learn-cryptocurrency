import React, { FC, memo } from 'react';
import { StyleSheet, Text, View } from 'react-native';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
});

interface IWeather {
  condition: condition;
  temp: number;
}

const Weather: FC<IWeather> = ({ condition, temp }) => {
  return (
    <View style={styles.container}>
      <Text>{temp}</Text>
    </View>
  );
};

export default memo(Weather);
