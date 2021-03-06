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
    'function-definition': {
        pattern: /\b(func|async)\b\s*[a-zA-Z_]\w*\b/,
        inside: {
            'keyword': /\b(func|async)\b/,
            'function-name': /[a-zA-Z_]\w*/,
        }
    },
    'type-definition': {
        pattern: /\b(type|interf)\b\s*[a-zA-Z_]\w*\b/,
        inside: {
            'keyword': /\b(type|interf|for|to)\b/,
            'class-name': /[a-zA-Z_]\w*/,
        }
    },
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
        // modifier
        {
            pattern: /\b(vol|closed|own|global)\b/
        },
        // special
        {
            pattern: /\b(new|make|from|import|delete|export|await|in|is|as|local|region)\b/
        }
    ],
    'static-access': {
        pattern: /[a-zA-Z_]\w*(::[a-zA-Z_]\w*)+/,
        inside: {
            'class-name': /[a-zA-Z_]\w*/, 
            'punctuation': /::/,
        }
    },
    'type-label': {
        pattern: /:\s*[a-zA-Z_]\w*\b/,
        inside: {
            'punctuation': /:/,
            'class-name': /[a-zA-Z_]\w*/, 
        }
    }, 
    'class-name': /[a-zA-Z_]\w*(?=[<{])/,
    'function': /[a-zA-Z_]\w*(?=\()/,
    'operator': /:?[+\-*\/%&\^!|><=\~\?]/,
    'punctuation': /[{}[\];(),.:@]/,
    'boolean': /true|false/,
    'constant': /\b(this|null|_|super)\b/,
};