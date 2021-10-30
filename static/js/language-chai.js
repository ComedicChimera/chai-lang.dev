Prism.languages.chai = {
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
            pattern: /\b\d+(\.\d+)?([eE]\d+)?[ul]*/
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
    'package-import': {
        pattern: /\b(from|import)\s+[a-zA-Z_]\w*(\.[a-zA-Z_]\w*)*/,
        inside: {
            'keyword': /\b(from|import)\b/,
            'entity': /[a-zA-Z_]\w*/, 
            'punctuation': /\./,
        }
    },
    'keyword': [
        // control flow
        {
            pattern: /\b(break|continue|return|for|if|elif|else|match|case|else|when|after|while|fallthrough|do|end)\b/
        },
        // declarations
        {
            pattern: /\b(def|async|type|class|space|oper|union)\b/
        },
        // variables
        {
            pattern: /\b(let|const)\b/
        },
        // data types
        {
            pattern: /\b(bool|string|rune|byte|nothing|u(8|16|32|64)|i(8|16|32|64)|f(32|64))\b/
        },
        // modifier
        {
            pattern: /\b(vol|closed)\b/
        },
        // special
        {
            pattern: /\b(from|import|pub|await|in|is|as|fn)\b/
        }
    ],
    'function': /[a-zA-Z_]\w*(?=\()/,
    'prop-access': {
        pattern: /(\.)[a-zA-Z_]\w*/,
        lookbehind: true,
        inside: {
            'property': /[a-zA-Z_]\w*/, 
        }
    },
    'package-access': {
        pattern: /[a-zA-Z_]\w*\./,
        inside: {
            'entity': /[a-zA-Z_]\w*/, 
            'punctuation': /\./,
        }
    },
    'type-label': {
        pattern: /(:)\s*[a-zA-Z_]\w*\b/,
        inside: {
            'entity': /[a-zA-Z_]\w*/, 
        },
        lookbehind: true,
    }, 
    'typed-variable': {
        pattern: /\b(\w+)(?:\s*(,)\s*(\w+))*\s*(:)/,
        inside: {
            'punctuation': /:,/,
            'variable': /[a-zA-Z_]\w*/, 
        },
    },
    'stmt-variable': {
        pattern: /(?!\))\s*\b(\w+)(?:\s*(,)\s*(\w+))*\s*([+\-*\\%]?=|<-)(?!=)/,
        inside: {
            'variable': /\w+/,
            'operator': /=|<-/,
            'punctuation': /,/
        }
    },
    'entity': /[a-zA-Z_]\w*(?=\{)|\b(List|Dict|Array|Seq|Iter|Option|Result|Future|Strand)\b/,
    'operator': /[+\-*\/%&\^!|><=\~\?]/,
    'punctuation': /[{}[\];(),.:@]/,
    'boolean': /\b(true|false)\b/,
    'constant': /\b(null|_)\b/,
};