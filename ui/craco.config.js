const CracoLessPlugin = require('craco-less');

module.exports = {
  plugins: [
    {
      plugin: CracoLessPlugin,
      options: {
        lessLoaderOptions: {
          lessOptions: {
            modifyVars: {
              // Custom theme
              // https://github.com/ant-design/ant-design/blob/master/components/style/themes/default.less
              // Colors
              '@primary-color': '@purple-6',
              '@info-color': '@purple-6',
              '@processing-color': '@purple-6',
              // Layout
              '@layout-header-background': '@purple-10',
              '@layout-body-background': '#fff',
              //
              '@border-radius-base': '5px',
            },
            javascriptEnabled: true,
          },
        },
      },
    },
  ],
};
