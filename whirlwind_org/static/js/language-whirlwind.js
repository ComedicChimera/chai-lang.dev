Prism.languages.whirlwind = {
    'comment': [
        {
            pattern: /\/\*.*\*\//s
        },
        {
            pattern: /\/\/.*/
        }
    ],
    'string': {
        pattern: /"(?:[^"\\']|\\.)*"/,
        greedy: true
    },
    'char': {
            pattern: /'(?:[^"\\']|\\.)*'/,
            greedy: true
    },
    'number': [
        // hex
        {
            pattern: /\b0x[0-9A-F]+/
        },
        // binary
        {
            pattern: /\b0b[10]+/
        },
        // numeric
        {
            pattern: /\b\d+(\.\d+)?/
        }
        
    ],
    'constant': /null|this|@[_a-zA-Z]\w*/,
    'variable': /\$[_a-zA-Z]\w*/,
    'keyword': [
        // control flow
        {
            pattern: /\b(break|continue|return|yield|for|do|if|elif|else|match|select|case|default|throw|with|except)\b/
        },
        // declarations
        {
            pattern: /\b(func|async|module|interface|union|constructor)\b/
        },
        // data types
        {
            pattern: /\b(int|bool|float|char|str|type|array|list|map|byte|long)\b/
        },
        // special
        {
            pattern: /\b(value|new|use|include|delete)\b/
        }
    ],
    'function': /[a-zA-Z_]\w*(?=\()/,
    'punctuation':/[{}[\];(),.:]/,
    'operator': /[+\-*\/%&\^!|><=?\$@]/,
    'boolean': /true|false/,
    'modifier': /\b(private|partial|extern|volatile|property)\b/,
};