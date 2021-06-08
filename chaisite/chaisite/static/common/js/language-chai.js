// Prism.languages.chai = {
window.language_chai = {
    'comment': [
        {
            pattern: /(#![^(\\\*)]*!#)/
        },
        {
            pattern: /#.*/
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
        pattern: /\bdef\b\s*[a-zA-Z_]\w*\b/,
        inside: {
            'keyword': /\bdef\b/,
            'function': /[a-zA-Z_]\w*/,
        }
    },
    'type-definition': {
        pattern: /\b(type|class|space)\b\s*[a-zA-Z_]\w*\b/,
        inside: {
            'keyword': /\b(type|class|space|for)\b/,
            'entity': /[a-zA-Z_]\w*/,
        }
    },
    'keyword': [
        // control flow
        {
            pattern: /\b(break|continue|return|yield|for|if|elif|else|match|else|when|nobreak|while|fallthrough|do|to|of)\b/
        },
        // declarations
        {
            pattern: /\b(def|async|type|class|space|oper|cons)\b/
        },
        // variables
        {
            pattern: /\blet\b/
        },
        // data types
        {
            pattern: /\b(bool|string|rune|byte|int|uint|any|nothing|u(8|16|32|64)|i(8|16|32|64)|f(32|64))\b/
        },
        // modifier
        {
            pattern: /\b(vol|closed)\b/
        },
        // special
        {
            pattern: /\b(make|from|import|delete|pub|await|in|is|as|fn|then)\b/
        }
    ],
    'prop-access': {
        pattern: /[a-zA-Z_]\w*(\.[a-zA-Z_]\w*)+/,
        inside: {
            'entity': /[a-zA-Z_]\w*/, 
            'punctuation': /\./,
        }
    },
    'type-label': {
        pattern: /:\s*[a-zA-Z_]\w*\b/,
        inside: {
            'punctuation': /:/,
            'entity': /[a-zA-Z_]\w*/, 
        }
    }, 
    "variable": {
        pattern: /\b(\w+)(?:\s*(,)\s*(\w+))*\s*(=|<-)(?!=)/,
        inside: {
            'variable': /\w+/,
            'operator': /=|<-/,
            'punctuation': /,/
        }
    },
    'entity': /[a-zA-Z_]\w*(?=[<{])/,
    'function': /[a-zA-Z_]\w*(?=\()/,
    'operator': /[+\-*\/%&\^!|><=\~\?]/,
    'punctuation': /[{}[\];(),.:@]/,
    'boolean': /true|false/,
    'constant': /\b(null|_|super)\b/,
};