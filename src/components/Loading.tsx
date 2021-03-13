import React, { FC, memo } from 'react';
import { StyleSheet, Text, View } from 'react-native';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'flex-end',
    paddingVertical: 96,
    paddingHorizontal: 32,
    backgroundColor: '#1976d2',
  },
  text: {
    color: '#fff',
    fontSize: 32,
  },
});

const Loading: FC = () => {
  return (
    <View style={styles.container}>
      <Text style={styles.text}>Getting Started with Weather Application</Text>
    </View>
  );
};

export default memo(Loading);
