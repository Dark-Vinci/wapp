module.exports = function (api) {
  api.cache(true);
  return {
    presets: ['babel-preset-expo'],
    plugins: [
      [
        'module:react-native-dotenv',
        {
          moduleName: '@env',
          path: '.env',
        },
      ],
      [
        'module-resolver',
        {
          alias: {
            '@components': './src/Components',
            '@containers': './src/Containers',
            '@store': './src/store',
            '@hooks': './src/hooks',
            '@network': './src/network',
            '@types': './src/types',
            '@/*': './src/*',
          },
        },
      ],
    ],
  };
};
