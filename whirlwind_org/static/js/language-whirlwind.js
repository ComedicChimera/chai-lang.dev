Prism.languages.whirlwind = {
    'comment': [
        {
            pattern: /(\/\*[^(\\\*)]*\*\/)/s
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
    'variant': {
        pattern: /\b(variant)\s*(<[\w ,]+>)\s*(\w+)\s*\{/,
        inside: {
            'function': /(\w+)\s*(?=\{)/,
            'keyword': /\b(variant|int|bool|float|char|str|array|list|map|byte|long|double)\b/,
            'punctuation': /[{,]/,
            'operator': /[<>]/
        }   
    },
    'keyword': [
        // control flow
        {
            pattern: /\b(break|continue|return|yield|for|do|if|elif|else|select|case|default|throw|except|where)\b/
        },
        // declarations
        {
            pattern: /\b(func|async|type|interface|template|struct|constructor|variant)\b/
        },
        // data types
        {
            pattern: /\b(int|bool|float|char|str|array|list|map|byte|long|double)\b/
        },
        // special
        {
            pattern: /\b(new|use|include|delete|cast|export)\b/
        }
    ],
    'function': /[a-zA-Z_]\w*(?=\()/,
    'punctuation':/[{}[\];(),.]/,
    'operator': /[+\-*\/%&\^!|><=?\$@:]/,
    'boolean': /true|false/,
    'modifier': /\b(private|partial|protected|volatile|static)\b/,
    'constant': /this/,
};