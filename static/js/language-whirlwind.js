Prism.languages.whirlwind = {
    'comment': [
        {
            pattern: /(\/\*[^(\\\*)]*\*\/)/
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
            pattern: /\b\d+(\.\d+)?[uld]*/
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
            pattern: /\b[us]?(int|bool|float|char|str|byte|long|double|short)\b/
        },
        // special
        {
            pattern: /\b(new|make|from|include|delete|export|await|then|is|as)\b/
        }
    ],
    'function': /[a-zA-Z_]\w*(?=\()/,
    'operator': /:?[+\-*\/%&\^!|><=\~\?]/,
    'punctuation': /[{}[\];(),.:#@]/,
    'boolean': /true|false/,
    'modifier': /\b(vol|static|dyn|own)\b/,
    'constant': /\b(this|null|_|super|value)\b/,
};