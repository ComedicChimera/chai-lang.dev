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
            'keyword': /\b(variant|int|bool|float|char|str|byte|long|double)\b/,
            'punctuation': /[{,]/,
            'operator': /[<>]/
        }   
    },
    'keyword': [
        // control flow
        {
            pattern: /\b(break|continue|return|yield|for|if|elif|else|select|case|default|when|after)\b/
        },
        // declarations
        {
            pattern: /\b(func|async|type|interf|struct|constructor|variant|operator|with)\b/
        },
        // variables
        {
            pattern: /\b(let|const)\b/
        },
        // data types
        {
            pattern: /\bu?(int|bool|float|char|str|byte|long|double)\b/
        },
        // special
        {
            pattern: /\b(new|make|from|include|delete|export|await|ref|then|val|is|as)\b/
        }
    ],
    'function': /[a-zA-Z_]\w*(?=\()/,
    'punctuation':/[{}[\];(),.]/,
    'operator': /[+\-*\/%&\^!|><=?\~]/,
    'boolean': /true|false/,
    'modifier': /\b(vol|static|own)\b/,
    'constant': /\b(this|null|_)\b/,
};