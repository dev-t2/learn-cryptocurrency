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
  temp: number;
}

const Weather: FC<IWeather> = ({ temp }) => {
  return (
    <View style={styles.container}>
      <Text>{temp}</Text>
    </View>
  );
};

export default memo(Weather);
