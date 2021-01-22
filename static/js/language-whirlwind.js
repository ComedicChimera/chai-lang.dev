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
            pattern: /\b0x[0-9A-Fa-f]+/
        },
        // binary
        {
            pattern: /\b0b[10]+/
        },
        // octal
        {
            pattern: /\b0o[0-7]+/
        },
        // numeric
        {
            pattern: /\b\d+[eE]?(\.\d+)?[ul]*/
        }
        
    ],
    'keyword': [
        // control flow
        {
            pattern: /\b(break|continue|return|yield|for|if|elif|else|match|else|case|default|when|nobreak|while|fallthrough|do|to|of|with)\b/
        },
        // declarations
        {
            pattern: /\b(func|async|type|interf|variant|operator)\b/
        },
        // variables
        {
            pattern: /\b(let|const)\b/
        },
        // data types
        {
            pattern: /\b[us]?(int|bool|float|rune|string|byte|long|double|short|nothing|any)\b/
        },
        // special
        {
            pattern: /\b(new|make|from|import|delete|export|await|in|is|as)\b/
        }
    ],
    'function': /[a-zA-Z_]\w*(?=\()/,
    'operator': /:?[+\-*\/%&\^!|><=\~\?]/,
    'punctuation': /[{}[\];(),.:@]/,
    'boolean': /true|false/,
    'modifier': /\b(vol|closed)\b/,
    'constant': /\b(this|null|_|super)\b/,
};