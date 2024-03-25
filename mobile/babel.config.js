module.exports = function (api) {
  api.cache(true);
  return {
    presets: ['babel-preset-expo'],
    plugins: [
      [
        'module-resolver',
        {
          alias: {
            '@components': './src/Components',
            '@containers': './src/Containers',
            '@screens': './src/Screen',
            "@store": "./src/store",
          },
        },
      ],
    ],
  };
};