module.exports = function (api) {
  api.cache(true);
  return {
    presets: ['babel-preset-expo'],
    plugins: [
      [
        'module-resolver',
        {
          alias: {
            '@components': './src/Component',
            '@containers': './src/Containers',
            '@store': './src/store',
            '@hooks': './src/hooks',
            '@/*': './src/*',
          },
        },
      ],
    ],
  };
};
