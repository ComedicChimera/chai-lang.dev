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
    'constant': /null|this|@[_a-zA-Z]\w*/,
    'variable': /\$[_a-zA-Z]\w*/,
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
            pattern: /\b(break|continue|return|yield|for|do|if|elif|else|match|select|case|default|throw|with|except)\b/
        },
        // declarations
        {
            pattern: /\b(func|async|module|interface|template|struct|constructor|variant)\b/
        },
        // data types
        {
            pattern: /\b(int|bool|float|char|str|array|list|map|byte|long|double)\b/
        },
        // special
        {
            pattern: /\b(value|new|use|include|delete|cast|export)\b/
        }
    ],
    'function': /[a-zA-Z_]\w*(?=\()/,
    'punctuation':/[{}[\];(),.]/,
    'operator': /[+\-*\/%&\^!|><=?\$@:]/,
    'boolean': /true|false/,
    'modifier': /\b(private|partial|protected|volatile|property|mut|uniform)\b/,
};