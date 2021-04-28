const MiniCSSExtractPlugin = require('mini-css-extract-plugin')
const path = require('path')

const mode = process.env.NODE_ENV || 'development'
const prod = mode == 'production'

// List of apps to be compiled using the multi-compiler
const apps = [
    {appName: 'whirlsite', staticDir: 'common'}, 
    {appName: 'home', staticDir: 'home'}
]

var config = {
    resolve: {
        alias: {
            svelte: path.resolve('node_modules', 'svelte'),
            util: 'util',
            sys: 'util',
            path: 'path-browserify',
            vm: 'vm-browserify',
            stream: 'stream-browserify',
            http: 'stream-http',
            os: 'os-browserify/browser',
            buffer: 'buffer',
            https: 'https-browserify',
            assert: 'assert',
            crypto: 'crypto-browserify',
            constants: 'constants-browserify'
        },
        extensions: ['.mjs', '.js', '.svelte'],
        mainFields: ['svelte', 'browser', 'module', 'main']
    },
    module: {
        rules: [
            {
                test: /\.svelte$/,
                use: {
                    loader: 'svelte-loader',
                    options: {
                        emitCss: true,
                        hotReload: true,
                        compilerOptions: {
                            customElement: true
                        }
                    },
                }
            },
            {
                test: /\.scss$/,
                use: [
                    prod ? MiniCSSExtractPlugin.loader : 'style-loader',
                    'css-loader',
                    'sass-loader'
                ]
            },
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: ['@babel/preset-env'],
                        plugins: [
                            '@babel/plugin-proposal-class-properties',
                            '@babel/plugin-transform-runtime'
                        ]
                    }
                }
            }
        ]
    },
    mode,
    plugins: [],
    externals: [
        {
            webpack: {
                commonjs: 'webpack',
                module: 'webpack'
            }
        }
    ],
}

module.exports = apps.map(app => {
    let name = app.name
    let staticDir = app.staticDir

    let uniqueConfig = {
        name: name,
        entry: {},
        output: {
            path: path.join(__dirname, `whirlsite/${name}/static/${staticDir}/dist/`),
            filename: `${name}.bundle.js`,
        }
    }

    uniqueConfig['entry'][name] = `./whirlsite/${name}/static/${staticDir}/src/app.js`

    let appConfig = Object.assign(uniqueConfig, config)
    appConfig['plugins'].push(
        new MiniCSSExtractPlugin({
            filename: `whirlsite/${name}/static/${staticDir}/dist/${name}.bundle.css`
        })
    )

    return appConfig
})